# Monitoring and Metrics by Mark Richman

## Console Metrics and Cloudwatch

* Metrics (data points over time)
* Dashboards
* Alarms
* Rules for events
* View Logs

### DynamoDB Metrics

* Consumed RCU
* Consumed WCU
* Provisioned RCU
* Provisioned WCU
* Read Throttle Events
* Successful Request Latency
* System Errors (HTTP 5XX status code, internal service error)
* Throttled Requests
* User Errors (HTTP 4XX status code)
* Write Throttle Events

## Alerts and Alarms

Alarms can bew created on metrics, taking action if the alarm is triggered

Alarm States:
* Insufficient: Not enough data to judge the state - Alarms often start here.
* Alarm: The alarm threshold has been breached
* OK: The threshold has not been breached.

Alarm Components:
* Metric: Data points over time being measured
* Threshold: Configurable value to trigger alarm
* Period: How long the threshold should be bad before an alarm is generated
* Action: What to do when alarm is triggered
    * SNS
    * Auto Scaling
    * EC2

### SNS - Simple Notification Service

Email notification

* Create a Topic
* Create a Subscription whitin a Topic
    * Set email protocol and pick an email

Messages sent to the topic will be directed to the configured email

### Creating a Metric

* Cloud Watch create metric
* Choose metric
* Statistic (say, average)
* Period (until an alarm is triggered)
* Pick a notification source (SNS topic)

### Exponencial Backoff Algorithm

DynamoDB SDK has an Exponencial Backoff algorithm for retrying operations

$(base * growthfactor)^(attempts-1 )$

It defaults to:
*  Base: 0.05
*  Growthfactor: 2
*  Attempts: 10


|Attempts | Delay(ms) |
|---|---|
|1|50|
|2|100|
|3|200|
|4|400|
|5|800|
|6|1600|
|7|3200|
|8|6400|
|9|12800|
|10| 25600|

### Batch Operations

Multiple operations at the same time.

The SDK uses the retry algorithm to make sure all operations are successful


