# Event Types

This document outlines the specific events one can subscribe to via a webhook.

Events are usually are composed of two parts: the event resource, and the event action. The "event resource" represents the kind of object that generated the event, while the "event action" represents the event's trigger. For example, for the event type `intervention.assigned` the event type is `intervention` and the event action is `assigned`.

## Event Anatomy

Each event consists of three top-level elements. They are:

Property | Value | Description
---------|-------|------------
`id` | String | Unique request identifier. This can be used to detect duplicated request, for example in case of failure.
`domain_id` | String | Unique domain identifier.
`events` | Array of ​Event object | An array of generated events.

You can see this structure in the following sample event notification triggered upon an assignation on an intervention:

```JSON
{
  "id":"bd13a9d9baa8c20cf93046cd",
  "domain_id":"48cc6703bdae1462ce06a555",
  "events":[
    {
      "type":"intervention.assigned",
      "id":"70d340997b8cd2c6f4dfee22",
      "user_id":"4f4f3a08a90ffb27ee000583",
      "resource":{
        "type":"intervention",
        "id":"5464b5c04d61639684110000",
        "metadata":{
          "custom_field_values":{
            "sample_field":null
          },
          "category_ids":[
            "4f3951557aa58d1462017a8f",
            "50895dbea90ffb3c35001ace"
          ],
          "closed_at":null,
          "deferred_at":null,
          "identity_id":"557f003c7765620fdc0002cc",
          "source_id":"56178fd27765625e06000a66",
          "thread_id":"565739986b65795289000029",
          "user_id":"52fcf5157aa58dd768000006"
        }
      },
      "issued_at":"2014-02-10T18:35:35.251Z"
    }
  ]
}
```

Every webhook can contain multiple events. Each event then conforms to the following structure:

Property | Value | Description
---------|-------|------------
`type` | String | The type of the event. Events MUST define this property. See "Event definition" below.
`id` | String | Unique event identifier. Events MUST define this property. This can be used to detect duplicate events in case of multiple send.
`resource` | Resource Object | It gives the resource of the event. Events MUST define this property.
`issued_at` | RFC-3339 date-time | The date and time at which the event happened. Events MUST define this property. This can be used to order the events client-side.
`user_id` | String | Unique user identifier. Will only appear for event triggered by an agent.
`action` | String | task.completed only, specifies the action leading to task completion, can be “deferred” or “closed”

### Resource object

Each event is triggered by changes to a resource, e.g. a resource is created, updated, deleted, etc. That resource is described in a standard way, following the structure below:

Property | Value | Description
---------|-------|------------
`type` | String | Type of resource. Resources MUST define this property. This gives information about the type of resource that must be queried from the API
`id` | Integer | Unique resource identifier. Resources MUST define this property.
`metadata` | Metadata Object | A resource MAY define some metadatas. The metadatas are additionals informations supplied to help by avoiding making extra queries to get them.

## Available Events

The events are sorted by resource, the resource can be:

### Interventions

#### Resource Types

