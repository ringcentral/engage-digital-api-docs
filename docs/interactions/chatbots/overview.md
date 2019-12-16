# Engage Chatbots

Engage Digital allows developers to create automated agents, or bots, to assist in processing client inquiries. Bots can help in the following ways:

* Pre-qualify a client inquiry
* Augment incoming inquiries with additional data about the customer or request
* Automate the handling of common customer/client inquiries or interactions

!!! tip "Before you begin..."
    This document aims to describe the process of integration of a multi-channel bot within Engage’s architecture. It does not describe the operating of the tools used for the integration nor does it explain the life-cycle of messages within Engage Digital. A thorough knowledge of Engage tools and of its operating is necessary to comprehend the entire process described in this documentation.
    
    Below is a list of documentation that we recommend you read carefully beforehand:
    
    * [Engage Webhook API](https://developers.ringcentral.com/engage/guide/webhooks) - Functionality that allows automatic events concerning the message’s life-cycle from Engage Digital Platform to be sent to a predefined url
    
    * [Engage REST API](https://developers.ringcentral.com/engage/guide/basics) -  Engage Digital’s API which allows to retrieve data stored in it.

## Bot Architecture

Engage Bots work primarily through what we call a "proxy," or an intermediary service that brokers the interaction between Engage Digital's API and an AI engine. The proxy is responsible for processing incoming events from Engage, determining the appropriate course of action to take by interfacing with the bot's AI, and finally formulating responses to the customer when applicable. The components of an Engage Bot integration are therefore the following:

1. **Engage Digital**
2. **Bot Proxy**
3. **Bot AI**
    
You can see the relationship between these components clearly in the diagram below. 

<img src="../dimelo_chatbot_overview.png" class="img-fluid">

## Bot Interaction Flow

In an ideal interaction within this architecture, the following steps will be followed. 

1. Engage Digital receives a message posted by a customer. 
2. Engage Digital posts an event notification via a webhook to notify the bot of the newly posted message. 
3. The Bot Proxy fetches additional information if necessary via Engage's REST API.
4. The Bot Proxy sends all the gathered information to its Bot AI.
5. The Bot AI formulates an answer according to the data it has received.
6. The Bot AI transmits its generated answer back to the Bot Proxy.
7. The Bot Proxy formats the answer.
8. The Bot Proxy submits the formated answer via the Engage REST API.
9. Finally, the answer is transferred to the customer via Engage.

The diagram belows shows this interaction sequence as well.

<img src="../dimelo_chatbot_success.png" class="img-fluid">

### Interventions

Not pictured above is the optional step of signaling that a response is required from a human agent, referred to as an "intervention." The responses a bot generates do not automatically create an intervention. Instead, when human interaction is required, an [intervention](../interventions/) can be created via the API. This provides operators with better analytics about the bot's process and allows traceability of interventions in the CRM.

## Handling Failures

The failure scenario corresponds to the process during which the AI is not able to generate an answer to the message sent by the proxy. When this happens, the following workflow is recommended:

1. Engage Digital receives a message posted by a customer. 
2. Engage Digital posts an event notification via a webhook to notify the bot of the newly posted message. 
3. The Bot Proxy fetches additional information if necessary via Engage's REST API.
4. The Bot Proxy sends all the gathered information to its Bot AI.
5. The Bot AI fails to generate an answer.
6. The Bot AI notifies the Bot Proxy of its failure to generate a response.
7. The Bot Proxy determines an appropriate way to process the failure.
8. The Bot Proxy changes the category assigned to the chat thread to a category visible to human agents. 
9. A Human Agent responds to the customer request.
10. The Human Agent's answers are transmitted to the customer. 

The diagram below describes this process as well. 

<img src="../dimelo_chatbot_failure.png" class="img-fluid">

## Best Practices

* Engage recommends using bots on private channels or at least on a third-party source connected via an SDK. However Engage does not recommend its upload on public platforms such as Facebook, Twitter, etc because automatic messages are often banned.

