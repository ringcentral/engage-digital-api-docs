# Survey SDK Quick Start

In this Quick Start, we are going to help you start using the Survey SDK. Let's get started.

In order to be able to use a survey from any kind of third party supplier in Engage Digital you will need to implement a bridge to make the communication between Engage Digital and your survey supplier possible.


## Survey creation in Engage Digital

To create a survey on Engage Digital you need to navigate to **Digital->Surveys**

<img class="img-fluid" width="363" src="../../../img/survey-sdk-ed-survey-menu.png">

Click **Add** to create a new survey and select **Survey SDK**.

Enter the follow required fields:

| Field | Description |
|-|-|
| **Active** | Make sure this box is checked to indicate your survey is active. |
| **Name** | Give your survey a name (for display purpose only). |
| **Base URL of the bridge** | This is the root url that will be used by Engage Digital to perform HTTP requests (to fetch the configuration and the responses). |
| **Secret key** | Secret key that will be sent in a `X-SDK-SECRET` header by Engage Digital. |
| **Verify secret key** | If this option is enabled Engage Digital will make sure that the bridge response contains the `X-SDK-SECRET` header and that its value is matching the one configured in Engage Digital. |

!!! warning
    To be able to save the survey your bridge must **at least** implement the survey configuration part (see [Configuration fetching](../configuration-fetching)).
