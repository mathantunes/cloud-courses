# DynamoDB and Scenario Introduction by Mark Richman

## History

**2004 AWS** had issues with SQL (Oracle).
AWS built their own proprietary DB.
Most of the applications were Key-value.

Create purposed key-value database to support it.

**Ref Dynamo: Amazon's Highly Available Key-value Store**

## DynamoDB

* Non Relational managed database service
* Supports both key-value and document data models
* Really fast
* Unlimited throughput and storage
* Automatic scaling up or down
* Handles trillions of requests per day
* ACID transaction support
* On-demand backups and point-in-time recovery
* Encryption at rest
* Data is replicated across multiple availability zones
* Highly available

### Where does it fit

* **Amazon RDS** (Relation database service) (aurora, psql, mysql, mariadb, oracle or sql)
* **Amazon ElastiCache** managed redis or memcache in memory data store
* **Amazon Neptune** graph database
* **Amazon Redshift** petabyte-scale data warehouse service
* **Amazon QLDB** ledger database providing a cryptographically verifiable transaction log
* **Amazon DocumentDB** MongoDB compatible database service

## Pinhead Records

Web application of record selling.
Has issues with scalability. (python and mysql stack)

### Migration Plan

#### Version 0

* Relational model in MySQL
* Limited database optimizations
* Limited application-level caching
* No indexes
* Inefficient queries
* Images stored on local filesystem
* User accounts in database

#### Version 1

* Naive migration from MySQL to DynamoDB
* Three DynamoDB tables mimicking the relational structure (Bad Idea so far)
* Images are moved to S3 with URI in a DynamoDB attribute
* No indexes
* User accounts in DynamoDB (Bad Idea so far)

#### Version 2

* Optimizations
* Better table structure (single hierarchical table)
* Indexes
* Transactions
* User accounts in DynamoDB

#### Version 3

* Federated web identity (Cognito)
* Fine-grained policies
* Lambda triggers
* Improved security
* DynamoDB Accelerator (DAX)

## Scenario Discussion

### Version 0 can be described as

#### WebServer

* No replication.
* Single point of failure

#### FileSystem

* Images stored inside the project (server filesystem)

#### SQL Schema

* Tables
    * Album (can have one or more track)
    * Artist (can have one or more albums)
    * Track
    * User (Ids and passwords)

Inflexible database/no caching