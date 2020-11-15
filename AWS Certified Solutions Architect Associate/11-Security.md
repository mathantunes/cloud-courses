# AWS Security by Mark Richman

## Reduce Security Threats

* NACLs to deny inbound from certain IPs
* Host-based firewall
  * Another layer of defense
  * Linux/Windows packages on EC2 Instances
* WAF (Web Application Firewall)
  * Attached to **ALB and CloudFront**
  * IP Blocking and Filtering
  * Block SQL Injection
  * XSS

ALB traffic terminates on Load Balancer and therefore the best way to block is with ACL

CloudFront passes in its IP and not the User IP to ALB, therefore the best way to secure is with a **WAF**

## Key Management Service (KMS)

Regional secure key management and encryption and decryption

Manage Customer Master Keys (CMKs)

* Ideal for S3 Objects
* DB Passwords
* API Keys
* Up to 4 KB in size
* Pay per API Calls
* Audit capability using CloudTrail
* **FIPS 140-2 Level 2**

### AWS Managed CMK

Free; Used by default if you pick encryption in most services.
Track usage
Lifecycle and permissions are AWS Managed

### Customer Managed CMK

Configure Lifecycle and permissions for usage

Key rotation is important

### AWS Owned CMK

Used by AWS on a shared basis across many account.
Can't view or audit these

### Symmetric

Same key for Encryption and Decryption

**AES-256**

* Never leaves AWS unencrypted
* Must call the KMS APIs to use
* AWS Services integrated with KMS use symmetric CMKs
* Encrypt, decrypt and re-encrypt data
* Generate data key, key pairs and random byte strings
* Import own key material

### Asymmetric

* Mathematically related public/private key pair

**RSA and Elliptic-curve cryptography**

* Private key never leaves AWS unencrypted
* Must call KMS APIs to use private key
* Download public key and use outside AWS

### Key Policies

Creating a CMK, it is possible to provide a Key Policy.

Defines the actions allowed to take on KMS API

### CLI

```sh
# aws/s3 is an alias for KMS created and managed by AWS to secure S3 bucket
aws kms create-key --description "test CMK"
# returns KeyID
aws kms create-alias --target-key-id KEYID --alias-name "alias/acgdemo"
aws kms list-keys
# alias will be displayed
# key policy to the root account full access
echo "secret message" > secret.txt
aws kms encrypt --key-id "alias/acgdemo" --plaintext file://secret.txt --output text --query CiphertextBlob | base64 --decode > encrypted.txt # B64 encoded data
aws kms decrypt --ciphertext-blob fileb://encrypted.txt --output text --query Plaintext | base64 --decode # does not neet a key-id parameter

## For data larger than 4KB
aws kms generate-data-key --key-id "alias/datakeydemo" --key-spec AES_256 # store the CiphertextBlob
```

## CloudHSM (Hardware Security Modules)

* Dedicated hardware security module solution
* **FIPS 140-2 Level 3**
* Manage your own keys
* Single tenant, dedicated harwade, multi AZ cluster
* No access to the AWS-managed component
* Runs within a VPC within an account
* Industry-standard APIs - **NO AWS APIs**
  * **PKCS#11**
  * **JCE**
  * **CNG**
* Keys are **IRRETRIEVABLE**

`AWS CLOUDHSM VPC Cluster -> Project ENIs in a chosen VPC -> Communicate with App`

* *Regulatory compliance requirements*
* *FIPS 140-2 Level 3*

## Systems Manager Parameter Store

* Component of AWS Systems Manager (SSM)
* Serverless
* Store Configuration and Secrets
  * Passwords
  * Database Connection Strings
  * License codes
  * API Keys
* Can be stored encrypted (KMS) or plaintext
* Separate data from source control
* Store parameters in **Hierarchies**
* Track Versions
* Set TTL to expire values

*Hierarchi of up to 15 levels*

/prod/db/mysql/db-string

**Integrates with CloudFormation**

## Secrets Manager

Rotate, Manage and Retrieve credentials

* Similar to Systems Manager Parameter Store
* *Charge per secret stored and per 10,000 API Calls*
* Automatically rotate secrets
* *Generate Random Secrets*

## AWS Shield

Protect against DDoS

* Shield Standard
  * When using WAF with cloudfront or ALB
  * No cost on shield
  * Layer 3 and 4 attacks
  * SYN/UDP floods
  * Reflection Attacks (*Source IP attack is spoofed*)
  * Stopped a 2.3 Tbps DDoS attack for 3 days in Feb/2020 in AWS
* Shield Advanced
  * 3000 per month per organization
  * Enhanced protection for EC2, ELB, CloudFront, Global Accelerator, Route 53
  * Business and Enterprise support customers get 24/7 access to the DDoS response teams
  * DDoS cost protection

## Web Application Firewall (WAF)

Monitor HTTP(S) requests to CloudFront, ALB, or API Gateway
* Control access to content
* Configure filtering rules to allow/deny traffic
  * IPs
  * Query String params
  * SQL Injection
  * XSS
* Behaviors
  * Allow all requests but specified
  * Block all requests but specified
  * Count requests that match properties specified
  * Properties
    * Origin IP
    * Origin Country
    * Request Size
    * Values in headers
    * String in request matching regex patterns

## AWS Firewall Manager

Centrally configure and manage firewall rules across an AWS Organization

* WAF Rules
  * ALB
  * API Gateway
  * CloudFront
* AWS Shield Advanced Protections
  * ALB
  * ELB Classic
  * EIP
  * Cloudfront
* Enable security groups for EC2 and ENIs
