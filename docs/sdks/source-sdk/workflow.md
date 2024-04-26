# Introduction to Channel SDK

The SDK source offers two working types:

1. [Polling](../polling) (not mandatory)
2. [Send API](../send-api) (not mandatory)

If the Polling synchronization fails, you will not be able to send or receive messages.

Send API allows you to send `actions` in realtime. This allows you to respond quickly and you do not have to wait for the polling request to import messages into the RingCentral RingCX Digital console.

Both methods are avalaible and it's possible to use them together, or independently.

## Example
<img class="img-fluid" width="824" src="../../../img/source-sdk-workflow.png">

This example allows communication between different entities. Here, a protocol platform realizes the communication between an external chat solution and the RingCentral RingCX Digital console. Although [Polling](../polling) is in place and active, the platform uses the [Send API](../send-api) to communicate with the platform in real time. This configuration therefore uses the two modes of communication (both polling and real time sending) together.
