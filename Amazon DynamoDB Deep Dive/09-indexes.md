# Table Indexes by Mark Richman

Index is a **data structure** that contains a subset of attributes from a table, along with an **alternate key** to **support query** operations. You can retrieve data from index using a query. **A table can have multiple secondary indexes**

## Global Secondary Index

Index with a partition key and a sort key that can be different from those on the base table. The primary key of a GSI can be either simple (partition key) or composite (partition and sort key). **Strong consistency is NOT supported**

## Local Secondary Index

Index that has the same partition key as the base table but a different sort key. The primary key of a LSI must be composite (partition and sort key). 

**Must be created at the time of table creation**

When performing queries, it is possible to choose from:

* Table primary key
* Global Index primary key
* Local Index sort key

## Creating an Index

### When a table already exists (GSI)

* Choose a Partition Key, optional Sort Key
* Projected Attributes
    * ALL: All table attributes are projected into the index
    * KEYS ONLY: Only the index and primary keys are projected into the index
    * INCLUDE: Only the specified table attributes are projected into the index

**Update Table witrh the GlobalSecondaryIndexUpdates**
```json
{
    "Create": {
        "IndexName": "album_id-idx",
        "KeySchema": [
            {
                "AttributeName": "album_id",
                "KeyType": "HASH"
            }
        ],
        "Projection": {
            "ProjectionType": "ALL"
        }
    }
}
```

**Query index**
```python
response = table.query(
    IndexName="artist_id-idx", # Index being queried
    KeyConditionExpression=Key("artist_id").eq(1), # Artist Id
    ProjectionExpression="title, price" # fields to return
)

# Returns all title, price from albums where artist_id is 1 = Queen
```

## Sparse Indexes

If the index configured attributes do not exist on certain items, the items are considered sparse and therefore do not apply to the query. 

## Importing Tables Using AWS Database Migration Service (DMS)

It is an EC2 instance that takes a source endpoint and a target endpoint.

Schedule a task to migrate.

* Create Replication Instance
    * Name
    * Description
    * Instance Class (EC2 class)
    * Engine Version
    * Allocated Storage
    * VPC
    * Multi Availability Zone
    * Public Access 
* Source Endpoint
    *   Identifier
    *   Engine (MYSQL)
    *   Get IP for Server (ServerName)
    *   Port
    *   SSL
    *   UserName
    *   Password
* Create IAM Role
    * DMS Role to **AmazonDynamoDBFullAccess**
    * Get Role ARN
* Target Endpoint
    * Identifier
    * Engine (DynamoDB)
    * Role ARN
* Database Migration Task
    * Identifier
    * Replication Instance (newly created)
    * Source Endpoint
    * Target Endpoint
    * Migration Type (if migrating into production environment)
    * Start as soon as it is created
    * Target table preparation mode (drop or not)
    * Include LOB columns? (Binary large objects)
    * Enable Validation (make sure migration is OK)
    * Cloudwatch logs
    * Table Mappings
        * Guided UI or JSON editor
        * Selection Rule
            * Schema  
            * Table Name
            * Action (Include All)
        * Transformation Rule
            *  Target (table)
            *  Schema Name
            *  Table Name
            *  Action (add suffix as -imported)
