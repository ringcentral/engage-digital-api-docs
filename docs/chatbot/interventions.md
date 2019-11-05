# Bot Interventions

It is possible to create an intervention via API upon the first bot answer to a customer inquiry. This allows traceability in the CRM in the case where the bot manages the entirety of the conversation without turning to an agent. This allows also to boost statistics linked to interventions from the initial processing by the bot. In order to undertake this action, you must setup the following request:

Description of the API request:

## Create an Intervention

This method creates a new intervention or reopens an existing one. When successful, Engage will render the intervention to the user. Otherwise, Engage returns an error (422 HTTP code).

### Query Parameters

| Parameter | Description |
|-|-|
| `content_id` | The ID of the conversation thread to recategorize |
| `access_token` | API identification token |

### Request Parameters

| Parameter | Description |
|-|-|


**Sample Request**

`curl -X POST "https://[DOMAIN].api.engagement.dimelo.com/1.0/interventions?access_token=yyyyyyyyyyy&content_id=c93e3586250ff60181b6c2f0"`

**Sample Response**

```json
{
  "id":"3f55c8330da4144afd1c6728",
  "created_at":"2012-05-21T01:15:28Z",
  "updated_at":"2012-05-21T01:19:49Z",
  "source_id":"f18c81948c137d86dac77216",
  "thread_id":"9c9903dc3d559a8801eb5441",
  "content_id":"c93e3586250ff60181b6c2f0",
  "deferred_at":"2012-05-21T01:18:49Z",
  "identity_id":"8a8deed44623a4c44268c266",
  "comments_count":1,
  "closed":false,
  "closed_at":"2012-05-24T02:00:32Z",
  "custom_fields":{
    "external_id":"342901"
  },
  "category_ids":[
    "4d0fb475b242228032cbdf6d",
    "59248c4dae276a021cb296d2"
  ],
  "user_id":"d033e22ae348feb5660fc214",
  "user_replies_count":1,
  "user_reply_in_average":84959,
  "user_reply_in_average_bh":63000,
  "user_reply_in_average_count":1,
  "first_user_reply_id":"573446514379728247000001",
  "first_user_reply_in":0,
  "first_user_reply_in_bh":0,
  "last_user_reply_in":0,
  "last_user_reply_in_bh":0,
  "status":"Fermée"
}
```

## Closing an Intervention

In the case where you would have opened a bot intervention via the API, you can also close an intervention via the API when the support is ended. This will allow you to close the intervention when the bot is fully-capable of managing the entirety of the conversation or when the customer disconnects at some point during the conversation. To close an intervention, consult the instructions below. 

### Path Parameters

| Parameter | Description |
|-|-|
| `id` | The ID of the intervention to close |

### Request Parameters

| Parameter | Description |
|-|-|
| `access_token` | API identification token |

**Sample Request**

`curl -X PUT "https://[DOMAIN].api.engagement.dimelo.com/1.0/interventions/3f55c8330da4144afd1c6728/close?access_token=yyyyyyyyyyy"`

**Sample Response**

```json
{
  "id":"3f55c8330da4144afd1c6728",
  "created_at":"2012-05-21T01:15:28Z",
  "updated_at":"2012-05-21T01:19:49Z",
  "source_id":"f18c81948c137d86dac77216",
  "thread_id":"9c9903dc3d559a8801eb5441",
  "content_id":"c93e3586250ff60181b6c2f0",
  "deferred_at":"2012-05-21T01:18:49Z",
  "identity_id":"8a8deed44623a4c44268c266",
  "comments_count":1,
  "closed":false,
  "closed_at":"2012-05-24T02:00:32Z",
  "custom_fields":{
    "external_id":"342901"
  },
  "category_ids":[
    "4d0fb475b242228032cbdf6d",
    "59248c4dae276a021cb296d2"
  ],
  "user_id":"d033e22ae348feb5660fc214",
  "user_replies_count":1,
  "user_reply_in_average":84959,
  "user_reply_in_average_bh":63000,
  "user_reply_in_average_count":1,
  "first_user_reply_id":"573446514379728247000001",
  "first_user_reply_in":0,
  "first_user_reply_in_bh":0,
  "last_user_reply_in":0,
  "last_user_reply_in_bh":0,
  "status":"Fermée"
}
```

!!! note "Errors"
    If the intervention is already being closed, it will return a 409 error.
    
    To be able to close an intervention, it must meet the following criteria otherwise a 403 will be raised:
    
    * Intervention MUST NOT already be closed
    * Intervention MUST have agent replies
    * Access-Token agent MUST be the owner of the intervention or have the permission to edit permissions
    * Access-Token agent MUST have read access on the source