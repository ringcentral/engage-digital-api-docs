# Introduction to Polling

The basic principle of the polling method is to call a defined URL periodically in order to fetch the new data from your application and send the new data created on RingCentral RingCX Digital. This periodic polling ensures that content is synchronized between RingCX Digital and your platform, ensuring we offer a coherent and cohesive view of the data to support both agents and clients. Currently, we will crawl at a rate of about:

* *1 minute* intervals.
* *3 minutes* intervals. If there are no implemented 'list' actions.

But it can take longer between crawls. We only guarantee the minimum amount of time.

This method is designed to allow for changes and progressive implementation of the features (see [Workflow](../polling/#workflow) below). Thus, you can start by implementing the few basic mandatory actions, which will be enough for us to read the contents (RSS-like functionality). At any time, more functionality can dynamically (without intervention from our side) be added or removed.

## Workflow

* Each time we **start** a synchronization run we will send an initial [implementation.info](../action-details) action to query the capabilities of the implementation.
* We look at the response and determine how to proceed depending on the response. Let's say for this example that the implementation info response looks like this:

``` json
{
  "objects": {
    "messages": ["list", "show", "create", "destroy"],
    "threads": ["list", "show", "create", "destroy", "sticky"]
  },
  "options": []
}
```

* We send the `threads.list` action to import all new threads.
* We send the `messages.list` action to import all new messages.
* We send any number of `messages.create` if we have any new messages.
* We send any number of `messages.destroy` if our agents deleted any messages.

Here is an example of the first few requests and their expected response, with the HMAC signing stripped and whitespace added for legibility:

### Initial request:

``` json
{
  "action": "implementation.info",
  "time": "2019-10-01T17:18:40Z"
}
```

### Response: 

``` json
{
  "objects": {
    "messages": ["list", "show", "create", "destroy", "spam", "unspam"],
    "threads": ["list", "show", "create", "destroy"]
  },
  "options": ["messages.unthreaded"]
}
```

### Response was ok so we proceed with the synchronization process:

``` json
{
  "action": "messages.list",
  "params": { "since_id": "2352342" },
  "time": "2019-10-01T17:18:40Z"
}
```

### Response: 

``` json
[
    {
        "id": "2352343",
        ...
    },
    ...
]
```
