# About Rich Link Structured Message

This structured message allows to decorate a link with an image, title and subtitle. See [Channel capabilities](../#channel-capabilities) to know on which channel you can use this structured message.

## Request Example

```bash
curl -X POST "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents"
```

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "structured_content": {
    "type": "rich_link",
    "attachment_id": "<attachment_id>",
    "attachment_fallback_id": "<attachment_id>",
    "title": "Ringcentral, Inc.",
    "subtitle": "Cloud Business Communications",
    "url": "github://github.com/ringcentral",
    "url_fallback": "https://github.com/ringcentral",
    "url_text": "Github"
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
| **`structured_content.type`** | String | Type of the structured message. Must be "rich_link". |
| **`structured_content.url`** | String |Url of the rich link. Limited to 2048 characters. Should only have http and https schemes. If the device does not support structured messages, this field will be sent as the message. |
| **`structured_content.url_fallback`** | String | **Optional**. Fallback in case the url is invalid. Limited to 2048 characters. Only http and https schemes. |
| **`structured_content.url_text`** | String | **Optional**. Text that will be displayed instead of the hostname of the url. Automatically gets populated as the hostname of the url if blank. Limited to 80 characters. |
| **`structured_content.title`** | String | Title of the rich link. Limited to 350 characters. |
| **`structured_content.subtitle`** | String | **Optional**. Subtitle of the rich link. Limited to 1000 characters. |
| **`structured_content.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the rich link with an image. Should be public. Should be jpg, jpeg or png. Should be less than 5MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only jpg, jpeg, png formats. Maximum size of 5 MB. [Upload attachments](../../../basics/uploads) for you own custom images. |

## Example: Apple Messages for Business - Rich Link

<img class="img-fluid" width="350" src="../../../img/structured-messages-rich-link-apple-biz.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Type | Description |
|-|-|-|
| **`structured_content.title`** | String | Truncated to 200 characters. |
| **`structured_content.attachment_id`** | String | Supports jpg, jpeg, png, and mp4 formats. Supports private attachments. Supports up to 300MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.url_text`** | String | Ignored property. |
| **`structured_content.subtitle`** | String | Ignored property. |


## Example: Apple Messages for Business - iMessage App
Rich Link structured messages can be used to create iMessage apps.

<img class="img-fluid" width="350" src="../../../img/structured-messages-rich-link-apple-abm-imessage-app.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Type | Description |
|-|-|-|
| **`structured_content.url`** | String | A URL string containing data that the Messages app sends to the iMessage app. It must start with '?' |
| **`structured_content.app_id`** | String | The App Store identifier of the iMessage app. |
| **`structured_content.app_name`** | String | The name of the iMessage app. |
| **`structured_content.app_icon`** | String | A Base64-encoded string representing the app icon of the iMessage app. |
| **`structured_content.bid`** | String | The bundle identifier of the iMessage app. Must start with 'com.apple.messages.MSMessageExtensionBalloonPlugin:' |
| **`structured_content.subtitle`** | String | **Optional**. Subtitle of the rich link.  |
| **`structured_content.use_live_layout`** | String | **Optional**. A Boolean that determines whether the Messages app should use Live Layout. The default is true. |
| **`structured_content.url_text`** | String | Ignored property. |
| **`structured_content.attachment_id`** | Ignored property. |

## Example: Facebook Messenger

Facebook Messenger does not *natively* support rich links. So a rich link can either be converted to a Generic Template or a Button Template:

* If the subtitle and the attachment_id are blank, the rich link will become a Button Template in Facebook Messenger with a single button being a "url button" that will redirect to the rich link url. The url_text will be the title of the button.
* Otherwise, the rich link will become a Generic Template in Facebook Messenger with no buttons.


<img class="img-fluid" width="503" src="../../../img/structured-messages-rich-link-fb.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Type | Description |
|-|-|-|
| **`structured_content.attachment_id`** | String | Supports bmp, gif, jpg, jpeg, png formats. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.url_text`** | String | Ignored property when converted to a Generic Template. Truncated to 20 characters when converted to a Button Template. |
| **`structured_content.title`** | String | Truncated to 80 characters. |
| **`structured_content.subtitle`** | String | Truncated to 80 characters. |

## Example: RingCX Messaging

<img class="img-fluid" width="472" src="../../../img/structured-messages-rich-link-engage.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Specificity |
|-|-|
| **`structured_content.attachment_id`** | Supports bmp, gif, jpg, jpeg, png formats. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images.<br>On RingCX Messaging Web, if the width of the image is bigger than the height, it will be displayed with a 5:3 ratio. Otherwise, a 1:1 ratio will be used.<br>Minimal recommended size with a 1:1 ratio: **258x258**<br>Minimal recommended size with a 5:3 ratio: **258x155** |
| **`structured_content.url`** | Deep links are supported. |
| **`structured_content.target`** | **Optional**. **Ignored** when the `url` is empty. Behavior applied when clicking on the url. Can be **webview** to open url on a [webview](../webview) above the chat, `open` to open url in new tab, or `current` to open url in current tab. Defaults to **open** when not specified. On iOS and Android, `current` value opens a full sized webview.|
| **`structured_content.webview_height`** | **Optional**. **Ignored** when `target` is other than `webview`. Size of the webview used to open the link. Can be `full`, `tall` or `compact`. Defaults to `full` if unset.|

## Example: Google Business Messages

Google Business Messages does not *natively* support rich links. That’s why the rich link is converted to a Rich Card and the link is converted to a single button that will redirect to the url when clicked. The url_text property is used to display the title of the button.

<img class="img-fluid" width="376" src="../../../img/structured-messages-rich-link-google-biz.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Specificity |
|-|-|
| **`structured_content.title`** | Truncated to 200 characters. |
| **`structured_content.url_text`** | Truncated to 25 characters. |
| **`structured_content.attachment_id`** | Supports private attachments. |

## Example: Instagram Messaging

Instagram Messaging does not *natively* support rich links. So a rich link can  be converted to a Generic Template

<img class="img-fluid" width="245" src="../../../img/structured-messages-rich-link-ig-dm.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Specificity |
|-|-|
| **`structured_content.attachment_id`** | Supports bmp, gif, jpg, png formats. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.url_text`** | Truncated to 20 characters. |
| **`structured_content.title`** | Truncated to 80 characters. |
| **`structured_content.subtitle`** | Truncated to 80 characters. |
