# Bots API

The Bots API is our way to integrate bots into RingCX Digital chat channels. This guide is meant to provide integrators a comprehensive overview of the available features for them to be able to create an implementation.

Currently it's only available on Messenger channels, whose associated documentation on Facebook is available [here](https://developers.facebook.com/docs/messenger-platform/handover-protocol).
## Setup

First you will need to choose which Facebook app is primary (RingCX or BOT app), based on that, handover protocol behavior won’t be exactly the same, see aforementioned Facebook documentation for more information.

You can change the app’s role for the source page whenever you want, you need to go on the Facebook’s page -> Settings -> Advanced messaging:

<img src="../../img/bots-api-set-primary-app.png" class="img-fluid">

<br>If your page uses New Pages Experience you will have to switch to your page’s profile. Then you go in Settings & privacy -> Settings -> New Pages Experience -> Advanced messaging:

<img src="../../img/bots-api-set-primary-app-new-experience.png" class="img-fluid">

After you’re done configuring that you will need to ensure that the bots uses well our new Bots api endpoint

## Using the endpoint

### Format

JSON response when trying to proceed to handover respects following format:

```json
{
  "status": 200,
  "message": "Success",
}
```

### Handover from bot to agent

#### Request example

```bash
curl -X POST "https://domain-test.api.engagement.dimelo.com/1.0/bots/handover?access_token=<access_token>&identity_foreign_id=<foreign_id>&from=bot&to=agent&type=messenger"
```

### Query Parameters

| Parameter | Description |
|-|-|
| `access_token` | API identification token |
| `identity_foreign_id` | The foreign_id of the identity for which you want to do the handover |
| `from` | The app that has thread control before handover |
| `to` | The app that should get thread control after handover |
| `type` | The type of channel |

When calling this endpoint, RingCX Facebook app will use Facebook’s request_thread_control api to ask for thread control, see [Facebook's documentation](https://developers.facebook.com/docs/messenger-platform/handover-protocol/conversation-control#request-control).

The Bot (or the app controlling the thread) app will then receive a webhook from Facebook. When the webhook is received , it's important for the bot app to call Facebook’s pass_thread_control endpoint to give agent an ability to control the thread, see [Facebook's documentation](https://developers.facebook.com/docs/messenger-platform/handover-protocol/conversation-control#pass-control-to-another-app) for more details.

### Handover from agent to bot

#### Request example

```bash
curl -X POST "https://domain-test.api.engagement.dimelo.com/1.0/bots/handover?access_token=<access_token>&identity_foreign_id=<foreign_id>&from=agent&to=bot&type=messenger&app_id=<bot_app_id>"
```

### Query Parameters

| Parameter | Description |
|-|-|
| `access_token` | API identification token |
| `identity_foreign_id` | The foreign_id of the identity for which you want to do the handover |
| `from` | The app that has thread control before handover |
| `to` | The app that should get thread control after handover |
| `type` | The type of channel |
| `app_id` | The id of the Facebook's app that controls the Bot |

This call will only pass if there is currently no intervention engaged with the linked identity.

If the call passes, ED Facebook app will use the pass_thread control Facebook API to give thread control back to the Bot app

## HOP features

On Facebook's side, when a customer sends a message to a Facebook page the thread control will go to the primary app. When no primary app is setup the Messenger thread will stay in IDLE mode

When an agent tries to send a message via RingCX while another app has thread control we will either:

* Request thread control if Bot app is primary
* Take thread control if RingCX app is primary

After that we will retry the synchronization of the content

When the intervention is closed, we will release thread control

## Workflow

### Bot primary

<img src="../../img/bots-api-workflow-bot-primary.png" class="img-fluid">

### RingCX primary

<img src="../../img/bots-api-workflow-ed-primary.png" class="img-fluid">

### No specified primary

<img src="../../img/bots-api-workflow-no-primary.png" class="img-fluid">

### When agent interrupts bot

<img src="../../img/bots-api-workflow-agent-interrupts-bot.png" class="img-fluid">
