# Data Resilience, Security and Encryption by Mark Richman

## Global Tables

Fully managed, multi-master, multi-region replication solutions

* High performance, globally distributed applications
* Low latency reads and writes to locally available tables
* Multi-region redundancy for disaster recovery (DR) or high availability (HA) applications
* Easy to setup and no application rewrites required
* Replication latency typically under one second

Setup

* On demand capacity or Auto Scaling
* Global Tables tab
* Enable Streams
* Add Region
    * Choose region
    * Create replica 

## Encryption

### In Transit

* AWS and Client -> HTTPS Protocol
* DynamoDB and other AWS Services

### At Rest

When accessing the table, it encrypts

* Primary Key
* LSIs and GSIs
* Streams
* Global Tables
* Backups
* DAX Clusters

There is no **unencrypted** DynamoDB.

It is either:
* DEFAULT: CMK (Customer Master Key) **AWS Owned and Managed**
* KMS (Serverside encryption using AWS Managed CMK) [**aws/dynamodb**]

### DynamoDB Encryption client

**Not available on all SDKs**

Encrypt before sending to dynamodb

Create a CMK Key ID

On KMS Service:
* Customer Managed Keys
*  Create key
*  Choose Role
*  Copy Id
*  Use on client code

## VPC Endpoints

When accessing DynamoDB from EC2 Instance, use VPC Endpoint to route internally.

No Exposure to internet.

` EC2 -> Router -> VPC Endpoint -> DynamoDB `

VPC

* Endpoints
* Create Endpoint
    * dynamodb
    * check access policy
    * route dynamodb endpoint to vpc endpoint 
