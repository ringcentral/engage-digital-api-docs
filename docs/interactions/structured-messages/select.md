# About Select Structured Messages

This structured message provides a list from which the customer can pick some elements. See [Channel capabilities](../structured-messages/#channel-capabilities) to know on which channel you can use this structured message.

## Request Example

```bash
curl -X POST "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents"
```

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
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

### Primary Parameters

| API Property | Type | Description |
|-|-|-|
| **`source_id`** | String | **Optional**. ID of the source. Most interactions are in reply to a message being sent to the agent. In these cases, the source ID is not required. |
| **`in_reply_to_id`** | String | ID of the message being replied to. |
| **`body`** | String | The select structured message body. If the device does not support structured messages, this field will be sent as the message. |
| **`structured_content`** | Object | Payload of the structured message. |
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "select". |
| **`structured_content.items`** | Array | An array of items representing the options presented to the customer. A maximum of 13 items is supported. |
| **Item Settings** | | |
| **`items.title`** | String | The title of the item. Limited to 80 characters. |
| **`items.payload`** | String | **Optional**. Payload that will be returned with the structured response. Limited to 1000 characters. |

## Example: Apple Business Chat (List Picker)

This list picker structured content has multiple specific fields. So here’s an example payload format.

<img class="img-fluid" width="300" src="../../../img/structured-messages-select-apple-biz.png">

### JSON Body

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "body": "Pick some options",
  "structured_content": {
    "type": "select",
    "subtitle": "We have great options",
    "attachment_id": "<attachment_id>",
    "attachment_fallback_id": "<attachment_id>",
    "sections": [
      {
        "title": "section 1",
        "multiple_selection": true,
        "identifier": "1"
      },
      {
        "title": "section 2",
        "multiple_selection": false,
        "identifier": "2"
      }
    ],
    "items": [
      {
        "section_identifier": "1",
        "title": "option 1-1",
        "attachment_id": "<id>",
        "attachment_fallback_id": "<attachment_id>",
      },
      {
        "section_identifier": "1",
        "title": "option 1-2",
        "subtitle": "subtitle"
      },
      { "section_identifier": "2",
        "title": "option 2-1",
        "payload": "payload"
      },
      { "section_identifier": "2",
        "title": "option 2-2"
      }
    ]
  }
}
```

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Type | Description |
|-|-|-|
| **Structured Content Settings** | | |
| **`structured_content.subtitle`** | String | **Optional**. The subtitle field.<br>Limited to 512 characters.  |
| **`structured_content.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the structured message with an image. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only jpg, jpeg, png formats. Maximum size of 5 MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.sections`** | Array | **Optional**. An array of sections in which the items will be organized.<br>Limited to 10 elements.<br>If blank, every item will be part of the same section. |
| **Section Settings** | | |
| **`structured_content.sections.title`** | String | **Optional if there's only a single section**. The title of the section.<br>Limited to 24 characters. |
| **`structured_content.sections.identifier`** | String | Identifier of the section that will be used to organize items in the section.<br>Limited to 200 characters. |
| **`structured_content.sections.multiple_section`** | Boolean | **Optional**. Allows the section to be multi selectable. False by default. |
| **Item Settings** | | |
| **`structured_content.items.section_identifier`** | String | **Optional if there's no sections**. The identifier of the section where the item is.<br>If there's no sections, the section_identifier field should be removed.<br>Each section must have at least 1 item.<br>Limited to 200 characters. |
| **`structured_content.items.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the item with an image. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only jpg, jpeg, png formats. Maximum size of 5 MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items.subtitle`** | String | **Optional**. The subtitle of the item. Limited to 512 characters. |

## Example: Facebook Messenger (Quick Replies)

The following example uses Facebook Messenger with quick replies.

<img class="img-fluid" width="628" src="../../../img/structured-messages-select-fb.png">

### JSON Body

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "body": "What do you wish?",
  "structured_content": {
    "type": "select",
    "items": [
      {
        "title": "Option 1",
        "payload": "first_option",
        "attachment_id": "<attachment_id>",
        "attachment_fallback_id": "<attachment_id>",
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

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Type | Description |
|-|-|-|
| **`items.title`** | String | The title of the item. *Truncated to 20 characters.* |
| **`items.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the item in the list. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`items.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only jpg, jpeg, png formats. Maximum size of 5 MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`items.payload`** | String | Automatically gets populated as a random hex if blank. |

## Example: Engage Messaging (Quick Replies)

The following example uses Engage Messaging with quick replies.

<img class="img-fluid" width="600" src="../../../img/structured-messages-select-engage.png">

The Engage Messaging quick replies support both the centering of items and the disabling of the text input:

### JSON Body

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "body": "What do you wish?",
  "structured_content": {
    "type": "select",
    "center_items": true,
    "disable_text_input": true,
    "items": [
      {
        "title": "Option 1",
        "payload": "first_option"
      },
      {
        "title": "Option 2",
        "payload": "second_option"
      }
    ]
  }
}
```

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Type | Description |
|-|-|-|
| **`structured_content.center_items`** | Booelan | **Optional**. Allows to center the items. False by default. Centers the items when at “true”. Aligns the items on the right when at “false”. |
| **`structured_content.disable_text_input`** | Booelan | **Optional**. Allows to disable the text input when the items are displayed. False by default. Disables the text input when at “true”. Also disables the ability to reply via the push notification on mobile. |

