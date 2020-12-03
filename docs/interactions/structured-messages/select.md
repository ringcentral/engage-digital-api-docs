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
| **`structured_content.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the structured message with an image. [Upload attachments](../../../basics/uploads) for you own custom images. Supports private attachments.  |
| **`structured_content.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only jpg, jpeg, png formats. Maximum size of 5 MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.sections`** | Array | **Optional**. An array of sections in which the items will be organized. If blank, every item will be part of the same section. |
| **Section Settings** | | |
| **`structured_content.sections.title`** | String | The title of the section.  |
| **`structured_content.sections.multiple_section`** | Boolean | **Optional**. Allows the section to be multi selectable. False by default. |
| **`structured_content.sections.identifier`** | String | Identifier of the section that will be used to organize items in the section. |
| **Item Settings** | | |
| **`structured_content.items.section_identifier`** | String | **Optional**. The identifier of the section where the item is. If the sections field is blank, the section_identifier field should be removed.  |
| **`structured_content.items.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the item with an image. [Upload attachments](../../../basics/uploads) for you own custom images. Supports private attachments. |
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
| **`items.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the item in the list. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`items.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only jpg, jpeg, png formats. Maximum size of 5 MB. [Upload attachments](../../../basics/uploads) for you own custom images. |

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
