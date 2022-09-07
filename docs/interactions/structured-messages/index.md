# About Structured Messages

Engage Digital provides a REST API (using JSON bodies) to create structured messages from third party applications. This provides a richer messaging experience for customers.

<img class="img-fluid" width="396" src="../../../img/structured-messages-demo.png">
<img class="img-fluid" width="398" src="../../../img/structured-messages-demo-template.png">

<a target="new" href="https://ringcentral-tutorials.github.io/engage-digital-structured-messages-demo/index.html" class="btn btn-primary">Try it out</a>

## Types of Structured Messages

Engage Digital supports the following types of structure messages:

  | Type | Description |
  |-|-|
  | **`select`** | This structured message provides a list from which the customer can pick some elements. |
  | **`rich_link`** | This structured message allows to decorate a link with an image, title and subtitle. |
  | **`template`** | This structured message provides some buttons associated with a message. |
  | **`carousel`** | This structured message allows to group multiple templates in the same message. |
  | **`time_select`** | This structured message provides a list of dates from which the customer can choose an appointment. |
  | **`apple_pay`** | This structured message provides a convenient way to ask for a payment from the customer. |
  | **`authenticate`** | This structured message provides a way to send an OAuth2 authentication request to the customer. |

## Channel Capabilities

Engage Digital provides an omnichannel API to create structured messages. Since not every channel supports every type of structured message, this comparison grid allows to have a general overview of the different features and limitations of each channel.

* **<span style="color:green">Yes</span>**: Supported
* **<span style="color:orange">Not yet</span>**: Supported by channel but not by Engage Digital yet
* N/A: Not supported by channel

  |  | **select** | **rich_link** | **template** | **carousel** | **list** | **time_select** | **apple_pay** | **authenticate** |
  |-|-|-|-|-|-|-|-|-|
  | **Apple Messages for Business** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | N/A | N/A | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** |
  | **Facebook Messenger** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | N/A | N/A | N/A |  N/A |
  | **Engage Web Messaging** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | N/A | N/A | N/A |  N/A |
  | **Engage iOS Messaging** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | N/A | N/A | N/A |  N/A |
  | **Engage Android Messaging** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | N/A | N/A | N/A |  N/A |
  | **Google Business Messages** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | N/A | N/A | N/A |  N/A |
  | **Channel SDK** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | N/A | N/A | N/A |  N/A |
  | **Twitter Direct Messages** | **<span style="color:orange">Not yet</span>** | **<span style="color:orange">Not yet</span>** | **<span style="color:orange">Not yet</span>** | N/A | N/A | N/A | N/A |  N/A |
  | **Viber** | **<span style="color:orange">Not yet</span>** | **<span style="color:orange">Not yet</span>** | **<span style="color:orange">Not yet</span>** | **<span style="color:orange">Not yet</span>** | N/A | N/A | N/A |  N/A |
  | **WhatsApp** | N/A | N/A | **<span style="color:green">Yes</span>** | N/A | **<span style="color:green">Yes</span>** | N/A | N/A |  N/A |
  | **Instagram Messaging** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** | N/A | N/A | N/A | N/A |

## Channel Compatibilities

The Engage Digital omnichannel REST API strives to make a single payload compatible on every channel. However, each channel has their own features and limitations, this is why this omnichannel REST API will implicitly ignore or limit the different fields to adapt to the specified channel.

This section describes the generic compatibility logic of the API. Check the following section for the features and limitations specific to every channel.