## Example: Google Business Messages

<img class="img-fluid" width="321" src="../../../img/structured-messages-select-google-biz.png">

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Specificity |
|-|-|
| **`structured_content.items.title`** | Truncated to 25 characters. |
| **`structured_content.items.payload`** | Automatically gets populated as a random hex if blank. |

## Example: WhatsApp (List Messages)

The following example uses WhatsApp with list messages.

<img class="img-fluid" width="398" src="../../../img/structured-messages-select-whatsapp-1.png">
<img class="img-fluid" width="398" src="../../../img/structured-messages-select-whatsapp-2.png">

### JSON Body

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "body": "Hello! What do you wish?",
  "structured_content": {
    "type": "select",
    "title": "Welcome to the store!",
    "footer": "We're always happy to offer you the best options!",
    "button": "See options",
    "sections": [
      {
        "title": "section 1",
        "identifier": "1"
      },
      {
        "title": "section 2",
        "identifier": "2"
      }
    ],
    "items": [
      {
        "section_identifier": "1",
        "title": "Option 1",
        "payload": "first_option",
        "description": "The first option"
      },
      {
        "section_identifier": "2",
        "title": "Option 2",
        "payload": "second_option",
        "description": "The second option"
      },
      {
        "section_identifier": "2",
        "title": "Option 3",
        "payload": "third_option",
        "description": "The third option"
      }
    ]
  }
}
```

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Type | Description |
|-|-|-|
| **Structured Content Settings** | | |
| **`structured_content.items`** | Array | Truncated to 10 elements. |
| **`structured_content.button`** | String | **Optional**. The button text field.<br>Limited to 20 characters.<br>*Truncated to 20 UTF-16 code units.*<br>"See options" by default. |
| **`structured_content.title`** | String | **Optional**. The title text field.<br>Limited to 60 characters.<br>*Truncated to 60 UTF-16 code units.* |
| **`structured_content.footer`** | String | **Optional**. The footer text field.<br>Limited to 60 characters.<br>*Truncated to 60 UTF-16 code units.* |
| **Section Settings** | | |
| **`structured_content.sections`** | Array | **Optional**. An array of sections in which the items will be organized.<br>Limited to 10 elements.<br>If blank, every item will be part of the same section. |
| **`structured_content.sections.title`** | String | **Optional if there's only a single section**. The title of the section.<br>Limited to 24 characters.<br>*Truncated to 24 UTF-16 code units.* |
| **`structured_content.sections.identifier`** | String | Identifier of the section that will be used to organize items in the section.<br>Limited to 200 characters. |
| **Item Settings** | | |
| **`structured_content.items.title`** | String | The item title field.<br>Truncated to 24 characters.<br>*Truncated to 24 UTF-16 code units.* |
| **`structured_content.items.payload`** | String | **Optional**. The item payload field.<br>Limited to 200 characters.<br>Automatically gets populated as a random hex if blank. |
| **`structured_content.items.description`** | String | **Optional**. The item description text field.<br>Limited to 72 characters.<br>*Truncated to 72 UTF-16 code units.* |
| **`structured_content.items.section_identifier`** | String | **Optional if there's no sections**. The identifier of the section where the item is.<br>If there's no sections, the section_identifier field should be removed.<br>Each section must have at least 1 item.<br>Limited to 200 characters. |

##  Example: Instagram Messaging (Quick Replies)

The following example uses Instagram Messaging with quick replies.

<img class="img-fluid" width="600" src="../../../img/structured-messages-select-ig-dm.png">

### JSON Body

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "body": "Please make a choice",
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

### Properties Unique to this Channel

Primary parameters are used by default, however, some parameters are unique or overwritten by parameters specific to this example.

| API Property | Type | Description |
|-|-|-|
| **`items.title`** | String | The title of the item. *Truncated to 20 characters.* |
| **`items.payload`** | String | Automatically gets populated as a random hex if blank. |
