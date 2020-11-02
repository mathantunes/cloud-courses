# Implementing High-Performance Architectures by Mark Richman

## Static Data Dumps (Streams to Static File)

Use Case:

When an order is created, run a lambda function from dynamodb streams to store albums on the leaderboards table and produces an index.html static website that can be seen online.

* Create a bucket on S3
    * Properties -> Enable Static Website Hosting
        * Index.html
        * Copy endpoint
        * Permissions -> Unblock public access
        * Bucket Policy -> Allow GetObject 
      * 
```json
// Policy for S3
{
  "Version": "2012-10-17",
  "Statement": [{
    "Sid": "PublicReadForGetBucketObjects",
    "Effect": "Allow",
    "Principal": "*",
    "Action": ["s3:GetObject"],
    "Resource": ["arn:aws:s3:::example-bucket/*"]
  }]
}
```

* Lambda Role
    * S3, DynamoDB, DynamoDBPipeline
* Create Lambda

```python
import decimal
import os
import boto3
from boto3.dynamodb.types import TypeDeserializer
from botocore.exceptions import ClientError
BUCKET = os.environ["BUCKET"] # env
table = boto3.resource("dynamodb").Table("orders_leaderboard")
s3 = boto3.resource("s3")
td = TypeDeserializer()
def lambda_handler(event, context):
    try:
        for record in event["Records"]:
            data = record["dynamodb"].get("NewImage")
            d = {}
            for key in data:
                d[key] = td.deserialize(data[key])
            title = d["album"]["title"]
            artist = d["album"]["artist_name"]
            update_leaderboard(title, artist)
        upload_index()
    except ClientError as e:
        print(f'ERROR: {e.response["Error"]["Code"]}: {e.response["Error"]["Message"]}')
        raise
def update_leaderboard(album_title, artist_name):
    try:
        table.update_item(
            Key={"album": album_title, "artist": artist_name},
            UpdateExpression="SET order_count = order_count + :val",
            ExpressionAttributeValues={":val": decimal.Decimal(1)},
        )
    except ClientError:
        try:
            table.put_item(
                Item={
                    "album": album_title,
                    "artist": artist_name,
                    "order_count": decimal.Decimal(1),
                }
            )
        except ClientError as e:
            raise
def upload_index():
    try:
        html_head = """
            <html>
                <head>
                    <style>
                    table, th, td {
                    border: 1px solid black;
                    }
                    </style>
                </head>
                <body>
                <h1>Pinehead Records: Top Orders</h1>
                <table>
                    <thead>
                        <tr>
                            <th>Album</th>
                            <th>Artist</th>
                            <th># Sold</th>
                        </tr>
                    </thead>
                    <tbody>
                """
        table_body = ""
        html_foot = "</tbody></table></body></html>"
        items = table.scan()["Items"]
        items.sort(key=lambda x: x["order_count"], reverse=True)
        for item in items:
            table_body += (
                "<tr>"
                + "<td>"
                + item["album"]
                + "</td> "
                + "<td>"
                + item["artist"]
                + "</td> "
                + "<td>"
                + str(item["order_count"])
                + "</td> "
                + "</tr>"
            )
        html = html_head + table_body + html_foot
        s3.Bucket(BUCKET).put_object(
            Key="index.html", Body=html.encode("utf-8"), ContentType="text/html"
        )
    except ClientError as e:
        raise
```

## DynamoDB Accelerator (DAX)

In memory cache (10x Performance Improvement)

Read through and Write through DynamoDB

DAX is an EC2 instance inside the VPC.

Available in 3 AZs

Exposes the same APIs for Read and Modify.

Does not support CreateTable...

Create DAX Cluster

* Name
* Node Type (EC2 Class)
* Cluster Size (1 to 10 nodes **default 3**)
* Enable Encryption
* IAM Role to interact with DynamoDB
* SubnetGroup -> Nodes deployed to this subnet group
    * Distributed in the subnets and AZs
* Security Group -> Allow inbound on 8111 
* Parameter group -> Configuration to all nodes in the group
* Maintenance Window (if needed)

## SQS Write Buffer

Process and Receive high volumes of messages

**Asynchronous processing**

* Orders are published to SQS **Orders** Queue.
* A Lambda trigger is invoked from SQS Triggers
* The Lambda calls PutItem into DynamoDB
* On Success, the lambda deletes message from SQS Queue, otherwise it is reprocessed.

SQS Queue Types: 
* Standard Queue (Unordered and At-Least-Once delivery)
* FIFO Queue -> (First in first out, exactly-once delivery)

Create Lambda SQS Policy:
* Access Logs
* SQS API
* DynamoDB API

```json
// Policy
{
  "Version": "2012-10-17",
  "Statement": [{
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*"
    },
    {
      "Action": [
        "dynamodb:PutItem"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:dynamodb:us-east-1:123456789012:table/orders"
    },
    {
      "Action": [
        "sqs:Describe*",
        "sqs:Get*",
        "sqs:List*",
        "sqs:DeleteMessage",
        "sqs:ReceiveMessage"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:sqs:us-east-1:123456789012:orders"
    }
  ]
}
```

```python
import decimal
import json
import os
import boto3
from botocore.exceptions import ClientError
QUEUE_NAME = os.environ["QUEUE_NAME"]
DYNAMODB_TABLE = os.environ["DYNAMODB_TABLE"]
sqs = boto3.resource("sqs")
dynamodb = boto3.resource("dynamodb")
def lambda_handler(event, context):
    queue = sqs.get_queue_by_name(QueueName=QUEUE_NAME)
    while True: # This is terrible!
        for message in queue.receive_messages(MaxNumberOfMessages=10): # Why get messages if there is a records from event ?!?!?!?!?!
            item = json.loads(message.body, parse_float=decimal.Decimal)
            table = dynamodb.Table(DYNAMODB_TABLE)
            try:
                response = table.put_item(Item=item)
                print("Wrote message to DynamoDB:", json.dumps(response))
                message.delete() # If this was implemented properly, there would be no need to manually delete the message
                print("Deleted message:", message.message_id)
            except ClientError as e:
                print(
                    f'{e.response["Error"]["Code"]}: {e.response["Error"]["Message"]}'
                )
            else:
                print(response)
```

