# Introduction to Engage Webhooks

!!! note "Dimelo is Engage"
    This guide is for RingCentral Engage Digital, formerly known as Dimelo. The product names in the documentation have not been updated yet.

Dimelo's Digital Platform provides a Webhook API to be notified when events occur. This API is almost realtime and provides an alternative to polling the REST API. The purpose is to be able to trigger event-based behavior as soon as an event occurs (intervention assigned, deferred, canceled...). This way you will be able to build gamification mechanism, live dashboard, synchronizing contents ...

The overall mechanism is based on ​[Pubsubhubbub protocol 0.3​](https://pubsubhubbub.github.io/PubSubHubbub/pubsubhubbub-core-0.3.html), and rely on HTTP request containing event payload being sent to a consumer Endpoint URL.

## Development guidelines

We strongly advise developers consuming this API to respect the following guidelines:

* API input/output is standard [JSON](https://www.ietf.org/rfc/rfc4627.txt) object encoded using UTF-8
* Always access a field using its name, not using its rank in the document
* Never rely on the ordering of elements in a JSON object unless specified
* Always use a JSON parser, as the object is subject to change (e.g. no regexp parsing)
* Expect new fields to appear without notice in event structure, and make sure that this can be handled by your code. We will preserve existing field though for backward compatibility as long as we can.

## Getting Started

The following resources will help you navigate how to setup and begin utilizing webhooks in your code.

* [Create a Webhook](../create/)
* [Receiving Webhooks](../receive/)
* [Event Reference](../events/)