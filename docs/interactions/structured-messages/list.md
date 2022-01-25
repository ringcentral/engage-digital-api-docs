# About List Structured Messages

This structured message provides a list from which the customer can pick some elements. See [Channel capabilities](../structured-messages/#channel-capabilities) to know on which channel you can use this structured message.

## Request Example

```bash
curl -X POST "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents"
```

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "structured_content": {
    "type": "list",
    "title": "List title",
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
| **`structured_content`** | Object | Payload of the structured message. |
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "select". |
| **`structured_content.title`** | String | Title of the list structured message. Limited to 1024 characters. |
| **`structured_content.items`** | Array | An array of items representing the options presented to the customer. A maximum of 13 items is supported. |
| **Item Settings** | | |
| **`items.title`** | String | The title of the item. Limited to 80 characters. |
| **`items.payload`** | String | **Optional**. Payload that will be returned with the structured response. Limited to 1000 characters. |

## Example: Apple Business Chat (List Picker)

This list picker structured content has multiple specific fields. So here’s an example payload format.

<img class="img-fluid" width="300" src="../../../img/structured-messages-list-apple-biz.png">

### JSON Body

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "structured_content": {
    "type": "select",
    "title": "Pick some options",
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
| **`structured_content.title`** | String | Title of the list structured message. Limited to 1024 characters. |
| **`structured_content.subtitle`** | String | **Optional**. The subtitle field.<br>Limited to 512 characters. |
| **`structured_content.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the structured message with an image. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only jpg, jpeg, png formats. Maximum size of 5 MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.sections`** | Array | **Optional**. An array of sections in which the items will be organized.<br>Limited to 10 elements.<br>If blank, every item will be part of the same section. |
| **Section Settings** | | |
| **`structured_content.sections.title`** | String | **Optional if there's only a single section**. The title of the section.<br>Limited to 24 characters. |
| **`structured_content.sections.identifier`** | String | Identifier of the section that will be used to organize items in the section.<br>Limited to 200 characters. |
| **`structured_content.sections.multiple_section`** | Boolean | **Optional**. Allows the section to be multi selectable. False by default. |
| **Item Settings** | | |
| **`structured_content.items.section_identifier`** | String | **Optional if there's no section**. The identifier of the section where the item is.<br>If there's no section, the section_identifier field should be removed.<br>Each section must have at least 1 item.<br>Limited to 200 characters. |
| **`structured_content.items.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the item with an image. Supports private attachments. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only jpg, jpeg, png formats. Maximum size of 5 MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.items.subtitle`** | String | **Optional**. The subtitle of the item. Limited to 512 characters. |

## Example: WhatsApp (List Messages)

The following example uses WhatsApp with list messages.

<img class="img-fluid" width="398" src="../../../img/structured-messages-select-whatsapp-1.png">
<img class="img-fluid" width="398" src="../../../img/structured-messages-select-whatsapp-2.png">

### JSON Body

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "structured_content": {
    "type": "select",
    "header": "Welcome to the store!",
    "title": "Hello! What do you wish?",
    "subtitle": "We're always happy to offer you the best options!",
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
| **`structured_content.subtitle`** | String | **Optional**. The subtitle field.<br>*Truncated to 60 UTF-16 code units.* |
| **`structured_content.header`** | String | **Optional**. The header field.<br>Limited to 60 characters.<br>*Truncated to 60 UTF-16 code units.* |
| **Section Settings** | | |
| **`structured_content.sections`** | Array | **Optional**. An array of sections in which the items will be organized.<br>Limited to 10 elements.<br>If blank, every item will be part of the same section. |
| **`structured_content.sections.title`** | String | **Optional if there's only a single section**. The title of the section.<br>Limited to 24 characters.<br>*Truncated to 24 UTF-16 code units.* |
| **`structured_content.sections.identifier`** | String | Identifier of the section that will be used to organize items in the section.<br>Limited to 200 characters. |
| **Item Settings** | | |
| **`structured_content.items.title`** | String | The item title field.<br>Truncated to 24 characters.<br>*Truncated to 24 UTF-16 code units.* |
| **`structured_content.items.payload`** | String | **Optional**. The item payload field.<br>Limited to 200 characters.<br>Automatically gets populated as a random hex if blank. |
| **`structured_content.items.description`** | String | **Optional**. The item description text field.<br>Limited to 72 characters.<br>*Truncated to 72 UTF-16 code units.* |
| **`structured_content.items.section_identifier`** | String | **Optional if there's no section**. The identifier of the section where the item is.<br>If there's no section, the section_identifier field should be removed.<br>Each section must have at least 1 item.<br>Limited to 200 characters. |
