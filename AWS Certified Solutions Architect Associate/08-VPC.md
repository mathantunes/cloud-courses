# AWS VPCs by Mark Richman

## VPC Overview

Virtual data center in the cloud

Virtual Private Cloud lets privision a **logically** isolated section of the AWS Cloud where you can launch AWS resources in a virtual network. You have complete control over virtual networking environment, including selection of IP address range, creation of subnets and configuration of route tables and network gateways.

Customize network configuration for VPC.

Subnets for public and private pieces of the system

Leverage Multiple layers of security including:
* Security Groups
* Network Access Control Lists

*Can extend corporate datacenter with VPC connected to on-prem network*

Inside a region:

**Public Subnet**

`Internet Gateway -> Route Table -> Network ACL -> Public Subnet -> Security Group -> Instance`

**Private Subnet**

`Virtual Private Gateway -> Route Table -> Network ACL -> Private Subnet -> Security Group -> Instance` 

**1 Subnet = 1 AZ**

**Corporate Subnets are usually 10.0.X.0/16 (Largest subnet /16 in VPC)**

### Features

* Launch instances into a subnet
* Assign custom IP address ranges in each subnet
* Route tables between subnets
* Internet gateway and attach to VPC
* Security control
* Instance security groups
* Subnet netwoek access control lists (ACLs)

### Default VPC vs Custom VPC

* Default VPC is user friendly, allowing to immediately deploy instances.
* All subnets have a route out to the internet
* Each EC2 instance has both a public and a private IP address

### VPC Peering

* Allows to connect one VPC with another via a direct network route using private IP addresses.
* Instances behave as if they were on the same private network
* Peer VPC's with other AWS accounts or same account
* Peering is in a STAR CONFIGURATION. 1 Central VPC peers with 4 others **NO TRASNSITIVE PEERING** `Peering is 1 to 1, there is no way to go across a peer to access another VPC`
* Peer across REGIONS

## Create Own Custom VPC

* Your VPCs
* Create VPC
  * IPv4 CIDR Block (10.0.0.0/16) 
  * Tenancy (Default -> Share underlying hardware with other AWS customers)
  * *Route Table is automatically created*
  * *Network ACL is automatically created*
  * *Security Group is automatically created*
* Create Subnets
  * Name Subnet (10.0.1.0 - us-east-1a cloudguru) 
  * Pick VPC (cloudguru)
  * Pick AZ 
  * IPv4 CIDR block (10.0.1.0/24)
  * *Auto-assign public IPv4 is default to **NO***
  * *AWS Reserves 5 IP addresses*
    * *Network Address*
    * *Router Address*
    * *DNS Server Address*
    * *One more for future use*
    * *Broadcast Address* 
  * Modify Auto-assign public IPv4 to YES
* Create Internet Gateway
  * Pick name (IGW cloud guru)
  * Attach to VPC
  * *Only One ATTACH per VPC*
* DO NOT Alter Route Table
  * *Always keep MAIN ROUTE TABLE PRIVATE*  
  * Have a separate route table for public access
* Create Route Table for Public Access
  * Edit ROUTE
    * Add Route
      * Destination: 0.0.0.0/0
      * Target: Internet Gateway (IGW cloud guru)   
  * **ANY SUBNET ASSOCIATED WITH THIS ROUTE TABLE WILL BECOME PUBLIC**
  * Associate Subnet (10.0.1.0 - us-east-1a cloudguru)
    * *It is now dissociated from MAIN ROUTE TABLE* 
* Create EC2 Instance
  * Set VPC and Public Subnet 
  * Set VPC and Private Subnet

  <img src="assets/ips.PNG">

* Private instances should have its own Security Group
  * Inbound Rules
    * HTTP Allow another security group or allow **IP range** (10.0.1.0/24)
* Actions -> Security -> Change Security Group of the Private subnet instance
### Examp Tips

* *When creating a VPC from scratch, a default route table, network ACL and default security group are created*
* *NO SUBNETS OR INTERNET GATEWAYS ARE CREATED BY DEFAULT*
* *AZ names are randomized*
* *Amazon Reserves 5 IP addresses within subnets*
* *Only 1 Internet Gateway per VPC*
* *Security Groups Can't Span VPCs*

## NAT Instances & NAT Gateways

Network Address Translation

Enable Private subnet instances to download software on internet

### NAT Instances

Single EC2 instances

*Legacy*

* Create EC2 Instance
  * Community AMI for NAT
  * Public subnet
  * Must disable Source/Destination checks (acts like a gateway to internet gateway)
* Update MAIN ROUTE TABLE
  * Set Destination 0.0.0.0/0 -> Target NAT instance
* Call update on Private instance after SSH from Public instance

*Single Virtual Machine that can get easily overwhelmed -> Single point of failure*

### NAT Gateway

Highly available gateway that allows private instances to access the internet without being public

Create NAT Gateway -> Pick a subnet
* Update MAIN ROUTE TABLE
  * Destination 0.0.0.0/0 -> Target NAT Gateway 

* *Redundant inside AZ*
* *Preferred by the enterprise*
* *From 5 to 45 Gbps*
* *No need to patch O.S.*
* *Not associated with security groups*
* *Automatically assigned a public ip address*
* *Update MAIN ROOT TABLE*

**If you have resources in multiple AZs sharing one NAT Gateway, if the NAT Gateway AZ is down, resources in other AZ lose internet access**

*Prevent it by creating a NAT Gateway per AZ*


