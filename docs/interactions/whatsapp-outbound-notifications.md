# Digital WhatsApp Outbound Notifications

Use the RingCX Digital API to send a business-initiated WhatsApp notification from an approved message template. Outbound notifications let an integration start a WhatsApp conversation for use cases such as order updates, appointment reminders, and account notifications.

RingCX sends the notification asynchronously through the WhatsApp channel. The API response represents the content created in RingCX; it does not confirm delivery to the recipient.

## Before you begin

You need:

- A configured and active WhatsApp channel.
- A WhatsApp message template that is approved and available on that channel.
- An [API access token](../basics/access-token.md) associated with a user who has **Read** and **Initiate discussion** permissions on the WhatsApp channel.
- The recipient's phone number in international format, beginning with `+` or `00`.

Use the API hostname assigned to your RingCX Digital account. The examples below use `https://[YOUR DOMAIN].api.digital.ringcentral.com`.

## Outbound notification workflow

1. Retrieve the WhatsApp channel and its available message templates.
2. Select a template name and language.
3. Supply values for any variables defined by the template.
4. Create the outbound content.
5. Use content events or the content resource to monitor subsequent synchronization status.

## Retrieve available templates

Retrieve the WhatsApp channel before sending a notification:

```bash
curl --request GET \
  --url "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/content_sources/{source_id}" \
  --header "Authorization: Bearer {access_token}" \
  --header "Accept: application/json"
```

The `template_messages` array contains the approved templates available for outbound notifications. Each entry identifies one template and language combination.

```json
{
  "id": "{source_id}",
  "name": "Customer notifications",
  "type": "whats_app",
  "template_messages": [
    {
      "name": "order_update",
      "language": "en",
      "components": [
        {
          "type": "body",
          "text": "Hello {{1}}, your order {{2}} has shipped."
        }
      ]
    }
  ]
}
```

Use the `name` and `language` values exactly as returned. The component definitions show which values the notification must provide.

## Send a notification without variables

Create outbound content with `POST /1.0/contents`. A template without variables does not require a `components` value.

```bash
curl --request POST \
  --url "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents" \
  --header "Authorization: Bearer {access_token}" \
  --header "Content-Type: application/json" \
  --data '{
    "source_id": "{source_id}",
    "to": "+14155550123",
    "template_name": "delivery_notice",
    "template_language": "en"
  }'
```

Do not include `body` or `in_reply_to_id` when initiating a WhatsApp outbound notification. The message body and supported interactive elements come from the approved template.

## Supply template variables

Use `components` when the selected template contains variables. Component types and their parameters must match the template definition returned by the channel.

For a body such as `Hello {{1}}, your order {{2}} has shipped.`, send one `body` component with the values in placeholder order:

```bash
curl --request POST \
  --url "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents" \
  --header "Authorization: Bearer {access_token}" \
  --header "Content-Type: application/json" \
  --data '{
    "source_id": "{source_id}",
    "to": "+14155550123",
    "template_name": "order_update",
    "template_language": "en",
    "components": [
      {
        "type": "body",
        "parameters": [
          {
            "type": "text",
            "text": "Alex"
          },
          {
            "type": "text",
            "text": "RC-10042"
          }
        ]
      }
    ]
  }'
```

### Component parameters

| Template element | Component value | Parameter value |
| --- | --- | --- |
| Body variable | `type: body` | One `type: text` parameter for each placeholder, in placeholder order. |
| Text header variable | `type: header` | One `type: text` parameter for each placeholder, in placeholder order. |
| Image header | `type: header` | One `type: image` parameter containing `image.attachment_id`. |
| Dynamic URL button | `type: button`, `sub_type: url` | One `type: text` parameter containing the dynamic URL suffix. |
| Quick-reply button | `type: button`, `sub_type: quick_reply` | One `type: payload` parameter containing the value returned when the recipient selects the reply. |

Button indexes are zero-based strings and must identify the corresponding button in the template.

## Send an image-header notification

First [upload the image](../basics/uploads.md), then provide the returned attachment ID in the header component. WhatsApp outbound notification headers support JPG, JPEG, and PNG images.

```json
{
  "source_id": "{source_id}",
  "to": "+14155550123",
  "template_name": "product_delivery",
  "template_language": "en",
  "components": [
    {
      "type": "header",
      "parameters": [
        {
          "type": "image",
          "image": {
            "attachment_id": "{attachment_id}"
          }
        }
      ]
    }
  ]
}
```

## Request parameters

| Parameter | Type | Required | Description |
| --- | --- | --- | --- |
| `source_id` | String | Yes | ID of the WhatsApp channel that sends the notification. |
| `to` | String | Yes | Recipient's phone number in international format, beginning with `+` or `00`. |
| `template_name` | String | Yes | Name of an available template returned in `template_messages`. |
| `template_language` | String | Yes | Language code for the selected template, exactly as returned in `template_messages`. |
| `components` | Array | Conditional | Values required by the selected template. Omit it when the template has no variables or media header. |

## Handle the response

A successful request returns `200 OK` and the created content resource. Sending to WhatsApp happens asynchronously. Monitor the content's `synchronization_status` and `synchronization_error` fields, or subscribe to the [`content.exported` webhook event](../webhooks/events.md), to track processing after creation.

Common request failures include:

| HTTP status | Cause |
| --- | --- |
| `403` | The user associated with the access token cannot initiate discussions on the WhatsApp channel. |
| `422` | A required field is missing, the template name or language is unavailable, or the supplied components do not match the template. |

Validate template availability immediately before sending when templates can change independently of your integration. Retry transient synchronization failures with backoff, but correct authorization and validation errors before retrying.

## Related resources

- [Obtain an access token](../basics/access-token.md)
- [Upload files](../basics/uploads.md)
- [Content webhook events](../webhooks/events.md)
- [API reference](https://developers.ringcentral.com/engage/api-reference/)
