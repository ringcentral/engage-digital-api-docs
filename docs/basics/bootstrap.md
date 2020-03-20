# Bootstrapping a Developer System

To do development for Engage Digital, it is useful to set up an agent which can receive and send messages and a source channel to receive messages from and post messages to.

This tutorial covers setting up a basic email source and and using Engage Digital to receive and send email, which you can develop against.

**Configuration**

Setting up the agent and source channel allows Engage Digital to receive and respond to messages.

1. Enable the "Create a source" Permission
1. Create a Community
1. Create an Email Source

**Handle Test Message**

Sending and responding to a test message allows your app to respond to message flows.

1. Send an Inbound Email
1. Respond to the Inbound Email

## Configuration

### Enable the "Create a source" Permission

Before you can create a source, make sure you have the "Create a source" Permission.

1. Check the role for your user by goiong to the "Admin" > "AGENTS MANAGEMENT" > "Agents" page and look up the Role namefor your agent.
1. Navigate to "Admin" > "AGENTS MANAGMENT" > "Roles" and click on the edit Pencil icon for your user's Role Name.
1. Under the "ADMIN - MANAGE MESSAGES" section, verify the "Create a source" permission is checked. If it is not, add it and click "Save".

### Creating a Community

Before you can create a source, a Community is needed to attach the Source channel to. If you don't already have a Community, use the following instructions to create one.

1. In the Enage Digital Console, go to "Admin" > "MESSAGES MANAGEMENT" > "Communties".
1. In the upper right corner, click "+" and then "Email" to bring up the "Create a community page".
1. Type a name into the "Name" field such as "Email".
1. Leave "Active" checked and click "Save".

### Create an Email Source

To use Engage Digital, a message needs to b configured to send and receive messages. Engage Digital Supports many different Source messaging channels, including custom ones.

To bootstrap your system we will describe how to set up a simple Email Source as it is one of the simplest to set up. This will enable you to send and receive messages. See thee Engage Digital documentation for setting up other message sources.

In your Engage Digital console, perform the following steps:

1. Navigate to: "Admin" > "MESSAGES MANAGEMENT" > "Sources"
1. Click the "+" in the upper right corner and then select "Email" to bring up the "Create a source" page.
1. Under "GENERAL SETTINGS" > "Name", enter a name for this source. A common approach is to use the email address you select in "POSTMARK SETTINGS" > "Email Address" mentioned below.
1. Under "GENERAL SETTINGS" > "Community", select the desired community, which will be the one you created in thee previous section.
1. Under "POSTMARK SETTINGS" > "Email Address". You can use the default email address. For quick setup, either use the default email address or just use the SMTP username and use the existing SMTP hostname. For example, you can use an email like the following: "mycompany@email.us1.digital.ringcentral.com". It is possible to use a custom email domain but that is outside the scope of this tutorial. See the Engage Digital documentation for more.
1. Click "Save" at the bottom of the page.

To verify the settings, click on the source name in the esource list and verify that the "GENERAL SETTINGS" > "Active" checkbox is clicked, and that the "POSTMARK SETTINGS" > "Enable" checkbox is clicked.

## Handle Test Message

### Send an Inbound Email

Use any email client to send an email to the one you chose above. For example:

`mycompany@email.us1.digital.ringcentral.com`

### Respond to the Inbound Email

