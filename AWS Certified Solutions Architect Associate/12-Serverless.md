# AWS Serverless by Ryan Kroonenburg

## Lambda

`Data Center -> IAAS -> PAAS -> Containers -> Serverless (FaaS)`

Ultimate Abstraction Layer (Function as a Service)

Compute service. AWS manages and provisions the servers that will run the code.

* Can be used in an event-driven based on other AWS Services
* Compute service in response to HTTP requests (API Gateway Calls)

Trigger Example
`USER -> S3 Bucket -> S3 CreateObject Trigger -> Lambda`

HTTP API Gateway
`USER -> HTTP Request -> API GATEWAY -> Lambda -> API Gateway -> USER`

### Architecture

Traditional:
`USER -> ELB/R53 -> EC2 -> DB`

Serverless: 
`USER -> API Gateway -> Lambda -> DB`

* *Supports many programming languages*
* *First 1 million requests are free. $0.20 per 1 million requests*
  * *Duration pricing -> rounded up to the nearest 100ms* 
* *No Servers*
* *Continuous Scaling -> Scales OUT*
* *Super Cheap*
* *Architectures can get extremely complicated. AWS X-Ray allows you to debug serverless*
* *Lambda can operate globally*
* *Understand triggers*

## Serverless Website Architecture

* Route53 DNS
* Get Static webpage from S3
* Dynamic Content -> API Gateway -> Lambda

Possible Lambda Triggers
* API Gateway
* AWS IoT
* Application Load Balancer
* Cloudwatch Events
* Cloudwatch Logs
* Code Commit
* Cognito Sync Trigger
* DynamoDB
* Kinesis
* S3
* SNS
* SQS

## Alexa Skill

When talking to an Alexa device, a serverless function is executed in order to respond

* Skill Service
  * AWS Lambda
* Skill Interface
  * Invocation Name
  * Intent Schema
  * Slot Type
  * Utterances

### Creating

* S3 Bucket
* Add Text file to it
* Amazon Polly schedule a conversion to mp3
* Create Lambda from a template
* Create Alexa Skill on AWS developer
* Update Lambda template to read audio

## SAM (Serverless Application Model)

CloudFormation extension optimized for serverless applications

* Supports anything CloudFormation supports
* Run locally on Docker
* Package and deploy using CodeDeploy

```sh
sam init # to create a template
sam build
sam deploy --guided # for interview like deploy
```

## Elastic Container Service (ECS)

Docker Containers
* Isolation without overhead of VMs
* Containerized applications are portable and offer consistent environment

* Docker Image: `App + App Runtime + Libs`
* Docker Engine (Docker Runtime)
* Host OS
* Infrastructure

ECS is a managed container **ORCHESTRATION SERVICE**
* Create clusters to manage fleets of container deployments
* ECS manages EC2 or FARGATE instances
* Schedules containers for optimal placement
* Defines rules for CPU and memory requirements
* Monitors resource utilization
* Deploy, update, roll back containers
* FREE **FOR REAL**
* Integrates with VPC, SG, EBS volumes
* ELBs
* CloutTrail, CloudWatch

### Cluster

Logical collection of ECS resources (EC2 or Fargate instances)

### Task Definition

Defines an application.

Similar to a DOCKERFILE but for running containers in ECS.

### Container Definition

Inside a task definition, it defined the individual containers a task uses.
Controls CPU and Memorry allocations and port mappings

### Task

Single running copy of any containers defined by a task definition.

One working copy of an application

### Service

Allows task definitions to be scaled by adding tasks.

Defines min and max values

### Registry

Storage for container images (Elastic Container Registry, DockerHub)

Used to download images to create containers

### Fargate

**Serverless** compute engine for containers
* Eliminates need to provision and manage servers
* Specify and pay for resources per application
* Works with both ECS and EKS
* Each workload runs in its own **kernel**
* **Isolation** and Security
* *USE EC2 instead of Fargate when: Compliance Requirements, Require broader customization, GPU required application*

### EKS

Elastic Kubernetes Services

* K8s is an open-source software that laets you deploy and manage containerized applications at scale.
* Same toolset on-premises and in cloud
* Containers are grouped in **PODS**
* Like **ECS**, supports both **EC2** and **Fargate**
* *Use EKS if: already use K8s, want to migrate to AWS*

### ECR

* Managed docker container registry in AWS.
* Store, manage and deploy images.
* Integrated with ECS and EKS
* Works with on-premises deployments
* Highly available
* Integrated with IAM
* Pay for storage and data transfer (similar to S3)

### ECS + ELB

* Distribute traffic evenly across tasks in your service
* Supports ALB, NLB, CLB
* Use NLB or CLB to route TCP
* Use ALB to route HTTP(S)
  * Dynamic host port mapping
  * path-based routing
  * Priority rules
  * ALB is **recommended**

### ECS Security

* Instance Roles
  * Applied policy to all tasks running on that EC2 instance
* Task Roles
  * Task based roles (per task basis)



