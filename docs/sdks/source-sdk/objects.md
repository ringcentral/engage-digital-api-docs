# Intro to SDK Objects

Currently the SDK defines 5 types of objects:

* Messages: publicly available messages exchanged between any number of users. Examples: blog comments, forum posts.
* PrivateMessages: messages between 2 or more users that cannot be seen by others. Examples: direct messages on Twitter.
* Threads: provide a way to structure the messages. Examples: blog posts, forum threads. Not mandatory (eg. Twitter).
* Users: every other object type must have an author. This is the only object type that has no actions available.
* Status: provides a way to update messages and private messages, like marking them as read. Used only in [Send API](../send-api)

## Messages

We ignore extra fields and you only need to provide fields marked as required. Also, we consider a nulled field, an empty set, and the absence of a field as the same thing. Contents without required fields will be ignored.

| Name | Type | Description |
|-|-|-|
| actions | Array | Possible actions for this content. Example: `["show", "reply"]` |
| attachments | Array | ***Required unless body is present***. An array of objects containing the url or a base64 content and the associated filename for each attachment**. Example: `[{ "url": "https://www.ringcentral.com/brand.png" }]` or `[{ "content": "iVBORw0KGgoAAAANSUhEUgAAAfUAAAH2CAIAAAD", "filename": "master.jpg" }]` |
| author | Hash | ***Required*** An User object. |
| body | String | ***Required unless attachments are present.*** Maximum size is 20k characters. |
| categories | Array | Example: `["TV", "Internet"]` |
| created_at | DateTime | Supported formats: `2012-02-10`, `2012-10-01T17:18:40Z` (ISO 8601) |
| context_data | Hash | A hash having the (expected) form: {'fieldname' => 'value'}. *The keys must be [a-zA-Z-]*. Example: `{"country": "Germany", "age": "34", "fruits": ["Apple", "Pear"]}` |
| display_url | String | A full URI where anyone can see the specific comment. |
| format | String | html or text. Default on html if not specified. |
| id | String | ***Required*** |
| in_reply_to_id | String | It has to be a valid content ID if present. |
| ip | String | Example: `1.2.3.4` |
| language | String | Example `en` |
| structured_content | Hash | ***Optional.*** Please keep in mind that an author used for structured content creation must be marked as a "puppet" identity and configured as a controlled identity in Engage Digital. Example: [Structured contents](../structured-contents). |
| structured_reply | Hash | ***Optional.*** Can be used when object is a reply to a structured content with a payload. Example: `{ payload: "DEFINED_PAYLOAD" }` |
| thread_id | String | ***Required*** If the option `messages.unthreaded` is set this field will be ignored. |
| updated_at | DateTime | Supported formats: `2012-02-10`, `2012-10-01T17:18:40Z` (ISO 8601) |

** For .mp4 attachments you can disambiguate the type ("audio" or "video") by specifying it in the object. Example: `{ "url": "https://www.ringcentral.com/song.mp4", "type": "audio" }`

When we post a message, we'll send the following attributes:

| Name | Type | Description |
|-|-|-|
| attachments | Array | Example: `[{ "url": "https://www.ringcentral.com/brand.png", type: "image" }]` |
| author_id| String | |
| body | String | |
| categories | Array | |
| created_at | DateTime | Example: `2012-10-01T17:18:40Z` |
| format | String | Example: "text" |
| in_reply_to_id | String | |
| sender_name | String | Example: "Jane from RingCentral" |
| structured_content | Hash | ***Optional.*** Example: [Structured contents](../structured-contents). |
| thread_id | String | Unless the `messages.unthreaded` option is set |
| updated_at | DateTime | Example: `2012-10-01T17:18:40Z` |

## PrivateMessages

We ignore extra fields and you only need to provide fields marked as required. Also, we consider a nulled field, an empty set, and the absence of a field as the same thing. Contents without required fields will be ignored.

Please note that private messages don't use threads, but are organized based on `in_reply_to_id`.

| Name | Type | Description |
|-|-|-|
| actions | Array | Possible actions for this content. |
| attachments | Array | ***Required unless body is present.*** An array of objects containing the url or a base64 content and the associated filename for each attachment**. Example: `[{ "url": "https://www.ringcentral.com/brand.png" }]` or `[{ "content": "iVBORw0KGgoAAAANSUhEUgAAAfUAAAH2CAIAAAD", "filename": "master.jpg" }]`. |
| author | Hash | ***Required*** An User object. |
| body | String | ***Required unless attachments are present.*** Maximum size is 20k characters. |
| categories | Array | Example: `["TV", "Internet"]` |
| created_at | DateTime | Supported formats: `2012-02-10`, `2012-10-01T17:18:40Z` (ISO 8601) |
| context_data | Hash | A hash having the (expected) form: `{'fieldname' => 'value'}`. *The keys must be [a-zA-Z-]*. Example: `{"country": "Germany", "age": "34", "fruits": ["Apple", "Pear"]}`. |
| display_url | String | A full URI where anyone can see the specific comment. |
| format | String | html or text. Default on html if not specified. |
| id | String | ***Required*** |
| in_reply_to_id | String | It has to be a valid content ID if present. |
| ip | String | Example: `1.2.3.4`. |
| language | String | Example `en`. |
| recipient | Hash | ***Required*** An User object. |
| structured_content | Hash | ***Optional.*** Please keep in mind that an author used for structured content creation must be marked as a "puppet" identity and configured as a controlled identity in Engage Digital. Example: [Structured contents](../structured-contents). |
| structured_reply | Hash | ***Optional.*** Can be used when object is a reply to a structured content with a payload. Example: `{ payload: "DEFINED_PAYLOAD" }` |
| updated_at | DateTime | Supported formats: `2012-02-10`, `2012-10-01T17:18:40Z` (ISO 8601) |

