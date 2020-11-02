# Advanced Security by Mark Richman

## Federated and Fine-Grained Access Control

DynamoDB configuration to access control.

No need to trust application logic for access control.

### Social Identity Providers

Represented by a token provided from the Social Identity Providers.

AWS STS provides AWS Credentials from the token tied to a particular **ROLE**

This role has access to only certain actions (defined to the **USER**)

The Credentials are evaluated through **Policies from IAM**

* Possible to control the Primary key which the Role is allowed to access

```json
{
  "Version": "2012-10-17",
  "Statement": [
      {
          "Sid": "FullAccessToUserItems",
          "Effect": "Allow",
          "Action": [
              "dynamodb:GetItem",
              "dynamodb:BatchGetItem*",
              "dynamodb:Query",
              "dynamodb:PutItem",
              "dynamodb:UpdateItem",
              "dynamodb:DeleteItem",
              "dynamodb:BatchWriteItem"
          ],
          "Resource": [ "arn:aws:dynamodb:us-east-1:ACCOUNT:TABLE/MYTABLE" ],
          "Condition": {
              "ForAllValues:StringEquals": {
                  "dynamodb:LeadingKeys": [
                      "${www.amazon.com:user_id}" // Federated User ID
                  ],
                  "dynamodb.Attributes": [ // Only these attributes (Not Filters)
                      "userid",
                      "artist_name",
                      "album_title",
                  ]
              }
          }
      },
  ]
}
```

**STS API Has an assume role function to get keys and secrets to interact with AWS from a certain ROLE**

## Auditing Admin Access Using CloudTrail

Primary services for logging services.

Logs any AWS action performed (90 Days).

Filters:
* Event Source (i.e. DynamoDB)
* Event Name

Create a Trail for more specific logs
* Trail Name: DynamoDB Trail
* Apply to all regions?
* Management Events (ALL, READ-ONLY, WRITE-ONLY)
* S3 Storage Location