# Introduction to Structured Messages

Structured Messages provide a richer messaging experience for customers. You can learn more using the [Structured Messages](../../../interactions/structured-messages).

## Prerequisites

* In order to enable structured messages support, '[view.messaging](../options/#viewmessaging)' option must be implemented.
* An author used for structured message creation must be marked as a "puppet" identity and configured as a controlled identity in Engage Digital. More details can be found in the [Users](../objects/#users) object.

## Actions

The following actions are supported by messages and private messages containing structured message:

1. For messages: create / publish / unpublish / delete.
2. For private messages: create / delete.

## Types of Structured Messages

| Type | Description |
|-|-|
| select | This structured message provides a list from which the customer can pick some elements. |
| rich_link | This structured message allows you to decorate a link with an image, title and subtitle. |
| template | This structured message provides some buttons associated with a message. |
| carousel | This structured message allows you to group multiple templates in the same message. |

## Creating a Select Structured Message

### Primary Parameters

| API Property | Type | Description |
|-|-|-|
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "select". |
| **`structured_content.items`** | Array | An array of items representing the options presented to the customer. A maximum of 13 items is supported. |
| **Item Settings** | | |
| **`items.title`** | String | The title of the item. Limited to 80 characters. |
| **`items.payload`** | String | **Optional**. Payload that will be returned with the structured response. Limited to 1000 characters. |

### Example Payload Format

```json
{
  "body": "What do you wish?",
  "structured_content": {
    "type": "select",
    "items": [
      {
        "title": "Option 1",
        "payload": "first_option"
      },
      {
        "title": "Option 2",
        "payload": "second_option"
      },
      {
        "title": "Option 3"
      }
    ]
  }
}
```

## Creating a Rich Link Structured Message

### Primary Parameters

| API Property | Type | Description |
|-|-|-|
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "rich_link". |
| **`structured_content.url`** | String | URL of the rich link. Limited to 2048 characters. Should only have http and https schemes. |
| **`structured_content.url_text`** | String | **Optional**. Text that will be displayed instead of the hostname of the URL. Automatically gets populated as the hostname of the url if blank. Limited to 80 characters. |
| **`structured_content.title`** | String | Title of the rich link. Limited to 350 characters. |
| **`structured_content.subtitle`** | String | **Optional**. Subtitle of the rich link. Limited to 1000 characters. |
| **`structured_content.attachment_url`** | String | **Optional**. Existing attachment id used to decorate the rich link with an image. Should be public. Should be jpg, jpeg or png. Should be less than 5MB. [Upload attachments](../../../basics/uploads) for your own custom images. |

### Example Payload Format

```json
{
  "structured_content": {
    "type": "rich_link",
    "attachment_url": "https://www.ringcentral.com/images/logo.jpg",
    "title": "Ringcentral, Inc.",
    "subtitle": "Cloud Business Communications",
    "url": "github://github.com/ringcentral",
    "url_text": "Github"
  }
}
```

## Create a Template Structured Message

### Primary Parameters

| API Property | Type | Description |
|-|-|-|
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "template". |
| **`structured_content.title`** | String | Title of the template. Limited to 350 characters. |
| **`structured_content.subtitle`** | String | **Optional**. Subtitle of the template. Limited to 1000 characters. |
| **`structured_content.url`** | String | **Optional**. Allows to redirect to this URL when clicking on the image or message of the template. Limited to 2048 characters. Should only have http and https schemes. |
| **`structured_content.url_text`** | String | **Optional**. Text that will be displayed instead of the hostname of the URL. Automatically gets populated as the hostname of the url if blank. Limited to 80 characters. |
| **`structured_content.attachment_url`** | String | **Optional**. Link to attachment id used to decorate the template with an image. Should be public. Should be jpg, jpeg or png. Should be less than 5MB. [Upload attachments](../../../basics/uploads) for your own custom images. |
| **Item Settings** | | |
| **`structured_content.items.type`** | String | The type of the button. Can be **url** or **reply**. |
| **`structured_content.items.title`** | String | The title of the item. Limited to 80 characters. |
| **`structured_content.items.payload`** | String | **Optional** when the type is "**reply**". **Ignored** when the type is "**url**". Payload that will be returned with the structured response. Limited to 1000 characters. |
| **`structured_content.items.url`** | String | **Required** when the type is "**url**". **Ignored** when the type is "**reply**". URL to which the user will be redirected to when clicking on the item. Limited to 2048 characters. Should only have http and https schemes. |

### Example Payload Format

```json
{
  "structured_content": {
    "type": "template",
    "attachment_url": "https://www.ringcentral.com/images/logo.jpg",
    "title": "Ringcentral, Inc.",
    "subtitle": "Cloud Business Communications.",
    "url": "github://github.com/ringcentral",
    "url_text": "Github",
    "items": [
      { "title": "Go to website", "type": "url", "url": "github://github.com/ringcentral" },
      { "title": "Ok", "type": "reply" },
      { "title": "Give me more", "type": "reply", "payload": "more" }
    ]
  }
}
```

## Create a Carousel Structured Message

### Primary Parameters

| API Property | Type | Description |
|-|-|-|
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "carousel". |
| **`structured_content.items`** | Array | List of items representing the templates presented to the customer. A maximum of 10 items is supported. |
| **Item Settings** | | |
| **`structured_content.items.title`** | String | The title of the template. Limited to 350 characters. If the device does not support structured messages, this the first item's field in the list, will be sent as the message.|
| **`structured_content.items.subtitle`** | String | **Optional**. The subtitle of the template. Limited to 1000 characters. |
| **`structured_content.items.url`** | String | **Optional**. Allows to redirect to this url when clicking on the image or message of the template. Limited to 2048 characters. Should only have http and https schemes. |
| **`structured_content.items.url_text`** | String | **Optional**. Text that will be displayed instead of the hostname of the url. Automatically gets populated as the hostname of the url if blank. This field is only displayed if the url field is present. Limited to 80 characters. |
| **`structured_content.items.attachment_url`** | String | **Optional**. Existing attachment id used to decorate the template with an image. Should be public. Should be jpg, jpeg or png. Should be less than 5MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only be jpg, jpeg or png. Maximum size of 5MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items.items`** | Array | List of items representing the buttons presented to the customer. A maximum of 4 items is supported. |
| **Items within Item Settings** | | |
| **`structured_content.items.items.type`** | String | The type of the button. Can be **url** or **reply**. |
| **`structured_content.items.items.title`** | String | The title of the item. Limited to 80 characters. |
| **`structured_content.items.items.payload`** | String | **Optional** when the type is "**reply**". **Ignored** when the type is "**url**". Payload that will be returned with the structured response. Limited to 1000 characters. |
| **`structured_content.items.items.url`** | String | **Required** when the type is "**url**". **Ignored** when the type is "**reply**". URL to which the user will be redirected to when clicking on the item. Limited to 2048 characters. Should only have http and https schemes. |

### Example Payload Format

```json
{
  "structured_content": {
    "type":"carousel",
    "items": [
      {
        "attachment_url": "https://www.ringcentral.com/images/logo.jpg",
        "title":"Ringcentral, Inc.",
        "subtitle":"Cloud Business Communications.",
        "items": [
          {
            "title":"Go to website",
            "type":"url",
            "url": "github://github.com/ringcentral",
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
        "attachment_url": "https://www.ringcentral.com/images/logo.jpg",
        "title":"Ringcentral, Inc.",
        "url": "github://github.com/ringcentral",
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