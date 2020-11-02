# Working with Items by Mark Richman

## Partitions, Partition and Sort Keys

The partition key or partition and sort key are used to uniquely write or reade an item.

**ConditionalCheckFailedException is thrown if one tries to write 2 items with the same keys**

**Any other attribute from an item is optional and may take any data type**

A Partition holds 10GB of data and supports up to 3000 RCU or 1000 WCU.

### Consisten Hashing

* Key (K)
* Storage Nodes (N)
* Selected Nodes ($Hash(K) mod N$)

i.e.

* K = "Mark"
* Hash(K) = 123123123
* Hash(K) mod 255 = 52
* Node -> S1

## Performance Units RCU and WCU

### RCU

One RCU represents one strongly consistent read request per second or two eventually consistent read requests, for an item up to 4KB in size.

**Transactional read requests require 2 RCUs for items up to 4KB**

**Filtered query or scan results consume full read capacity**

**For an 8KB item size:** 

  * Strongly consistent -> 2 RCUs
  * Eventually consistent -> 1 RCU
  * Transactional -> 4 RCUs

#### Eventually vs. Strongly Consistent Read

Eventually consistent reads might include stale data

* Eventually -> Might include stale data
* Strongly -> Reads are always up to date but are subject to network delays

### WCU

One WCU represents one write per second for an item up to 1KB in size.

Therefore: **For a 3KB item size: Standard -> 3 WCUs; Transactional -> 6 WCUs**

**Transactional write requests require 2 WCUs for items up to 1KB**

#### Write vs. Transactional Write

Writes are eventually consistent whitin one second or less


### Privisoned Throughput Calculations

$60 records/min = 1 record/s$

$recordsize = 1.5KB$  and $1 WCU = 1KB/s$

**Therefore, a WCU setting of 2 is required on the table**

## Consistency Model (Strongly vs. Eventually)

Eventually consistency happens when reading (GetItem) from DynamoDB:

* The DynamoDB Request Router will pick a Storage Node from the 3 availability zones that replicate the data from DynamoDB.
* There is a chance that the node being read is not up to date in respect to the leader storage node.

**Eventually consistent reads are 50% the "cost" of strongly consistent**

Strongly consistent reads are a 2 out of 3 architecture between the 3 Storage Nodes from DynamoDB

## Scans and Queries

### Scans

Scans will return all items and attributes

**Scans are less efficient**

Narrow down the results by filtering

```sh
aws dynamodb scan --table-name artist \
    --filter-expression "#n = :name" \
    --expression-attribute-values '{ ":name": { "S": "Dream" } }' \
    --expression-attribute-names '{ "#n": "name" }'
```

**Scan is eventually consistent by default**

```sh
# Consistent read scan
aws dynamodb scan --table-name artist \
    --filter-expression "#n = :name" \
    --expression-attribute-values '{ ":name": { "S": "Dream" } }' \
    --expression-attribute-names '{ "#n": "name" }'
    --consistent-read
```

```sh
# Limit to three on scan
aws dynamodb scan --table-name artist \
    --limit 3
```

#### Pagination

```sh
# Pagination of 10
aws dynamodb scan --table-name artist \
    --page-size 10
# Responds a LastEvaluatedKey as long as there is more items on the request
```

### Queries

Find items based on primary key values (partition key or partition and sort key)

 ```sh
aws dynamodb query --table-name artist \
    --key-condition-expression "id = :id" \
    --expression-attribute-values '{ ":id": { "N": "123" } }'
# returns instantly   
```

## PutItem Operations

### PutItem

* Creates a new item
* Replaces existing item with the same key

### UpdateItem

* Creates a new item if the specified key does not exist, otherwise modifies the existing item

### DeleteItem

* Deletes the item with the specified key

### Write Conflicts

i.e.
* Two users get the item
* Both try to PutItem
* Write conflict if both write. (last one wins)

## Batch Operations

### BatchGetItem

* Returns attributes for multiple items from multiple tables
* Request using primary key
* Returns up to 16MB of data, up to 100 items
* **If any item is left out, returns on UnprocessedKeys**

**Much faster than serial GetItem**

### BatchWriteItem

Write array of json objects

* Puts or deletes multiple items in multiple tables
* Writes up to 16MB of data, up to 25 put or delete requests
* Get unprocessed items exceeding limits via **UnprocessedItems**
* Conditions are **NOT** supported for performance reasons
* Threading may be used to write items in parallel
