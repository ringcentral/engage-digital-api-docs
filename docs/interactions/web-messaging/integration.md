# Integration in the Target Website

The codes listed below have to be inserted in the HTML code of the target website, on each page of the website.

!!! warning
    The chat can only be shown on pages in which the codes listed below are present. Having the codes on only a few pages, if the visitors starts a conversation and browses to a page in which the integration codes is not present, he will not see the chat window until he gets to a page with the codes. The chat session is not lost, it just can’t be shown.

## Minimal Integration

In the Engage Messaging source configuration, get the HTML code from the “Code to include” section. Here is an example:

``` javascript
<script>
    (function(d) {
        var cm = d.createElement('scr' + 'ipt'); cm.type = 'text/javascript'; cm.async = true;
        cm.src = ('https:' == d.location.protocol ? 'https://' : 'http://') + 'the-domain.dimelochat.com/chat the-source-token/loader.js';
        var s = d.getElementsByTagName('scr' + 'ipt')[0]; s.parentNode.insertBefore(cm, s);
    }(document));
</script>
```

This code has to be inserted in the HTML source of the target website. It can be added anywhere in the HTML source.

!!! Required
    Provide the exact code from the source configuration.

## Identity Integration (optional)

We offer a way to register some information about the user. These data are then sent to the server to create and fill an identity.

### Identity Data

The data are given with a JavaScript object (a hash), with the field's name as keys, and the field's value as values.

Example:

``` json
{
    "screenname": "John Doe",
    "avatar_url": "https://my-domain.com/avatars/54579fed6d617868b3010000/john-doe.jpg",
    "firstname": "John",
    "lastname": "Doe",
    "email": "john.doe@my-domain.com",
    "uuid": "54579fed6d617868b3010000",
    "extra_values": {
        "customer_id": "54579fed6d617868b3010000",
        "my_extra_value": "A123456"
    }
}
```

There are predefined fields, such as `firstname`, `lastname`, `screenname`, `company`, `email`, `uuid` or `avatar_url`. Each of them is optional.

