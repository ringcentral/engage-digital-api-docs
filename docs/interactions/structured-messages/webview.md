# Webview

The webview is a convenient way of opening a custom made webpage. It can be chosen by setting the optional target field to `webview` when sending a Structured Message containing a link element, such as [Rich Link](../rich-link), [Carousel](../carousel), or [Template](../template).

The webview is a small window that sits atop the chat.

# Design

<img class="img-fluid" width="419" src="../../../img/web-messaging-webview-portrait.png">

The webview colors can be changed with the `$webview-header-background-color`, and `$webview-header-text-color`. See the [Customization](./customization) page for details.

If you need a smaller window to open, you can also set the `webview_height` parameter to fit your need, you can see the design for `tall` and `compact` webviews in the following visuals.

<img class="img-fluid" width="419" src="../../../img/web-messaging-webview-tall.png">
<img class="img-fluid" width="419" src="../../../img/web-messaging-webview-compact.png">



# Use case

The webview is not just a simple window opening, it can be used to display a custom survey, to select a room or a seat on a train or even a more attractive way to select a subscription. For each action taken on the page, a custom response can be sent for the agent side, that can confirm a seat selection or express gratitude to a client indicating satisfaction.

# Integration

The page on the webview should be hosted on your side, and setup to accept display on other domains. Check  [Webview SDK](../../structured-messages/webview-sdk) for precisions on the integration.
