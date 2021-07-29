# About Webview SDK

This SDK is meant to be used on pages linked to by items with a webview target.

## Integration

You can load the SDK by adding the folling `script` tag to the page.

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
  })(document, 'scr' + 'ipt', 'https://<YOUR_DOMAIN>.engagement.dimelo.com/webview/sdk.js', 'ed_webview_sdk_loader');
</script>
```

Once loaded the SDK will attempt to call the `rcedWebviewSdkLoaded` function if defined. It will be called with the SDK instance

```javascript
window.rcedWebviewSdkLoaded = function(sdk) {
  // do something with sdk
};
```

## Available methods

The SDK instance responds to the following methods:
### supportFeature

This method takes a string corresponding to a feature name as a parameter and returns `true` if the channel that launched the webview supports it or `false` otherwise.

Even if a feature is not supported, an empty method with its name is still available.

Usage:
```javascript
function rcedWebviewSdkLoaded(sdk) {
  if (sdk.supportFeature('close')) {
    // logic that uses the feature
  }
}
```

### close

This method will close the webview if the channel that launched the webview supports it or do nothing otherwise.

This method doesn't return anything.

Usage:
```javascript
function rcedWebviewSdkLoaded(sdk) {
  sdk.close();
}
```