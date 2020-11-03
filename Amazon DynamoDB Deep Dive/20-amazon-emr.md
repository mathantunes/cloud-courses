# Amazon EMR by Mark Richman

## Amazon EMR

Hive Tables to perform SQL queries in DynamoDB

Create Cluster:
* Core Hadoop
* Instance Type (EC2 Class)
* Number of Instances (3)
* Role to DynamoDB and S3
* SSH to access master node (must be accepted by security group)

### Add DynamoDB Integration

Setup (On HiveQL):

* External Hive Table -> Map to DynamoDB

```sql
-- DynamoDB table 'pinehead_records_s3'
CREATE EXTERNAL TABLE pinehead (
  type string,
  id bigint,
  album_art string,
  artist_id bigint,
  format string,
  name_title string,
  price double,
  sku string,
  year bigint,
  album_id bigint,
  number string,
  length bigint)
STORED BY 'org.apache.hadoop.hive.dynamodb.DynamoDBStorageHandler' 
TBLPROPERTIES ("dynamodb.table.name" = "pinehead_records_s3",
"dynamodb.column.mapping" = "type:type,id:id,album_art:album_art,artist_id:artist_id,format:format,name_title:name_title,price:price,sku:sku,year:year,album_id:album_id,number:number,length:length");
```

**Now we can query on SQL**

```sql
SELECT name_title, year 
FROM   pinehead 
WHERE  type = 'album' 
       AND artist_id = 303 
GROUP  BY name_title, year 
ORDER  BY year DESC;
```

It starts a **HADOOP** Map/Reduce Job on the Cluster

### Add S3 Integration

Create an External Hive Table referencing CSV on S3.

```sql
-- Loads orders.csv from S3 bucket, skipping header row
CREATE EXTERNAL TABLE Orders_S3(customer_id bigint, album_id bigint)
ROW FORMAT DELIMITED FIELDS TERMINATED BY ',' 
LOCATION 's3://dynamodb-deep-dive-static-data-dump/orders/'
TBLPROPERTIES ("skip.header.line.count"="1");
```

### Join Both DynamoDB and S3

```sql
SELECT o.customer_id, p.name_title, p2.name_title 
FROM Orders_S3 o
JOIN pinehead p ON o.album_id=p.id 
JOIN pinehead p2 ON p.artist_id = p2.id
WHERE p.type = 'album'
AND p2.type = 'artist';
```

## Amazon ElasticSearch Service

Use DynamoDB Streams to trigger a lambda and store data in ElasticSearch Service

Create a New Elasticsearch Domain
* Deployment type (development and testing)
* Version
* Domain Name (Cluster Name)
* Instance Type (EC2 Class)
* Storage Config (defaults)
* VPC Access (VPC, Subnet and Security Groups)
* Security Group must have **22 SSH** and **443 HTTPS**

On DynamoDB, Enable Streams
* Lambda Role with Elasticsearch, EC2, dynamodb, cloudwatch permissions
* Lambda Function (Must have 3rd party libraries loaded)
* Configure Lambda VPC, Subnet and Security Group
* Configure Trigger
* 
```python
import boto3
import requests
from requests_aws4auth import AWS4Auth
credentials = boto3.Session().get_credentials()
awsauth = AWS4Auth(
    credentials.access_key,
    credentials.secret_key,
    "us-east-1",  # Replace with your region
    "es",
    session_token=credentials.token,
)
host = "https://vpc-your-amazon-es-domain.region.es.amazonaws.com"
url = host + "/lambda-index/lambda-type/"
headers = {"Content-Type": "application/json"}
def lambda_handler(event, context):
    count = 0
    for record in event["Records"]:
        id = (
            record["dynamodb"]["Keys"]["type"]["S"]
            + "_"
            + record["dynamodb"]["Keys"]["id"]["N"]
        )
        if record["eventName"] == "REMOVE":
            r = requests.delete(url + id, auth=awsauth)
        else:
            document = record["dynamodb"]["NewImage"]
            r = requests.put(url + id, auth=awsauth, json=document, headers=headers)
        print(r.text)
        count += 1
    return str(count) + " records processed."
```

### Kibana Dashboard

Access through endpoint in Elasticsearch

Create Index Pattern

* Name Index
* Time Filter (no need for now)
* Access discover and start filtering (**KQL Syntax**)