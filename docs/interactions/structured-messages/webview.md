# Webview

The webview is a convenient way of opening a custom made webpage. It can be chosen by setting the optional `target` parameter to `webview` when sending a Structured Message containing a link element, such as [Rich Link](../rich-link), [Carousel](../carousel), or [Template](../template).

The webview is a small window that sits atop the chat. It can be customized in size through the `webview_height` parameter.

## Use Case

The webview is not just a simple window opening, it can be used to display a custom survey, to select a room or a seat on a train or even a more attractive way to select a subscription. For each action taken on the page, a custom response can be sent for the agent side, that can confirm a seat selection or express gratitude to a client indicating satisfaction.

## Recommended Usage

It should be noted that no default picker or page is provided and this feature is intended to be customizable, not out of the box. No default response after the customer has finished taking actions on the page is provided either.

What is recommended is setting the conversation's `data-id` inside of the url query string in order to provide a custom page taylored to the customer (e.g., displaying a confirmation instead of a choice the second time the customer clicks on it), avoiding him doing several times the procedure.

You can also use this `id` to send a `post` request to Engage Digital API to send a message to the customer after he has taken all actions on the page to confirm or give a summary of his choices. Passing any parameter the client has already given the agent through the query string is possible, such as a ticket number or an email.

## Integration

The page on the webview should be hosted on your side, and setup to accept display on other domains. Check [Webview SDK](../../structured-messages/webview-sdk) for precisions on the integration.

## Channels supporting this feature

* **<span style="color:green">Yes</span>**: Supported
* **<span style="color:orange">Not yet</span>**: Supported by channel but not by Engage Digital yet
* N/A: Not supported by channel

| | **webview** | **webview_height** |
|-|-|-|
| **Apple Business Chat** | N/A | N/A |
| **Engage Android Messaging** | **<span style="color:orange">Not yet</span>** | **<span style="color:orange">Not yet</span>** |
| **Engage iOS Messaging** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** |
| **Engage Web Messaging** | **<span style="color:green">Yes</span>** | **<span style="color:green">Yes</span>** |
| **Facebook Messenger** | **<span style="color:orange">Not yet</span>** | **<span style="color:orange">Not yet</span>** |
| **Google Business Messages** | N/A | N/A |
| **Channel SDK** | **<span style="color:orange">Not yet</span>** | **<span style="color:orange">Not yet</span>** |
| **Twitter Direct Messages** | N/A | N/A |
| **Viber** | N/A | N/A |
| **WhatsApp** | N/A | N/A |

## Engage Messaging Web specific customization

<img class="img-fluid" width="419" src="../../../img/web-messaging-webview-portrait.png">

The webview colors can be changed with the `$webview-header-background-color`, and `$webview-header-text-color`. See the [Customization](../../web-messaging/customization) page for details.

Engage Messaging is one of the sources allowing you to set a smaller window to open. To so do, set the `webview_height` parameter to fit your need, you can see the design for `tall` and `compact` webviews in the following visuals.

<img class="img-fluid" width="419" src="../../../img/web-messaging-webview-tall.png">
<img class="img-fluid" width="419" src="../../../img/web-messaging-webview-compact.png">
