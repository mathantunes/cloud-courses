# DynamoDB Architecture by Mark Richman

## Partitions

Allocation of storage for table backed by SSD.
Replicated between multiple AZ in a region.

Increases number of partitions automatically.

## Table

Collection of data

* Name
* Primary Key
    * Partition Key
    * Partition Key and Sort Key -> Composite primary key

**Unlimited size**

Partition key determines how the data is distributed across all of the nodes in DynamoDB storage.

## Performance

Throughput capacity depends on read/write capacity.

* On-demand (automatic scale) (charged by reads and writes)
* Provisioned (provide number of reads and writes RCU/WCU)
    * May throw trottling exception

**Max 40000 RCU and 40000 WCU**

### ARN

arn:aws:dynamodb:REGION:ACCOUNT:table/TABLENAME

## Items

A table may contain multiple items.

An item is a unique group of attributes, similar to rows or records in a SQL DB (limited to 400KB)

## Attribute

Fundamental data element, similar to fields or columns in RDBMS.

### Data Types

Scalars:

* Number
* String
* Binary
* Boolean
* Null

Document:

* List (items can have different values)
* Map
* Set (list always of the same data type)
