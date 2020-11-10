# AWS Route 53 by Mark Richman

## DNS 101

**Domain Name System**

Convert human friendly domain name to IP addresses.

### IPv4

32 bit field and has over 4 billion different addresses

### IPv6

Created to solve the depletion issue and has an address space of 128 bits - 340 undecillion addresses

### Top Level Domains

* .com
* .edu
* .gov
* .co.**uk**
* .ch 

Top level domain names are controlled by the IANA in a root zone database.

### Domain Registrars

Authority that can assign domain names

ICANN forces uniqueness -> Registers domain names on WhoIS

### Start Of Authority Record (SOA)

Record stores information aboutÇ

* Name of the server that supplied the data for the zone
* The administrator of the zone
* Current version of the data file
* Default number of seconds for the **TTL** file on resource records

### NS Records (Name Server Records)

Used by top level domain servers to direct traffic to the Content DNS Server which contains the authoritative DNS records

`USR Calls a DNS -> Browser searches the TOP LEVEL SERVER for the DNS -> Gets an NS endpoint -> NS Record -> SOA -> DNS Record`

#### A Record

Fundamental type of DNS Record

Translates to an IP address.

#### C Name

Canonical Name can be used to resolve one domain name to another

`google.com -> mobile version m.google.com`

#### Alias Records

**Route53 specific**

Used to map resource record sets in the hosted zone to *Elastic Load Balancers*, *CloudFront distributions* or *S3 buckets* that are configured as websites

Alias records work like a **CNAME** record by mapping one DNS to another *target* DNS name

*Key difference - A CNAME can't be used for **NAKED** domain names*

### TTL

The length that a DNS record is cached on either the Resolving Server or user PC.

The lower the time to live, the faster changes to DNS records take to propagate throughout the internet

### Exam Tips

* *ELBs do not have pre-defined IPv3 addresses; Resolve to them using a DNS name*
* *Understand the different betwqeen Alias Record and CNAME -> **ALWAYS PICK ALIAS RECORD***
* DNS Types
  * *SOA Records*
  * *NS Records*
  * *A Records*
  * *CNAMES*
  * *MX Records* 
  * *PTR Records*

## Register a domain name - LAB

Networking & Content Delivery -> Route53

* Domain Registration
  * Pick a domain name (`hellocloudgurus2019.com`)
  * Registrant Contact (`Personal Information`) 
  * Takes up to 3 days to register (`normally 2 hours is enough`)

* After registration, DNS is on Hosted Zone
* Provision EC2 instances

### Exam Tips

* *Buy domain names directly with AWS*
* *Take up to 3 days to register depending on the circumstances*

## Route53 Routing Policies

* Go to Hosteed Zone
* Create Record Set
  * Name Record Set
  * Type A
  * Values: IP Servers
  ```text
  IP1
  IP2
  IP3
  ```
  * Lower TTL to 1 minute (so browser DNS cache is expired in 1 minute)

### Simple

Only have one record with multiple IP addresses.

Specify multiple values in a record, Route53 returns all values to the user in a **random order**

*Simplest Routing*

*One record with one or multiple IPs. If multiple, Route 53 responds randomly*

### Weighted

Split traffic based on different weights assigned.

Create separate *Record Sets* for each IP and Specify the Weight.

`10% RegionX 90% RegionY`

Route53 responds an IP based on the weights

*Option to configure health check for Route53 to ignore IPs of Bad Instances*

### Latency-Based

Route traffic based on the lowest network latency for end user

Create a latency resource record set in each region.

Route53 selects the lowest latency IP for the USER.

### Failover

when creating an active/passive set up.

Route53 will monitor the primary IP with health check. If it's down, routes to passive

### Geolocation

Choose where traffic will be sent based on geographic location of USERS

`European users -> EC2 Instance in EU`

### Geoproximity

Route traffic to resources based on the geographic location of users and resources.

Optionally choose to route more traffic or less to a given resource by specifying a value, known as **bias**. A bias expands or shinks the size of the geographic region from which traffic is routed to a resource

**ONLY AVAILABLE ON ROUTE53 TRAFFIC FLOW**

### Multivalue Answer

Route53 to return multiple values, such as IP addresses. Specify multiple values for almost any record.

It allows health check for each record set.

`Basically simple routing with health check`

*Multiple record sets with health checks, fails over to the next IP*


