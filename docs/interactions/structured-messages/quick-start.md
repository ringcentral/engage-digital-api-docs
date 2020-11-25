style: quick-start
no_breadcrumb:true

# Structured Message Quick Start

Let's get started building a structured message. A structured message consists of elements that enhance the customer experience. In this Quick Start, we will walk you through creating a structured message within a matter of minutes. Let's get started.

## Configure an Engage Messaging Channel
In this quick start, we begin with the Engage Messaging channel. To create the Engage Messaging channel, login as and "Admin" to your Engage portal.

<img class="img-fluid" width="540" src="../../../img/structured-messages-channel-tab.png">

On the top right, is a button to create a new channel. Click this and choose "Engage Messaging" for our example.

<img class="img-fluid" width="280" src="../../../img/structured-messages-new-channel-button.png">

Give this channel a name. For our example here, we'll call it "Engage Messaging Example".  While here, let's enable web messaging by checking the "Enable Web Messaging" box.

<img class="img-fluid" width="800" src="../../../img/structured-messages-create-channel.png">

## Set Agent Permissions to the Channel

For agents to interact with messages from a channel, make sure to set the proper permissions for that agent. In this instance, you are making sure your agent has read and reply permissions. Just check both boxes to make sure your agent can receive and respond to messages.

<img class="img-fluid" width="730" src="../../../img/structured-messages-edit-permissions.png">

## Configure Channel Mode

There are two types of channel modes: asynchronus and realtime. For our testing purposes, we want realtime, but for some interactions, the channel mode could be asynchronus (when responding to email for example). Find the channel mode settings under Routing, and then select Channel Modes to configure.

<img class="img-fluid" width="540" src="../../../img/structured-messages-channel-mode.png">

Now let's add our Engage Messaging Example to the realtime channel. First click the pencil icon to edit the realtime channel.

<img class="img-fluid" width="800" src="../../../img/structured-messages-channel-modes-edit.png">

Then add our Engage Messaging Example in the Channels field by selecting it from the dropdown.

<img class="img-fluid" width="800" src="../../../img/structured-messages-channel-modes-realtime.png">

Now let's go back to our Engage Messaging Example channel listed in the channels section. Click the pencil icon to edit this channel and look for the JavaScript code under "Chat connection", in the "Code to include" field. We will use this code in the next section to embed a messaging window on a web page.

## Create an Engage Chat Window on a Web Page

Now let's create an HTML web page with the JavaScript code to include. The code should look something like the following.

### Code to include
```JavaScript
<script>
  (function(d) {
    var cm = d.createElement('scr' + 'ipt'); cm.type = 'text/javascript'; cm.async = true;
    cm.src = 'https://example.chat.digital.ringcentral.com/chat/8e5321aaeb72373cde85a9bb/loader.js';
    var s = d.getElementsByTagName('scr' + 'ipt')[0]; s.parentNode.insertBefore(cm, s);
  }(document));
</script>
```

We'll want to embed this code segment into an HTML page so we can start a Engage Messaging chat window. Here's some sample code to create your own web page with.

### Sample HTML
```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <title>Engage Messaging</title>

        <!–– add for styles ––>
        <style>
          #messaging-container {
            width: 80%;
            height: 500px;
            margin: 20px auto;
          }
        </style>

        <script>
            //script from the source
            (function(d) {
              var cm = d.createElement('scr' + 'ipt'); cm.type = 'text/javascript'; cm.async = true;
              cm.src = 'https://example.chat.digital.ringcentral.com/chat/8e5321aaeb72373cde85a9bb/loader.js';
              var s = d.getElementsByTagName('scr' + 'ipt')[0]; s.parentNode.insertBefore(cm, s);
            }(document));
        </script>
    </head>
    <body>
        <h1>Embedded Mode</h1>
        <p>web page</p>
        <!–– add container ––>
        <div id="messaging-container"></div>
    </body>
</html>
```

With this web page, you'll find the chat window hidden in the bottom right of the page as a blue circle with a couple of bubbles in it. Click on this and you can begin typing messages to your agent. To begin the conversation, start with a simple "Hi". That message will be routed to the agent so make sure you click on "Agent" and then "Routing mode" to see the message.

### Customer Chat
<img class="img-fluid" width="500" src="../../../img/structured-messages-customer-chat.png">

### Agent View
<img class="img-fluid" width="320" src="../../../img/structured-messages-agent-view.png">

## Sending a Structured Message

Here's the point you've been building up to, sending a structured message using an API call. You'll need two things to send a structured message:

* API Access Token
* Reply to ID

### Obtain Access Token

The first thing you need to do is obtain an API Access Token if you do not already have one. Access tokens can be created and/or accessed via the "API access token" area in the "Dev tools" area of your Engage portal under the "Admin" tab.

??? tip "How to generate an Engage API access token"

    1. Login to your Engage portal and click on the "Admin" menu located in the top, horizontal menu.

    2. Select "Dev tools" towards the bottom of the left hand menu.

    3. Select "API access tokens" from the drawer that pops out.

    3. You should see a list of access tokens if any have been provisioned. Select the token, or click the "Add" button to create a new one.

    4. Finally, enter a label/description for your token, and select an Agent on which the token will act on behalf of. Make sure the token is "enabled" and click "Save."

          ![API access token](../../../img/api-token-engage.png)

Make note of the access token generated as you will need it later.

### Find the Reply ID

To find the reply ID, you need to do a little inspect and execute some code. In this example we will inspect and open the console in Google Chrome. Paste the following code in the console window and hit "enter".

```javascript
document.querySelectorAll(".content.chat-content")[0].getAttribute("data-id").split(":")[1]
```

<img class="img-fluid" width="800" src="../../../img/agent-inspect-console.png">

### Post Structured Message

Now post the following code to create the structured message reply.  Note that the {access_token} and {reply_id} are the two elements you just found.

```bash
curl -X POST "https://[YOUR DOMAIN].api.digital.ringcentral.com/1.0/contents" -H "Authorization: Bearer {access_token}"
```

```json
{
	"in_reply_to_id": {reply_id},
	"body": "Hi, what do you want to buy?",
	"structured_content": {
		"center_items": false,
		"disable_text_input": false,
		"type": "select",
		"items": [
			{ "title": "Laptop" }, { "title": "Desktop PC" }, { "title": "Mobile phone" }, { "title": "Webcamera" }, { "title": "Notebook" }, { "title": "Pocket PC" }, { "title": "iWatch" }, { "title": "Other watch" }
		]
	}
}
```

Your output should look like the following.

<img class="img-fluid" width="500" src="../../../img/structured-messages-select-engage.png">

## Need Help?

Having difficulty? Feeling frustrated? Receiving an error you don't understand? Our community is here to help and may already have found an answer. Search our community forums, and if you don't find an answer please ask!

<a target="_new" href="https://forums.developers.ringcentral.com/search.html?c=72&includeChildren=true&f=&type=question+OR+kbentry+OR+topic&redirect=search%2Fsearch&sort=newest&q=interactions">Search the forums &raquo;</a>

## What's Next?

When you have successfully made your first API call, it is time to take your next step towards building a more robust Engage application.

<a class="btn btn-success btn-lg" href="https://developers.ringcentral.com/engage/api-reference/">Explore the REST API &raquo;</a>
