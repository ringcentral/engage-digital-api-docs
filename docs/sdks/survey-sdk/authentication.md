# Request/Response Authentication

From a technical standpoint, a Survey SDK bridge is an HTTP endpoint that accepts `GET` requests and responds with JSON.


## Request Authentication

Along with the requests to fetch the survey configuration and the responses, RingCX Digital will also send a specific `X-SDK-SECRET` header containing the survey's **secret key** (see [Quick Start](../quick-start#survey-creation-in-engage-digital)) to allow you to make sure that the request actually comes from RingCX Digital.

!!! note
    Please note that the `X-SDK-SECRET` header will be sent by RingCX Digital regardless of the survey's **verify secret key** setting.


## Response Authentication

When sending a request to a Survey SDK bridge, RingCX Digital will check both the presence and the validity of the `X-SDK-SECRET` header in the bridge HTTP response.

If the header is not present or if it doesn't match the survey's **secret key**, the entire response will be considered invalid and will therefore be ignored.

!!! note
    Please note that this verification mechanism can be skipped by turning the survey's **verify secret key** setting off in RingCX Digital.


## Realtime Authentication

When receiving a `POST` request from a Survey SDK bridge, RingCX Digital will check both the presence and the validity of the `X-SDK-SECRET` header in the bridge HTTP request.

If the header is not present or if it doesn't match the survey's **secret key**, the entire request will be considered invalid and will therefore be ignored.

!!! note
    Please note that this verification mechanism can be skipped by turning the survey's **verify secret key** setting off in RingCX Digital.
