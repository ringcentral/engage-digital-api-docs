# About Apple Authenticate Structured Messages

This structured message provides a way to send an OAuth2 authentication request to the customer. See [Channel capabilities](../structured-messages/#channel-capabilities) to know on which channel you can use this structured message.

## Prerequisites
* Set up OAuth URL and Token URL in "End User Authentication” on your Business Chat Account.

```bash tab="Request"
curl -X POST "https://[YOUR DOMAIN].api.engagement.dimelo.com/1.0/contents"
```

```json tab="JSON Body"
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "body": "Please Log In",
  "structured_content": {
    "type": "authenticate",
    "attachment_id": "<attachment_id>",
    "response_body": "Response",
  }
}
```

### Primary Parameters

| API Property | Type | Description |
|-|-|-|
| **`source_id`** | String | ID of the source. |
| **`in_reply_to_id`** | String | ID of the message being replied to. |
| **`body`** | String | The authenticate structured content body. |
| **`structured_content`** | Object | Payload of the structured message. |
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "authenticate". |
| **`structured_content.response_body`** | String | **Optional**. Field to be the body of the user’s response. If absent, default is the structured message body. |
| **`structured_content.attachment_id`** | Object | **Optional**. Existing attachment id used to decorate the authenticate structured message with an image.<br>Should be jpg, jpeg or png.<br>Should be less than 5MB. |

### Webhook

The customer will receive a message including the body and a button that will open the OAuth login page. Once authenticated, the token will be sent back by Apple and will be accessible through the Engage Digital webhook of the imported message in the "authenticate” part. The token is sent back. See following page for an example of the webhook payload format.

```json tab="Webhook Payload Format"
{
  "id":"0b44eb19f120e3230942572f",
  "domain_id":"48cc6703bdae1462ce06a555",
  "events": [
    {
      "type":"content.imported",
      "id":"5c3f6bbdd02c8e3a585fb9a3",
      "issued_at":"2019-01-16T17:37:01.192Z",
      "resource": {
        "type":"apple_business_chat/authenticate_response",
        "id":"5c3f6bbcd02c8e3a585fb9a1",
        "metadata": {
          "approval_required":false,
          "author_id":"5c3721a4d02c8e735192254a",
          "body":"This is an authentication test",
          "body_input_format":"text",
          "category_ids":[],
          "creator_id":null,
          "date":"2019-01-16",
          "first_in_thread":"false",
          "foreign_categories":[],
          "foreign_id":"e70e270e-cdcb-4c3c-8e77-97e17b72a62f",
          "has_attachment":false,
          "intervention_id":"5c372253d02c8e6eabc2a365",
          "in_reply_to_author_id":"5c371f92d02c8e6eabc2a23b",
          "in_reply_to_id":"5c3f6ba5d02c8e3a585fb99c",
          "language":"en",
          "source_id":"5c371f92d02c8e6eabc2a237",
          "status":"assigned",
          "thread_id":"5c3721a4d02c8e7351922551",
          "thread_title":"Thread Title",
          "created_from":"synchronizer",
          "private":true,
          "context_data":{},
          "authenticate": {
            "token":"ba371cc8c511e0d9114b8e17777ebc4ad5e23811144341f0ca4726d67ff4b19e",
            "status":"authenticated"
          }
        }
      }
    }
  ]
}
```

## Example: Apple Business Chat (Apple Pay)

Nothing specifically unique as this is an Apple Business Chat specific structured message type.

<img class="img-fluid" width="350" src="../../../img/structured-messages-apple-auth-apple-biz.png">
