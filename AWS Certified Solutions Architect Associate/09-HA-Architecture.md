# AWS High Availability Architecture by Mark Richman

## Elastic Load Balancers

Designed to balance the network load across instances

### Application Load Balancer

Suited for balancing **HTTP** and **HTTPS** traffic.

Operate at Layer 7 and are **application-aware**.

Allow advanced routing configuration

### Network Load Balancer

Load balancing of TCP traffic

Extreme performance is required.

Operates at Layer 4

Capable of handling millions of requests per second with ultra-low latency

### Classic Load Balancer

**Legacy ELB** for HTTP HTTPS applications.

Cheaper

**If application does not respond, LB responds 504 - GATEWAY TIMEOUT**

*Use X-Forwarded-For header to access User IP on instance*

`USER (IP1) -> ELB (IP2) -> Instance(receives IP2 and IP1 on X-Forwarded-For`

## Load Balancers & Health Checks - LAB

Launch EC2 Instances (2 instances, in different AZs)

Inside EC2 Console
* Access Load Balancer
* Create Load Balancer
  * Classic Load Balancer
  * Pick VPC (default)
  * Internal (it can be a private ELB, we want a public)
  * Port (80)
  * Security Group (default)
  * Health Checks
    * HTTP
    * Port 80
    * Path (index.html)
    * Response timeout
    * Interval (period)
    * Unhealthy threshold (number of bad responses to consider a failure)
    * Healthy threshold (number of good responses to consider healthy)
  * Add Instances
  * Enable Cross-Zone Load Balancing (in multiple AZs CHECK IT)
  * Enable Connection Draining

*We do not assign an IP address. AWS generates a DNS Name*

*NAME-ACCOUNT.AZ.ELB.AMAZONAWS.COM*

accessing the DNS will route traffic to EC2 instances as expected.

### Target Group

Group of instances so that the ELB can use the group to target routing and health checks

* Create Target Group
  * Name
  * Type (Instance, IP or Lambda Function)
  * Protocol (HTTP:80)
  *  VPC
  *  Health Check Settings
     * **Same as before**
  * Add Targets
    * Pick EC2 Instances

* Create Application Load Balancer
  * Name
  * Internet-Facing
  * HTTP:80
  * Every AZ
  * Security Group
  * Pick Target Group

*It allows to create specific rules using header and path to route*

## Advanced Load Balancer Theory

### Sticky Sessions

Classic Load Balancer routes each request independently to the registered EC2 instance with the smallest load.

Sticky sessions allow you to bind user's session to a specific EC2 instance. -> All requests from the user during the session are sent to the same instance

*Can be used in CLB*

*Can be used to Target Group Level in ALB*

### Cross Zone Load Balancing

`User -> R53 (spliting 50%AZ1 50%AZ2) -> Each Region has a CLB -> EC2 instances`

ELBs allowed to send traffic between AZs

### Path Patterns

Create a listener with rules to forward requests based on the URL path.

Path-based routing.

```
i.e.
www.myurl.com routes to AZ1
www.myurl.com/images routes to AZ2
```

## Auto Scaling

Components:
* Groups
  *  Logical Component
  *  Webserver group or application group
* Configuration Templates
  * Launch template configuration for EC2 instances
  * Specify AMI ID, instance type, key pair, sg, block device mapping
* Scaling Options
  * Ways of scaling the Auto Scaling groups
  * A condition (dynamic scaling)
  * Scheduled scaling

Options:
* Maintain current instance levels at all times
  * Minimum of 10 instances all times.
  * Periodic health check (if any is unhealthy, it is terminated and another one is launched)
* Manual
  * Specify the change in the maximum, minimum or desired capacity of the Auto Scaling group
  * Auto Scaling manages the termination/launching of instances
* Scheduled
  * Function of date and time to increase or decrease of instances 
* Demand
  * More popular
  * Define parameters to control the scaling process
  * `i.e. CPU around 50%`
* Predictive Scaling
  * Use previous performance

## Launch Configurations & Auto Scaling Groups - LAB

Access EC2 console

* Create Launch Configuration
  * Configure Basic EC2 settings
* Create Auto Scaling Group
  * Pick Launch Configuration
  * Group Size (start with X instances)
  * VPC
  * Subnets
  * Health Checks
  * Policies
    * Keep at initial size
    * Create Scaling Policies
      * Between X and Y instances
      * Conditions (80% CPU)
      * Instance Warm Up (60s)

## High Availability Architecture

*Always plan for failure*

**Check out Netflix Simian ARMY**

Redundancy is important -> AZ failover, Region failover

* *Design for failure*
* *Multiple AZs and Regions*
* *Multi AZ for disaster recovery in RDS*
* *Read Replicas for RDS*
* *Scaling Out - Auto Scaling groups to add instances*
* *Scaling Up - Increase resources inside EC2 -> T2 to C5*
* *Consider cost element*
* *Understand s3 storage classes*

## Building a fault tolerant word press site

`USER -> R53 -> ELB -> EC2 Instances -> RDS Instances`
`USER -> R53 -> ELB -> EC2 Instances -> S3 -> CloudFront`

EC2 in auto scaling group

EC2 and RDS in multiple AZs

* Create Buckets
* Create CloudFront Distribution
  * Web Distribution
  * Domain Nome (s3 bucket)
* Create Security Groups
  * Allow MySQL
* Create RDS
  * MySQL
* Create Role (to Read from S3)
* Launch EC2 instance with Wordpress and Role
* Create ALB
* Create Target Group
  * Register EC2 instance (only read nodes)
* Create R53
  * Point to ALB 
* Create Launch Configuration
* Create Auto Scaling Group

*Failover RDS -> Reboot with Failover*

Automate with CloudFormation
* Pick the WP blog template (LOL!)

## Elastic Beanstalk - LAB

*Cloudformation is a JSON or YAML IaC*

Beanstalk is a simpler version (GUI)

**Found under COMPUTE**

* Create Elastic Beanstalk application
  * Name
  * Platform (PHP for WordPress)
  * Sample App or Custom Code

It automatically provisioned many AWS resources (SG, S3, EC2)

It allows to edit sample configuration *Allows Auto Scaling*

## High Availability with Bastion Host

* Launch Bastion instances in different AZs 
  * it **requires** a **Network Load Balancer** since it utilizes SSH or RDP
  * Expensive Solution
* Auto Scaling group with a minimum/maximum of 1 Bastion
  * Use an Elastic IP Address 
  * Auto Scaling provisions a new Bastion in another subnet
  * **Downtime for recreating the Bastion Host**

## On-premise Strategies with AWS

* Database Migration Service
  * Move DBs to and from AWS
* Server Migration Service
  * Incremental replication of on-premises server in to AWS
  * Backup, multi-site strategy
* AWS Application Discovery Service
  * Plan migration projects by gathering information about their on-premises data centers
  * Builds utilization/dependency map of the environment
* VM Import/Export
  * Migrate existing applications in to EC2
  * Export AWS VMs
* Download Amazon Linux 2 as an ISO
  * Download image as ISO for VMWare