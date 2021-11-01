# The Azure Connection

## Logic Apps

Used to **schedule, orchestrate, automate** processes and workflows in Azure

Integration platform as a service.

### Create

* Login Azure
  * Logic App Service
  * Fill wizard

* CLI
  * `az logic workflow create ...`
    * --definition -> JSON workflow file

* View workflow file in the *App Logic designer*
  * Represents the components of the workflow
  * Triggers
  * Version
  * Parameters
  * Outputs

### Custom Template

Mean to tell Logic Apps what to do

`Trigger -> Condition -> Actions`

* Microsoft has many templates already
* Crafting a custom template
  * Trigger
  * Condition
  * Action

**Can be done with drag and drop lol**

### Custom Connector

Connections to sources that are not yet integrated into Azure.

Will make use of REST or SOAP APIs to allow communication

* Access Logic Apps Custom Connector
  * Name it
  * Some more metadata
  * Edit connection params
  * Choose endpoint
    * Supports OPEN API
  * Authentication
    * None
    * Basic ?
    * API Key
    * OAuth 2.0
  * Can also include actions and triggers here
    * Validate endpoint and things like that

## API Management

Versatile front end layer for the APIs.

Can publish to internal, external and partnet developers.

Hides access to backend

Broken down into 3 portals

* Dev
  * learn about the api
  * get documentation
  * manage api keys
  * playground
  * analytics
* Admin
  * management console
  * define and import schema
  * package api
  * manage users
  * policies
  * analytics
* Proxy
  * Forwards calls to the backend api
  * calls routes appropriately
  * verify keys, tokens or credentials
  * enforce policies
  * cache backend responses
  * log calls and metadata

**Product** is how the API is defined.
* A Product can have 1 or more APIs, a Title, description and terms of use.
* Types
  * Open
    * Anonymous
  * Protected
    * Must authenticate
    * Subscribe

### Policies

Collection of statements that run sequencially and can alter the behavior of the API

Will be applied on request and on response

JSON file
* Inbound policies
* Backend ??
* Outbound
* OnError

## Event Grid

Event storage solution that can help gather event data from many sources.

Pub-sub model.

Many azure services can be configured to send events to event grid.

The course showed an example with BlobStorage account configuration for event grid.
* Configured topic name and some other random stuff
* Pick an `endpoint` as destination
  * Chose webhook to point to a web app running
  * Entered the url for the web app
  * LOL

### Events

What is being registered

### Source

Service responsible for generating the event

### Topics

Grouping to where the event was sent

### Subscriptions

Used to receive event notifications

### Handlers

Reacting handle on the subscribed service.

Subscribers are called handlers here lol

## Event Hub

Stream data rather than store it

Alternative to Event Grid.

More of a data analytics usage

**Real time event** management and response

`Event Producers -> HTTP(S)/AMQP -> Partitions in EventHub -> Subscriber Pulls what they need`

Store data for no more than 7 days (similar to apache Kafka)

Use **Event Hub CAPTURE** to store for longer than 7 days
* Can store in blob storage
* Azure data lake storage

## Notification Hub

Managed solution to send push notifications to multiple platforms.

Supports mobile tech - iOS, Android with Firebase Cloud Messaging and Windows notification service

## Azure Service Bus

Specific for message brokering

Queuing, Pub/Sub and advanced integrations

Multiple communication protocols, data contracts, trust domains or networks

Used for synchronous transactions such as inventory management

**Suports FIFO if using Managed Sessions**

## Azure Queue Storage

Part of azure storage accounts
Managed and secured by storage accounts

Stored by HTTP and HTTPS calls
64kb max -> As many as u have of storage

* Async processing
* Large volume

**DO NOT SUPPORT FIFO**
