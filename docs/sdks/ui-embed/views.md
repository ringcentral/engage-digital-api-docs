# Views

Here's a list of the different views available with the embedded UI.

## Task View

The task view URL doesn't have variable parameters. The currently connected user inbox will be displayed.

```html
<iframe src="https://{your-domain}.digital.ringcentral.com/tasks?view=no-header" width="1400" height="1000"></iframe>
```

<img src="../../ui-embed/task-view.png" class="img-fluid">

## Task Permalink View

The task permalink view displays a single task and doesn't allow navigating to other parts of the interface (other tasks, pull view, admin, etc…). Any task related action (transfer, defer, etc...) can be taken from the web console or handled through [RingCentral Engage Digital REST API](https://developers.ringcentral.com/engage/digital/api-reference/routing) since it's not available in this view.

It has one variable parameter which is the id of the task which will be displayed. Example:

!!!note
    `taskId` can be retrieved from [Get All Tasks API](https://developers.ringcentral.com/engage/digital/api-reference/Tasks/getAllTasks)

```html
<iframe src="https://{your-domain}.digital.ringcentral.com/tasks/{taskId}/permalink" width="1400" height="1000"></iframe>
```

<img src="../../ui-embed/task-permalink-view.png" class="img-fluid">

The task won't be accessible if it's not in the agent inbox.

!!!warning
    In task permalink view, the “average handling time” stat is affected. Indeed, in this view, the intervals of time the tasks are considered as selected are synchronized when the iframe gains or loses focus.
    
    For example, if the iframe loses the focus but the task is not completed, and 10 minutes later the task is completed without gaining the focus in-between, those 10 minutes won't be considered as an interval where the task was selected.


## Pull View

The pull view URL doesn't have variable parameters. The currently connected user pull view will be displayed.

```html
<iframe src="https://{your-domain}.digital.ringcentral.com/home?view=no-header" width="1400" height="1000"></iframe>
```

<img src="../../ui-embed/pull-view.png" class="img-fluid">

## Permalink View

!!!note
    `contentThreadId` can be retrieved from [Get All Content Threads API](https://developers.ringcentral.com/engage/digital/api-reference/Threads/getAllThreads)

!!!note
    `contentId` can be retrieved from [Get All Contents API](https://developers.ringcentral.com/engage/digital/api-reference/Contents/getAllContents)


The permalink view displays a single thread and doesn't allow navigating in folders or other thread. It's also the view agents gets by default when they don't have permission to access the pull view and click on a permalink. It works with the content thread and contents URL and just requires query parameter `view=permalink`. Example:

```html
<iframe src="https://{your-domain}.digital.ringcentral.com/content_threads/{contentThreadId}?view=permalink" width="1400" height="1000"></iframe>

<iframe src="https://{your-domain}.digital.ringcentral.com/contents/{contentId}?view=permalink" width="1400" height="1000"></iframe>

<img src="../../ui-embed/permalink-view.png" class="img-fluid">
```

## User Supervision View

!!!note
    `userId` can be retrieved from [Get All Users API](https://developers.ringcentral.com/engage/digital/api-reference/Users/getAllUsers)

The supervised user id need to be provided in the user supervision URL:

```html
<iframe src="https://{your-domain}.digital.ringcentral.com/supervision/users/{userId}?view=no-header" width="1400" height="1000"></iframe>
```

<img src="../../ui-embed/user-supervision-view.png" class="img-fluid">

## Analytics View

```html
# dashboard without header
<iframe src="http://{your-domain}.digital.ringcentral.com/stats?dashboard=1&view=no-header" width="1400" height="1000"></iframe>

# dashboard without header and without sidebar
<iframe src="http://{your-domain}.digital.ringcentral.com/stats?dashboard=1&view=no-header,no-sidebar" width="1400" height="1000"></iframe>
```

For analytics views, you can add the `no-sidebar` value to the `view` parameter in order to choose whether or not you want to display the sidebar. If you want to specify the two values, they should be comma separated (e.g. `view=no-header,no-sidebar`). The `no-sidebar` value will prevent agents to go to another statistic.

Here's a view without header but with the sidebar:

<img src="../../ui-embed/analytics-view.png" class="img-fluid">

Here's a view without header and without the sidebar:

<img src="../../ui-embed/dashboard-view.png" class="img-fluid">

## Specific Statistic View

For the different specific statistics, the statistic identifier need to be provided through the `identifier` parameter:

!!!note
    `statistic-identifier` can only be found when you open an analytics page from web console, and the link looks like:
    
    **https://{your-domain}.digital.ringcentral.com/stats?begin_date={begin-date}&end_date={end-date}&period%5Btype%5D={period-type}&contents_status={contents-status}&identifier=`{statistic-identifier}`&advanced_analytics_type={adanced-analytics-type}**


```html
<iframe src="http://{your-domain}.digital.ringcentral.com/stats?identifier={statistic-identifier}&view=no-header,no-sidebar" width="1400" height="1000"></iframe>
```

<img src="../../ui-embed/specific-analytics-view.png" class="img-fluid">
