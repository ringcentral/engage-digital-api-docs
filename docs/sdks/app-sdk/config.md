# Configuration and Development

## Creating a New App SDK Application

By default, App SDK application management is disabled. You’ll have to request your RingCX Digital project manager to grant you the permission to manage them.

In RingCentral's RingCX Digital administration section, you’ll get the **« Dev tools »** menu section and **« Developed applications (SDK) »** menu item in the administration and be able to create new ones.

<img class="img-fluid" width="100%" src="../../../img/app-sdk-config.png">

Click on the **Add** button to start adding your own application.

<img class="img-fluid" width="593" src="../../../img/app-sdk-add-new.png">

An application is defined by the following attributes:

| Attribute | Description|
|-|-|
| `Name` | An unique name that describes the application. This name is mandatory. |
| `Description`​ | A short description of what the application does. |
| `Mode` | The application current mode (by default: development). See available modes below for more details. |
| `Developers` |​ The list of agents that are able to develop this application. 25 developers can be added at maximum. |
| `Callback URL​` | The callback used for OAuth authentication (more details in OAuth authentication section). |
| `Configuration URL` | If present, this URL is used to configure each SDK installation. See installation section for more information. |
| `Source code URL(s)` | The URL(s) of the JavaScript code you wrote that you have to host on a public server. The created code will be automatically imported in all RingCX Digital application pages. |

Available application modes are:

| Mode | Description|
|-|-|
| `Development` |​ In this mode, code is integrated directly, and therefore, updated on each refresh of the RingCentral RingCX Digital page. The code is, however, not executed in a sandbox. JavaScript code is included in page only for users that are in the application's developers list. |
| `Pre-production​` | In this mode, JavaScript code is integrated in a sandbox (only some methods are available) and updated every hour. JavaScript code is included in the page only for users that are in the application's developers list. |
| `Production​`| In this mode, JavaScript code is integrated in a sandbox (only some methods are available) and updated every hour. JavaScript code is included for all RingCentral RingCX Digital agents. |

The best practice is to develop the application in ​*development​* mode. When the application seems to be ready, it should be switched to ​*pre-production​* mode. Then if everything is working correctly, it has to be switched to *production* ​mode. Note that to have live refresh of code, it must be in *development* mode.

## Installation

When created, the application **m​ust**​ be installed in order to make it available on the domain. Installed applications can be found under **« Applications »** in the **« Installed Apps »** menu section.

<img class="img-fluid" width="1041" src="../../../img/app-sdk-install.png">

Click on the **Install** button, choose the application you want to install, and confirm your selection. By default, only *certified* and *local* applications are available in the list. Other applications could be installed if you get its *installation link*.

Once installed, the application can be configured (if the configuration URL has been specified).  Look for the gear icon to configure your application.

<img class="img-fluid" width="940" src="../../../img/installed-app-config.png">

If the configure button is clicked, an iframe overlay is rendered with the application’s configuration URL plus current domain name as parameter. Example if application configuration URL is http://example.org​ and current domain name is g​oogle, ​the iframe URL will be: http://example.org?domain=google.​

The configuration overlay window aims to set domain dependent settings of the application.

!!! info "Note"
    This URL **​must**​ be authenticated, see OAuth section for more information.

## Permissions

An App SDK application could be granted a set of permissions. API calls may require a permission to be set before an API call is allowed. A 403 error will be returned if the application or current user is not granted the proper permission.

By default all permissions are disabled. To enable one or many of them click on the **edit permissions** icon as shown.

<img class="img-fluid" width="1007" src="../../../img/app-sdk-edit-permission.png">

This will bring a pop-up window where you can check the permissions you want your application to have.

<img class="img-fluid" width="645" src="../../../img/app-sdk-permissions.png">

If you change the application's permissions, installations are required to update their permission. Navigating to the installations administration section you will see the following message:

<img class="img-fluid" width="1020" src="../../../img/app-sdk-update-permissions.png">

Upon clicking the update permissions icon, you will need to confirm you are ready to update the permission of the installed application.

<img class="img-fluid" width="584" src="../../../img/app-sdk-confirm-update-permissions.png">

## IFrame Messaging

The RingCX Digital App SDK offers iframe rendering features. The main inconvenience with iframes is the communication with the parent window (RingCX Digital).

