# Configuration Fetching

When creating a Survey SDK in RingCX Digital you need to specify the **base URL of the bridge** field (see [Getting Started](../quick-start)) which will be used to perform an HTTP request to fetch the survey configuration.

Upon survey creation, RingCX Digital will perform a `GET /configuration` call on your bridge to fetch the survey configuration (e.g. if **base URL of the bridge** is `https://survey.bridge.com` then RingCX Digital will perform a `GET https://survey.bridge.com/configuration`).

## Response example

Here's an example of a valid Survey SDK configuration:
``` json
{
  "id": "12345",
  "link": "https://link.tothesurvey.com",
  "status": "Opened",
  "questions": [
    {
      "id": "main_question",
      "text": "Are you satisfied with our customer service?",
      "main": true,
      "choices": [
        {
          "text": "Very satisfied",
          "value": "5"
        },
        ...
      ]
    },
    {
      "id": "multiple_choices_question",
      "text": "What can we do to improve our service?",
      "main": false,
      "choices": [
        {
          "text": "Faster response time"
        },
        {
          "text": "Better follow-up"
        }
      ]
    },
    {
      "id": "free_text_question",
      "text": "Any other suggestion?",
      "main": false,
      "choices": []
    },
    ...
  ]
}
```


### Detailed response fields description

| Field | Type | Mandatory | Description |
|-|-|-|-|
| id | String | NO | Unique identifier of the survey. |
| link | String | **YES** | Link to access the survey. |
| status | String | NO | Status of the survey (display only). |
| questions | Array | **YES** | List of the questions belonging to the survey, see [Question object description](#question-object-description). |


### Question object description

| Field | Type | Mandatory | Description |
|-|-|-|-|
| id | String | **YES** | Unique identifier of the question, the entire question will be ignored if not present. |
| text | String | **YES** | Question text, the entire question will be ignored if not present. |
| main | Boolean | NO | Indicates whether this question is the survey's main question (a survey must have a single main question), will default to `false` if not present. |
| choices | Array | NO | List of the available choices for the question (can be omitted for free text question), see [Choice object description](#choice-object-description). |


### Choice object description

| Field | Type | Mandatory | Description |
|-|-|-|-|
| text | String | **YES** | Choice text, the entire choice will be ignored if not present. |
| value | String | NO | Value of the choice (e.g. "5", which is equivalent to a "Very satisfied" text value). |