| Limitation/Feature | Compatibility Description |
|-|-|
| Max number of items | Only processes the channel max number of items and ignores the other items. This applies for both the number of items in a template and the number of templates in a carousel.<br><br>For example, if the channel supports a maximum of 4 items and the payload has 10 items, the API will only keep the first 4 items and discard the last 6 items. |
| Max number of characters of a String | Truncates the String to the channel max number of characters and adds an ellipsis.<br><br>For example, if the channel supports a maximum of 80 characters and in the payload, the field is 100 characters, the API will truncate the field to 80 characters with an ellipsis. |
| Channel specific property | Ignores the property on channels not supporting it. Only the channel supporting the property processes it.<br><br>For example, the Engage Messaging channel supports the “center_items” property on the select structured content. If this property is also used to create a FB Messenger quick reply structured content, it will be ignored. |
| Payload property presence | When the payload property (optional by default) is required by a channel, the API will generate a default value to be sent.<br><br>For example, Google Business Messages requires the presence of the “payload” property. When blank, the API will generate a random value to be sent instead. |
| Deep links (url) | If the channel does not support deep links, we provide a url_fallback field to be used instead.<br><br>The url_fallback is validated **if the url is present**. The validations are:<ul><li>Scheme must be http or https</li><li>Max length of 2048</li></ul>For example, Engage Messaging supports deep links in its url fields. But when sending a payload with deep links on Facebook Messenger, the url_fallback field will be used if present, or return an error. |
| Non-standard attachment type | If the channel does not support the attachment, we provide an attachment_fallback_id field to be used instead.<br><br>The attachment_fallback_id is validated **if the attachment_id is present**. The validations are:<ul><li>Must be public</li><li>Must be less than 5MB</li><li>Must be a jpg, jpeg or png</li></ul>For example, Apple Messages for Business supports mp4 as its rich_link attachment. But when sending a payload with a mp4 attachment on Engage Messaging , the attachment_fallback_id field will be used if present, or return an error. |
| Structured content type not existing in the channel | If a structured content type does not exist in a channel, the API will try to best represent it by converting it to another type the channel supports.<br><br>See the channel specificities section for a more detailed behavior on this conversion.<br><br>For example, the rich link type does not exist in Google Business Messages. So an Engage Digital Rich Link is converted to a Google Business Messages Rich Card with a single button representing the url. |

## Channel Specificities

### Apple Messages for Business

* [List Picker (select)](../structured-messages/select/#example-apple-messages-for-business-quick-replies)
* [Rich Link](../structured-messages/rich-link/#example-apple-messages-for-business)
* [Time Picker](../structured-messages/time-select/#example-apple-messages-for-business-time-picker)
* [Apple Pay](../structured-messages/apple-pay/#example-apple-messages-for-business-apple-pay)
* [Apple Auth](../structured-messages/apple-auth/#example-apple-messages-for-business-apple-authenticate)

### Facebook Messenger

* [Quick Replies](../structured-messages/select/#example-facebook-messenger-quick-replies)
* [Rich Link](../structured-messages/rich-link/#example-facebook-messenger)
* [Template](../structured-messages/template/#example-facebook-messenger)
* [Carousel](../structured-messages/carousel/#example-facebook-messenger)

### Engage Messaging

* [Quick Replies](../structured-messages/select/#example-engage-messaging-quick-replies)
* [Rich Link](../structured-messages/rich-link/#example-engage-messaging)
* [Template](../structured-messages/template/#example-engage-messaging)
* [Carousel](../structured-messages/carousel/#example-engage-messaging)

### Google Business Messages

* [Suggested Replies](../structured-messages/select/#example-google-business-messages)
* [Rich Link](../structured-messages/rich-link/#example-google-business-messages)
* [Rich Card](../structured-messages/template/#example-google-business-messages-rich-card)
* [Rich Card Carousel](../structured-messages/carousel/#example-google-business-messages-rich-card-carousel)

### WhatsApp

* [List Messages](../structured-messages/select/#example-whatsapp-list-messages)
* [Reply Buttons](../structured-messages/template/#example-whatsapp-reply-buttons)

### Instagram Messaging

* [Quick Replies](../structured-messages/select/#example-instagram-messaging-quick-replies)
* [Template](../structured-messages/template/#example-instagram-messaging)
* [Rich Link](../structured-messages/rich-link/#example-instagram-messaging)
* [Carousel](../structured-messages/carousel/#example-instagram-messaging)

## Channel SDK

Channel SDK supports only public attachments. Before sending a payload to third-party endpoint attachment_id is replaced by “attachment_url”. All other capabilities are inherited from the base validation described in this document.
