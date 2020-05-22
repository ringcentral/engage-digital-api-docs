# Implementing an Engage Chatbot

Below is the detail of the different information you will have to provide while initializing the multi-channel bot, as well as the outlines of each part and the condition tests to be done at the moment of launching of the integration.

!!! tip "Testing process"
    During the testing process, it is recommended to use a source for the test. This allows to test the integration of the multi-channel bot on Engage’s side in the real conditions of production without impacting all agents already using Engage Digital.

After you familiarize yourself with [the various flows for processing customer inquiries via a bot](../overview/), we can distill the process of creating an Engage Bot down even further into the following steps.

1. Build a service to receive notifications of new messages.
2. Build an AI to determine how to handle an incoming message.
3. Take action on the message

Let's explore each of these steps in greater detail. 

## Receive a new message notification

Engage Digital will send an event notifying you of a new message posted by a customer to a URL via a "webhook." The process of creating and validating a webhook endpoint is discussed in greater detail within [Creating Webhooks](../../../webhooks/create/).

Once the Webhook has been created and validated, Engage will begin transmitting via HTTP POST event messages corresponding to the channels associated with the corresponding webhook and endpoint URL of your choice. 

**Sample New Message Notification**

```json
{
  "id":"50df05692c64ac7799f09a98",
  "domain_id":"48cc6703bdae1462ce06a555",
  "events":[
    {
      "type":"content.imported",
      "id":"5950328914bf8a22305f446a",
      "issued_at":"2017-06-25T22:00:41.339Z",
      "resource":{
        "type":"answers/private_message",
        "id":"5950328814bf8abf7f79e3cb",
        "metadata":{
          "approval_required":false,
          "author_id":"536cbd9e7aa58d58f3000050",
          "body":"This is just one more test email, sent automatically",
          "body_input_format":"text",
          "creator_id":null,
          "date":"2017-06-26",
          "first_in_thread":false,
          "foreign_categories":[

          ],
          "foreign_id":"1056855",
          "has_attachment":false,
          "intervention_id":"541014eea90ffb3f600000ac",
          "in_reply_to_author_id":"4dcceb1f2f1a692a2a000057",
          "in_reply_to_id":"541015010f4ca111df0000b0",
          "source_id":"4f97fbea7aa58d073900344f",
          "status":"assigned",
          "thread_id":"541014e17aa58d8ccf000024",
          "thread_title":"This is a nagios check thread, please don't delete!",
          "created_from":"synchronizer",
          "private":true
        }
      }
    }
  ]
}
```

When your service receives a new message notification, it must respond with an HTTP status code of 200, otherwise Engage will consider the message delivery a failure, and will attempt to resend the webhook again. 

!!! warning "Webhook Disablement"
    After 10 unsuccessful delivery attempts of a notification via a webhooks (spread out over 24 hours) the Webhook will be automatically disabled.

## Process the incoming message

Upon receiving a new message, it is up to your bot or service to determine the best way to respond to the customer. This is where you engage whatever workflows are most appropriate for your service and the inquiry posted by the customer.

When your service has determined the best course of action to take, it should post a message back to the customr. 

## Take action

Your bot can take any number of actions as it relates to incoming message. The most common of which being to post a reply. Here are all the actions a bot might take in response to an incoming message:

* [Post a reply](../replies/)
* [Transfer to a Human Agent](../transfers/)
* [Close the thread](../threads/)


