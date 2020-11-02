# Advanced Index Usage by Mark Richman

## Selective Write Sharding

Small number of partitions with a lot of WCU

**Add Random number to the end of the partition key value to expand partition space**

Scenario:

Only two options:
* Candidate_A
* Candidate_B

To optmize sharding, create Candidate_A_{RANDOMNUMBER}.

To aggregate the total of votes, a lambda function can sum up all the Candidate_{X}_{RANDOMNUMBER}

Periodically through CloudWatch Event

### Creating the system

* Define a role for lambda with DynamoDBFullAccesswithDataPipeline and DynamoDBFullAccess
* Create Lambda Function

```python
"""
Run periodically, sum the values for each candidate segment, saving back
to the table as "Candidate A: total" and "Candidate B: total"
"""
import boto3
from botocore.exceptions import ClientError
dynamodb = boto3.resource("dynamodb")
table = dynamodb.Table("votes")
def lambda_handler(event, context):
    try:
        items = table.scan()["Items"]
        a = 0
        b = 0
        for i in items:
            if i["segment"].startswith("Candidate A_"):
                a = a + i["votes"]
            if i["segment"].startswith("Candidate B_"):
                b = b + i["votes"]
        table.update_item(
            Key={"segment": "Candidate A"},
            UpdateExpression="set votes = :v",
            ExpressionAttributeValues={":v": a},
        )
        table.update_item(
            Key={"segment": "Candidate B"},
            UpdateExpression="set votes = :v",
            ExpressionAttributeValues={":v": b},
        )
    except ClientError as e:
        print(f'{e.response["Error"]["Code"]}: {e.response["Error"]["Message"]}')
    else:
        print("Aggregate totals updated")
```
* Add Trigger
    *  CloudWatch Events
    *  Schedule Expression
        * rate(1 minute) 

## Aggregation with Streams

Same voting data. But this time using lambda streams to increment the votes table data.

**No CloudWatch events**

Trigger from DynamoDB Stream from raw_votes table.

```python
import decimal
import boto3
from boto3.dynamodb.types import TypeDeserializer
from botocore.exceptions import ClientError
table = boto3.resource("dynamodb").Table("votes")
td = TypeDeserializer()
def lambda_handler(event, context):
    sum_a = 0
    sum_b = 0
    for record in event["Records"]:
        data = record["dynamodb"].get("NewImage")
        d = {}
        for key in data:
            d[key] = td.deserialize(data[key])
        if d["candidate"] == "Candidate A":
            sum_a = sum_a + 1
        if d["candidate"] == "Candidate B":
            sum_b = sum_b + 1
    update_sum("Candidate A", sum_a)
    update_sum("Candidate B", sum_b)
    print("Successfully processed {} records.".format(len(event["Records"])))
def update_sum(candidate, sum):
    try:
        table.update_item(
            Key={"segment": candidate},
            UpdateExpression="SET votes = votes + :val",
            ExpressionAttributeValues={":val": decimal.Decimal(sum)},
        )
    except ClientError as e:
        print(f'{e.response["Error"]["Code"]}: {e.response["Error"]["Message"]}')
        raise
    else:
        print("Aggregate votes updated")
```