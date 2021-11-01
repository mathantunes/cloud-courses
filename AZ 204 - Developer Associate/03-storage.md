# Cosmos DB

Microsoft's primary NoSQL. -> Flexible, scalable and reliable

Cosmos is PaaS -> User concerns code and security

Can be configured to use multiple APIs:
* SQL
* MongoDB
* Cassandra
* Gremlin
* Table Storage (Azure)

Highly available (99.999%)
Globally distributed
Low latency

## Getting started

* CosmosDB Account
  * Foundation unit -> root
* CosmosDB Database
  * Management unit for containers relating to the db/namespace
* CosmosDB Container
  * Scalability unit for both throughput and storage

### Choosing API

* SQL API
  * Migrate from SQL platform
  * Data stored as **JSON**
  * Can be queried using SQL-esqueries
  * Dedicated SDK for this API
* MongoDB
  * Best when already using nosql
  * Seamless import and migration
* Table storage
  * Advanced form of Azure Storage Tables
  * Global replication and high availability
* Cassandra
  * Quick conversion from Apache Cassandra
  * Not much gain besided the scalability and global replication
  * Can use native CQL
* Gremlin
  * Fully managed graphdb
  * Global replication, quick scaling and high available

## CosmosDB pricing

* Provisioned throughput
  * Cost and performance based on pre-defined threshold, scale as needed after
  * Scales in 100 RU/s each step
  * Used in predictable workloads
* Serverless
  * Cost and performance based on what is used
  * Automatic scale based on necessary
  * Best for inconsistent and unpredictable workloads
  * **Still in preview LOL**

## Availability

* Spread data in other regions
* Rely on azure replication
  * 4 copies of all the data at all time in the region setup for the **COSMOS ACCOUNT**

## Consistency

Somehow we have 5 types

* Boundless Staneless
  * Same region will be readily available
  * Other regions will have delay
    * Strong but with tolerated delay (**WTF**)
* Session
  * For duration of a session
  * All requests from that session will be transactionally consistent
  * Other sessions do not wait for commits
* Consistent prefix
  * Lazy updates
  * In order
* Eventual
  * Updates happen eventually but **out of order** (**WTF**)
* Strong
  * All writes are spread across all regions at the same time
  * Low performance, high latency but consistent

**Both use Request Units per Second (RU/s)** somewhat like in AWS

## Programming Cosmos DB

### CosmosDB Containers

Table to cosmosdb
* Are scalable
* Stores data, functions, triggers, stored procedures and so on
* Can be accessed by SDK

`az cosmosdb sql container create `
`New-AzCosmosDBSqlContainer`

### Partitioning

* Physical
  * Each partition allows for 10k RU/s
  * Each partition stores 50GB
* Logical
  * Can impact the physical partition
  * PartitionKey is used to sort the containers
  * Choose partitionkey for best grouping

### Scaling

DB Layer or container layer

* DB Layer 
  * scaling has to be set @ deployment time
  * uses throughput pool for all DB instances
* Container Layer
  * Allows scaling after deployment
  * Individual to each container
  * Can be done manually or with autoscale
  * Manual
    * Pick RU/s
  * Autoscale
    * Pay for the highest RU/s in use for the hour

### Server-side programming

Mostly for SQL API

* Container level mostly
  * Used for Stored Procedures, functions and triggers
  * Written in **JAVASCRIPT** (LOL)
  * Can be executed through:
    * portal
    * **JS Integrated query API**
    * CosmosDB SQL SDK

* Good for:
  * Procedural logic -> Batches through SP
  * Atomic transactions
  * Performance -> low latency
  * Encapsulation -> standard data access

# Blob Containers

Object-oriented storage

* Great for unstructured data
* Videos, music, files, logs, images
* Long-term retention
* SDKs for many languages

Hierarchy

* Azure Storage Accounts
  * Blob Containers (Organizes blobs)

## Types of blobs

* Block
  * General usage (Text and binary)
  * Can be managed individually
  * Can store up to 4.75 TiB
* Append
  * Optimized to be frequently appended
* Page
  * Used for random accessed files
  * Used for files such as Virtual Hard Disk (VHD files)
  * Up to 8 TiB

## Types of Storage Accounts

* General Purpose V2
  * Basic account for all types of storage
  * Most versatile and recommended
* General Purporse V1
  * To maintain legacy systems
* BlockBlobStorage Account
  * Optmized for Block and Append blobs
  * Mostly for high transaction rates
  * Large amounts of small objects
  * Low latency
* FileStorage Account
  * Enterprise
  * Premium file share
  * Support to File Store only
* BlobStorage Account
  * Legacy type

**Naming -> numbers and lowercase only and UNIVERSALLY UNIQUE**

## Access Tiers

Optimization for different access patterns

Can easily switch between tiers.

* Hot
  * Frequent access required
  * DEFAULT TIER
  * Cheap transfer rates
  * Expensive store
* Cool
  * Data accessed once every 30+ days
  * Expensive transfer rates
  * Cheap store
* Archive
  * Compliance, legal
  * Cheap store
  * Need to 're-hydrate' before access -> may take hours

<img src=./assets/blobendpoints.png/>

## Move data

Azure console, SDK or scripts

* Azure
  * One time process
* Scripted
  * `azcopy copy storageUrl storage2Url` -> nice to copy all of it
  * `az storage blob copy start` -> move one or multiple
  * Powershell is terrible here

## Meta data

Rest operations to create, find, manipulate and remove blobs

* PutBlob
* GetBlob
* GetBlobProperties
* SetBlobProperties
* GetBlobMetadata
* SetBlobMetadata
* DeleteBlob

```C#
GetPropertiesAsync()
```

They all return **XML** data LOL

## Lifecycle of blob

Move to different access tier based on props

Changing from hot to cold based on access
Deleting based on TTL

<img src="./assets/lifecyclepolicy.PNG"/>
