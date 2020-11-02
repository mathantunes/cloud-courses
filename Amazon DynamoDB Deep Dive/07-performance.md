# Table Performance by Mark Richman

## Provisioned vs. On-Demand capacity modes

### Provisioned

There is a minimum required capacity to be configured (as low as 1)

Setting a low capacity will lead to throttling

**Pricing is per region**

### On-Demand

No minimum configuration

No over or under provisioning database

**Pricing per million WCU/RCU** 

Throttling can occur if you double traffic in less than half an hour.
Scaling dynamoDB takes time.

## Auto Scaling

Configured on Provisioned capacity

Separated by Read and Write capacity scaling

Process:

* Configure Scaling
* DynamoDB generates consume metrics on cloudwatch
* When exceeds, cloudwatch generates an alarm
* Application Auto Scaling is triggered and evaluates the scaling policy
* Issues an updateTable request to dynamoDb, to update provisioned throughput

Properties:

* Target Utilization
* Minimum provisioned capacity (min 1)
* Maximum provisioned capacity (max 40000)

