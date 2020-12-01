# Serverless

## Lambda

AWS provisions and manages the servers to run lambda

* Event driven compute service ran in response to events
  * Node, Java, Python, C#, Go, Ruby
  * 1 million requests are free. $0.2 per 1 million requests thereafter
  * Duration is calculated from the time the code begins executing until it returns, rounded up to the nearest 100ms

## Version Control with Lambda

When using versioning, you can publish one or more versions of your lambda function.

* Each lambda function version has a unique ARN.
* Each version is immutable
* Only the $LATEST can be mutated

* Qualified ARN -> The function ARN with the version suffix
* Unqualified ARN -> The function ARN without the version suffix

* Alias -> Name that points to a certain version of the lambda
  * i.e. PROD alias uses version 2, it is possible to promote version 3 to PROD

### Using

Access function Qualifiers.

* Publish as a version (version 1)
  * ARN:1
  * Becomes immutable
* Access $LATEST version
  * Publish as a new version (version 2)
* Create Alias
  * Version PROD alias
  * Points to version 1
* Create splits between traffic
  * New Alias
  * Version 3 (75%)
  * Version 2 (25%)

## Step Functions

Allows to visualize and test serverless applications. It provides a graphical console to arrange and visualize the components of the applications as a series of steps.

* Sequential Steps
* Branching Steps
* Parallel Steps

### Demo

* Access Step Functions (Application Integration section)
* Create State Machine
  * Sample Projects (Job Poller)
  * Cloudformation Stack with Lambda and Batch

## X-Ray

Collects data about the requests that the application serves, and provides tools you can use to view, filter and gain insights into that data to identify issues and opportunities for optimization.

`X-Ray SDK -> X-Ray Daemon -> X-Ray API -> X-Ray Console`

* X-Ray SDK
  * Interceptor to HTTP requests
  * Client handlers to instrument
* Integration
  * ELB
  * Lambda
  * API Gateway
  * ELC
  * Elastic Beanstalk

### Demo

* Deploy elastic beanstalk java web app
  * Generate sample traffic
  * Access X-Ray console
    * Service map showing requests

## API Gateway Advanced

* Use the API Gateway Import API feature to import an API from an external definition file into API Gateway (Swagger/OpenAPI)
 
* Access API Gateway (Network and Content Delivery section)
  * Import from Swagger file

### API Throttling
 
* Limits the steady-state request reate to 10k requests per second
* Maximum concurrent requests is 5000 requests across all APIs within an AWS Account
* 429 too many requests response
* SOAP Webservice Passthrough
  * [REF](https://www.rubix.nl/blogs/how-configure-amazon-api-gateway-soap-webservice-passthrough-minutes)