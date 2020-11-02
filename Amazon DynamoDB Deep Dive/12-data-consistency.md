# Data Consistency and Management by Mark Richman

## Conditional and Update Expression

To protect agains write conflicts.

Conditional updates allow the put/update/delete operation to continue if a condition is met:

* Check attribute value
* Check if attribute exists

**Atomic counter** can be used to increment or decrement the value of an existing attribute without interefing with other write requests

```python
table.update_item(
    Key={"id": 123}, # PARTITION KEY
    UpdateExpression="set price=:p", # UPDATE EXPRESSION WITH PLACEHOLDER :P
    ExpressionAttributeValues={":p": decimal.Decimal(1.5)}, # REPLACED BY 1.5
    ConditionExpression=Attr("price").not_exists(), # EVALUATE TO TRUE TO CONTINUE
    ReturnValues="UPDATED_NEW",
)
```

**Atomic Counter**

```python
table.update_item(
    Key={"id": 123}, # PARTITION KEY
    UpdateExpression="set in_stock = in_stock - :val", # UPDATE EXPRESSION WITH PLACEHOLDER :val
    ExpressionAttributeValues={":val": 1}, # REPLACED BY 1
    # No Condition expression
    ReturnValues="UPDATED_NEW",
)
```

**Conditional Writes consume WCU according to the size of the item, even if it fails**

## ACID Transactions

**Atomic Writes and Isolated Reads**

### TransactWriteItems

* Up to 25 items or 4MB
* Evaluate Conditions
* Perform only if all conditions are simultaneously **true**

### TransactGetItems

* Up to 25 items or 4MB
* Return a consistent isolated snapshot of all items

It prepares and commits transactions (double WCU or RCU)

### Constraints

* Same Partition key or across partition keys
* Same Table or across tables
* Same region **ONLY**
* Same account **ONLY**
* DynamoDB tables **ONLY**
* Unlimited concurrent transactions
* Low latency
* Scales horizontally
* Design to avoid hot keys

### Failures

* Precondition failure
* Insufficient capacity (throttling)
* Transactional conflicts
* Transaction still in progress
* Service error (unlikely)
* Malformed request
* Permission

```python
# Orders table
# Albums table

client.transact_write_items(
    TransactItems=[
        {
            "Update": { # Decrement the album table stock value
                "TableName": "album",
                "Key": {
                    "id": {
                        "n":"123"
                    }
                },
                "UpdateExpression": "set in_stock = in_stock - :val",
                "ConditionExpression": "in_stock > :zero",
                "ExpressionAttributeValues": {
                    ":val": { "N": "1" },
                    ":zero": { "N": "0" },
                }
            }
        },
        {
            "Put": { # Generate an order
                "TableName": "orders",
                "Item": {
                    "id": { "S": "123123123" },
                    "amount": { "N": "123.12" },
                    # more data for order
                }
            }
        }
    ]
)
```

## TTL Time to Live

Reduce storage costs to relevant items only.

Does not compute WCU/RCU.

When enabled, a background job checks for the ttl attribute to verify if the item is expired or not.

**Eventually** removed from GSI and LSI too

i.e.

session table

```json
{
    "id": 1,
    "ttl": 1571415073
}
```

Creating:

* Manage TTL
    * TTL Attribute (defined as ttl on example) 
    * Preview TTL (see if logic checks out)

Use FilterExpression so only unexpired items are fetched on Scan

**EPOCH TIME FORMAT ONLY**
