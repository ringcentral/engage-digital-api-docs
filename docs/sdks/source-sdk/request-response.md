# Introduction to Request/Response

From a technical standpoint, an SDK implementation is an HTTP endpoint that accepts POST requests and responds in a specified format.

The SDK communication consists of JSON strings sent in the request and response body and a hex encoded [HMAC SHA-512](https://en.wikipedia.org/wiki/HMAC) signature of this body sent as a header.

Please bear in mind that we encode and process content using only UTF-8. Also, all dates and timestamps are expressed and parsed using the [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.

## Request Format

All our requests are POST requests whose body will consist of an JSON object specifing an action, optionally some parameters for that action and the time when we made the request. Here is an example of the aforementioned object:

```json
{
  "action": "messages.list",
  "params": {
    "since_id": "2523423"
  },
  "time": "2012-10-01T17:18:40Z"
}
```

We will also send a specific `X-SMCCSDK-SIGNATURE` header containing the hex encoded [HMAC SHA-512](https://en.wikipedia.org/wiki/HMAC) signature for the request body. You will need to verify this signature and discard any requests that fail this validation. [See below](#complete-example) for a complete example of handling a request and sending the response.

In case you are unable to authenticate request via `X-SMCCSDK-SIGNATURE` header, you can authenticate the request via a `signature` query parameter containing the hex encoded [HMAC SHA-512](https://en.wikipedia.org/wiki/HMAC) (we recommend you to use this method only if the first one does not work).

Please bear in mind that even with the signed requests we highly recommend using an HTTPS channel.

## Response Format

When interpreting the response we will take into account 3 things:

* [HTTP status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes): anything else outside the `2xx` is considered an error.
* response body: we expect a JSON object, with a specific structure for each action type. We explain this structure on the [Actions details](../action-details) page.
* a `X-SMCCSDK-SIGNATURE` header that should contain the hex encoded [HMAC SHA-512](https://en.wikipedia.org/wiki/HMAC) signature for the response body.

## Complete Example

Below we have written an example for handling an initial request sent by the SDK. The request body will be:

```json
{ "action": "implementation.info", "time": "2012-10-01T17:18:40Z" }
```

and the request signature header (see the example for the secret key):

```http
X-SMCCSDK-SIGNATURE: 826b61e7939505b2e773ef43a2aad53ec0385dd9d783fbd1c8fea00d0e2a3e2fb0ae0a5b2eb342356b61c41b5f19baec4c1f7e7e37a5b486fe9b593942017ff9
```

=== "PHP"

    ```php
    <?php

    function sign_string($str) {
        // The secret key/API access token we provided
        $secret = '3YJZzqMJ5Ec7i2JGvnt8TgvleD7dtpwpmag4S6MuRA2GQdfvV4STIsxDRJ4fEjO8';
        return hash_hmac('sha512', $str, $secret);
    }

    function respond($response) {
        $body = json_encode($response);
        $signature = sign_string($body);
        header('X-SMCCSDK-SIGNATURE: ' . $signature);
        die($body);
    }

    function respond_error($message) {
        header('Status: 400 Bad Request');
        die($message);
    }

    function handle_implementation_info() {
        respond(array(
            'objects' => array('messages.list', 'messages.show', 'messages.create'),
            'options' => array()
        ));
    }

    // read the request body
    $request_body = file_get_contents('php://input');

    // read the signature header we sent
    $signature = $_SERVER['HTTP_X_SMCCSDK_SIG'];

    // compute the signature of the request body using the secret key
    $expected_signature = sign_str($request_body);

    // the signatures don't match - somebody tampered with the content
    if ($signature !== $expected_signature)
        respond_error('Invalid signature');

    $request = json_decode($request_body);
    $action  = $request['action'];

    // processing the request
    if ($action == 'implementation.info')
        handle_implementation_info();

    // other actions here, like:
    // if ($action == 'messages.list')
    //     handle_messages_list();

    // if the action was not handled it means the request was incorrect
    respond_error('Invalid action');

    ?>
    ```

=== "Ruby" 

    ```ruby
    # Assuming you are using Rails
    require 'openssl'

    def sign_string(str)
        # The secret key/API access token we provided
        secret = '3YJZzqMJ5Ec7i2JGvnt8TgvleD7dtpwpmag4S6MuRA2GQdfvV4STIsxDRJ4fEjO8';
        OpenSSL::HMAC.hexdigest OpenSSL::Digest::SHA512.new, secret, str
    end

    def respond(response)
        body = ActiveSupport::JSON.encode response
        response.headers['X-SMCCSDK-SIGNATURE'] = sign_string(body)
        render :text => body
    end

    def handle_implementation_info
        respond({
            :objects => {'messages' => ['list', 'show', 'create']},
            :options => []
        })
    end

    # read the request body
    request_body = request.raw_post

    # read the signature header we sent
    signature = request.headers['x-smccsdk-signature']

    # compute the signature of the request body using the secret key
    expected_signature = sign_string(request_body)

    # the signatures don't match - somebody tampered with the content
    if signature != expected_signature
        raise 'Invalid signature'

    request = ActiveSupport::JSON.decode request_body
    action  = request['action']

    case action
    when 'implementation.info'; handle_implementation_info
    # Other actions here, like
    # when 'messages.list'; handle_messages_list
    # If the action was not handled it means the request was incorrect
    else; raise 'Invalid action'
    end
    ```