# Receiving Events via a Webhook

By default, endpoints will not receive any events, as they are subscribed to none. To know more about the subscription process, see [Creating Webhooks](../create/).
Once subscribed to events, endpoints will receive updates in the following fashion:

* Each event update is sent using an HTTP POST request to the endpoint URL.
* The endpoint URL MUST return with a status code 200 to acknowledge the reception of the event.

## Event Delivery Guarantees

This Webhook API implementation tries to deliver the event notification request reasonably fast. We may buffer notification for efficiency reasons.

Engage cannot guarantee the delivery of every event (see [Endpoint Unavailability](#endpoint-unavailability)). However, Engage will retry delivery during the first 24 hours after the event was triggered. Engage will retry until the endpoint received them correctly by responding with an HTTP status code of 200. 

Engage cannot guarantee the order in which the events arrive. Engage does however provide a timestamp in every event for the implementor to understand the exact point in time the event was triggered. 

Engage cannot guarantee that events will be sent only once. A unique identifier for each event is provided so that events can be deduplicated by the implementor if unicity is a constraint.

## Endpoint Unavailability

When an event notification request to an endpoint is not responsed by an HTTP 200 OK status (network issue, endpoint is down ...), the request will be retried. The request will be retried during 12 hours almost 10 times with waiting time growing.

To give an order of magnitude of the waiting time before each retry (those data are for information only and are susceptible to change):

| Request | Time interval | 
|-|-|
| Try #1  | ~ 1 sec |
| Try #2  | ~ 15 sec | 
| Try #4  | ~ 5 min |
| Try #5  | ~ 10 min | 

It means that when the first request fails the next retry is processed after 1 sec, then 15 seconds after the last retry... If it’s still failing this request will never be sent afterward.

!!! warning "Automatic Disablement"
    If your endpoint's error rate is too high over a long period of time, we MAY have to disable delivery of event notifications. You will need to fix the problem before re-enabling the event notifications in the admin interface.

## Implementation Considerations

While the rate of event updates is usually reasonable, it can greatly increase depending on external factors. Therefore, it is recommended that the endpoint URL does as little as possible, ideally just storing the event information in a queue for background processing.

Event updates will be sent at least once. In some cases, the same event MAY be sent several times, therefore the endpoint SHOULD be idempotent. In order to help achieve idempotency, each event update has a unique UUID, that will persist if sent more than once.

## Diagram - Webhook events workflow

Below are some diagrams describing specific workflows for sending webhook events.

## Task Life-cycle

[![](../dimelo_webhook_task-life-cycle_800x.png)](../dimelo_webhook_task-life-cycle_full.png)
