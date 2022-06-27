# Registering a Webhook

To register an endpoint to begin receiving webhooks, using the Engage Digital admin interface under the "Webhooks" section under "Advanced Settings."

<img class="img-fluid" src="../dimelo_webhook_register_800x.png">

In details:

<img class="img-fluid" src="../dimelo_webhook_register-details.png">

## Fields

| Field Name | Description |
|-|-|
| **API access token** | You can only use the API access token belonging to your user, unless you have the "Manage API access token" permission. | 
| **Name** | User-friendly name of the given webhook that can be used to differentiate it among all others. The field is mandatory in the UI and optional when creating a webhook through the API. If it's not provided when the webhook is created in API, this field will be prefilled with the hostname from the URL + associated token name. |
| **Active** | This checkbox toggles the webhook on and off, allowing you to halt the flow of event notifications temporarily or permanently. 
| **Staging Environment ** | A toggle that configures which environment, staging or production, from which event notifications will be sent. |
| **Ignore SSL validity** | A toggle to turn on and off SSL cert verification. This is often needed in testing environments for which SSL certs may not be created by a certified authority. | 
| **URL** | The URL to which events will be posted. This URL must be accessible to RingCentral Engage. We strongly advise to use HTTPS for this endpoint. The URL is unique by API access token. |
| **Verification Token** | Any arbitrary string used during the validation phase of the webhook setup process. We recommend the use of a randomly generated string. |
| **Secret** | A secret that will be served as a `X-Dimelo-Secret` HTTP header with every request (including the validation request). This field must only contain ASCII characters and be less than 256 characters. The header won’t be present if there’s no secret configured. Changing this field will trigger a validation request. |
| **Registered Events** | A list of all **registrable events**, only selected events will be used by your endpoint. At least one event must be selected to create your webHook. |
| **Sources filtering strategy** |  A simple rule for determining which sources will be subscribed to, e.g. "all sources except" or "no source except." By default, events associated to any source are broadcast. |
| **Sources** | A list of sources for which events should or should not be broadcasted, depending on the "filtering strategy" used. |

!!! tip "Multiple Webhooks "
    Rather than creating a single endpoint listening to all the events, it is possible to create a number of specialized endpoint, listening to only certain events.

!!! note "Duplicate Events"
    If more than one endpoint are registered to a specific event, event updates will be sent to every registered endpoint.

When submitting your settings, if the required fields are not filled out, it will display an error message, or if it’s valid it will save the data and a subscription validation will be triggered. In case the subscription failed, ​your settings will be updated and this endpoint will be disabled​, you will have to go back on the API administration page to re-enable your endpoint.

## Webhook Validation

The Webhook API doesn’t use authentication in the conventional sense. Instead, Engage validates each registered URL/endpoint prior to directing events to it, unless the webhook has a status of "inactive." Validating will commence immediately upon registering your webhook, so be sure you have setup your endpoint to respond to the validation request prior to registering it. 

### Validation Request

Webhook validation is an asynchronous process. To validate a webhook, Engage will transmit a GET request on the endpoint URL with the following parameters:

| Parameter | Description |
|-|-|
| `hub.mode` | The string "subscribe" is passed in this parameter. |
| `hub.challenge` | A random string. |
| `hub.verify_token` | The verify token entered when setting up the enabled URL. |

When the endpoint receives this request, it should do all of the following:

1. Verify that the content-type set to application/json in the header.
2. Check the `hub.mode` is set to "subscribe." This confirms we are registering an endpoint URL.
3. Check the `hub.verify_token`'s validity, to ensure the request comes from the Webhook API.
4. Set the response body to only contain the value of `hub.challenge`.
5. Respond with an HTTP status code of 200.

## Webhook auto disconnection

To improve webhook management, Engage has a specific mechanism that will control if sent webhook has been acknowledged or not. If the endpoint returns error in response to the webhook within the given threshold (e.g. if 75% of requests are failed), we'll send the email notifying that something went unexpected. After that, until the situation recovers, we will stop sending webhooks except for ones that are required for a periodical check of the endpoint. Once everything normalizes, the webhooks sending will be auto-recovered.



