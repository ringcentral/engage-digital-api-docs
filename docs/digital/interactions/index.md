# Introduction to Engage Interactions

<div class="jumbotron pt-1">
  <h3 class="display-5">Getting Started with Interactions</h3>
  <p class="lead">Within the Engage Platform, Interactions encompass the functionality relating to the exchange of information between agents and customers. It is through the Interaction API that developers can post messages back into 3rd-party sources.</p>
  <p>We invite all developers to try out our Interactions API by writing a simple app to list threads that agents can engage with in almost no time at all. Get started using a Quick Start in any of the following languages:</p>
  
  <a href="quick-start/node/" class="btn btn-light qs-link">Node JS &raquo;</a>
  <a href="quick-start/php/" class="btn btn-light qs-link">PHP &raquo;</a>
  <a href="quick-start/python/" class="btn btn-light qs-link">Python &raquo;</a>
  <a href="quick-start/ruby/" class="btn btn-light qs-link">Ruby &raquo;</a>
</div>

The Engage Interactions API encompasses the functionality necessary to manage the communication between agents and customers.

## Core Concepts

### Sources

A "source" refers to the origination of a customer interaction. Sources are handles to third-party services, external to Engage. Source represent on which platform a given interaction is taking place. Example sources include: Twitter, Facebook, YouTube, etc.

### Threads

When a customer initiates a conversation via a source, this results in the creation of a thread. Associated with each node within a thread is some **content** and an **identity** to identify what was said/shared and who posted it.

### Interventions

When an agent chooses to engage with a customer within a thread, they create an "intervention." An intervention is in many ways synonymous with a "ticket" or "case." An intervention remains in an open status until the agent has successfully resolved the issue and set the disposition accordingly.

## Engage Bots

An Engage Bot is a specific kind of Engage integration that help developers build automated interactions. A bot follows this simple flow:

1. Bot is notified via a webhook of a new thread.
2. Bot fetches content associated with the interaction.
3. Bot creates an intervention.
4. Bot posts messages to the customer associated with the intervention.

Chatbots can be used to perform any of the following functions:

* Augment a thread with more detail about the customer.
* Correspond with the customer.
* Resolve the associated intervention.
* Transfer the intervention to a human agent for resolution.

<a class="btn btn-primary" href="./chatbots/overview/">Build an Engage Chatbot &raquo;</a>
