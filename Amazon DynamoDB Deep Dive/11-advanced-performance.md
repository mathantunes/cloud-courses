# Advanced Performance and Scaling Considerations by Mark Richman

## Offloading Large Attribute Values to S3

Since the max is 400KB, you can:

* Compress (store binary)
* S3 and add reference (path) on DynamoDB

Scans on dynamodb with binary return **faster but less items** (16MB max)

Be very cautious when storing binary data to dynamodb.

S3 enables more functionalities if needed:

* Static File Hosting service to serve on web
* CloudFront distribution for caching to improve performance

## Hot and Cold Partition Imbalance

$CEIL((RCU / 3000) + (WCU / 1000)) = Partitions$

i.e.

Provisioned: 400WCU -> 4 partitions

Each partition can sustain up to 100 writes/s

Initially evenly distributed

Say, Partition #4 is receiving more load on 150 WCU and #1, #2, #3 are on 50 WCU.

Adaptive Capacity will apply a boost to partition #4 enabling it to consume the 150 WCU.

**Available by default**
