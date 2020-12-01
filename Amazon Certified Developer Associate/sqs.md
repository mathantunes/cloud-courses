# SQS

## SQS Delay Queues

Postpone delivery of new messages

* Messages sent to the **Delay Queue** remain invisible to consumers for the duration of the delay period.
* Up to 900 seconds

* For standard queues, changing the setting does not affect the delay of messages already in the queue. Only new messages
* For FIFO queues, this affects the delay of messages already in the queue

### Use Case

* Large distributed applications which may need to introduce delay in processing
* Allow for asynchronous updates on other parts of the distributed system

## Managing Large SQS Messages

* Over 256KB up to 2GB in size
* S3 to store the messages
* Use Amazon SQS Extended Client Library
  * Provides an API for S3 bucket and object operations
  * Specify messages to be stored in s3
  * Send a message which references a message object stored in s3
