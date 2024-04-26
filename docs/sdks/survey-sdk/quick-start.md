# Survey SDK Quick Start

In this Quick Start, we are going to help you start using the Survey SDK. Let's get started.

In order to be able to use a survey from any kind of third party supplier in RingCX Digital you will need to implement a bridge to make the communication between RingCX Digital and your survey supplier possible.


## Survey creation in RingCX Digital

To create a survey on RingCX Digital you need to navigate to **Digital->Surveys**

<img class="img-fluid" width="363" src="../../../img/survey-sdk-ed-survey-menu.png">

Click **Add** to create a new survey and select **Survey SDK**.

Enter the follow required fields:

| Field | Description |
|-|-|
| **Active** | Make sure this box is checked to indicate your survey is active. |
| **Name** | Give your survey a name (for display purpose only). |
| **Base URL of the bridge** | This is the root url that will be used by RingCX Digital to perform HTTP requests (to fetch the configuration and the responses). |
| **Secret key** | Secret key that will be sent in a `X-SDK-SECRET` header by RingCX Digital. |
| **Verify secret key** | If this option is enabled RingCX Digital will make sure that the bridge response contains the `X-SDK-SECRET` header and that its value is matching the one configured in RingCX Digital. |
| **Enable import of survey results via polling** | When this option is enabled RingCX Digital will regularly (every 15 minutes) make HTTP requests to fetch newly added responses, see [Responses Polling](../responses-polling). |
| **Enable realtime import of survey results** | Enabling this option will display the survey **Realtime endpoint URL** which you will be able to use to push new responses to RingCX Digital in realtime, see [Realtime Responses Import](../realtime-responses-import). |
| **Additional URL parameters** | List of conversation-related data that can be added to the survey link that will be sent to the end-user, see [Additional URL parameters](#additional-url-parameters). |

!!! warning
    To be able to save the survey your bridge must **at least** implement the survey configuration part (see [Configuration fetching](../configuration-fetching)).


### Additional URL parameters

This setting is allowing you to add various conversation-related data to the survey link that will be sent to the end-user (in order to analyze survey metrics per channel, category or any other available criteria).

Here's the list of each of the available option and their effect on the survey link:

| Displayed name in RingCX Digital | URL parameter name | Description |
| - | - | - |
| Channel ID | **c** | ID of the channel on which the conversation took place |
| Agent ID | **u** | ID of the agent that handled the conversation |
| Identity ID | **a** | ID of the identity |
| Skill category IDs | **sc** | List of IDs of the skill categories that were set on the conversation (separated by commas) |
| Analytics category IDs | **ac** | List of IDs of the analytical categories that were set on the conversation (separated by commas) |
| Disposition category IDs | **dc** | List of IDs of the disposition categories that were set on the conversation (separated by commas) |

!!! example
    If you selected **Channel ID** and that the survey link fetched in your configuration by RingCX Digital is `https://survey.com` then the link that is going to be generated will be `https://survey.com?i=mapping_key&c=channel_id`
