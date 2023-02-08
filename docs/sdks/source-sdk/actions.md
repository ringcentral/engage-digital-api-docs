# Introduction to Actions

All requests we send specify an `action` with optional parameters. The action name consists of an object and an operation applied to it, eg: `messages.list`.

Available actions are supplied in 2 contexts:

* on the initial request, you will have to specify all possible operations for each object. Eg: `"messages": ["list", "create"]`
* for messages and private_messages we will receive you **may** specify a list of actions (and custom actions) supported by that particular instance. Eg: `"actions": ["reply", "unpublish", "delete"]`

The reason for this is so that we know all implemented features for a specific collection, while being able to allow you to say when a particular action cannot be applied to a specific object instance. For example `unpublish` might be implemented for `messages`, but you only allow messages created by a certain category of users to be unpublished.

For details about the parameters and response for each action please the [Actions details](../action-details) page. If you want to better understand how actions work it would be better to read this page first.

## Standard vs Custom

There are a few operations which have a predefined meaning, in the sense that we expect a certain behavior to occur. These actions have a standard flow within Engage Digital and a specific interface change will occur (eg. the content will disappear for a `delete` action). The standard actions are:

* `create`
* `delete`
* `list`
* `publish`
* `show`
* `unpublish`
* `reply`

Any other operation will be considered a custom action and will result in a button showing up in the interface for the specific content to which it applies.

Please have a look on the [Actions details](../action-details) page for the complete list of parameters each action needs and the answer we except to receive.

## Collection vs Instance Level

Collection-level methods should only be specified in the `implementation.info` response and not in the `actions` field of an object (if you do, they will be ignored). They are called without specifing an object as a target. The list of actions is:

`implementation.info`

* `object.list`
* `object.show`
* `object.create`
* `object.reply`

Instance level actions should be specified in both the implementation info and for each particular instance of an object. These include the rest of standard actions and all the custom actions.

## Mandatory vs Optional

A valid implementation has to implement at least the following action:

`implementation.info`

`implementation.info` is initial action we send to discover the capabilities of your implementation.

`list` is **expected** for each object that you want to update via [Polling](../polling).
