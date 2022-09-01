# Realtime Responses Import

As explained in [Getting Started](../quick-start) you can enable the import of responses in realtime, by doing so the **Realtime endpoint URL** will be displayed in the survey configuration page. Please note that if the option is not enabled, Engage Digital will not accept any incoming HTTP call on this **Realtime endpoint URL**.

The **Realtime endpoint URL** accepts HTTP `POST` requests in order to import survey responses in realtime (see [Payload example](#payload-example) for more detail on the payload expected by Engage Digital).


## Payload example

```json
{
  "record": {
    "id": "12345_response_1",
    "submitted_at": "2022-01-10T10:00:00.000Z",
    "mapping_key": "response_mapping_key",
    "questions": [
      {
        "id": "main_question",
        "replies": [
          {
            "value": "5"
          }
        ]
      },
      {
        "id": "multiple_choices_question",
        "replies": [
          {
            "value": "Faster response time"
          },
          {
            "value": "Better follow-up"
          }
        ]
      },
      {
        "id": "free_text_question",
        "replies": [
          {
            "value": "Everything's already perfect"
          }
        ]
      },
      ...
    ]
  }
}
```

### Detailed response format

| Field | Type | Mandatory | Description |
|-|-|-|-|
| record | Hash | **YES** | Response data, see [Response object description](#response-object-description). |


### Response object description

| Field | Type | Mandatory | Description |
|-|-|-|-|
| id | String | **YES** | Unique identifier of the response. |
| submitted_at | Time | **YES** | Time the response was submitted at, the entire response will be ignored if not present. |
| mapping_key | String | **YES** | Mapping key that was sent as the `i` parameter in the survey link, see [Response mapping key](#response-mapping-key). |
| questions | Array | **YES** | List of questions and their associated replies, see [Question object description](#question-object-description). |


### Question object description

| Field | Type | Mandatory | Description |
|-|-|-|-|
| id | String | **YES** | Unique identifier of the question, the whole entry will be ignored if not present. |
| replies | Array | **YES** | List of replies for a given question (can have multiple entries), see [Reply object description](#reply-object-description). |


### Reply object description

| Field | Type | Mandatory | Description |
|-|-|-|-|
| value | String | **YES** | Value of the reply (in the [example above](#response-example), value 5 will be equivalent to "Very satisfied" as seen in the [configuration](../configuration-fetching#response-example)). |


## Response mapping key

When sending the link to take a survey, Engage Digital will add an additional `i` parameter in the URL to bind the response with the conversation:

<img class="img-fluid" width="363" src="../../../img/survey-sdk-response-mapping-key.png">

!!! warning
    Please note that the response will be ignored by Engage Digital if the `mapping_key` parameter is not present or if it cannot be matched with an existing conversation.
