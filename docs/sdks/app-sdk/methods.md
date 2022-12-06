# RingCentral Engage Digital App SDK JavaScript Methods

All RingCentral Engage Digital App SDK methods are prefixed by ​SMCC ​namespace. Sub-namespace(s) may also be present.

Static methods are described using a dot (.), and instance methods are described by a number sign (#​).

## SMCC

The SMCC namespace contains some regular routines.

### SMCC.locale()

Retrieves the locale of the SMCC user.

#### Returns

The current locale of SMCC user, example: "​ fr"​.

##### Type

String

### SMCC.onLoad(callback())

Invokes provided callback when document’s DOM is fully loaded.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A callback function that is executed when DOM is loaded. |

#### Returns

The ​SMCC​ namespace.

#### Type

An SMCC namespace instance

### SMCC.reload()

Reloads the current page.

#### Returns

The ​SMCC​ namespace.

##### Type

An SMCC namespace instance

### SMCC.url()

#### Example
```http
  http://domain-test.digital.ringcentral.com/admin/users
```

#### Returns

The current URL of the page. It includes the scheme, host, port, params, etc.

##### Type

String

### SMCC.version()

#### Returns

The current version of the App SDK, example: “​ 1.0.rc1”.​​

##### Type

String

## SMCC.API

SMCC.API namespace provides a way to make requests on REST API easily. Please refer first to REST API documentation for all available methods. The API request user is the same as the current user logged in.

### SMCC.API.open(options)

#### Parameters

| Name | Description |
|-|-|
| `options` | An hash that can contain following options: <ul><li>clientId​: The App SDK application key. This option is mandatory to get an access token for current user. Refer to OAuth section for more information.</li></ul> |

#### Returns

An instance of SMCC.API and then request could be performed.

##### Type

An ​SMCC.API​instance.

### SMCC.API.\[get,post,put,delete\](path, parameters)

SMCC.API provides 4 instance methods to reach API: ​#get​, #​post​, #​put ​& ​#delete. ​All of them accept the same arguments. They respectively make a `GET`​, ​`POST`​, `PUT`​ and `DELETE` ​request on REST API.

#### Parameters

| Name | Description |
|-|-|
| `path` | The path of the request (example: "​contents"​). Note: the first `/`​ ​is not mandatory. |
| `parameters` | The parameters of the request. |

!!! info "Parameters Interpolation"
    Path could use parameters to be interpolated. Example:

    ```javascript
    api.get("contents/:id", { id: contentId, baz: "foo" })
    ```

    Is equivalent to:

    ```javascript
    api.get("contents/" + contentId, { baz: "foo" })
    ```

    Parameters that are not interpolated in path are sent as regular parameters (here: ​`baz`)​.

#### Returns

A `promise` object.

The return value of `SMCC.API.get`, `SMCC.API.post`, etc. methods is a promise. Please refer to [SMCC.Promise](../methods/#smccpromise) section for more information.

Here is an example that describes how we can chain API requests:

```javascript
api.get("contents/:id", { id: contentId }).then(
  function(content) {
    return SMCC.API.get("identities/:id", { id: content["author_id"] });
  },
  function(error) {
    console.debug(error["message"]);
  }
).then(
  function(identity) {
    SMCC.UI.alert(identity["firstname"]);
  },
  function(error) {
    console.debug(error["message"]);
  }
);
```

!!! info "Note"
    The first callback **must** return the API request.

#### Example

The following example will make a GET request to ​`/1.0/contents` and render the first message’s body in the console.

```javascript
  var api = SMCC.API.open({ clientId: "53438e23421f49f424000002" });

  api.get("contents").done(function(response) {
    console.debug(response["records"][0]["body"]);
  });
```

### SMCC.API#version([version])

Get or set the API version for provided instance. By default, all API requests are locked to the latest SMCC REST API version (1​ .0)​.

#### Parameters

| Name | Description |
|-|-|
| `version` | A string that represents the new version to be set. If omitted the current version will be returned. |

#### Returns

Current SMCC API version.

##### Type

String

#### Example

Example to get the version:

```javascript
var api = SMCC.API.open({ clientId: "53438e23421f49f424000002" })

console.debug(api.version());
```

Example to set the version:

```javascript
api.version("2.0")
```

!!! info "Note"
    The version is changed only for instance other `​SMCC.API` objects returned via `SMCC.API.open` method will have default version.

## SMCC.Content

The `SMCC.Content` namespace contains methods to deal with messages.

### SMCC.Content.current()

#### Returns

Returns the instance of the content linked to the currently selected task or null if no task is selected. Returns null when outside of push mode.

##### Type

A content object or null.

### SMCC.Content.mark(options)

The `mark` ​method allows parts of message to be marked and action could be performed on it.

#### Parameters

| Name | Description |
|-|-|
| `options` | A hash that contains the following to pass to the method: <ul><li>`name`:​ A common name of the kind of word marked, example: "customer_id"​. This option is mandatory.</li><li>`pattern`:​ The pattern of the word to replace it should be a regexp, but string is accepted.</li><li>`replacement`:​ A function which take the word as argument and returns the replaced value. This option can be omitted to preserve the original value.</li><li>`onClick(value)`:​ A function that is invoked when user clicks on marked text. The function is invoked with value as first argument and with content instance as this​context.</li></ul> |

#### Returns

##### Type

`SMCC.Content` namespace

#### Example

Here is a complete example:

```javascript
SMCC.Content.mark({
  name: ​"​customer_id"​​,
  pattern: /#\d{4}/,
  replacement: function(part) {
    return part.replace(/^#/, ​"")​;
  },
  on Click: function(value) {
    console.debug(​"T​his is client "​​+ value + "​​on message with id ​"​+ this.id());
  }
});
```

Here is how the message looks like when not marked (ie: when #mark method is not invoked):
<img class="img-fluid" width="100%" src="../../../img/app-sdk-message-not-marked.png">

Here is how the message looks like when marked:
<img class="img-fluid" width="100%" src="../../../img/app-sdk-message-marked.png">

When the customer id is clicked, the message is rendered on the console:
<img class="img-fluid" width="700" src="../../../img/app-sdk-message-rendered.png">

### SMCC.Content.onArchive(callback())

Invokes provided callback when a message is archived by an agent.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when message is archived. This function is invoked with archived content instance as `this` context. |

#### Returns

##### Type

`SMCC.Content` namespace

#### Example

```javascript
SMCC.Content.onArchive(function() {
  SMCC.Window.open({
    iframeURL: ' http://example.com ?content_id' + this.id(),
    title: 'This is a test'
  });
});
```

### SMCC.Content.onCreate(callback())

Invokes provided callback when a message is created by an agent (a reply or a
discussion initiated).

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when message is created. This function is invoked with created content instance as `this` context. |

#### Returns

##### Type

`SMCC.Content` namespace

### SMCC.Content.onExtraActionsShow(callback())

Invokes provided callback when message’s extra actions drop down is displayed.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when extra actions drop down is displayed. This function is
invoked with created content instance as `this` context. |

#### Returns

##### Type

`SMCC.Content` namespace

Here is what are `extra actions`:

<img class="img-fluid" width="147" src="../../../img/app-sdk-extra-actions.png">

### SMCC.Content.onComposeContent(callback())

Invokes provided callback when a content composition is initiated by an agent.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when a discussion is initiated. This function is invoked with
created content instance as `this` context. |

#### Returns

##### Type

`SMCC.Content` namespace

### SMCC.Content.onReply(callback())

Invokes provided callback when a message is replied by an agent.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when message is replied. This function is invoked with
created content instance as `this` context. |

#### Returns

##### Type

`SMCC.Content` namespace

### SMCC.Content#addExtraActionButton(options)

This instance method adds an extra action button on message.

#### Parameters

| Name | Description |
|-|-|
| `options` | An hash that can contain following options: <ul><li>`label`: The label of the button. This option is mandatory.</li><li>`onClick()`: A callback that is invoked with `SMCC.Content` instance as this
context. |

#### Returns

##### Type

`SMCC.Content` instance object

#### Example

```javascript
SMCC.Content.onExtraActionsShow(function() {
  this.addExtraActionButton({
    label: "A test button",
    onClick: function() {
      console.debug("Test button clicked on message #" + this.id());
    }
  });
});
```

Renders a new test button in the drop down:

<img class="img-fluid" width="154" src="../../../img/app-sdk-test-button.png">

### SMCC.Content#id()

This instance method returns the id of the message.

#### Returns

##### Type

Object ID

### SMCC.Content#interventionId()

This instance method returns the ID of the message’s intervention or null when the
message is not linked to an intervention.

#### Returns

##### Type

An object ID or null.

## SMCC.ContentForm

The `SMCC.ContentForm` namespace contains methods to deal with the message reply form.

<img class="img-fluid" width="645" src="../../../img/app-sdk-reply-form.png">

### SMCC.ContentForm.insert(text)

Insert *text* at the current cursor position in the message reply form.

#### Parameters

| Name | Description |
|-|-|
| `text` | A string that contains text to insert in the content body. |

### SMCC.ContentForm.read()

Fetch the text contained in the message reply form

#### Returns

The text contained in the message reply form

##### Type

String

### SMCC.ContentForm.insert(text)

Replace content form value by *text*.

#### Parameters

| Name | Description |
|-|-|
| `text` | A string that contains text to replace the content body. |

## SMCC.ContentThread

The `SMCC.ContentThread` namespace contains methods to deal with message threads.

### SMCC.ContentThread.current()

Returns the instance of the thread currently selected or null if no thread is selected. In
push mode it will return the thread linked to the current task.

#### Returns

##### Type

A thread object or null

### SMCC.ContentThread.onCategorize(callback())

Invokes provided callback when a thread is categorized or recategorized by an agent.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when content thread is categorized. This function is
invoked with content thread instance as `this` context. |

#### Returns

##### Type

`SMCC.ContentThread` namespace.

### SMCC.ContentThread.current()

This instance method returns the id of the thread.

#### Returns

##### Type

An object id

## SMCC.Id

The `SMCC.Id` namespace contains methods to deal with object’s IDs.

### SMCC.Id.valid(id)

Determines if provided id is valid format or not. Example: “48cc6703bdae1462ce06a555” is valid but “zzba670-bdae1462ce06a555” is not.

#### Parameters

| Name | Description |
|-|-|
| `id` | A string to check if it is a valid or invalid object’s id. |

#### Returns

##### Type

Boolean

## SMCC.IdentityDetails

The `SMCC.IdentityDetails` namespace represents the modal window when an identity is clicked on
the UI. Here is an example:

<img class="img-fluid" width="645" src="../../../img/app-sdk-identity-card.png">

### SMCC.IdentityDetails.list()

#### Returns

Returns an array of visible identity details windows. It returns an empty array if there is
no visible identity details window.

##### Type

Array

### SMCC.IdentityDetails.onLoad(callback())

Invokes provided callback when identity details modal window is loaded.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked with identity details instance as `this` context. See `addTab` example below for usage. |

#### Returns

##### Type

`SMCC.IdentityDetails` namespace

### SMCC.IdentityDetails.refreshAll()

Refresh all identity details window.

#### Returns

##### Type

`SMCC.IdentityDetails` namespace

### SMCC.IdentityDetails.show(identityId)

Renders the identity details window from the provided identity id.

#### Parameters

| Name | Description |
|-|-|
| `identityId` | An identity ID. |

#### Returns

##### Type

`SMCC.IdentityDetails` namespace

### SMCC.IdentityDetails#addTab(options)

This instance method adds a new tab on `this` identity details modal window with
provided options.

#### Parameters

| Name | Description |
|-|-|
| `options` | An option hash that can contain following options: <ul><li>`target`: The target type. Can be set to *iframe*, *popup* (opens the link in a popup), or *window* (opens the link in a new tab or a new window). If missing, a click on the tab will open the url in an iframe.</li><li>`iframeId`: Used for target: iframe in combination with `SMCC.Window.sendPostMessage`.<ul><li>Only alphanumeric case insensitive and -_ are allowed : *^[A-Za-z0-9-_]+$*. Also “parent” is a reserved keyword to send to the parent window.</li><li>If iframeId is invalid, a new window will still open without the iframe ID set.</li></ul></li><li>`url`: The complete url of the target using **https**. If missing the tab will show a loader until you call *setUrl*.</li><li>`icon`: The name of the icon to choose from [the list of available icons](http://domain-test.engagement.dimelo.com/icons/demo.html) (remove the prefix “icon-” part).</li><li>`label`: The label’s text rendered on tab control.</li><li>`title`: The title’s text (tooltip) rendered on tab control.</li><li>`width`: If popup is chosen for the target option, the width option will be used. If missing, the popup width will be half of the window width.</li><li>`height`: If popup is chosen for the target option, the height option will be used. If missing, the popup height will be half of the window height.</li><li>`default`: Passing true will focus the tab, last inserted tab has priority </li></ul>Note if label and image/icon options aren’t specified a default image is used. |

#### Returns

A tab object, with a *setUrl* method that allows to change the URL.

##### Type

Tab object

#### Example

```javascript
SMCC.IdentityDetails.onLoad(function() {
  this.addTab({
    label: " Test ",
    url: " https://example.com/identity?id= " + this.identityId(),
    iframeId: "crm-integration",
  });
  // or
  var tab = this.addTab({
    label: "Test",
    default: false,
  });
  tab.setUrl(" https://example.com/identity?id= " + this.identityId());
});
```

<img class="img-fluid" width="597" src="../../../img/app-sdk-identity-card-test-tab.png">

### SMCC.IdentityDetails#identityGroupId()

This instance method returns the identity group id associated to identity details
window.

#### Returns

##### Type

Object ID

### SMCC.IdentityDetails#identityId()

This instance method returns the identity id associated to identity details window.

#### Returns

##### Type

Object ID

### SMCC.IdentityDetails#interventionId()

This instance method returns the intervention id associated to identity details window.
Actually, this is the intervention of the message associated to identity.

#### Returns

##### Type

Object ID

### SMCC.IdentityDetails#refresh()

This instance method refresh `this` identity details window. It reloads entire window HTML.

#### Returns

##### Type

`SMCC.IdentityDetails` instance object

## SMCC.Intervention

The SMCC.Intervention namespace contains methods to deal with interventions.

!!! info "Note"
    Callbacks for interventions do not work in task mode.

### SMCC.Intervention.current()

#### Returns

Returns the instance of the intervention linked to the currently selected task or null if
no task is selected. Returns null when outside of push mode.

##### Type

Intervention object or null

### SMCC.Intervention.onAssign(callback())

Invokes provided callback when an intervention is assigned from an agent to another one.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when intervention is assigned. This function is invoked with
intervention instance as `this` context. |

#### Returns

##### Type

`SMCC.Intervention` namespace

#### Example

```javascript
SMCC.Intervention.onAssign(function() {
  SMCC.Window.open({
    iframeURL: ' http://example.com ?intervention_id' + this.id(),
    title: 'This is a test'
  });
});
```

### SMCC.Intervention.onCancel(callback())

Invokes provided callback when an intervention is cancelled on a message.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when intervention is cancelled. This function is invoked
with the intervention instance as `this` context. |

#### Returns

##### Type

`SMCC.Intervention` namespace

### SMCC.Intervention.onClose(callback())

Invokes provided callback when an intervention is closed on a message.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when intervention is closed. This function is invoked with
the intervention instance as `this` context. |

#### Returns

##### Type

`SMCC.Intervention` namespace

### SMCC.Intervention.onCreate(callback())

Invokes provided callback when an intervention is opened or assigned.

This method is equivalent to:
[SMCC.Intervention.onAssign(myCallback)](../methods/#smccinterventiononassigncallback),
[SMCC.Intervention.onOpen(myCallback)](../methods/#smccinterventiononopencallback)

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when intervention is opened or assigned. This function is
invoked with the intervention instance as `this` context. |

#### Returns

##### Type

`SMCC.Intervention` namespace

### SMCC.Intervention.onOpen(callback())

Invokes provided callback when an intervention is opened on a message by the current agent.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when intervention is opened. This function is invoked with
the intervention instance as `this` context. |

#### Returns

##### Type

`SMCC.Intervention` namespace

### SMCC.Intervention#id()

This instance method returns the id of the intervention.

#### Returns

##### Type

Object ID

### SMCC.Intervention#identityId()

This instance method returns the id of the intervention’s identity.

#### Returns

##### Type

Object ID

## SMCC.NICE.MAX

The `SMCC.NICE.MAX` namespace contains methods for the integration of NICE.MAX in ED.

### SMCC.NICE.MAX.onPostMessage(callback)

Registers a callback for messages received with *postMessage* and issuer *MAX* from an iframe. Use multiple times to register multiple callbacks.

#### Parameters

| Name | Description |
|-|-|
| `callback` | A callback called when a message from an allowed iframe is received. |

#### Returns

none

#### Example

```javascript
SMCC.NICE.MAX.onPostMessage(function(data, origin) {
  console.log(data.text, "received from", origin);
});
```

!!! info "Note"
    Works similarly to:
    [SMCC.Window.onPostMessage(callback)](../methods/#smccwindowonpostmessagecallback)

### SMCC.NICE.MAX.registerEvents(subscriptionTypes)

Registers a callback for messages received with postMessage from an iframe.

#### Parameters

| Name | Description |
|-|-|
| `subscriptionTypes` | This is a required array of strings. It is not case sensitive, since everything in the array will be normalized to lowercase. The values in the array are additive. Each option specifies a type of message that you will receive. There are a few possible valid entries: <ul><li>`all`: This will return everything.</li><li>`agent`: This will return anything that doesn't have a contact ID.</li><li>`contact`: This will return events limited either by the contact ID requested or the contact ID for the associated panel if the IFrame is connected to a skill or contact.</li><li>`contacts`: This will return all events that have a contact ID. If you are in a contact panel and want additional information with the panel's contact, add this field.</li><li>`sessioninfo`: This will return an agent's session token. If you are looking to make agent session specific API calls, register for this event.</li></ul>

!!! info "Note"
    Array can't be empty.

#### Returns

none

#### Example

```javascript
SMCC.NICE.MAX.registerEvents(["all"]);
```

### SMCC.NICE.MAX.unregisterEvents()

Unregister all callbacks linked to NICE events.

#### Parameters

none

#### Returns

none

#### Example

```javascript
SMCC.NICE.MAX.unregisterEvents();
```

## SMCC.Promise

`SMCC.Promise` provides a chainable utility object. Here is a good [blog](http://blog.parse.com/2013/01/29/whats-so-great-about-javascript-promises/) that describes how promises are working. Engage Digital SMCC
App SDK Implementation has been inspired by [jQuery](https://api.jquery.com/deferred.then/).

Let’s start with a complete example:
```javascript
function asyncEvent() {
  var deferred = SMCC.Promise.deferred();

  // Resolve after a random interval
  setTimeout(function() {
    deferred.resolve("hurray");
  }, Math.floor(400 + Math.random() * 2000));

  // Reject after a random interval
  setTimeout(function() {
    deferred.reject("sorry");
  }, Math.floor(400 + Math.random() * 2000));

  // Return the promise so caller can't change the Deferred
  return deferred.promise();
}

// Invokes the asyncEvent function and add two callbacks: one for
// when resolved, the other one when rejected.
asyncEvent().then(
  function(status) {
    console.debug(status + ", things are going well" );
  },
  function(status) {
    console.debug(status + ", you fail this time" );
  }
);
```

### SMCC.Promise.deferred()

#### Returns

Returns a new deferred object.

##### Type

`deferred` object

### deferred#promise()

#### Returns

Returns a new promise object, so the caller can't change the deferred.

##### Type

`promise` object

### deferred#reject([args])

Invokes attached `fail` callback.

#### Parameters

As many arguments you want. Those arguments are forwarded to `fail` callback.

#### Returns

##### Type

This `deferred` object

### deferred#resolve([args])

Invokes attached `done` callback.

#### Parameters

As many arguments you want. Those arguments are forwarded to `done` callback.

#### Returns

##### Type

This `deferred` object

### promise#done(callback[args])

Attaches a `done` callback to this promise.

#### Parameters

| Name | Description |
|-|-|
| `callback([args])` | A callback that accepts as many arguments as described in `deferred#resolve()`. |

#### Returns

##### Type

This `promise` object

### promise#fail(callback)

Attaches a `fail` callback to this promise.

#### Parameters

| Name | Description |
|-|-|
| `callback([args])` | A callback that accepts as many arguments as described in `deferred#reject()`. |

#### Returns

##### Type

This `promise` object

### promise#then(doneCallback[ , failCallback])

Attaches a `done` and an optional `fail` callback to this promise. `#then` method could be chained.

#### Parameters

| Name | Description |
|-|-|
| `doneCallback([args])` | A callback that accepts as many arguments as described in `deferred#resolve()`. |
| `failCallback([args])` | A callback that accepts as many arguments as described in `deferred#reject()`. |

#### Returns

##### Type

This `promise` object

## SMCC.Task

The `SMCC.Task` namespace contains methods to deal with tasks.

### SMCC.Task.current()

#### Returns

Returns the instance of the currently selected task or null if no task is selected. Returns
null when outside of push mode.

##### Type

A `task` object or null

### SMCC.Task#id()

#### Returns

This instance method returns the ID of the task.

##### Type

Object ID

### SMCC.Task.accept(taskId)

Accept the ringing task with the id matching the argument taskId.

#### Parameters

| Name | Description |
|-|-|
| `taskId` | The id of the task to accept. |

#### Returns

##### Type

An object ID or null

### SMCC.Task.onComplete(callback)

Invokes provided callback when a task is completed by the current agent.

#### Parameters

| Name | Description |
|-|-|
| `callback()` | A function that is invoked when intervention is closed. This function is invoked with intervention instance as `this` context.<br><br>In the callback, `this.action()` will give the name corresponding the triggered action. The action can be "delete_scheduled", "close", "publish", "unpublish", "complete", "deferred", "destroyed" or "archive". |

#### Returns

##### Type

`SMCC.Task` namespace

## SMCC.UI

The SMCC.UI namespace contains utility methods to deal with user interface.

### SMCC.UI.alert(message[, options])

Renders a message like standard javascript alert function.

#### Parameters

| Name | Description |
|-|-|
| `message` | The message that is displayed on dialog window. |
| `options` | An hash containing some options:<ul><li>`onClose()`: A callback that is invoked when dialog window is closed.</li><li>`onSuccess(value)`: A callback that is invoked when agent clicks on *OK* button.</li></ul>Note: onClose() callback is also invoked when dialog is *OK* button is clicked. |

#### Returns

##### Type

`SMCC.UI` namespace

#### Example

```javascript
SMCC.UI.alert("This is an example");
```

<img class="img-fluid" width="344" src="../../../img/app-sdk-ui-alert.png">

### SMCC.UI.confirm(message[, options])

Renders a message and ask agent confirmation like standard javascript confirm function.

#### Parameters

| Name | Description |
|-|-|
| `message` | The message that is displayed on dialog window. |
| `options` | An hash containing some options:<ul><li>`onClose()`: A callback that is invoked when confirm dialog window is closed.</li><li>`onConfirm()`: A callback that is invoked when agent clicks on *OK* button.</li></ul>Note: onClose() callback is also invoked when dialog is *OK* button is clicked. |

#### Returns

##### Type

`SMCC.UI` namespace

#### Example

```javascript
SMCC.UI.confirm("Do you confirm?", {
  onConfirm: function() {
    // DO SOMETHING HERE
  }
})
```

<img class="img-fluid" width="347" src="../../../img/app-sdk-ui-confirm.png">

### SMCC.UI.composeContent()

Renders the compose window for the current agent to create a new content.

#### Parameters

| Name | Description |
|-|-|
| `options` | An optional hash containing some options:<ul><li>`sourceId`: If this option is present, the compose window of the specified source will be open. Otherwise a source selection popup will be shown. *(optional)*</li><li>`categoryIds`: An array of categoryIds which will be used for reply assistant entries filtering. *(optional)*</li><li>`to`: The message recipient, if available for the source. *(optional)*</li><li>`cc`: The message copied recipient, if available for the source. *(optional)*</li><li>`bcc`: The message hidden recipient, if available for the source. *(optional)*</li><li>`title`: The message title, if available for the source. *(optional)*</li><li>`body`: The message body. *(optional)*</li></ul> |

#### Returns

none

#### Example

```javascript
SMCC.UI.composeContent();
```

<img class="img-fluid" width="151" src="../../../img/app-sdk-compose-content-1.png">

```javascript
SMCC.UI.composeContent({
  sourceId: '5af07f3cccc9ac2954237908',
  categoryIds: ['5a958fa47a8beb0ff47e2db7'],
  to: '+33123456789',
  body: 'Hello, world!'
});
```

<img class="img-fluid" width="600" src="../../../img/app-sdk-compose-content-2.png">

### SMCC.UI.executeSearch(query)

Execute a search on contents.

#### Parameters

| Name | Description |
|-|-|
| `query` | The query to execute formatted “attribute: value”. A composed search can be executed by joining multiple. |

#### Returns

##### Type

`SMCC.UI` namespace

#### Example

```javascript
SMCC.UI.executeSearch('active_and_assigned: true')
SMCC.UI.executeSearch('active_and_assigned: true categorized: true')
```

### SMCC.UI.flashMessage(message[, options])

Opens a flash message.

#### Parameters

| Name | Description |
|-|-|
| `message` | The message that is displayed on the flash message. |
| `options` | A hash containing some options:<ul><li>`sticky`: The message persists if set to true</li><li>`type`: The type of message. notice by default</li><li>`delay`: The display time in ms</li><li>`workerId`: The id of a worker to have a link to the corresponding job</li></ul> |

#### Returns

##### Type

`SMCC.UI` namespace

#### Example

```javascript
SMCC.UI.flashMessage(‘Message’, { delay: 500 })
```

### SMCC.UI.openAdvancedSearch(options)

Opens the advanced search modal box.

#### Parameters

| Name | Description |
|-|-|
| `options` | A hash containing some options:<ul><li>`search`: The initial search when the modal box is open. The format of the search input has to be used as “text: &lt;some text&gt; source: &lt;sourceId&gt;”.</li></ul> |

#### Returns

none

#### Example

```javascript
SMCC.UI.openAdvancedSearch({ c })
```

### SMCC.UI.prompt(message[, options])

Prompts to input text.

#### Parameters

| Name | Description |
|-|-|
| `message` | The message that is displayed on dialog window. |
| `options` | A hash containing some options:<ul><li>`onClose()`: A callback that is invoked when the window is closed.</li><li>`onSuccess(value)`: A callback that is invoked when an agent clicks on *OK* button.</li><li>`validate(callback(value))`: A callback that is invoked when an agent clicks on *OK* button. This callback is invoked with input value as first argument. It should return true if input value is correct. If input value is incorrect, it must return an error message has string.</li><li>`value`: An optional pre-filled value.</li></ul>Note: onClose() callback is also invoked when the dialog’s *OK* button is clicked. |

#### Returns

##### Type

`SMCC.UI` namespace

#### Example

```javascript
SMCC.UI.prompt('Please type your age', {
  onSuccess: function(value) {
    // DO SOMETHING HERE
  },
  validate: function(value) {
    return value.match(/^\d+$/) ? true : ('Invalid age: ' + value);
  }
})
```

<img class="img-fluid" width="347" src="../../../img/app-sdk-ui-prompt.png">

On invalid entry:

<img class="img-fluid" width="347" src="../../../img/app-sdk-ui-prompt-error.png">

## SMCC.Window

The SMCC.Window namespace represents a modal window that can be rendered from any page, example:

<img class="img-fluid" width="655" src="../../../img/app-sdk-window.png">

### SMCC.Window.acceptPostMessageOrigin(origin)

Allows RingCentral Engage to receive messages from *origin*. Use multiple times to allow many origins.

#### Parameters

| Name | Description |
|-|-|
| `hostname` | A string which is the origin. |

#### Returns

none

#### Example

```javascript
SMCC.Window.acceptPostMessageOrigin("https://www.google.com");
```

### SMCC.Window.onPostMessage(callback)

Registers a callback for messages received with *postMessage* from an iframe. Use multiple times to register many callbacks.

#### Parameters

| Name | Description |
|-|-|
| `callback` | A callback called when a message from an allowed iframe is received. |

#### Returns

none

#### Example

```javascript
SMCC.Window.onPostMessage(‘foo’, function(data, origin) {
  console.log(data.text, “received from”, origin);
});
```

!!! info "Note"
    The origin has to be whitelisted with `SMCC.Window.acceptPostMessageOrigin(“origin_hostname”)` to receive messages in RingCentral Engage. More info in the [iframe section](../config/#iframe-messaging).

### SMCC.Window.sendPostMessage(iframeId, namespace, message)

Send a postMessage to iframe.

#### Parameters

| Name | Description |
|-|-|
| `iframeId` | **Iframe to target**. If `iframeId` is the keyword "*parent*", it will send a `postMessage` to the parent window. |
| `namespace` | **Identify postMessage**. Namespace is always prefixed with -> *smcc*:<ul><li>`window:open -> smcc:window:open`</li></ul> |

#### Returns

##### Type

`SMCC.Window` namespace

#### Example

```javascript
SMCC.Window.sendPostMessage('crm-integration', 'window:open',{ testContent: 'Ring' });
```

### SMCC.Window.open(options)

Opens a new window with provided options.

#### Parameters

| Name | Description |
|-|-|
| `options` | An hash that can contain following options:<ul><li>`target`: The target type. Can be set to *iframe*, *popup*(opens the link in a popup), or *window*(opens the link in a new tab or a new window). If missing, a click on the button will open the url in an iframe.</li><li>`iframeId`: Used for *target: iframe* in combination with `SMCC.Window.sendPostMessage`.<ul><li>Only alphanumeric case insensitive and -_ are allowed : `^[A-Za-z0-9-_]+$`. Also “parent” is a reserved keyword to send to the parent window.</li><li>If iframeId is invalid, a new window will still open without the iframe ID set.</li></ul></li><li>`url`: The complete url of the target using **https**. If missing the window will be empty.</li><li>`title`: The title of the window.</li><li>`width`: If *popup* or *iframe* is chosen for the target option, the width option will be used.</li><li>`height`: If *popup* or *iframe* is chosen for the target option, the height option will be used.</li><li>`modal`: If set to false, the popup can be opened with other popup. Default to true.</li></ul> |

#### Returns

##### Type

`SMCC.Window` instance

#### Example

```javascript
SMCC.Window.open({
  iframeURL: 'https://example.com',
  title: 'This is a test',
  iframeId: 'crm-integration',
});
```

### SMCC.Window.close()

Closes all the windows previously opened with `SMCC.Window.open()`.

#### Returns

##### Type

`SMCC.Window` namespace

### SMCC.Window#close()

Closes the `SMCC.Window` instance on which the method is called.

#### Example

```javascript
window = SMCC.Window.open({
  iframeURL: 'https://example.com',
  title: 'This is a test'
});
window.close();
```

## SMCC.User

The `SMCC.User` namespace contains methods to deal with agents.

### SMCC.User.current()

#### Returns

Returns the current user logged in instance or nil if not authenticated.

##### Type

`SMCC.User` object or null

### SMCC.User#id()

#### Returns

This instance method returns the user’s id.

##### Type

Object ID

## SMCC.Utils

The `SMCC.Utils` namespace contains general utilities method, including those needed to deal with URL encoding.

### SMCC.Utils.decodeURI()

See documentation at [decodeURI()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/decodeURI)

### SMCC.Utils.decodeURIComponent()

See documentation at [decodeURIComponent()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/decodeURIComponent)

### SMCC.Utils.encodeURI()

See documentation at [encodeURI()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURI)

### SMCC.Utils.encodeURIComponent()

See documentation at [decodeURIComponent()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/encodeURIComponent)
