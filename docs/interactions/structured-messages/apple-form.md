# About Apple Form Structured Messages

This structured message provides a form made of a list of questions the customer can answer to. See [Channel capabilities](../#channel-capabilities) to know on which channel you can use this structured message.

## Request Example

```bash
curl -X POST "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents"
```

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "structured_content": {
    "type": "form",
    "title": "Report a product issue",
    "subtitle": "Please fill this form to report any issue you had with our product",
    "attachment_id": "<attachment_id>",
    "attachment_fallback_id": "<attachment_id>",
    "first_page_id": "page_1",
    "private": false,
    "show_summary": true,
    "splash": {
      "header": "Lorem ipsum",
      "text": "Lorem ipson dolor sit amet",
      "button_title": "Splash button",
      "attachment_id": "<attachment_id>",
    },
    "pages": [
      {
        "page_id": "page_1",
        "type": "datePicker",
        "header": "Retrieval Date",
        "title": "What date did you receive, or expect to receive the product?",
        "next_page_id": "page_2"
      },
      {
        "page_id": "page_2",
        "type": "select",
        "multiple_selection": true,
        "header": "Report an issue",
        "title": "What was the issue with your product?",
        "next_page_id": "page_3",
        "items": [
          {
            "id": "item_id_1",
            "value": "doesnt_work",
            "title": "The product does not work"
          },
          {
            "id": "item_id_2",
            "value": "damaged",
            "title": "The product is damaged"
          },
          {
            "id": "item_id_3",
            "value": "not_received",
            "title": "I haven't received the product"
          }
        ]
      },
      {
        "page_id": "page_3",
        "type": "picker",
        "header": "Compensation",
        "title": "What would you like to obtain as compensation?",
        "next_page_id": "page_4",
        "items": [
          {
            "id": "item_id_1",
            "value": "refund",
            "title": "A refund"
          },
          {
            "id": "item_id_2",
            "value": "new_product",
            "title": "A new product"
          },
          {
            "id": "item_id_3",
            "value": "replacement",
            "title": "A replacement"
          }
        ]
      },
      {
        "page_id": "page_4",
        "type": "input",
        "header": "More informations",
        "title": "If you have anything you'd like to tell us please do",
        "submit_form": true
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
| **`structured_content.type`** | String | Type of the structured message. Must be "form". |
| **`structured_content.title`** | String | The title of the structured message. |
| **`structured_content.subtitle`** | String | The subtitle of the structured message. |
| **`structured_content.first_page_id`** | String | `page_id` of the first page you want to display. |
| **`structured_content.private`** | Boolean | Indicates wether to mark the response as private in the `structured_reply` field of the form response. |
| **`structured_content.show_summary`** | Boolean | Displays a summary of the responses before submitting the form. Is `false` by default. |
| **`structured_content.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the form message with an image.<br>Supports private attachments.<br>Should be jpg, jpeg or png.<br>Should be less than 5MB. |
| **`structured_content.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only be jpg, jpeg or png. Maximum size of 5MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.splash`** | Object | **Optional**. Displays an introductive screen before the form questions. |
| **`structured_content.pages`** | Array | **Optional**. An array of objects representing each form question. |
| **Splash Settings** | | |
| **`structured_content.splash.header`** | String | **Optional**. Header of the splash screen. |
| **`structured_content.splash.text`** | String | **Optional**. Text in the splash screen. |
| **`structured_content.splash.button_title`** | String | **Optional**. Label of the button that allows to open the form. |
| **`structured_content.splash.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the splash with an image.<br>Supports private attachments.<br>Should be jpg, jpeg or png.<br>Should be less than 5MB. |
| **`structured_content.splash.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only be jpg, jpeg or png. Maximum size of 5MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **Page Settings** | | |
| **`structured_content.pages.page_id`** | String | Page identifier. |
| **`structured_content.pages.type`** | String | Page type, can be `select`, `picker`, `datePicker`, or `input`. |
| **`structured_content.pages.header`** | String | **Optional**. Header label for the page. |
| **`structured_content.pages.title`** | String | Title for the page. |
| **`structured_content.pages.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the page with an image.<br>Supports private attachments.<br>Should be jpg, jpeg or png.<br>Should be less than 5MB. |
| **`structured_content.pages.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only be jpg, jpeg or png. Maximum size of 5MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **`structured_content.pages.next_page_id`** | String | `page_id` of the page you want to be displayed after this one. In `select` pages without `multiple_selection` to  `true`, it should be at `item` level. |
| **`structured_content.pages.submit_form`** | Boolean | Should be `true` on the page that ends the form. |
| **`structured_content.pages.multiple_selection`** | Boolean | **Optional**. On `select` pages, should be `true` to allow selecting multiple items. |
| **`structured_content.pages.picker_title`** | String | **Optional**. On `picker` pages, adds a label on the picker element. |
| **`structured_content.pages.selected_item_index`** | Integer | **Optional**. On `picker` pages, is the default selected item. |
| **`structured_content.pages.hint_text`** | String | **Optional**. On `datePicker` and `input` pages, adds a text below the field. |
| **`structured_content.pages.items`** | Array | **Optional**. An array of objects representing each selectable items, is mandatory on `select` and `picker` pages. |
| **`structured_content.pages.options`** | Object | **Optional**. An object containing diverse options for the different page types. |
| **Item Settings** | | |
| **`structured_content.pages.items.title`** | String | Text that will actually be displayed for the item. |
| **`structured_content.pages.items.value`** | String | Item hidden value. |
| **`structured_content.pages.items.id`** | String | Item identifier. |
| **`structured_content.pages.items.next_page_id`** | String | **Optional**. Mandatory in `select` pages with `multiple_selection` missing or `false`. |
| **`structured_content.pages.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the item with an image.<br>Supports private attachments.<br>Should be jpg, jpeg or png.<br>Should be less than 5MB. |
| **`structured_content.pages.attachment_fallback_id`** | String | **Optional**. Fallback in case the attachment related to the attachment_id doesn’t meet the source requirements. Must be public. Only be jpg, jpeg or png. Maximum size of 5MB. [Upload attachments](../../../basics/uploads) for you own custom images. |
| **Options Settings** | | |
| **`structured_content.pages.options.date_format`** | String | *Optional*. Available for `datePicker` only. Uses [Unicode Date Format Patterns Standard](https://www.unicode.org/reports/tr35/tr35-31/tr35-dates.html#Date_Format_Patterns) |
| **`structured_content.pages.options.start_date`** | String | *Optional*. Available for `datePicker` only. Default value for the date picker. Must be ISO 8601 (`YYYY-MM-DD`) |
| **`structured_content.pages.options.maximum_date`** | String | *Optional*. Available for `datePicker` only. Maximum value for the date picker. Must be ISO 8601 (`YYYY-MM-DD`) |
| **`structured_content.pages.options.minimum_date`** | String | *Optional*. Available for `datePicker` only. Minimum value for the date picker. Must be ISO 8601 (`YYYY-MM-DD`) |
| **`structured_content.pages.options.label_text`** | String | *Optional*. Available for `datePicker` only. A string representing the text string to be shown next to date field. Defaults to text `Date`. |
| **`structured_content.pages.options.regex`** | String | *Optional*. Available for `input` only. A string representing a JSON encoded regular expression (regex) string to limit the type of input for input field to use. |
| **`structured_content.pages.options.placeholder`** | String | *Optional*. Available for `input` only. A text string used when there is no other text in the input text field. Default value are `Required` or `Optional`. |
| **`structured_content.pages.options.required`** | Boolean | *Optional*. Available for `input` only. A Boolean value that defaults to `false`. When set to `true`, the next button on page is disabled until the user provides input. |
| **`structured_content.pages.options.input_type`** | String | *Optional*. Available for `input` only. A string value that defaults to `singleline`. Other values are `multiline` or `singleline`. |
| **`structured_content.pages.options.prefix_text`** | String | *Optional*. Available for `input` only. A string value representing optional text shown next to the text field. This value defaults to an empty string. Only applies to inputType : singleline. |
| **`structured_content.pages.options.max_character_count`** | Integer | *Optional*. Available for `input` only. An integer value representing the field size in characters for singleline and multiline. The field size defaults to 30 characters for `singleline` and 300 characters for `multiline`. |
| **`structured_content.pages.options.keyboard_type`** | String | *Optional*. Available for `input` only. Type of keyboard to be shown. Available values are: `default`, `asciiCapable`, `numbersAndPunctuation`, `numberPad`, `phonePad`, `namePhonePad`, `emailAddress`, `decimalPad`, `UIKeyboardTypeTwitter`, `webSearch`. |
| **`structured_content.pages.options.text_content_type`** | String | *Optional*. Available for `input` only. A string value representing the keyboard and system information about the expected semantic meaning for the content that users enter. Available values are: `name`, `namePrefix`, `givenName`, `middleName`, `familyName`, `nameSuffix`, `nickname`, `jobTitle`, `organizationName`, `location`, `fullStreetAddress`, `streetAddressLine1`, `streetAddressLine2`, `addressCity`, `addressState`, `addressCityAndState`, `sublocality`, `countryName`, `postalCode`, `telephoneNumber`, `emailAddress`, `URL`, `creditCardNumber`, `username`, `password`, `newPassword`, `oneTimeCode`. |

## Example: Apple Messages for Business (Form Message)

Nothing specifically unique as this is an Apple Messages for Business specific structured message type.

### Splash

<img class="img-fluid" width="350" src="../../../img/structured-messages-apple-form-splash.png">

### Date Picker

<img class="img-fluid" width="350" src="../../../img/structured-messages-apple-form-date-picker.png">

### Select

<img class="img-fluid" width="350" src="../../../img/structured-messages-apple-form-select.png">

### Picker

<img class="img-fluid" width="350" src="../../../img/structured-messages-apple-form-picker.png">

### Input

<img class="img-fluid" width="350" src="../../../img/structured-messages-apple-form-input.png">

## Get the form responses

The form responses can be directly viewed in the conversation. If you wish to get them via API you can do so by accessing the `structured_reply` field. It is also included in the webhook of `content.imported` events.

```bash
curl -X GET "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents/633567aeb42e443f436ac389"
```

```json
{
  "id": "633567aeb42e443f436ac389",
  "created_at": "2022-09-29T09:38:54Z",
  "updated_at": "2022-09-29T11:25:08Z",
  "approval_required": false,
  "author_id": "63344dcbb42e44256eaf28cf",
  "body": "Report a product issue",
  "body_input_format": "text",
  "category_ids": [],
  "creator_id": null,
  "foreign_id": "ac7f81aa-f881-4857-85da-5f7a2b37cb13",
  "intervention_id": "63344dd1b42e4428b86b5c0d",
  "language": "en",
  "remotely_deleted": false,
  "source_id": "6038c1c69090ee8393c2f169",
  "source_url": null,
  "synchronization_status": "success",
  "status": "ignored",
  "thread_id": "63344dcbb42e44256eaf28d0",
  "in_reply_to_id": "63356768b42e443f08f32cde",
  "in_reply_to_author_id": "6038c1c69090ee8393c2f16e",
  "attachments_count": 0,
  "attachments": [],
  "auto_submitted": false,
  "body_formatted": {
    "text": "Report a product issue",
    "html": "<p>Report a product issue</p>"
  },
  "context_data": null,
  "created_from": "synchronizer",
  "private_message": true,
  "published": true,
  "source_type": "apple_business_chat",
  "structured_content_supported": true,
  "structured_content": {},
  "structured_reply": {
    "responses": [{
        "page_id": "page_1",
        "title": "What date did you receive, or expect to receive the product?",
        "items": [{
          "id": "page_1",
          "title": "09/29/2022",
          "value": "2022-09-28T22:00:00Z"
        }]
      },
      {
        "page_id": "page_2",
        "title": "What was the issue with your product?",
        "items": [{
            "id": "item_id_1",
            "title": "The product does not work",
            "value": "doesnt_work"
          },
          {
            "id": "item_id_3",
            "title": "I haven't received the product",
            "value": "not_received"
          }
        ]
      },
      {
        "page_id": "page_3",
        "title": "What would you like to obtain as compensation?",
        "items": [{
          "id": "item_id_2",
          "title": "A new product",
          "value": "new_product"
        }]
      },
      {
        "page_id": "page_4",
        "title": "If you have anything you'd like to tell us please do",
        "items": [{
          "id": "page_4",
          "title": "Thank you",
          "value": "Thank you"
        }]
      }
    ]
  },
  "type": "form_response",
  "synchronization_error": null,
  "capabilities_supported": [
    "apple_pay",
    "rich_link",
    "authenticate",
    "list",
    "time_select",
    "select"
  ]
}
```
