# About Webview SDK

This SDK is meant to be used on pages linked to by messages/items with a "webview" target.

This kind of target is useful to create specific interactions that are not provided by channels.

The Webview SDK provides a generic way to interact with supported channels that can open your page in a webview which allows you to use the same target on different channels without having to write multiple integrations.

## Supported channels

See the [webview documentation](../webview#channels-supporting-this-feature).

## Integration

You can load the SDK by adding the following `script` tag to your target page.
You'll need to replace `<YOUR_DOMAIN>` with your ED domain's name.
You'll need to replace `<YOUR_ED_HOST>` with your ED hostname, valid values include `engagement.dimelo.com` and `digital.ringcentral.com`.

```html
<script>
  (function (d, s, url, id) {
    d.addEventListener('DOMContentLoaded', function() {
      var js, fjs = d.getElementsByTagName(s)[0];

      if (d.getElementById(id)) return;

      js = d.createElement(s);
      js.id = id;
      js.src = url;
      js.async = true;
      fjs.parentNode.insertBefore(js, fjs);
    });
  })(document, 'scr' + 'ipt', 'https://<YOUR_DOMAIN>.<YOUR_ED_HOST>/webview/sdk.js', 'ed_webview_sdk_loader');
</script>
```

Once loaded the SDK will attempt to call the `rcedWebviewSdkLoaded` function, if defined. It will be called with the SDK instance as a parameter. You can also access the instance through `RCED.WebviewSDK.instance`

Here's a basic `rcedWebviewSdkLoaded` example:
```javascript
window.rcedWebviewSdkLoaded = function(sdk) {
  console.log(sdk === RCED.WebviewSDK.instance) // true
  // use sdk in your page
};
```

## Available methods

The SDK instance responds to the following methods:
### supportFeature

This method will tell you if a feature is supported by the opening channel.

This method takes a string corresponding to a feature name as a parameter and returns `true` if the channel that launched the webview supports it or `false` otherwise.

Even if a feature is not supported, an empty method with its name is still available.

Here's a basic example:
```javascript
window.rcedWebviewSdkLoaded = function(sdk) {
  if (sdk.supportFeature('close')) {
    // logic that uses the feature
  } else {
    sdk.close(); // does not crash, here `close` definition is equivalent to `function() {}`
  }
}
```

### close

This method will close the webview if the channel that launched the webview supports it or do nothing otherwise.

This method doesn't take any parameter and doesn't return anything.

Here's a basic example:
```javascript
window.rcedWebviewSdkLoaded = function(sdk) {
  sdk.close();
}
```

## Supported features by channels

Since not every channel supports every feature, you'll find support info in the following table:

* **<span style="color:green">Yes</span>**: Supported
* **<span style="color:orange">Not yet</span>**: Supported by channel but not by SDK yet
* N/A: Not supported by channel

| | **Engage Messaging Web** | **Engage Messaging iOS** | **Engage Messaging Android** | **Unsupported channels** |
|-|-|-|-|-|
| [**close**](#close) | <span style="color: green">Yes</span> | <span style="color: green">Yes</span> | <span style="color: orange">Not yet</span> | N/A |
| [**Automatic title update**](#engage-messaging-webios) | <span style="color: green">Yes</span> | <span style="color: green">Yes</span> | <span style="color: orange">Not yet</span> | N/A |

## Integrations specific behavior

Some integrations have extra behavior specific to their related channel. Those are the following.

### Engage Messaging Web/iOS

When opening your webview page in Engage Messaging web chat or in Engage Messaging iOS, the title element's text will become the chat webview header's text, and when you change your page's title element's `textContent`, it will change the chat's webview header text.
