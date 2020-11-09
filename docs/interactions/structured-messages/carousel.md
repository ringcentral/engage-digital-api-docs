# About Carousel Structured Messages

This structured message allows to group multiple templates in the same message. See [Channel capabilities](../structured-messages/#channel-capabilities) to know on which channel you can use this structured message.

## Request Example
```bash
curl -X POST "https://[YOUR DOMAIN].api.engagement.dimelo.com/1.0/contents"
```

```json
{
  "source_id":"<source_id>",
  "in_reply_to_id":"<in_reply_to_id>",
  "structured_content": {
    "type":"carousel",
    "items": [
      {
        "attachment_id":"<attachment_id>",
        "attachment_fallback_id": "<attachment_id>",
        "title":"Ringcentral, Inc.",
        "subtitle":"Cloud Business Communications.",
        "items": [
          {
            "title":"Go to website",
            "type":"url",
            "url": "github://github.com/ringcentral",
            "url_fallback": "https://github.com/ringcentral",
          },
          {
            "title":"Ok",
            "type":"reply"
          },
          {
            "title":"Give me more",
            "type":"reply",
            "payload":"more"
          }
        ]
      },
      {
        "attachment_id":"<attachment_id>",
        "title":"Ringcentral, Inc.",
        "url": "github://github.com/ringcentral",
        "url_fallback": "https://github.com/ringcentral",
        "url_text":"Github",
        "items": [
          {
            "title":"Go to website",
            "type":"url",
            "url":"https://github.com/ringcentral"
          },
          {
            "title":"Ok",
            "type":"reply"
          },
          {
            "title":"Give me more",
            "type":"reply",
            "payload":"more"
          }
        ]
      }
    ]
  }
}
```

### Primary Parameters

| API Property | Type | Description |
|-|-|-|
| **`source_id`** | String | ID of the source. |
| **`in_reply_to_id`** | String | ID of the message being replied to. |
| **`structured_content`** | Object | Payload of the structured message. |
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "carousel". |
| **`structured_content.items`** | Array | List of items representing the templates presented to the customer. A maximum of 10 items is supported. |
| **Item Settings** | | |
| **`structured_content.items.title`** | String | The title of the template. Limited to 350 characters. |
| **`structured_content.items.subtitle`** | String | **Optional**. The subtitle of the template. Limited to 1000 characters. |
| **`structured_content.items.url`** | String | **Optional**. Allows to redirect to this url when clicking on the image or message of the template. Limited to 2048 characters. Should only have http and https schemes. |
| **`structured_content.items.url_fallback`** | String | **Optional**. Fallback in case the url is invalid. Limited to 2048 characters. Only http and https schemes. |
| **`structured_content.items.url_text`** | String | **Optional**. Text that will be displayed instead of the hostname of the url. Automatically gets populated as the hostname of the url if blank. This field is only displayed if the url field is present. Limited to 80 characters. |
| **`structured_content.items.attachment_id`** | String | **Every template must either have an attachment or none at all**. Existing attachment id used to decorate the template with an image. Should be public. Should be jpg, jpeg or png. Should be less than 5MB. |
| **`structured_content.items.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only be jpg, jpeg or png. Maximum size of 5MB. |
| **`structured_content.items.items`** | Array | List of items representing the buttons presented to the customer. A maximum of 4 items is supported. |
| **Items within Item Settings** | | |
| **`structured_content.items.items.type`** | String | The type of the button. Can be **url** or **reply**. |
| **`structured_content.items.items.title`** | String | The title of the item. Limited to 80 characters. |
| **`structured_content.items.items.payload`** | String | **Optional** when the type is "reply". **Ignored** when the type is "url". Payload that will be returned with the structured response. Limited to 1000 characters. |
| **`structured_content.items.items.url`** | String | **Required** when the type is "url". **Ignored** when the type is "reply". Url to which the user will be redirected to when clicking on the item. Limited to 2048 characters. Should only have http and https schemes. |
| **`structured_content.items.items.url_fallback`** | String | **Optional**. Fallback in case the url is invalid. Limited to 2048 characters. Only http and https schemes. |

## Example: Facebook Messenger

<img class="img-fluid" width="592" src="../../../img/structured-messages-carousel-fb.png">

A carousel structured content is converted to a **Generic Template** with multiple elements in Facebook Messenger.

The same Template <**Generic Template**> specificities apply to every item of the carousel.

## Example: Engage Messaging

<img class="img-fluid" width="468" src="../../../img/structured-messages-carousel-engage.png">

### Properties Unique to this Channel

| API Property | Specificity |
|-|-|
| **`structured_content.attachment_id`** | Supports gif, jpg, jpeg, png formats. Supports private attachments.<br>On Engage Messaging Web, if the width of the **first** image is bigger than the height, every image will be displayed with a 5:3 ratio. Otherwise, a 1:1 ratio will be used.<br>Minimal recommended size with a 1:1 ratio: **258x258**<br>Minimal recommended size with a 5:3 ratio: **258x155** |
| **`structured_content.url`** | Deep links are supported. |

## Example: Google Business Messages (Rich Card Carousel)

<img class="img-fluid" width="419" src="../../../img/structured-messages-carousel-google-biz.png">

### Properties Unique to this Channel

| API Property | Specificity |
|-|-|
| **`structured_content.title`** | Truncated to 200 characters. |
| **`structured_content.url`** | Ignored property. |
| **`structured_content.url_text`** | Ignored property. |
| **`structured_content.items.title`** | Truncated to 25 characters. |
| **`structured_content.items.payload`** | Automatically gets populated as a random hex if blank. |
