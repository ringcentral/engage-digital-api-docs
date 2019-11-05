# Transferring to a Human Agent

If you encounter the circumstance in which the bot is no longer able to answer, the Bot Proxy should transfer the conversation to an [human] agent.

By default, conversations managed by a bot will be assigned to a category accessible only by the agent responsible for the bot. Therefore to change control of a conversation back a human agent, you will need to recategorize the conversionsation using a category accessible to agents. 

However, before you change the conversation's category, it is advised that you check to see if agents are available to provide the customer with a response. 

## Checking for an available agent

Before one transfers a conversation to a human agent, it is important that the Bot Proxy checks the availability of agents that will be responsible for responding to the customer. For example, if you need to escalate a customer request, but everyone is on the phone or otherwise unavailable, then the user is likely to have a poor experience. To check the availability status of all connected agents, perform a GET request on the `/status` resource. 

### Query Parameters

| Parameter | Description |
|-|-|
| `access_token` | API identification token | 

**Sample Request**

`curl "https://[DOMAIN].api.engagement.dimelo.com/1.0/status?access_token=yyyyyyyyyyy"`

**Sample Response**

```json
[
  {
    "agent_id":"4f4f3a08a90ffb27ee000583",
    "channels":[
      {
        "id":"55794bd9416472d4e8050000",
        "name":"Async",
        "status":"available",
        "busyness":"busy"
      },
      {
        "id":"55794bda416472d4e8060000",
        "name":"Realtime",
        "status":"available",
        "busyness":"unoccupied"
      }
    ],
    "custom_status":null
  }
]
```

Possible values for `status` are:

| Enum | Description |
|-|-|
| null | The agent is not connected on this channel |
| `available` | The agent is available (green/yellow/orange dot) | 
| `away` | The agent is away (red dot) |

Possible values for `busyness` are:

| Enum | Description |
|-|-|
| `unoccupied` | The agent don't have any task in progress on this channel |
| `ok` | The agent has less task in progress than the soft limit allows (green dot) |
| `busy` | The agent has more task than the soft limit allows (yellow dot) |
| `full` | The agent is at full capacity / hard limit (orange dot) |

The `custom_status` attribute represents the custom `away` status of the corresponding agent and can contain one of two possible values:

1. null — The agent is available or in the generic "Away" status.

2. `{ "id":"5582b1f4776562af9b000008" }` — The id of the custom status

### Determining an Agent's Availability

In order to determine if an agent is available or not, you need to verify that each of the following are true:

* `name` = Name of the channel containing the targeting source
* `status` = "available"
* `busyness` != full

You can safely ignore all statuses belonging to agents that are not members of intended response channel because only members of that channel are authorized to respond. 

## Recategorizing a thread

After you have confirmed that an agent is available to respond to the request, you can transfer the thread to the category in which agents are available. This is done as follows.

### Path Parameters

| Parameter | Description |
|-|-|
| `id` | The ID of the conversation thread to recategorize |

### Request Parameters

| Parameter | Description |
|-|-|
| `thread_category_ids[]` | The IDs of the categories setup on the conversation thread |
| `access_token` | API identification token |

**Sample Request**

`curl -X PUT -d "thread_category_ids[]=58ff349313047d2c&thread_category_ids[]=4d0fb475b242228a32cbd" "https://[DOMAIN].api.engagement.dimelo.com/1.0/content_threads/9c9903dc3d559a6801ec544 1/update_categories?access_token=yyyyyyyyyyy"`

**Sample Response**

```json
{
  "id":"9c9903dc3d559a6801ec5441",
  "source_id":"d19c81948c137d86dac77216",
  "title":"ADSL modem iss1ue",
  "interventions_count":1,
  "contents_count":4,
  "closed":false,
  "category_ids":[
    "4d0fb475b242228a32cbdf6d",
    "59248c4dae276a041cb296d2"
  ],
  "thread_category_ids":[
    "4d0fb475b242228a32cbdf6d"
  ],
  "extra_data":{
    "custom_my_number":123456,
    "trigger_id":"foo"
  },
  "foreign_id":"ab-2031",
  "created_at":"2012-05-18T14:24:44Z",
  "updated_at":"2012-05-21T18:10:12Z"
}
```

### Errors

* If token’s user does not have "read" on thread’s source a 404 HTTP response will be returned.
* If the thread is already being categorized, a 409 HTTP response will be returned.

