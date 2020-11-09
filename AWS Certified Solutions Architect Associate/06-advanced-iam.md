# AWS Advanced IAM by Mark Richman

## AWS Directory Service

* Family of managed services
* Connect AWS Resources with on-premises Active Directory
* Standalone directory
* Use existing corporate credentials
* SSO to any domain-joined EC2 instance

*Active Directory is an on-premises directory service with a hierarchical database of users and groups*

### AWS Managed Microsoft AD
**AD Compatible**

AD Domain controllers running Windows Server

* Reachable by applications in VPC
* Add DCs for HA and performance
* Exclusive access to DCs
* Extend existing AD to on-premises using **AD Trust**

AWS MANAGES:
* Multi AZ deployment
* Patch, monitor, recover
* Instance Rotation
* Snapshot and restore

Customer Manages:
* Users, Groups and Group Policies
* Standard AD Tools
* Scale out DCs
* Trusts (resource forest)
* Certificate Authorities (LDAPS)
* Federation

### Simple AD
**AD Compatible**

* Standalone managed directory
* Supports basic AD features
* Easier to manage (EC2)
* Linux workloads that need LDAP
* Does not support trusts
* Size
  * Small <= 500 users
  * Large >= 5000 users

### AD Connector
**AD Compatible**

* Directory gateway (proxy) for on-premises AD
* Avoid caching information in the cloud
* Allow on-premises users to log in to AWS using AD
* Join EC2 instances to existing AD domain
* Scale across multiple AD Connectors

### Cloud Directory
**NOT AD Compatible**

* Directory-based store for **Developers**
* Multiple hierarchies with hundreds of millions of objects
* Fully Managed Service

* Use Cases
  * Organizational Charts
  * Course Catalogs,...

### Cognito User Pools
**NOT AD Compatible**

Managed user directory fos SaaS applications
* Sign-up and sign-in for web or mobile
* Works with **SOCIAL MEDIA** creds

## IAM Policies

### Amazon Resource Name (ARN)

Uniquely identifies resources in AWS

`arn:PARTITION:SERVICE:REGION:ACCOUNTID:...`

`arn:aws:ec2::0123456789012:user/mark` -> (no region) Resource/ Qualifier

`arn:aws:ec2:us-east-1:0123456789012:instance/*` -> all instances from account/region

### Policies

* JSON document that defines permissions
* Structured as a list of statements

**Implicitly Denied, if not Explicitly ALLOWED**

*Explicit denies overrides explicit allows when multiple policies*

#### Statement Structure

* Sid: Human Readable Name
* Effect: Allow or Deny
* Action: array of string with "service:APICall" `dynamodb:Get*`
* Resource: resource the action is against `dynamodb table/mytable`

#### Identity Policies

* Permissions

#### Resource Policies

* Attached to AWS resources
* Access and actions allowed

### Permission Boundaries

* Delegate administration to toher users
* Prevent **PRIVILEGE ESCALATION** OR **UNNECESSARILY BROAD PERMISSIONS**
* Control Maximum Permissions an IAM Policy can grant

Use Cases:
* Developers creating roles for Lambda Functions
* Application owners creating roles for EC2 instances
* Admins creating AD HOC users

## Resource Access Manager (RAM)

* Resource Sharing between accounts

Currently available for 
* App Mesh
* Aurora
* CodeBuild
* EC2
* EC2 Image Builder
* License Manager
* Resource Groups
* Route 53

### Example

Launch EC2 instance in a shared subnet (between 2 accounts)

## Single Sign-On (SSO)

* Centrally manage access to AWS accounts and business applications
  * Office 365
  * GitHub
  * G Suite
  * ...

### Granular Account-Level Permissions

`Third Party Login -> SSO -> Users`

**SAML stands for Security Assertion Markup Language**

*If SAML, look for SSO as an answer*

## Advanced IAM Summary

* Active Directory

Connect AWS Resources with on-premises AD

* SSO to any domain-joined EC2 instance
* AWS Managed Microsoft AD
* AD Trust
* AWS vs. Customer Responsibilities
* Simple AD (basic AD without trusts)
* AD Connector (directory gateway)
* Cloud Directory (developers and hierarchical data)
* Cognito User Pool (social media identities)


* ARN
* IAM Policy Structure
  * Effect/Action/Resource
* Identity vs Resource Policies
* Policy Evaluation Logic
  * Deny overrides Allow
* AWS Managed vs. Customer managed
* Permission boundaries
  * Define maximum permissions an identity can have
* Resource Access Manager
  * Resource sharing between accounts
* Single sign-on
  * Centrally manage access
  * Office 365
  * Use existing identities
  * Account-level permissions
  * SAML    