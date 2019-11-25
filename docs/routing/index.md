# Introduction to Routing in Engage

<div class="jumbotron pt-1">
  <h3 class="display-5">Getting Started with Routing</h3>
  <p class="lead">Within the Engage Platform, work across an agent pool is managed by routing interactions to agents. The Routing API is therefore how work can be distributed across your agent pool.</p>
  <p>We invite all developers to try out our Routing API by writing a simple app to list a set of tasks representing open threads agents need to resolve in almost no time at all. Get started using a Quick Start in any of the following languages:</p>
  <a href="quick-start/node/" class="btn btn-light qs-link">Node JS &raquo;</a>
  <a href="quick-start/php/" class="btn btn-light qs-link">PHP &raquo;</a>
  <a href="quick-start/python/" class="btn btn-light qs-link">Python &raquo;</a>
  <a href="quick-start/ruby/" class="btn btn-light qs-link">Ruby &raquo;</a>
</div>

## Routing Paradigms

Within Engage there are two primary ways agents can determine what customer interactions to engage with:

### Inboxes

Agents can manage their list of outstanding interactions and interventions through the use of folders. Every folder in Engage is a smart folder - one whose contents are determined by a set of rules defined by the user.

Administrators can used these folders to monitor the set of interactions for an entire team, or manage interactions that exceed a certain age, or a set of interactions that meet other criteria.

Individual agents can use folders to keep track of the active interactions currently assigned to them, to monitor interactions with specific customers to respond to them more quickly, etc.

### Tasks

Agents can also keep track of the interactions they are responsible for resolving. Each task is associated with an intervention created with an associated thread. Tasks are resolved when the associated thread is set with the appropriate disposition.

Tasks allow agents to engage with customer much like they would interface with a todo list, giving them a list of outstanding interventions that need their specific attention.

Tasks are created automatically by the system when interventions are created and assigned to an agent.

!!! info "What is a "task" in Engage?"
    Today tasks refer exclusively to interventions that require a reply from the assigned agent. In the future, Engage may allow for tasks associated with other actions (e.g. "open a bug," or "issue a refund," etc.) to be created.