For information about UUID please see [User persistent identifier (aka UUID)](../#user-persistent-identifier-aka-uuid).

Other values can freely be registered in the `extra_values` hash with a key, value format. Note that we strongly recommend to register an ID in the extra values for the case of the visitor being a customer (see [related section](../#company-managed-uuid)).

### Non Signed Identity

This is the most common case. The JavaScript object must be given as is to the “_setIdentity” asynchronous method:

``` javascript
<script>
    var _chatq = _chatq || [];
    _chatq.push(["_setIdentity", {
        "screenname": "John Doe",
        "avatar_url": "https://my-domain.com/avatars/54579fed6d617868b3010000/john-doe.jpg",
        "firstname": "John",
        "lastname": "Doe",
        "email": "john.doe@my-domain.com",
        "uuid": "54579fed6d617868b3010000",
        "extra_values": {
            "customer_id": "54579fed6d617868b3010000",
            "my_extra_value": "A123456"
        }
    }]);
</script>
```

This code has to be inserted in the HTML source of the target website. It can be added anywhere in the HTML source. Note that the `var _chatq = _chatq || [];` is mandatory.

!!! Required
    Prepare and provide an example with the desired information for the integration.

### JWT Signed Identity

The server in RingCentral Engage currently supports the following JWT signature algorithms:

* "HS256": HMAC with SHA256
* "RS256": RSA with SHA256

The server in RingCentral Engage currently uses the following JWT extensions:

* “exp” : an expiration timestamp must be provided in the payload (via the “exp” key), and must be set to a future date (for example a date matching “now + 2 hours”). The expected format for this expiration date is a [Unix Timestamp](https://en.wikipedia.org/wiki/Unix_time) (integer). Our configured acceptable leeway is 10s

!!! warning "Required data"
    The **“uuid”** key with a valid value is **mandatory** in the JWT payload.
    When missing, the JWT will be rejected, a default identity will be used, and, if the **“Require JWT signed identities” JWT mode** is set on the community profile, the chat won't open.

#### Generation of the JWT

For the headers, you must have the “typ”, “alg” and “kid” keys. The “kid” contains the `Key ID` of the key configured in the Engage Messaging Community in RingCentral Engage and used for the signature.

``` json
{
    "typ": "JWT",
    "alg": "HS256",
    "kid": "the_key_id_of_the_jwt_key_used"
}
```

For the payload, here is an example. The expiration data must be added to the payload with the “exp” key.

``` json
{
    "exp": 507749400,
    "screenname": "John Doe",
    "avatar_url": "https://my-domain.com/avatars/54579fed6d617868b3010000/john-doe.jpg",
    "firstname": "John",
    "lastname": "Doe",
    "email": "john.doe@my-domain.com",
    "uuid": "54579fed6d617868b3010000",
    "extra_values": {
        "customer_id": "54579fed6d617868b3010000",
        "my_extra_value": "A123456"
    }
}
```

#### Integration

The calculated JWT must be given as a string to the “_setSignedIdentity” asynchronous method:

``` javascript
<script>
    var _chatq = _chatq || [];
    _chatq.push(["_setSignedIdentity", "eyJ0eXAiOiJKV1QiLCJhbGc...jdruLc9qbBWCWsQ"]);
</script>
```

This code has to be inserted in the HTML source of the target website. It can be added anywhere in the HTML source. Note that the `var _chatq = _chatq || [];` is mandatory.

!!! Required
    Create a JWT key in RingCentral Engage and provide both the Key ID and the Secret/Public Key. Prepare and provide an example with the information wanted for the JWT payload and a copy of this section to be given to the customer tech team.

#### Security

The **“Accept JWT signed identities” JWT mode** is an unsecure mode for production if you are using JWT to set your user's data since it also allows non-signed identities. This mode is **not recommended** when you have a JWT integration.

To be able to share the history cross device between web and mobile, the mobile SDK uses a JWT to authenticate the user through its UUID. But then if you allow non-signed identities on the web, there is a security risk since, without a JWT, identity UUID is not guaranteed to come from your system.

## Custom Variables Integration (optional)

We offer a way to register manually some custom variable, for the cases when `In page filling` is not activated. The markup code can be found in each custom variable configuration page. Note that names for custom variables must start with “custom_”.

Here is an example:

``` javascript
<script>
    var _chatq = _chatq || [];
    _chatq.push(["_fillCustomVariable", "custom_cart_amount", 58]);
</script>
```

This code can be inserted in the HTML source of the target website. It can be added anywhere in the HTML source. Or integrators may directly use the javascript code. Note that the `var _chatq = _chatq || [];` is mandatory.

!!! Required
    Provide the code needed for each Custom Variables defined in the configuration that doesn’t have the `In page filling` enabled.

## Fixed Buttons Integration

For each fixed button The code has to be inserted at the precise location where the button will be, in the HTML source of the target website.

Here is an example:

``` html
<div id="dimelo_chat_item_markup_b8d7e7c72b6f86af6e766559"
  class="dimelo_chat_item_markup"></div>
```

!!! Required
    Provide the code needed for each button defined in the configuration that has a “Fixed” positioning.

## Embedded Mode Integration

When the Engage Messaging source is configured in messaging mode (not Live Chat), the embedded mode can be used to integrate the chat in the target website.

Here’s the code allowing to do this:

``` html
<div id="messaging-container"></div>
<script>
    var _chatq = _chatq || [];
    _chatq.push(['_setConfiguration', { 'embedded': true, 'parentNodeId': 'messaging-container' }]);
</script>
```

The `embedded` param must be set to true.

The `parentNodeId` must be the ID of an existing element in the target website. The chat will be embedded in it. This element can be customized as needed to fit in the target website.

For example:

<img class="img-fluid" width="784" src="../../../img/web-messaging-chat-embedded.png">

### Iframe Integration

If adding the Engage Messaging JS in the target website is not an option, an iFrame integration is also possible. Just like with the `parentNodeId`, an `iframe` element can be added in the target website to have the embedded Chat.

The url of the iFrame has this pattern:

``` html
the-domain.dimelochat.com/chat/the-source-token/web?jwt_identity=eyJhbGc...2drA
```

The identity **must** be signed, otherwise the chat won’t start.

This integration has limited interactions and the identity’s data must be passed in the JWT.

The only other command available is the filling of the custom variables. This can be done by adding parameters in the URL. For example:

``` html
the-domain.dimelochat.com/chat/the-source-token/web?jwt_identity=eyJhbGc...2drA&custom_customer_id=1233456789
```

!!! warning
    With the iFrame integration, the domain needs to allow iFrame integration.

!!! warning
    With the embedded mode, the triggers are ignored and the chat is displayed as soon as possible. This means that any option linked to the triggers are unavailable in this mode.
