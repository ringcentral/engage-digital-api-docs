# Introduction to Action Details

This page lists the parameters we send for all actions and specifies what we expect to receive in the response.

Although we don't show it here for brevity reasons, don't forget that all requests and responses [must be signed](../request-response) and that you need to send a `2xx` HTTP status code for it be considered successful.

## Errors

### Server errors

One thing that all actions have in common is the structure of the response in case of a failure. Sending any other [HTTP status code](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes) than one in the `2xx` range will mean an error occurred. It will also be helpful if you also write a short explanation in the body of the response.

#### Example error response:

```html
HTTP/1.1 503 Service Unavailable

Optional message explaining the error
```

### Resource errors

For any received message, you can return an error:

#### Response

```json
{
  "error": {
    "fatal": false,
    "message": "Unable to save message."
  }
}
```

We'll always retry to send the message after a message level error. If you want to skip the content and stop the auto retry, set the `fatal` value to `true`.

## Standard Actions

### implementation.info

#### Request

```json
{
  "action": "implementation.info",
  "time": "2012-10-01T17:18:40Z"
}
```

#### Response

```json
{
    "objects": {
        "messages": ["create", "show", ...],
        "private_messages": [...],
        "threads": [...],
        "status": ["mark_as_read"]
    },
    "options": ["messages.no_title", ...],
    "locales": {
        "default": {"custom_action": "Default name"},
        "fr": {"custom_action": "Fr Name"}
    }
}
```

### create

#### Request

```json
{
    "action": "messages.create",
    "params": {
        "id": "1",
        "body": "Hello world!",
        ...
    }
}
```

#### Response

```json
{
    "id": "1",
    "body": "Hello world!",
    ...
}
```

!!! alert "Please Note"
    Right now we'll try to create a content 3 times, if anything goes wrong, before giving up.

### delete

#### Request

```json
{
  "action": "messages.delete",
  "params": {
    "id": "1"
  }
}
```

#### Response

```json
{
    "id": "1",
    "body": "Hello world!",
    ...
}
```

!!! alert "Please Note"
    Right now we'll try to destroy a content 3 times, if anything goes wrong, before giving up.

### list

We will pass the `id` of the last previously imported content as the `since_id` parameter. This will allow you to only respond with the new messages. Please bear in mind that the first time we send the request we won't pass any `since_id` parameter (we don't have any previously imported content) and it will be your choice how many contents to send in that initial run.

If you are limiting the number of contents, you must return the **oldest** ones (but still after `since_id`) so we can then query for the following contents by giving you another `since_id`.

**Results must be ordered in descending order**. So we know which content is the last to give as `since_id` for the next calls (IDs are not necessarily ascending)

#### Request

```json
{
  "action": "messages.list",
  "params": {
    "since_id": "1"
  }
}
```

#### Response

```json
[
    {
        "id": "2",
        "body": "Hello world!",
        ...
    },
    {
        "id": "1",
        "body": "Hello world!",
        ...
    },
    ...
]
```

### publish

#### Request

```json
{
  "action": "messages.publish",
  "params": {
    "id": "1"
  }
}
```

#### Response

```json
{
    "id": "1",
    "body": "Hello world!",
    ...
}
```

!!! alert "Please Note"
    Right now we'll try to publish a content 3 times, if anything goes wrong, before giving up.

### show

#### Request

```json
{
  "action": "messages.show",
  "params": {
    "id": "1"
  }
}
```

#### Response

```json
{
    "id": "1",
    "body": "Hello world!",
    ...
}
```

### unpublish

#### Request

```json
{
  "action": "messages.unpublish",
  "params": {
    "id": "1"
  }
}
```

#### Response

```json
{
    "id": "1",
    "body": "Hello world!",
    ...
}
```

!!! alert "Please Note"
    Right now we'll try to unpublish a content 3 times, if anything goes wrong, before giving up.

### reply

The reply action is used internally into Engage Digital by displaying the reply button or not, if the call to `implementation.info` returns a `reply` action for a collection, then we can only reply to contents of that collection that also have `reply` in their own actions. **You will never receive a request with that action.**

!!! alert "Please Note"
    You don't need to implement `reply` action to reply to a content. You can create any message with the `message.create` action. `Reply` action is used to respond to a **specific message.**

## Custom Actions

All non-standard actions will receive the same parameters: the content id and the id of the user who performed the action. The response will have to show the entire content so we can process all the changes that occured.

#### Request

```json
{
  "action": "messages.example",
  "params": {
    "id": "1",
    "user_id": "1"
  }
}
```

#### Response

```json
{
    "id": "1",
    "body": "Hello world!",
    ...
}
```