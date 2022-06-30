# Introduction to Send API

The Send API allows you to send contents instantly to Engage Digital. Its purpose is to build a near realtime API when the third party to be connected can support it. This API is optional and is a way to mitigate the slow import (~2 min) of the [polling approach](../polling).

It also allows you to mark agent contents as read, which is visible by agents when '[view.messaging](../options/#viewmessaging)' option is enabled.

## Requirements

To use the send API, you must define what actions can be processed by Engage Digital. Use the [polling method](../polling) to define them in response of the `implementation.info` action (more information can be found [here](../polling)).

As with the polling method, the send API allows you to implement actions progressively.

## Workflow

* Each time that we receive a call to the `/realtime/sdk/:source_id` endpoint we will look at the content of your query to determine that it's well prepared.
* The `source_id` url parameter must be filled with the Engage Digital source id that you target.
* For security reasons, the request body must always be signed as described [here](../request-response).
* The request body must be JSON encoded.

### JSON Encoded Request Body

```json
{
  "action": "messages.create",
  "params": {
    "id": "890",
    "actions": ["reply", "show", "list"],
    "author": {
      "id": "666",
      "created_at": "07/11/2019",
      "screenname": "John Doe"
    },
    "body": "Awesome post !",
    "format": "text",
    "thread_id": "42"
  }
}
```

* `action` and `params` keys are mandatory.
* `action` key defines the action that you want to perform on which resource type.
* `params` key content depends to the action and resource type.
* If the request runs smoothly, the response should be a HTTP 200 with ok as body. Otherwise consult the [Errors](../action-details/#errors) section.

## Actions

This section lists the parameters you must send for all actions and specifies what we return in the response.

### Supported Actions

| Resource | Actions |
|-|-|
| Messages | create |
| Private messages | create |
| Threads | create |
| Users | create |
| Status | mark_as_read |
| Typing | start stop |

### Create Example

#### Request

``` json
{
  "action": "messages.create",
  "params": {
    "id": "890",
    "author": {
      "id": "666",
      "created_at": "06/11/2017",
      "screenname": "John Doe"
    },
    "body": "Awesome post !",
    "thread_id": "42"
  }
}
```

#### Response

``` json
ok
```

### Mark as read Example

### Request

```json
{
  "action": "status.mark_as_read",
  "params": {
    "id": "890"
  }
}
```

#### Response

``` json
ok
```

### Typing Example

### Request

```json
{
  "action": "typing.start",
  "params": {
    "thread_id": "42",
    "in_reply_to_id": "890",
    "preview": "Hello world"
  }
}
```

```json
{
  "action": "typing.stop",
  "params": {
    "thread_id": "42",
    "in_reply_to_id": "890"
  }
}
```

#### Response

``` json
ok
```

### Errors

| Value | HTTP Code | Description |
|-|-|-|
| Bad request | 400 | The body is not JSON formatted or malformed. |
| Invalid request format | 400 | The request is malformed, a mandatory field is probably missing. |
| Invalid request size (max: 20971520 bytes) | 413 | The request body exceeds the limit of 20 megabytes, please limit the data submitted. |
| Source not found | 422 | The `source_id` does not exist in Engage Digital. |
| Invalid signature | 422 | The signature header is invalid, don't forget to sign (cf [Request-Response](../request-response)) your body request. |
| Invalid action | 422 | The action that you provided is invalid or not implemented in Send API. |
| Invalid resource format | 422 | A required field is missing in the params object that you provided. |
| Source doesn't support structured contents | 422 | The `view.messaging` option is not implemented. [Structured contents](../structured-contents) |
| Type is required for structured contents | 422 | A structured contents type is missing. [Structured contents](../structured-contents) |
| Source doesn't support this type of structured content | 422 | A structured contents type is not supported by source. [Structured contents](../structured-contents) |
| `structured_content`: "attribute_name" must not be empty | 422 | A mandatort contents attribute is missing. [Structured contents](../structured-contents) |
| The author must be puppetizable | 422 | The author is used for structured contents creation must be puppetizable. [Structured contents](../structured-contents) |
| too many typing requests, please try again in a few seconds | 429 | Too many typing requests were sent, retry later |

### Rate Limits

There is no particular limit for the time being.

## Complete Examples

=== "PHP"

    ```php
    <?php

    function sign_string($str) {
        $secret = '56e79ea426735090665b1c107c757d6d3a692c061335e4728899af27f215edf3'; // source access_token
        return hash_hmac('sha512', $str, $secret);
    }

    $content = [
    'action' => 'messages.create',
        'params' => [
        'id' => '890',
            'author' => [
                'id' => '42',
                'created_at' => '06/11/2017',
                'screenname' => 'Alice'
                ],
            'body' => 'Nice post!',
            'thread_id' => '888'
        ]
    ];

    $body = json_encode($content);
    $signature = sign_string($body);
    $source_id = '8da592bc6fcf8bae41b2';

    $session = curl_init("https://yourdomain.engagement.dimelo.com/realtime/sdk/{$source_id}");
    curl_setopt($session, CURLOPT_HTTPHEADER, ["X-SMCCSDK-SIGNATURE: {$signature}" ]);
    curl_setopt($session, CURLOPT_POST, true);
    curl_setopt($session, CURLOPT_POSTFIELDS, $body);
    curl_setopt($session, CURLOPT_RETURNTRANSFER, true);

    $response = curl_exec($session);
    $http_code = curl_getinfo($session, CURLINFO_HTTP_CODE);

    echo "{$http_code}: {$response}\n";

    ?>
    ```

=== "Ruby"

    ```ruby
    require 'faraday'
    require 'openssl'
    require 'json'

    body = {
    action: "messages.create",
    params: {
        id: "890",
        author: {
        id: "42",
        created_at: "06/11/2017",
        screenname: "Alice"
        },
        body: "Nice post !",
        thread_id: "888"
    }
    }.to_json


    access_token = '56e79ea426735090665b1c107c757d6d3a692c061335e4728899af27f215edf3' # source access_token
    signature = OpenSSL::HMAC.hexdigest(OpenSSL::Digest::SHA512.new, access_token, body)
    source_id = '8da592bc6fcf8bae41b2' # source id
    url = "https://yourdomain.engagement.dimelo.com/realtime/sdk/#{source_id}"

    response = Faraday.new(url: url).post do |req|
    req.body = body
    req.headers['X-SMCCSDK-SIGNATURE'] = signature
    end

    puts "#{response.status}: #{response.body}"
    ```
