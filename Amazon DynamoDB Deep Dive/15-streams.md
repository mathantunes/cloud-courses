# Streams and Triggers by Mark Richman

## Streams

DynamoDB writes a stream record with information about an item modification.

**May give Before and After state**

* Both Old and New
* Keys only Only the Key attributes modified
* New
* Old

Streams are almost real time processing

Use Cases:
* Cross-region replication
* Relationships across tables
* Messaging/Notification
* Aggregation/filtering
* Analytical reporting
* Archiving and auditing
* Search

## Triggers

PineHead Records Use Case

`Orders Table put item -> DynamoDB Streams -> OrderConfirmation Lamba function -> Simple Email Service -> Order Confirmation`

Create orders table

* Id
* Timestamp
* Email (**REQUIRED**)

Create lambda function with corresponding IAM Role

Access IAM:

* Roles
    * New Role
    * Lambda
    * Permissions
        * AWSLambdaDynamoDBExecutionRole
        * AmazonSESFullAccess

Create function

```python
import boto3
import json
import os

from boto3.dynamodb.types import TypeDeserializer
from botocore.exceptions import ClientError

ses = boto3.client("ses")
td = TypeDeserializer()

def lambda_handler(event, context):
    print("Received event: " + json.dumps(event, indent=2))
    for record in event["Records"]:
        data = record["dynamodb"].get("NewImage")
        d = {}
        for key in data:
            d[key] = td.deserialize(data[key])
        send_email(d)
    print("Successfully processed {} records.".format(len(event["Records"])))

def send_email(data):
    SENDER = os.environ["SENDER"]
    CHARSET = "UTF-8"
    SUBJECT = "Pinehead Records: Order Confirmation"
    BODY_TEXT = (
        "Pinehead Records: Order Confirmation\r\n"
        f"Order ID: {data['id']}"
        f"Album: {data['album']['title']} ({data['album']['year']})"
        f"Format: {data['album']['format']}"
        f"Amount: {data['amount']}"
    )
    BODY_HTML = f"""<html>
    <head></head>
    <body>
    <h1>Pinehead Records: Order Confirmation</h1>
    <ul>
        <li>Order ID: {data['id']}</li>
        <li>Album: {data['album']['title']}</li>
        <li>Format: {data['album']['format']}</li>
        <li><b>Amount: {data['amount']}</b></li>
    </ul>
    </body>
    </html>
    """
    try:
        response = ses.send_email(
            Destination={"ToAddresses": [data["email"]]},
            Message={
                "Body": {
                    "Html": {"Charset": CHARSET, "Data": BODY_HTML},
                    "Text": {"Charset": CHARSET, "Data": BODY_TEXT},
                },
                "Subject": {"Charset": CHARSET, "Data": SUBJECT},
            },
            Source=SENDER,
        )
    except ClientError as e:
        print(e.response["Error"]["Message"])
    else:
        print("Email sent! Message ID:"),
        print(response["MessageId"])
```

Create Trigger

* DynamoDB Table
* Batch Size (Largest number of records)
* Batch Window (0 to real time)
* Starting Postion (Latest)
