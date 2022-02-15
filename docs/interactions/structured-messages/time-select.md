# About Time Select Structured Messages

This structured message provides a list of dates from which the customer can choose an appointment. See [Channel capabilities](../#channel-capabilities) to know on which channel you can use this structured message.

## Request Example

```bash
curl -X POST "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents"
```

```json
{
  "source_id": "<source_id>",
  "in_reply_to_id": "<in_reply_to_id>",
  "body": "Choose your appointment",
  "structured_content": {
    "type": "time_select",
    "attachment_id": "<attachment_id>",
    "location": {
      "latitude": 48.874989,
      "longitude": 2.345589,
      "radius": 100,
      "title": "Engage Office"
    },
    "timeslots": [
      {
        "duration": 3600,
        "start_time": "2020-04-26T08:00+0200",
        "identifier": "identifier 1"
      },
      {
        "duration": 3600,
        "start_time": "2020-04-28T08:45+0200",
        "identifier": "identifier 2"
      },
      {
        "duration": 5,
        "start_time": "2020-04-28T22:45+0200",
        "identifier": "identifier 3"
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
| **`body`** | String | The time_select structured message body. |
| **`structured_content`** | Object | Payload of the structured message. |
| **Structured Content Settings** | | |
| **`structured_content.type`** | String | Type of the structured message. Must be "time_select". |
| **`structured_content.timeslots`** | Array | Represents the different options for the event. |
| **`structured_content.attachment_id`** | String | **Optional**. Existing attachment id used to decorate the time_select with an image.<br>Supports private attachments.<br>Should be jpg, jpeg or png.<br>Should be less than 5MB. |
| **`structured_content.location`** | Object | **Optional**. Represents the location of the event. |
| **Time Slot Settings** | | |
| **`structured_content.timeslots.duration`** | Integer | The duration of the event in seconds. |
| **`structured_content.timeslots.start_time`** | String | The start time of the event using the format "2020-04-26T08:00+0200". |
| **`structured_content.timeslots.identifier`** | String | The identifier of the time slot. |
| **Location Settings** | | |
| **`structured_content.location.latitude`** | Number | The latitude of the location. |
| **`structured_content.location.longitude`** | Number | The longitude of the location. |
| **`structured_content.location.radius`** | Integer | The radius around latitude/longitude of the location. |
| **`structured_content.location.title`** | String | The title of the location. |

## Example: Apple Business Chat (Time Picker)

Nothing specifically unique as this is an Apple Business Chat specific structured message type.

<img class="img-fluid" width="350" src="../../../img/structured-messages-time-select-apple-biz.png">
