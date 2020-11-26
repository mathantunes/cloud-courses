# Practice Tests

* What is the maximum response time for a Business Level 'production down' Support Case?
    * 1 Hour
* Virtual-host-style URLs are the recommended way to access S3 buckets
* DynamoDB allows for the storage of large text and binary objects, but there is a limit of 400 KB
* The resource owner refers to the AWS account that creates the Amazon S3 resource.

* Dedicated Hosting EC2 - The tenancy of an instance can only be change between variants of ‘dedicated' tenancy hosting. It cannot be changed from or to default tenancy hosting.
* Use the AWS CLI to modify the Instance Placement attribute of each instance and the VPC tenancy attribute of the VPC

* DynamoDB Strong Consistency is more expensive
* There's no way to send delay instructions for Spot interruptions.
* When a new message is added to the SQS queue, it will be hidden from consumer instances for a fixed period.



* Amazon Kinesis Data Streams for real-time events with a partition key for each device. Use Amazon Kinesis Data Firehose to save data to Amazon S3

* API in one VPC and Consumer in another:
  * Configure a VPC peering connection between the two VPCs. Access the API using the private address
  * Configure **PRIVATE LINK** connection for the API into the client VPC. Access the API using the PrivateLink address


* Use an Elastic Load Balancer, a multi-AZ deployment of an Auto-Scaling group of EC2 Spot instances (primary) running in tandem with an Auto-Scaling group of EC2 On-demand instances (secondary), DynamoDB.



## hybrid **environment**

* The main requirements to drive this selection are overall **cost** considerations and the ability to reuse existing internet connections
  * AWS VPN (Internet Based)
  * WRONG - AWS DIRECT CONNECT IS NOT INTERNET BASED

* What can you do to stop these instances when they are idle for long periods?
  * CloudWatch alarm that is triggered when the average **CPU utilization percentage has been lower than 10 percent for 4 hours**, and stops the instance

* The warehouse is **50TB**
  * AWS Snowball **Edge**

* AWS service can help you optimize your AWS environment by giving recommendations to reduce cost
  * **TRUSTED ADVISOR**

* Sell the reserved instances on the **Reserved Instance Marketplace**.

* Bastion Setup `Provision -> SecurityGroup Ingress on 22 -> IPs allowed -> Private Key -> Internet Gateway -> Route Table -> Route on Route Table`

## NACLs

* The **default** configuration of the default NACL is **Allow**, and the default configuration of a **custom** NACL is **Deny**
* What is true about the default network ACL?
  * You can **add** or **remove** rules from the default network ACL

* You can assign up to **five** **security** **groups** to the **instance**

* What can be used to provide the subcontractors with short-lived access tokens that act as temporary security credentials to the company AWS account?
  * AWS STS AWS **Security Token Service**

* which AWS service provides integration with **Puppet**
  * AWS OpsWorks

* You can only have **20 EC2 instances per region**. -> Not a hard limit, talk to AWS to higher
* vCPU-based on-demand instance limit per region

* You want to attempt a warm attach. What does this mean?
  * Attach the ENI to an instance when it is **stopped**.
  * HOT ATTACH IS RUNNING
  * COLD ATTACH IS WHEN CREATING

* You suspect that one of the AWS services your company is using has gone down. How can you check on the status of this service?
  * AWS Personal Health Dashboard

## Auto Scaling

* Reuse some software licenses and therefore need to use dedicated hosts on EC2 instances in your Auto Scaling Groups. What step must you take to meet this requirement?
  * Create the Dedicated Host EC2 instances, then add them to an existing Auto Scaling Group.
* Which AWS service can meet this need by exporting data from DynamoDB and importing data into DynamoDB?
  * Elastic Map Reduce
* What is an appropriate metric for auto scaling with SQS?
  * backlog per instance

## Disaster Recovery Plans

* Backup and restore (RPO in hours, RTO in 24 hours or less)
  * Restore when necessary.
* Pilot Light (RPO in minutes, RTO in hours)
  * Maintain a minimal version of an environment always running the most critical core elements of system
  * Rapid provision a full-scale production environment around the critical core
* Warm Standby (RPO in seconds, RTO in minutes)
  * Maintain a scaled-down version of a fully functional environment always running
  * Critical systems are fully duplicated always on but scaled down
* Multi-region active-active (RPO is none or seconds, RTO in secods)
  * Workload deployed to actively serve traffic from multiple AWS Regions
  * Requires sync users and data across the Regions
  * Direct traffic to healthy zones