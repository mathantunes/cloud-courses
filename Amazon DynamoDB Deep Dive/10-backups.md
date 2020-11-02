# Backups by Mark Richman

## Point-in-time Recovery

Helps protect dynamodb tables from accidental writes and deletes.

Maintain incremental backups

Default **DISABLED**

Many configuration are erased:
* Auto Scaling policies
* AWS IAM Policies
* CloudWatch metrics and alarms
* Tags
* Streams
* TTL Settings
* Point-in-time Recovery settings

## On-Demand Backup and Restore

Create a backup manually

After available, can be used to restore the dynamodb table

## Scheduled backups using Lambda

**TODO**