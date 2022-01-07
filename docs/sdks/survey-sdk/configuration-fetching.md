# Configuration Fetching

When creating a Survey SDK in Engage Digital you need to specify the **Base URL of the bridge** field (see [Getting Started](../quick-start)) which will be used to perform an HTTP request to fetch the survey configuration.

Your bridge should respond to `GET /configuration` with a `Content-Type: application/json` (e.g. if **Base URL of the bridge** is `https://survey.bridge.com` then Engage Digital will perform a `GET https://survey.bridge.com/configuration`).

Here's an example of a valid Survey SDK configuration:
``` json
{
  "id": "12345",
  "link": "https://link.tothesurvey.com",
  "status": "Opened",
  "questions": [
    {
      "id": "question_1",
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
    ...
  ]
}
```


### Survey object description

| Field | Type | Mandatory | Description |
|-|-|-|-|
| id | String | **YES** | Unique identifier of the survey. |
| link | String | **YES** | Link to access the survey. |
| status | String | NO | Status of the survey (display only). |
| questions | Array | **YES** | Survey questions. |


### Question object description

| Field | Type | Mandatory | Description |
|-|-|-|-|
| id | String | **YES** | Unique identifier of the question (the entire question will be ignored if not present). |
| text | String | **YES** | Question text (the entire question will be ignored if not present). |
| main | Boolean | NO | Indicates whether this question is the survey's main question (a survey should have at least one main question). |
| choices | Array | NO | List of the available choices for the question (can be omitted for free text question). |


### Choice object description

| Field | Type | Mandatory | Description |
|-|-|-|-|
| text | String | **YES** | Unique identifier of the survey. |
| value | String | **YES** | Link to access the survey. |