** For .mp4 attachments you can disambiguate the type ("audio" or "video") by specifying it in the object. Example: `{ "url": "https://www.ringcentral.com/song.mp4", "type": "audio" }`

When we post a private message, we'll send the following attributes:

| Name | Type | Description |
|-|-|-|
| attachments | Array | Example: `[{ "url": "https://www.ringcentral.com/brand.png", type: "image" }]` |
| author_id| String | |
| body | String | |
| categories | Array | |
| created_at | DateTime | Example: `2012-10-01T17:18:40Z` |
| format | String | Example: "html" |
| id | String | |
| in_reply_to_id | String | |
| ip | String | |
| sender_name | String | Example: "Jane from RingCentral" |
| recipient_id | String | |
| structured_content | Hash | ***Optional.*** Example: [Structured contents](../structured-contents). |
| updated_at | DateTime | Example: `2012-10-01T17:18:40Z` |

## Threads

We ignore extra fields and you only need to provide fields marked as required. Also, we consider a nulled field, an empty set, and the absence of a field as the same thing. Contents without required fields will be ignored.

| Name | Type | Description |
|-|-|-|
| actions | Array | Possible actions for this content. Example: `["show", "reply"]`. |
| author | Hash | ***Required*** *unless no_content option activated*. |
| body | String | ***Required*** *unless no_content option activated*. Maximum size is 20k characters. |
| categories | Array | Example: `["TV", "Internet"]` |
| created_at | DateTime | Supported formats: `2012-02-10`, `2012-10-01T17:18:40Z` (ISO 8601) |
| context_data | Hash | A hash having the (expected) form: `{'fieldname' => 'value'}`. *The keys must be [a-zA-Z-]*. Example: `{"country": "Germany", "age": "34", "fruits": ["Apple", "Pear"]}`. |
| display_url | String | A full URI where anyone can see the specific comment. |
| format | String | html or text. Default on html if not specified. |
| id | String | ***Required*** |
| ip | String | Example: `1.2.3.4`. |
| title | String | |
| updated_at | DateTime | Supported formats: `2012-02-10`, `2012-10-01T17:18:40Z` (ISO 8601) |

When we post a thread, we'll send the following attributes:

| Name | Type | Description |
|-|-|-|
| author_id| String | |
| body | String | |
| categories | Array | |
| created_at | DateTime | Example: `2012-10-01T17:18:40Z` |
| format | String | Example: "text" |
| id | String | |
| sender_name | String | Example: "Jane from RingCentral" |
| title | String | |
| updated_at | DateTime | Example: `2012-10-01T17:18:40Z` |

## Users

We ignore extra fields and you only need to provide fields marked as required. Also, we consider a nulled field, an empty set, and the absence of a field as the same thing. Contents without required fields will be ignored.

| Name | Type | Description |
|-|-|-|
| id | String | ***Required***, ID of the author on the third party system (yours). User with the same ID will be considered to be the same identity in Engage Digital. |
| firstname | String | Firstname of the user if you can provide it. |
| lastname | String | Lastname of the user if you can provide it. |
| screenname | String | ***Required*** Name we display for this user. |
| avatar_url | String | |
| email | String | Confirmed email of the user, can be used to send email PM if no PM mechanism is supported. |
| home_phone | String | Home phone number. |
| mobile_phone | String | Mobile/cellular phone number. |
| created_at | DateTime | ***Required*** Supported formats: `2012-02-10`, `2012-10-01T17:18:40Z` (ISO 8601). |
| updated_at | DateTime | ***Required*** Supported formats: `2012-02-10`, `2012-10-01T17:18:40Z` (ISO 8601). |
| puppetizable | Boolean | Whether Engage Digital can use this user as identity to create message. False by default. |
| gender | String | Accepted values are "`man`" or "`woman`". |
| context_data | Hash | Data specific to the User, that will be stored in Engage Digital at the identity and identitygroup level. Expect a hash having the (expected) form: `{'field_name' => 'value'}`. *The keys must be [a-zA-Z-]*. Example: {`"country": "Germany", "age": "34", "fruits": ["Apple", "Pear"]}`. |

## Status

We ignore extra fields and you only need to provide fields marked as required. Used only in [Send API](../send-api)

| Name | Type | Description |
|-|-|-|
| id | String | ***Required***, ID of the message or private message on the third party system (yours). This is the ID you return on `message.create` and `private_message.create` actions |