The RingCX Digital App SDK offers via ​`window.postMessage​` method the feature to send messages from an iframe to the RingCX Digital window. You will get more details on how `window.postMessage` works in this [technical review](https://developer.mozilla.org/en-US/docs/Web/API/Window/postMessage).

### Sending a Message to RingCX Digital Window

On the iframe side, here is an example of how to send a message to the RingCX Digital window:

``` javascript
    window.parent.postMessage({
      name: "smcc:foo",
      data: { identityId: "4f0aa52d656a3d75867f784b" }
    }, "http://example.digital.ringcentral.com");
```
As you see, the `postMessage` method takes two (mandatory) arguments:

| Argument | Description|
|-|-|
| `object` | A javascript object that describes the message. This object contains the following properties: <ul><li>`name`:​ The name of the event (mandatory). This is a message identifier that describes the action performed on RingCX Digital side. The name has to be prefixed by “smcc:”, like “smcc:baz”</li><li>`data​`: An optional hash of arguments that message can accept.</li></ul> |
| `url` | The URL of the RingCX Digital application. For example: `"​​http://foo.digital.ringcentral.com​"`​, where `f​oo​` is your RingCX Digital domain’s name.

To receive messages in RingCX Digital, first you have to whitelist the iframe sender origin:

``` javascript
    SMCC.Window.acceptPostMessageOrigin("https://example.com");
    SMCC.Window.onPostMessage("foo", function(data, origin) {
      console.log("New data:", data, "from:", origin);
    });
```

## OAuth Authentication

RingCX Digital App SDK provides many ways to render iframes. But, most of the time those iframes must be authenticated. The RingCX Digital’s interface offers a standard [​OAuth 2.0](https://oauth.net/2/)​ provider to request access tokens and get RingCX Digital's current user information.

### Prerequisites

First of all, you must get your application’s `k​ey` ​and `s​ecret`. Navigate to **Dev tools**, and then **Developed applications (SDK)** to get them:

<img class="img-fluid" width="377" src="../../../img/app-sdk-key-secret.png">

The ​`key` ​and ​`secret` ​are mandatory to make OAuth requests. In order to be secure, iframe authentication callbacks should be in HTTPS. Note that all OAuth requests **m​ust**​ be in HTTPS.

You must also configure a callback URL for the application to a valid one. This will be explained later.

### Authentication Process

If you are not familiar with the OAuth 2.0 protocol, here is how authentication works:

1. User agent is not authenticated on application, a redirection is made on the RingCX Digital OAuth provider with the application key and a redirect URL.
2. If the application key is valid, user agent is redirected to the provided redirect URL with authorization code as a parameter. Otherwise the user agent is redirected to the redirect URL with an error parameter.
3. Then, the application must make a `POST` request with returned code and application secret to get an OAuth access token.
4. with this access token, the iframe can get the current user infos via REST `/1.0/users/me​` API method.

### Request an OAuth Authorization Code

An OAuth authorization code is required before trying to get an access token. User agent **must**​ be redirected to the RingCX Digital authorization URL to get it.

#### URL

``` http
    https://<domain-name>.digital.ringcentral.com/oauth/authorize
```

| Parameter | Required? | Description |
|-|-|-|
| `client_id` | **Required** | This parameter **must** be the application key |
| `redirect_uri` | **Required** | This parameter must be a valid HTTP(s) URI. Note that the scheme, host, port and path must be equal to application’s callback URL (see configuration section). |
| `response_type` | **Required** | This OAuth standard parameter determines the kind of response. Its value **must** be "code". |

#### Response

If every parameter is correct, user agent is redirected to redirect URI with a ​"code"​ parameter:

``` http
    https://redirct-uri.com?code=xxx
```

If authorization fails (user can’t authenticate or does not accept application), user agent is redirected to redirect URI with an ​`error` ​parameter:

``` http
    https://redirect-uri.com?error=authentication_failed
```

!!! info "Note"
    * The authorization code is unique and random. It also expires quickly.
    * Other redirect URI query parameters are preserved.
    * If user agent isn’t authenticated when an authorization code is requested, the login form is rendered and authorization is performed when logged in.

If one of the provided parameters is invalid or missing, a `400 HTTP` error is returned with a JSON error as body. Here are all available error responses:

| Response Body | Description |
|-|-|
| <pre lang="json">{<br> "error": "missing_parameter",<br> "message": 'Missing required parameter: "client_id"',<br> "status": 400<br>}</pre> | `client_id` parameter has been omitted. |
| <pre lang="json">{<br> "error": "invalid_client_id",<br> "message": "Invalid client ID",<br> "status": 400<br>}</pre> | `client_id` parameter has been provided but does not match an application. |
| <pre lang="json">{<br> "error": "missing_parameter",<br> "message": 'Missing required parameter: "redirect_uri"',<br> "status": 400<br>}</pre> | `redirect_uri` parameter has been omitted. |
| <pre lang="json">{<br> "error": "invalid_redirect_uri",<br> "message": "Invalid redirect URI",<br> "status": 400<br>}</pre> | `redirect_uri` parameter has been provided but URI is malformed (invalid scheme, etc.). |
| <pre lang="json">{<br> "error": "invalid_redirect_uri",<br> "message": "Redirect URI scheme does not match callback URL scheme",<br> "status": 400<br>}</pre> | `redirect_uri` parameter has been provided but URI scheme does not match callback URL scheme on application. |
| <pre lang="json">{<br> "error": "invalid_redirect_uri",<br> "message": "Redirect URI port does not match callback URL host",<br> "status": 400<br>}</pre> | `redirect_uri` parameter has been provided but URI host does not match callback URL host on application. |
| <pre lang="json">{<br> "error": "invalid_redirect_uri",<br> "message": "Redirect URI port does not match callback URL port",<br> "status": 400<br>}</pre> | `redirect_uri` parameter has been provided but URI port does not match callback URL port on application. |
| <pre lang="json">{<br> "error": "invalid_redirect_uri",<br> "message": "Redirect URI path does not match callback URL path",<br> "status": 400<br>}</pre> | `redirect_uri` parameter has been provided but URI path does not match callback URL path on application. |
| <pre lang="json">{<br> "error": "missing_parameter",<br> "message": 'Missing required parameter: "response_type"',<br> "status": 400<br>}</pre> | `response_type` parameter has been omitted. |
| <pre lang="json">{<br> "error": "invalid_response_type",<br> "message": 'Invalid response type, it must be "code"',<br> "status": 400<br>}</pre> | `response_type` parameter is present but its value isn’t ​`code`. |

### Request an OAuth Access Token

To get an access token you must first get an authorization code (see above). Then when the user agent is redirected to the application (redirect URI), a request must be made **by the application** to get an access token.

#### Request

```bash
POST https://<domain-name>.digital.ringcentral.com/oauth/token
```

#### Response

If all provided parameters are valid, a `200 HTTP` response is returned and the access token is rendered in the response body as JSON:

```json
{
  "access_token": "06eaea7ac0e0b4b7ac622e81"
}
```

Otherwise, if one of the provided parameters is invalid or missing, a `400 HTTP` error is returned with a JSON error as body. Here are all available error responses:

| Response Body | Description |
|-|-|
| <pre lang="json">{<br> "error": "missing_parameter",<br> "message": 'Missing required parameter: "client_id"',<br> "status": 400<br>}</pre> | `client_id` parameter has been omitted. |
| <pre lang="json">{<br> "error": "invalid_client_id",<br> "message": "Invalid client ID",<br> "status": 400<br>}</pre> | `client_id` parameter has been provided but does not match an application. |
| <pre lang="json">{<br> "error": "missing_parameter",<br> "message": 'Missing required parameter: "client_secret"',<br> "status": 400<br>}</pre> | `client_secret` parameter has been omitted. |
| <pre lang="json">{<br> "error": "invalid_client_secret",<br> "message": "Invalid client secret",<br> "status": 400<br>}</pre> | `client_secret` parameter has been provided but does not match an application secret. |
| <pre lang="json">{<br> "error": "missing_parameter",<br> "message": 'Missing required parameter: "code"',<br> "status": 400<br>}</pre> | `code` parameter has been omitted. |
| <pre lang="json">{<br> "error": "invalid_code",<br> "message": "Invalid code",<br> "status": 400<br>}</pre> | `code` parameter has been provided but does not match a valid OAuth authorization code. |
| <pre lang="json">{<br> "error": "expired_code",<br> "message": "Expired code",<br> "status": 400<br>}</pre> | `code` parameter is correct but the OAuth authorization code is expired. |
| <pre lang="json">{<br> "error": "missing_parameter",<br> "message": 'Missing required parameter: "grant_type"',<br> "status": 400<br>}</pre> | `grant_type` parameter has been omitted. |
| <pre lang="json">{<br> "error": "invalid_grant_type",<br> "message": 'Invalid grant type, it must be "authorization_code"',<br> "status": 400<br>}</pre> | `grant_type` parameter is present but its value isn't `authorization_code`. |

### Getting Access Token User Information

When getting an OAuth access token, you’ll be able to get information for this token. An access token is associated to a RingCX Digital agent. To get the agent’s ID, a REST API request (`​/1.0/users/me`​) must be made by an application.​ P​lease refer to RingCX Digital REST API documentation for more information.
