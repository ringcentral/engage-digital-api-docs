# About Carousel Structured Messages

This structured message allows to group multiple templates in the same message. See [Channel capabilities](../#channel-capabilities) to know on which channel you can use this structured message.

## Request Example
```bash
curl -X POST "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents"
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
| **`source_id`** | String | **Optional**. ID of the source. Most interactions are in reply to a message being sent to the agent. In these cases, the source ID is not required. |
| **`in_reply_to_id`** | String | ID of the message being replied to. |
| **`structured_content`** | Object | Payload of the structured message. |
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "carousel". |
| **`structured_content.items`** | Array | List of items representing the templates presented to the customer. A maximum of 10 items is supported. |
| **Item Settings** | | |
| **`structured_content.items.title`** | String | The title of the template. Limited to 350 characters. If the device does not support structured messages, this the first item's field in the list, will be sent as the message.|
| **`structured_content.items.subtitle`** | String | **Optional**. The subtitle of the template. Limited to 1000 characters. |
| **`structured_content.items.url`** | String | **Optional**. Allows to redirect to this url when clicking on the image or message of the template. Limited to 2048 characters. Should only have http and https schemes. |
| **`structured_content.items.url_fallback`** | String | **Optional**. Fallback in case the url is invalid. Limited to 2048 characters. Only http and https schemes. |
| **`structured_content.items.url_text`** | String | **Optional**. Text that will be displayed instead of the hostname of the url. Automatically gets populated as the hostname of the url if blank. This field is only displayed if the url field is present. Limited to 80 characters. |
| **`structured_content.items.attachment_id`** | String | **Every template must either have an attachment or none at all**. Existing attachment id used to decorate the template with an image. Should be public. Should be jpg, jpeg or png. Should be less than 5MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only be jpg, jpeg or png. Maximum size of 5MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items.items`** | Array | List of items representing the buttons presented to the customer. A maximum of 4 items is supported. |
| **Items within Item Settings** | | |
| **`structured_content.items.items.type`** | String | The type of the button. Can be **url** or **reply**. |
| **`structured_content.items.items.title`** | String | The title of the item. Limited to 80 characters. |
| **`structured_content.items.items.payload`** | String | **Optional** when the type is "reply". **Ignored** when the type is "url". Payload that will be returned with the structured response. Limited to 1000 characters. |
| **`structured_content.items.items.url`** | String | **Required** when the type is "url". **Ignored** when the type is "reply". Url to which the user will be redirected to when clicking on the item. Limited to 2048 characters. Should only have http and https schemes. |
| **`structured_content.items.items.url_fallback`** | String | **Optional**. Fallback in case the url is invalid. Limited to 2048 characters. Only http and https schemes. |

## Example: Facebook Messenger

<img class="img-fluid" width="592" src="../../../img/structured-messages-carousel-fb.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Specificity |
|-|-|
| **`structured_content.items.attachment_id`** | Supports bmp, gif, jpg, jpeg, png formats. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items.url_text`** | Ignored property. |
| **`structured_content.items.items`** | A maximum of 3 items is supported. |
| **`structured_content.items.title`** | Truncated to 80 characters. |
| **`structured_content.items.subtitle`** | Truncated to 80 characters. |
| **`structured_content.items.items.title`** | Truncated to 20 characters. |
| **`structured_content.items.items.payload`** | Automatically gets populated as a random hex if blank. |

## Example: Engage Messaging

<img class="img-fluid" width="398" src="../../../img/structured-messages-carousel-engage.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Specificity |
|-|-|
| **`structured_content.items.attachment_id`** | Supports gif, jpg, jpeg, png formats. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images.<br>On Engage Messaging Web, if the width of the **first** image is bigger than the height, every image will be displayed with a 5:3 ratio. Otherwise, a 1:1 ratio will be used.<br>Minimal recommended size with a 1:1 ratio: **258x258**<br>Minimal recommended size with a 5:3 ratio: **258x155** |
| **`structured_content.items.url`** | Deep links are supported. |
| **`structured_content.items.target`** | **Optional**. **Ignored** when the `url` is empty. Behavior applied when clicking on the url. Can be `webview` to open url on a [webview](../webview) above the chat, `open` to open url in new tab, or `current` to open url in current tab. Defaults to `open` when not specified. On iOS and Android, `current` value opens a full sized webview.|
| **`structured_content.items.webview_height`** | **Optional**. **Ignored** when `target` is other than `webview`. Size of the webview used to open the link. Can be `full`, `tall` or `compact`. Defaults to `full` if unset.|
| **`structured_content.items.items.target`** | **Optional** when the type is `url`. **Ignored** when the type is `reply`. Behavior applied when clicking on the item link. Can be `webview` to open on a [webview](../webview) above the chat, `open` to open in new tab, or `current` to open in current tab. Defaults to `open` when not specified. On iOS and Android, `current` value opens a full sized webview.|
| **`structured_content.items.items.webview_height`** | **Optional**. **Ignored** when `target` is other than `webview`. Size of the webview used to open the link. Can be `full`, `tall` or `compact`. Defaults to `full` if unset.|

## Example: Google Business Messages (Rich Card Carousel)

<img class="img-fluid" width="419" src="../../../img/structured-messages-carousel-google-biz.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Specificity |
|-|-|
| **`structured_content.items.title`** | Truncated to 200 characters. |
| **`structured_content.items.url`** | Ignored property. |
| **`structured_content.items.url_text`** | Ignored property. |
| **`structured_content.items.attachment_id`** | Supports private attachments. |
| **`structured_content.items.items.title`** | Truncated to 25 characters. |
| **`structured_content.items.items.payload`** | Automatically gets populated as a random hex if blank. |

## Example: Instagram Messaging

<img class="img-fluid" width="245" src="../../../img/structured-messages-carousel-ig-dm.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Specificity |
|-|-|
| **`structured_content.items.attachment_id`** | Supports bmp, gif, jpg, png formats. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items.url_text`** | Ignored property. |
| **`structured_content.items.items`** | A maximum of 3 items is supported. |
| **`structured_content.items.title`** | Truncated to 80 characters. |
| **`structured_content.items.subtitle`** | Truncated to 80 characters. |
| **`structured_content.items.items.title`** | Truncated to 20 characters. |
| **`structured_content.items.items.payload`** | Limited to 1000 characters. Automatically gets populated as a random hex if blank. |
