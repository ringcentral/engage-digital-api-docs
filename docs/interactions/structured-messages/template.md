# About Template Structured Message

This structured content provides some buttons associated with a message. See [Channel capabilities](../structured-messages/#channel-capabilities) to know on which channel you can use this structured message.

## Request Example

```bash
curl -X POST "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents"
```

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "structured_content": {
    "type": "template",
    "attachment_id": "<attachment_id>",
    "attachment_fallback_id": "<attachment_id>",
    "title": "Ringcentral, Inc.",
    "subtitle": "Cloud Business Communications.",
    "url": "github://github.com/ringcentral",
    "target": "open",
    "url_fallback": "https://github.com/ringcentral",
    "url_text": "Github",
    "items": [
      { "title": "Go to website", "type": "url", "url": "github://github.com/ringcentral", "url_fallback": "https://github.com/ringcentral", "target": "webview" },
      { "title": "Ok", "type": "reply" },
      { "title": "Give me more", "type": "reply", "payload": "more" }
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
| **`structured_content.type`** | String | Type of the structured message. Must be "template". |
| **`structured_content.title`** | String | Title of the template. Limited to 350 characters. If the device does not support structured messages, this field will be sent as the message. |
| **`structured_content.subtitle`** | String | **Optional**. Subtitle of the template. Limited to 1000 characters. |
| **`structured_content.url`** | String | **Optional**. Allows to redirect to this url when clicking on the image or message of the template. Limited to 2048 characters. Should only have http and https schemes. |
| **`structured_content.url_fallback`** | String | **Optional**. Fallback in case the url is invalid. Limited to 2048 characters. Only http and https schemes. |
| **`structured_content.url_text`** | String | **Optional**. Text that will be displayed instead of the hostname of the url. Automatically gets populated as the hostname of the url if blank. This field is only displayed if the url field is present. Limited to 80 characters. |
| **`structured_content.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the template with an image. Should be public. Should be jpg, jpeg or png. Should be less than 5MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only jpg, jpeg, png formats. Maximum size of 5 MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items`** | Array | List of items representing the buttons presented to the customer. A maximum of 4 items is supported. |
| **Item Settings** | | |
| **`structured_content.items.type`** | String | The type of the button. Can be **url** or **reply**. |
| **`structured_content.items.title`** | String | The title of the item. Limited to 80 characters. |
| **`structured_content.items.payload`** | String | **Optional** when the type is "reply". **Ignored** when the type is "url". Payload that will be returned with the structured response. Limited to 1000 characters. |
| **`structured_content.items.url`** | String | **Required** when the type is "url". **Ignored** when the type is "reply". Url to which the user will be redirected to when clicking on the item. Limited to 2048 characters. Should only have http and https schemes. |
| **`structured_content.items.url_fallback`** | String | **Optional** Fallback in case the url is invalid. Limited to 2048 characters. Only http and https schemes. |

## Example: Facebook Messenger

<img class="img-fluid" width="466" src="../../../img/structured-messages-template-fb.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example. If the template only has a title and items, it will be converted to a Button Template in Facebook Messenger.

**Button Template**

| API Property | Specificity |
|-|-|
| **`structured_content.title`** | Limited to 640 characters. |
| **`structured_content.items`** | A maximum of 3 items is supported. |
| **`structured_content.items.title`** | Truncated to 20 characters. |
| **`structured_content.items.payload`** | Limited to 1000 characters. |

Otherwise, it will be converted to a Generic Template in Facebook Messenger.

**Generic Template**

| API Property | Specificity |
|-|-|
| **`structured_content.attachment_id`** | Supports bmp, gif, jpg, jpeg, png formats. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.url_text`** | **Ignored** property. |
| **`structured_content.items`** | A maximum of 3 items is supported. |
| **`structured_content.title`** | Truncated to 80 characters. |
| **`structured_content.subtitle`** | Truncated to 80 characters. |
| **`structured_content.items.title`** | Truncated to 20 characters. |

## Example: Engage Messaging

<img class="img-fluid" width="398" src="../../../img/structured-messages-template-engage.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Specificity |
|-|-|
| **`structured_content.attachment_id`** | Supports gif, jpg, jpeg, png formats. Supports private attachments. [Upload attachments](../../../basics/uploads) for your own custom images.<br>On Engage Messaging Web, if the width of the image is bigger than the height, it will be displayed with a 5:3 ratio. Otherwise, a 1:1 ratio will be used.<br>Minimal recommended size with a 1:1 ratio: **258x258**<br>Minimal recommended size with a 5:3 ratio: **258x155** |
| **`structured_content.url`** | Deep links are supported. |
| **`structured_content.target`** | **Optional**. **Ignored** when the `url` is empty. Behavior applied when clicking on the url. Can be `webview` to open url on a [webview](../web-messaging/webview) above the chat, `open` to open url in new tab, or  `current` to open url in current tab. Defaults to `open` when not specified. Can also be set on `item` with type `url`.|
| **`structured_content.webview_height`** | **Optional**. **Ignored** when `target` is other than `webview`. Size of the webview used to open the link. Can be `full`, `tall` or `compact`. Defaults to `full` is unset. Can also be set on `item` with type `url`.|

## Example: Google Business Messages (Rich Card)

<img class="img-fluid" width="419" src="../../../img/structured-messages-template-google-biz.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Specificity |
|-|-|
| **`structured_content.title`** | Truncated to 200 characters. |
| **`structured_content.url`** | Ignored property. |
| **`structured_content.url_text`** | Ignored property. |
| **`structured_content.items.title`** | Truncated to 25 characters. |
| **`structured_content.items.payload`** | Automatically gets populated as a random hex if blank. |
| **`structured_content.attachment_id`** | Supports private attachments. |

## Example: WhatsApp (Reply Buttons)

<img class="img-fluid" width="398" src="../../../img/structured-messages-template-whatsapp.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Type | Specificity |
|-|-|-|
| **If subtitle is present** | | |
| **`structured_content.title`** | String | Truncated to 60 characters. |
| **`structured_content.subtitle`** | String | Truncated to 963 characters if an attachment is present. |
| **If subtitle is empty** | | |
| **`structured_content.title`** | String | Limited to 1000 characters. Used as the message body. |
| **Structured Content Settings** | | |
| **`structured_content.attachment_id`** | String | Supports jpg, jpeg, png, mp4 formats. Supports private attachments. [Upload attachments](../../../basics/uploads) for your own custom images or videos. Should be less than 64 MB. WhatsApp supports videos with only H.264 and AAC codecs and a single audio stream. |
| **`structured_content.footer`** | String | Limited to 60 characters. |
| **`structured_content.url`** | String | **Ignored** property. |
| **`structured_content.url_fallback`** | String | **Ignored** property. |
| **`structured_content.url_text`** | String | **Ignored** property. |
| **Item Settings** | | |
| **`structured_content.items`** | Array | Truncated to 3 elements. |
| **`structured_content.items.type`** | String | Only `reply` is supported. |
| **`structured_content.items.title`** | String | Truncated to 20 charactes. |
| **`structured_content.items.payload`** | String | Limited to 256 characters.<br>Automatically gets populated as a random hex if blank. |
| **`structured_content.items.url`** | String | **Ignored** property. |
| **`structured_content.items.url_fallback`** | String | **Ignored** property. |