Type | Description
-----|------------
`intervention.assigned` | When an agent is assigned to an intervention
`intervention.canceled` | When an agent cancels an intervention
`intervention.closed` | When an agent marks the intervention as closed
`intervention.deferred` | When an intervention is deferred for a period of time (the `intervention will be reactivated later)
`intervention.opened` | When an intervention is started by engaging a customer
`intervention.reopened` | When an intervention is reopened (by receiving a new message from the identity or manually by an agent, api ...)
`intervention.custom_fields_updated` | When custom fields of an intervention are updated
`intervention.user_updated` | When the agent of an intervention is changed (re-assigned)
`intervention.reactivated` | When an intervention deferred time is finished
`intervention.recategorized` | When categories on an intervention are changed

#### Metadata

Property | Value | Description
---------|-------|------------
`source_id` | String | Unique identifier for the source (facebook, email...)
`thread_id` | String | Unique identifier for the thread
`user_id` | String | Unique identifier for the agent of an intervention
`identity_id` | String | Unique identifier for the identity which created the content
`category_ids` | Array | Array of unique identifier by categories
`closed_at` | String | The date when the intervention was closed
`deferred_at` | String | The date when the intervention was deferred
`custom_field_values` | Custom field Object | It gives the custom fields linked to the object

### Intervention Segments

#### Resource Types

Type | Description
-----|------------
`intervention_segment.summary_generated` | When a conversation summary is generated for an intervention segment
`intervention_segment.summary_submitted` | When a conversation summary is submitted by an agent
`intervention_segment.summary_generation_failed` | When an error occurred during the conversation summary generation

#### Metadata

Property | Value | Description
---------|-------|------------
`intervention_id` | String | Unique identifier for the intervention associated with this event
`content_source_id` | String | Unique identifier for the content source (e.g., Facebook, Email, etc.)
`user_id` | String | Unique identifier for the agent/user who triggered this event
`content_thread_id` | String | Unique identifier for the conversation thread
`intervention_segment_id` | String | Unique identifier for a specific segment/part of the intervention
`summary` | String | Summary of the intervention segment


### Tasks

#### Resource Types

Type | Description
-----|------------
`task.assigned` | When a new task is assigned to one or several agents
`task.completed` | When an agent marks a task as completed
`task.created` | When a task is created
`task.destroyed` | When contents are ignored or interventions are deferred or closed in folders view, related tasks are destroyed.
`task.expired_from_step` | When a task couldn't be delivered to any of the assigned agents for longer than the timeout defined for the topology step. It then continue to the next step.
`task.expired_from_workbin` | When a task was abandoned by an agent in his workbin automatically taken back by RingCX Digital to be routed to other agents.
`task.recategorized` | When a content is recategorized, the task associated to this content gets the same categories.
`task.resume` | When an agent resumes handling of a task from his history or deferred tasks
`task.supervisor_assigned` | When a task is assigned to an agent by a supervisor and the option “Bypass queue and assign to agent” has been checked
`task.taken` | When an agent accepts a ringing task
`task.transferred` | When an agent or an administrator transfers a task to another agent
`task.undelivered` | When a task is moved into the undelivered queue

#### Metadata

Property | Value | Description
---------|-------|------------
`created_at` | date-time | The date and time at which the task was created
`channel_id` | String | Unique identifier for the channel
`source_id` | String | Unique identifier for the source associated to the task's content
`priority` | Float | Priority of the task
`content_id` | String | Unique identifier for the content linked to the task
`intervention_id` | String | Unique identifier for the intervention linked to the task
`agent_ids` | Array | Array of unique identifier of agents assigned to the task
`identity_id` | String | Unique identifier of the content's author identity
`action` | String | task.completed only, specifies the action leading to task completion, can be “deferred” or “closed”
`queue` | String | Task’s current workbin
`thread_id` | String | Unique identifier for the thread linked to the task
`type` | String | task.taken only, specifies if the agent requested a task or took a ringing one. Can be “ring” or “request”
`language` | String | Language of the task's content
`segment_index` | Integer | Number starting at 1 and incremented each time a task is closed, defered or moved to another agent's workbin.

### Push Agents

#### Resource Types

Type | Description
-----|------------
`push_agent.accept_task` | When an agent accepts a ringing task
`push_agent.availability_change` | When an agent's status changes
`push_agent.busyness_change` | When an agent’s busyness changes (according to the number of tasks in their workbin)
`push_agent.connected` | When an agent arrives on the push view
`push_agent.disconnected` | When an agent leaves the push view
`push_agent.reconnected` | When an agent reconnects after a network issue
`push_agent.request_task` | When an agent is busy but hasn’t reached the hard capability and requests an additional task

#### Metadata

Property | Value | Description
---------|-------|------------
`channels` | Array | Array containing data for each channel in push mode. Each channel has the following properties
`id` | String | Unique identifier for the channel
`name` | String | Name of the channel
`status` | String | Status of the push_agent on the channel
`busyness` | String | Busyness of the push_agent on the channel
`custom_status` | Hash | Id of the custom_status used by the agent

### Content

#### Resource Types

Type | Description
-----|------------
`content.approved` | When a content has been approved for publication. **This does not mean that the content has been published yet. Notifications about exports are done with content.exported.**
`content.discussion_initiated` | When a discussion has been started. **This does not mean that the content has been published yet. Notifications about exports are done with content.exported.**
`content.exported` | When a new content has been exported from RingCX to the source. **This does not mean that the export succeeded. The synchronization_status field needs to be checked.**
`content.imported` | When a new content has been imported from the source to RingCX.
`content.replied` | When someone replied to another content. **This does not mean that the content has been published yet. Notifications about exports are done with content.exported.**
`content.thread_auto_closed` | When an Engage Messaging thread was automatically locked.
`content.thread_closed` | When a thread is locked.
`content.thread_opened` | When a thread is unlocked.
`content.update_exported` | When a content has been update_exported (and the update has been propagated to the external source).

!!! warning "New content events"
    To be notified of new content from customers you should watch for content.imported but be careful as you will receive more than just customer contents.

    In particular the ​`content.imported` event means imported from a RingCX Digital point of view, it notifies all interactions coming directly from the synchronized source. For example you publish a post on a managed facebook page without using RingCX Digital, a `content.imported` webhook will be triggered. For the Chat automatic welcome message (initial message) of the conversation will also be imported. So it’s not only customer message, this is a bit more complicated and some filtering may be required. To ignore those messages you will filter based on the status field to segregate actionable contents from customer vs. source itself.

    Similarly, to be notified when a content has been exported from RingCX Digital to the source, you should watch for ​`content.exported` (new content) which will also notify automatic messages like survey or auto response and ​`content.update_exported` for update on existing content.

    Finally, if you need more granular or earlier notification about contents creation, you can subscribe to ​`content.discussion_initiated` (new discussion), ​content.replied (reply on a message) and `content.approved​` (content has been approved).

    It can be useful to understand the basic workflow of contents:

    1. A new content is created from RingCX Digital: it can trigger `content.discussion_initiated` if a new thread has been initiated by an agent (outgoing content). Otherwise, regular replies to contents triggers content.replied.​ Those content may not be exported yet depending on approval settings.
    2. If the content requires to be approved (​`approval_required` attribute is ​`true​`), it has to wait for an approbation. Once the approbation is received, content.approved​ is triggered.
    3. The content can now be exported to the external source. Successful and Unsuccessful export operations trigger the ​`content.exported` event for new contents, and ​`content.update_exported` event for existing events that have been edited.

#### Metadata

Property | Value | Description
---------|-------|------------
`id` | String | ID of the content
`approval_required` | Boolean | Whether the content require approval before being published
`author_id` | String | ID of the author of the content
`body` | String | Body of the content
`body_input_format` | String | Format of the body (html, text)
`created_from` | String | From where the content has been created (interface, API, ...)
`creator_id` | String | ID of the creator of the content
`date` | String | Date at which the content was created
`first_in_thread` | Boolean | Whether the content is the first of the associated thread
`foreign_categories` | Array of Strings | Foreign categories associated to the content
`foreign_id` | String | ID of the content on the external website
`has_attachment` | Boolean | Whether the content has one or multiple attachments
`intervention_id` | String | ID of the intervention associated to the content
`language` | String | Language of the content
`type` | String | Type of the content
`in_reply_to_author_id` | String | ID of the author of the content for which this content is a reply
`in_reply_to_id` | String | ID of the content for which this content is a reply
`private` | Boolean | Whether the content is a private message
`source_id` | String | ID of the source associated to the content
`source_type` | String | Type of the source associated to the content
`status` | String | Status of the content
`synchronization_status` | String | For content.exported and content.update_exported event, the status of the synchronization (successful or unsuccessful).
`thread_id` | String | ID of the thread that owns the content
`thread_title` | String| Title of the thread that owns the content.
`capabilities_supported` | Array of Strings | List the types of structured messages that can be used to reply to this type of message.

### Identities

#### Resource Types

Type | Description
-----|------------
`identity.merged` | When an identity is merged to other identities.
`identity.unmerged` | When an identity is unmerged from other identities.

#### Metadata

Property | Value | Description
---------|-------|------------
`new_identity_group_id` | String | ID of the new Identity Group attached to the identity after the merge/unmerge.
`old_identity_group_id` | String | ID of the Identity Group which was previously attached to the identity before the merge/unmerge.

### Surveys

#### Resource Types

Type | Description
-----|------------
`survey_response.imported` | When a response to a survey has been imported.

#### Metadata

Property | Value | Description
---------|-------|------------
`submitted_at` | date-time | Date/time response was submitted (EST/EDT or GMT -5/GMT -4)
`main_indicator` | Integer | Main question result
`main_indicator_scaled` | Float | Main question result scaled to 0..1
`intervention_id` | String | ID of the intervention related to this survey
`survey_id` | String | ID of the survey
`source_id` | String | ID of the source related to this survey
`user_id` | String | ID of the user related to this survey
`questions` | Array of Hashes | List of questions containing the displayed `text` of the question, its `foreign_id` (from the third party survey supplier) and their `replies` in the form of an Array. Each one of the `replies` entry contains two keys: `text` and `value` of the customer reply.
`response_foreign_id` | String | Internal Response ID fetched from the provider.
`answers` | Hash | `[Deprecated]` Please use `questions` instead. Survey question and responses
`foreign_id` | String | `[Deprecated]` Please use `id` instead. Internal Response ID generated by RingCX Digital.

### Custom Fields

Property | Value | Description
---------|-------|------------
`key` | String | To have more information you need to look in RingCX Digital custom field admin, it depends on your setup
`value` | String | To have more information you need to look in RingCX Digital custom field admin, it depends on your setup
