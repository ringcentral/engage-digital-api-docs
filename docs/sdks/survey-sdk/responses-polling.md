# Responses Polling

Engage Digital will poll your survey responses by performing a `GET /responses` on your bridge every 15 minutes along with a `since` parameter corresponding to the the most recent response fetched's `submitted_at` field.

This parameter will be sent as milliseconds in utc time (e.g. if **base URL of the bridge** is `https://survey.bridge.com` and your most recent response was submitted on `2022-01-10T10:00:00.000Z` then Engage Digital will perform a `GET https://survey.bridge.com/responses?since=1641808800000`).

## Response example

Here's an example of a valid response to the responses polling:
```json
{
  "id": "12345",
  "records": [
    {
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
    },
    ...
  ]
}
```


### Detailed response format

| Field | Type | Mandatory | Description |
|-|-|-|-|
| id | String | NO | Unique identifier of the survey. |
| records | Array | **YES** | List of responses, see [Response object description](#response-object-description). |


### Response object description

| Field | Type | Mandatory | Description |
|-|-|-|-|
| id | String | **YES** | Unique identifier of the response. |
| submitted_at | Time | NO | Time the response was submitted at. |
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
