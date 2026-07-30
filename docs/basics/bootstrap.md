# Bootstrapping a Developer System

To develop with RingCX Digital, set up an agent that can receive and send messages and a source that can receive and publish messages.

This tutorial covers setting up a basic email source and using RingCX Digital to receive and send email.

**Configuration**

Setting up the agent and source allows RingCX Digital to receive and respond to messages.

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

RingCX Digital supports many messaging channels, including custom channels.

The following steps describe how to configure an email source so that you can send and receive messages. See the RingCX Digital documentation for information about configuring other source types.

In your RingCX Digital console, perform the following steps:

1. Navigate to: "Admin" > "MESSAGES MANAGEMENT" > "Sources"
1. Click the "+" in the upper right corner and then select "Email" to bring up the "Create a source" page.
1. Under "GENERAL SETTINGS" > "Name", enter a name for this source. A common approach is to use the email address you select in "POSTMARK SETTINGS" > "Email Address" mentioned below.
1. Under "GENERAL SETTINGS" > "Community", select the desired community, which will be the one you created in thee previous section.
1. Under "POSTMARK SETTINGS" > "Email Address," use the default email address or the SMTP username and existing SMTP hostname. For example: `mycompany@email.us1.digital.ringcentral.com`. You can also configure a custom email domain.
1. Click "Save" at the bottom of the page.

To verify the settings, click on the source name in the esource list and verify that the "GENERAL SETTINGS" > "Active" checkbox is clicked, and that the "POSTMARK SETTINGS" > "Enable" checkbox is clicked.

## Handle Test Message

### Send an Inbound Email

Use any email client to send an email to the one you chose above. For example:

`mycompany@email.us1.digital.ringcentral.com`

### Respond to the Inbound Email

1. Navigate to "FOLDERS" > "New Messages" to view inbound messages. Select the message from the queue you wish to respond to and click "Engage".
1. Click "Engage" to move the message to the "Global Inbox".
1. Click the "Global Index" folder, select the message, and click "Reply" to bring up the Email Reply window. Type a reply and click "Send".
