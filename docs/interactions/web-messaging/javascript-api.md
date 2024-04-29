# Javascript API

RingCentral RingCX Messaging offers an asynchronous JavaScript API for the customer to be able to make any integration on the target website.

## Events

### List of Events

The chat manager supports the following publicly subscribable events:

|**Event name**|**Data**|**Details**|
| - | - | - |
|button_shown|trigger, mode & item|Fired when a button is shown|
|chat_engaged|trigger|Fired when the visitor engages the conversation, i.e. sends a first message|
|chat_shown|trigger|Fired when a chat client is created and shown|
|chat_closed||Fired when a chat client is closed and destroyed (when the visitor definitely leaves the chat client)|
|invitation_shown|trigger, mode & item|Fired when an invitation is shown|
|item_shown|trigger, mode & item|Fired when a button or an invitation is shown|
|message_received|trigger|Fired each time the visitor receives a message|
|message_sent|trigger|Fired each time the visitor sends a message|
|started||Fired when the chat manager is loaded and started|
|trigger_activated|trigger & mode|Fired when a trigger meets all its condition for activation and tries to take its action|
|trigger_executed|trigger & mode|Fired when a trigger was activated and then is able to take its action, e.g. the item to be shown is found on the page.|

For most cases, a “trigger_activated” event is followed right after by a “trigger_executed” event. The sole case for which trigger execution can be halted after the activation is for triggers that have to show a fixed button. If the markup for placing the fixed button is not found on the page, the trigger can’t be executed.

#### Event's Data

The data column tells what specific information an event handler would receive when the event is triggered.

The “mode” data gives . Example of data for “mode”:

``` json
{
    "mode": "online"
}
```

Example of data for “trigger”:

``` json
{
    "trigger_id": "54be84696d61786d2f160000",
    "trigger_label": "Home Button",
    "trigger_human_id": "Home Button #54be84696d61786d2f160000",
    "trigger_continuation": false
}
```

The “trigger_continuation” key is true if the event is the result of a continuation, false otherwise (for example, a chat client that reopens himself after a page change would fire a “chat_shown” event with “trigger_continuation” being true).

The “trigger_id” key may be null for the chat internal continuation triggers. 

Example of data for “item”:

``` json
{
    "item_id": "55a7d1986d6178585b000002",
    "item_label": "Default floating button - dark",
    "item_human_id": "Default floating button - dark #55a7d1986d6178585b000002",
    "item_type": "button"
}
```

Data always contains an “event” key with the name of the event as value.

### Registering an Event Handler

To register for an event, you must give a callback function to the “_onEvent” asynchronous method:

``` javascript
<script>
    var _chatq = _chatq || [];
    _chatq.push(['_onEvent', 'chat_engaged', function (data) {
        // Do something
    }]);
</script>
```

The callback function may accept one argument, “data”, which will be a plain Javascript object.

### Usage

#### How to Track a Chat Engagement Funnel

You may want to track the behavior of the visitors regarding the usage of the chat.

Here is an example of an events chain that is triggered when a visitor engages a conversation and does receive help:

``` html
trigger_executed → chat_engaged → message_received
```

#### Common Question That Can Be Answered

We look here for the events fired during the whole visit.

* **Q1:** *Was the visitor offered any assistance by chat?* 

At least one “trigger_executed” event was fired with a “trigger_id” data being not null, and a “mode” data being “online”.

* **Q2:** *Has the visitor tried to get assistance by chat?*

At least one “chat_engaged” event was fired.

* **Q3:** *Has the visitor refused assistance by chat, while having been offered one?*

At least one “trigger_executed” event was fired (see **Q1** for the conditions) AND no “chat_engaged” event was fired.

* **Q4:** *Has the visitor received any assistance by chat?*

At least one “message_received” event was fired.

* **Q5:** *Has the service failed to provide assistance by chat, while the visitor used the chat?*

At least one “chat_engaged” event was fired AND no “message_received” event was fired.

## Actions

### Add Message

The add message action allows you to create a customer or agent message in the chat from your JavaScript code. These messages will be sent to the server as automated messages. Example:

``` javascript
<script>
    var _chatq = _chatq || [];
    _chatq.push(["_addMessage", { body: "Hello, how can I help?", agent: true }])
</script>
```

This API can be used for example to make buttons/links inserting custom data into the chat (ex: product details) or to pre-fill the chat with initial data like a previous conversion with a virtual agent or responses from a qualification form/wizard.

!!! warning
    The `addMessage` action only works when the chat is already open, if you intend to use it to pre-fill the chat we recommend using it in the "chat_shown" event as shown below. This also allows you to check the trigger used to open the chat and change the behavior depending on which one it is.

``` javascript
<script>
    var _chatq = _chatq || [];
    _chatq.push(['_onEvent', 'chat_shown', function (trigger) {
        if (trigger.trigger_continuation) return; // Avoids executing on each page refresh
        _chatq.push(["_addMessage", { body: "bot message #1", agent: true }])
        _chatq.push(["_addMessage", { body: "customer message" }])
        _chatq.push(["_addMessage", { body: "bot message #2", agent: true }])
    }]);
</script>
```

This code should be placed alongside the chat loading script to make sure the chat_show event is registered early enough.

Here are the available options to the `addMessage` action:

|**Param**|**Type**|**Default**|**Details**|
| - | - | - | - |
|body|String|""|The body of the message to add|
|agent|Boolean|false|Whether the message will appear to be from the agent or from the customer|
|engage|Boolean|false|For customer messages only, determines whether the chat should be considered "engaged" (and thus creates an interaction in RingCentral RingCX) or wait for the next real customer message|

### Page Change

The page change action allows you to inform the server that the customer has changed the page. This is typically useful for **single-page applications**. Example:

``` javascript
<script>
    var _chatq = _chatq || [];
    _chatq.push(["_pageChange"]);
</script>
```

The page change action has the following effects:

* The triggers are re-initialized. This means that if a trigger has already been executed, it will be able to be re-executed after the page change.
* The triggers can use the new page information to check if it matches or not.
* If a trigger matches, there is a new check on the server to see if there is an agent available.
* The new URL is visible in RingCentral RingCX.

!!! warning
    If a trigger is executed and thus shows an element (button, invitation or Chat), this element will stay on the screen after a page change, even if there is no trigger that matches the new page.
