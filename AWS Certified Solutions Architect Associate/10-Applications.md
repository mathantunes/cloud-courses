# AWS Applications by Mark Richman

## SQS

Simple Queue Service - Oldest service in AWS

Message Queue for async processing

`USER -> Upload to S3 -> Lambda -> SQS -> EC2 Fleet POOL`

**Decouple** components of application so they run independently

**SQS is pull-based for EC2 instances**

**Retention is between 1 minute to 14 days -> default is 4 days**

*Visibility timeout is the amount of time that the message is invisible in the SQS queue after a reader picks the message **Max is 12 hours***

**Long polling does not return a response until there is a message in the queue**

Messages contain up to 256KB of text in any format.
*It can be above, but it will use s3 to store data*

* Standard queue (default)
  * Nearly-unlimited number of transactions per second.
  * At least once delivery
  * Out of order messages
* FIFO queue (First In First Out)
  * Exactly once delivery
  * Order is preserved
  * Allows message groups (to allow multiple ordered message groups within a single queue)
  * Up to 300 transactions per second

## Simple WorkFlow Service

Coordinate work across distributed application components.

SWF enables applications to be designed as a coordination of tasks

**Task** represents invocations of various processing steps (steps can be code, human actions, scripts...)

**Amazon utilizes it in their Warehouse with human interaction for shipping**

* *SWF executions can last up to 1 year*
* *Task oriented API*
* *Tasks are assigned only once*
* *Keeps track of all the tasks and events in the application*

**SWF Actors**
* Starters -> An application that can initiate a workflow
* Decider -> Control the flow of activity tasks in a workflow
* Activity -> Carry out the activity tasks

## SNS

Send notifications from the cloud

Highly scalable, flexible and cost-effective capabilityu to publish messages

* Push notification (iOS, Android, Google...)
* SMS Text Messages
* Email
* HTTP endpoint

Group multiple recipients using **Topics**

* *All messages are stored redundantly across multiple AZs*
* *Push-based delivery (NO POLLING)*
* *Simple API*
* *Flexible message delivery*
* *Pay as you go model*

## Elastic Transcoder

Media Transcoder -> Convert media files to different formats

To formats for Smartphones, PCs, Tablets, etc..

* *Pay based on the minutes and quality of transcode*

`USER Upload Media to S3 -> Lambda -> Elastic Transcoder -> S3`

## API Gateway

Fully managed service that makes easy for developers to publish, maintain, monitor and secure APIs at any scale

`USER -> API Gateway => Lambda / EC2/ DynamoDB`

* *Distributes traffic to EC2, Lambda, WebApp...*
* Exposes HTTPS endpoints to define a restful API
* Connect to services like Lambda & DynamoDB
* Send each API endpoint to a different target
* Efficient with low cost
* Scales effortlessly
* Track and control usage by API Key
* Throttle requests to prevent attacks
* Connect to CloudWatch to log requests
* Maintain multiple versions of the API

### Configuring

* Define an API (Container)
* Define Resources and nested Resources (URL PATHS)
* Select HTTP Methods
* Security
* Target (EC2, Lambda, DDB)
* Set request/response transformations

### Deploy

Deploy API to a Stage

* Uses API Gateway domain by default
* Can use custom domain
* Now supports AWS Certificate Manager: Free SSL/TLS certs

### Exam Tips

* *Cache endpoint's response to a specific TTL*
* CORS (Origin Policy cannot be read at the remote resource)
* *Low cost and Auto Scales*
* *Throttle API Gateway to prevent attacks*
* *Log results to CloudWatch*

## Kinesis 101

Streaming Data is data generated continuously from different sources.

**Small data sizes (KBs)**
* i.e. Stock Prices, Games, Purchases

Kinesis is a platform on AWS to send streaming data to.

Easy to load and analyze streaming data.

Types:
* Streams
  * `Data Source -> Stream to Kinesis -> Stored in a Shard -> Data Consumers (EC2)`
  * 24 hours to 7 days retention
  * Consists of SHARDS:
    * Shards are 5 transactions per second for reads up to a maximum total data read rate of 2MB per second
    * Up to 1000 records per second for writes, up to a maximum total data wwrite rate of 1 MB per second
  * Data capacity of the stream is a function of the number of shards specified for the stream (*Sum of Shards capacities*)
* Firehose
  * `Data Source -> Kinesis FH -> Optional Processing -> S3/Redshift/ElasticSearch`
  * No Shards
  * No Data Persistance
* Analytics
  * Analyse data on the fly and stores data on s3/redshift/elasticsearch
  * Analyses data on Fiherose and Streams

## Web Identity Federation & Cognito

Give users access to AWS resources after they have successfully authenticated with an identity provided (AWS, fb, google)

User receives an auth code from Web ID provided which can be traded for temprary AWS security credentials

**Amazon Cognito** is a Web Identity Federation service
* Sign up and Sign-in to apps
* Access for guest users
* Acts as an identity broker between app and web ID providers
* Synchs user data for multiple devices
* Recommended for all mobile apps in AWS
* `App -> Login with Facebook (TOKEN) -> TOKEN TO COGNITO -> AWS Creds to USER`

### User Pools

User directories used to manage sign-up and sign-in functionality for mobile and web apps.

Users can sign-in directly to the User Pool, or using FB, Amazon, Google.

Cognito acts as an Identity Broker between identity provider and AWS. **Generates JWT**

**User Data**

### Identity Pools

Temporary AWS credentials to access AWS services like s3 or DynamoDB.

**Grant to AWS resource**

### Synchronization

Tracks the association between user identity and the various different devices they sign-in from.

Uses **Push Synchronization** to push updates and synchronize users data across multiple devices. Cognito uses **SNS** to send a notification to all the devices associated with a given user identity

## Event Processing Patterns

### Event Driven Architecture Patterns

* Pub/Sub messaging
  * Decoupled systems running in response to events
    * SNS Topic `Publisher -> SNS Topic -> Subscribers`

### Dead-Letter Queue

* SNS
  *  SQS queue holds messages that failed to be delivered
* SQS
  * After *MaxReceiveCount* configuration, it will send the messages to a DLQ
* Lambda
  * Result from failed async invocations;
  * Will Retry twice and send to either an SQS Queue or SNS topic DLQ

`USER -> S3 -> Lambda (Retried but failed) -> DLQ -> Failure Notification Processor (Lambda)`

### Fanout Pattern

Use SNS Topic to fan out to multiple SQS queues

`Publisher -> SNS Topic -> SQS QUEUES (Subscribed to Topic) -> Lambda`

