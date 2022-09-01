# About Apple Pay Structured Messages

This structured message provides a convenient way to ask for a payment from the customer. See [Channel capabilities](../#channel-capabilities) to know on which channel you can use this structured message.

## Prerequisites
* An apple pay merchant ID.
* A way to request a payment session payload from Apple.
* A correctly configured Messages for Business account.

## Request Example

```bash
curl -X POST "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents"
```

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "body": "Please pay for the requested items",
  "structured_content": {
    "type": "apple_pay",
    "attachment_id": "<attachment_id>",
    "payment_gateway_url": "https://example.com/paymentGateway",
    "merchant_session": {
      <your merchant session payload>
    },
    "merchant_identifier": "merchant.com.example.abc",
    "merchant_capabilities": ["supports3DS"],
    "supported_networks": ["visa", "mastercard"],
    "merchant_name": "ABC Example Merchant ID",
    "country_code": "FR",
    "currency_code": "EUR",
    "items": [
      { "amount": "5.00", "label": "Carottes", "type": "final" },
      { "amount": "3.00", "label": "Tomates", "type": "final" },
      { "amount": "1.00", "label": "Oignons", "type": "final" }
    ],
    "total": { "amount": "9.00", "label": "Panier", "type": "final" }
  }
}
```

### Primary Parameters

| API Property | Type | Description |
|-|-|-|
| **`source_id`** | String | ID of the source. |
| **`in_reply_to_id`** | String | ID of the message being replied to. |
| **`body`** | String | The apple_pay structured message body. |
| **`structured_content`** | Object | Payload of the structured message. |
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "apple_pay". |
| **`structured_content.payment_gateway_url`** | String | The url of the payment gateway matching the one configured in your Messages for Business account. |
| **`structured_content.merchant_session`** | Object | The complete payload of the merchant session request from Apple. |
| **`structured_content.merchant_identifier`** | String | The merchant identifier configured in your Messages for Business account. |
| **`structured_content.merchant_capabilities`** | Array | Capabilities for the merchant id (example: ["supports3DS”]). |
| **`structured_content.supported_networks`** | Array | Support payment networks (example: ["visa”, "mastercard”, "amex”]). |
| **`structured_content.merchant_name`** | String | The merchant name as configured in Apple’s merchant configuration. |
| **`structured_content.country_code`** | String | Two-letter ISO 3166 country code. |
| **`structured_content.currency_code`** | String | Three-letter ISO 4217 currency code for the payment. |
| **`structured_content.total`** | Object | The total price of the items the customer is going to purchase. |
| **`structured_content.items`** | Array | **Optional**. Items the customer is going to purchase. |
| **`structured_content.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the apple_pay with an image.<br>Supports private attachments.<br>Should be jpg, jpeg or png.<br>Should be less than 5MB. |
| **Total Settings** | | |
| **`structured_content.total.amount`** | String | The total price. |
| **`structured_content.total.label`** | String | Label of the total price. |
| **`structured_content.total.type`** | String | Must be set to “final”. |
| **Items Settings** | | |
| **`structured_content.items.amount`** | String | Price of the individual item. |
| **`structured_content.items.label`** | String | Name of the individual item. |
| **`structured_content.items.type`** | String | Must be set to “final”. |

## Example: Apple Messages for Business (Apple Pay)

Nothing specifically unique as this is an Apple Messages for Business specific structured message type.

<img class="img-fluid" width="350" src="../../../img/structured-messages-apple-pay-apple-biz.png">
