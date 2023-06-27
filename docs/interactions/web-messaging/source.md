# Configuring the Source

To add a chat to an existing RingCentral Engage instance, you need to create an Engage Messaging source. Right after its creation, edit it to access advanced configuration.

## Avatar

You may choose an avatar used for the main identity that the agents use when posting messages.

Note that the avatar is displayed:

* in a 35x35px circle in the chat window header
* in a 24x24px circle in the chat window conversation on the left of the last message of every block of agent messages
* in a 35x35px circle in RingCentral Engage

We recommend a 100x100px image that fits being displayed rounded, to support high DPI screens.

When no avatar is configured in the source, in the chat window, there is no avatar and in RingCentral Engage, a stylized avatar is used like any other identity without avatar.

## Sender Name

The displayed name of the agent in the chat window, next to each of its messages, is configurable via the `Sender name` field.

Here is an example that specify the rendering of the sender name for some among the most common cases, for an agent whose name is Alicia Johnson:

| Sender name configured | Rendering | Strategy |
|-|-|-|
| {{firstname}} {{lastname}} | Alicia Johnson | Agent centric |
| {{firstname}} | Alicia | Agent centric |
| {{firstname}} from PurpleShop | Alicia from PurpleShop | Agent centric |
| Natalie from PurpleShop | Natalie from PurpleShop | Persona |
| PurpleShop Team | PurpleShop Team | Impersonal |

!!! information
    The chat is a one-to-one channel that emphasizes the fact that the visitor speaks to a particular agent. Depending on the communication strategy, the customer can place their agents on the spotlight by displaying their first name, or opt for a unified persona that represents the customer care service.

## Integrations

RingCentral Engage Messaging supports natively several basic integrations with trackers for example. These are optional. They are made with the chat public Javascript API and one can copy and adapt them freely.

Available integrations:

* Google Analytics Classic (asynchronous mode - ga.js - the old way)
* Google Universal Analytics (analytics.js)

The integrations reports the following events (based on [JS events](../javascript-api/#events)) directly to the Google Analytics available on the page:

| Event name | Data | Details |
|-|-|-|
| chat_proposed | trigger | Fired when a trigger was activated and able to execute its actions (e.g. the item to be shown is found on thepage) |
| chat_engaged | trigger | Fired when the visitor engages the conversation, i.e. sends a first message |
| chat_shown | trigger | Fired when a chat client is created and shown |
| message_received | trigger | Fired each time the visitor receives a message |
| message_sent | trigger | Fired each time the visitor sends a message |

## Other Settings

### Queuing

The `Queuing factor` field, defaulting to "0", controls the extra slots that are allowed in the waiting queue. With "0", no extra slots are allowed. With a positive decimal number, such as "0.25" for example, and 4 busy agents with soft limit at 2, the chat would allow 2 more discussions to be queued even if no agents are available (4 * 2 * 0.25), the total number of slotswould be 10.

### Timed Messages (Live chat mode only)

The `Impossibility message delay` field, defaulting to "8m" (i.e. 8 minutes), controls the delay before a message is automatically shown to the visitor to state that it was impossible to reach an agent to handle the engaged conversation. One of the effects of this message is to end the conversation too.

These fields accept "chronic durations" (i.e. a series of value plus suffix corresponding to an unit: "2m 25s" for 2 minutes and 25 seconds). Set to 0 to disable.

### Enable Navigation Contents

The `Enable navigation contents` field, defaulting to true, can be unchecked to keep the server from creating a Navigation record displayed among the messages for each page seen by the visitor. Note that the current page of the visitor remains visible (and always gets updated) on the Messaging session panel on the right of the Task agent view.

### Welcome Message

The `Welcome messages` field can be filled to display a message on a new conversation **if** there is no automatic message from a trigger. For example:

<img class="img-fluid" width="413" src="../../../img/web-messaging-welcome.png">

### Automatic Email Request

On web messaging mode only, an automatic request for email can be sent to the visitor if we don’t already have their email address:

<img class="img-fluid" width="323" src="../../../img/web-messaging-enable-email.png">

This message will be displayed to the visitor only after their first message on a new conversation:

<img class="img-fluid" width="303" src="../../../img/web-messaging-enter-email.png">

Then after the visitor submits their email, this message will be displayed:

<img class="img-fluid" width="303" src="../../../img/web-messaging-email-thanks.png">

This automatic message will only be sent with a grace period of 2 hours.

### Back to Chat Link

On web messaging mode only, an email notification containing unread messages is sent to the user’s email if he hasn’t read them in 15 minutes. The link to go back to the chat and see the messages can be configured in the source under the "Notifications" section:

<img class="img-fluid" width="469" src="../../../img/web-messaging-chat-link.png">

If blank, the link will be the last visited page.

### Time for Thread Inclusion and Threads Auto Closing

This parameter is mainly used in the mobile single-threaded part of this connector. When receiving a new message from the customer, if his last message was within the time for thread inclusion, then it will be added to the last conversation. If not, it will create a new conversation.

!!! warning
    This parameter is also used for auto closing the messaging threads. The messaging thread will be automatically closed by the application if the inactivity in the thread is at least the max between this parameter and 4 days. This is to ensure we do not leave open threads forever.

### Enable Outbound Messaging

This setting is available for web messaging, iOS messaging and Android messaging. It allows to initiate discussions on behalf of an agent through Engage Digital's API.

It can be configured under the "Advanced settings" section:

<img class="img-fluid" width="276" src="../../../img/web-messaging-enable-outbound-messaging.png">

Enabling this option will unlock the `Initiate discussion` permission for this source. You will then be able to add it to agents you want to send messages on behalf of.

## Related Configuration

### Engage Messaging Specific Permissions

You must ensure that each of the agents that will handle the messaging conversation has the `Read` and `Reply` permissions on the Engage Messaging source.

You must also ensure they have the `Initiate discussion` permission if you enabled the outbound messaging option and want to send outbound messages on their behalf through the API.

You must ensure that the Roles of the administrators/managers that will configure the messaging (triggers, items & custom variables) contain the `Manage Engage Messaging configuration` permission.

### Task View (Mandatory for Live Chat mode)

The agents handle the messaging conversation via the task view (a.k.a. the Push mode).

You must ensure that there is a topology. When you create a new one, it will be populated with a set of standard default values.

You must add the Engage Messaging source to the `Realtime` channel.

You must ensure that the role of the agents contains the `View tasks` permission.

### Custom Language

You can add custom translations or modify the existing ones by using this field. To add new languages, use the ISO country code as the main key. See [Customizing the Languages](../customization/#customizing-the-languages) for the JSON format that needs to be used.

View a list of all [translations on many languages](https://docs.google.com/document/d/1ly9bR9q7VSBOx4PdeUpp9HSt5PyxFW3r_T_lPdJOmgc/edit).
