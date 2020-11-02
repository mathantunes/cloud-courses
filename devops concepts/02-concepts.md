# DevOps Concepts by Nick Triantafillou

## Key terms

### Agile

Iterative software development approach

**DevOps is considered agile beyond software development team**

### AWS

Amazon Web Services

### Container

Standalone, executable piece of software

### Continuous delivery CD

Ongoing release of software to production

### Continuous integration CI

Ongoing programming, building and testing of code

### Infrastructure as Code

Defining infrastructure with programming code

### Microservices

Application architecture broken in smaller pieces

### OpenSource

Free license and software code released

### Pipeline

Set of connected processes where the output is the input for the next

### Serverless

Running services without having to deploy infrasctructure

### Source Code Repository

Upload and track the history of a code base

### Unit Testing

Breaking your application down into small parts to test that each feature works

## CI/CD

### Benefits

* Builds faster
* Decrease code review time
* Automatic test
* Faster fault isolation
* Additional deployment features

### Exemple AWS

**AWS CodeCommit (like github)** Save code as a repo
triggers a deploy stage on **AWS CodeDeploy** on a running container on **EC2**.

The whole pipeline is managed on **AWS CodePipeline**.

## DevOps Tools

### Git

* Free and open-source version control system
* Stores the entire history of your code

### AWS CodeCommit

* Fully managed source control service for Git
* Secure and encrypted
* Highly available
* Easy integration with other AWS Services
  
### AWS CodeBuild

* Fully managed continuous integration service
* Compiles code, runs tests, and produces software packages
* Scales automatically and can process multiple builds at the same time (concurrently)

### AWS CodeDeploy

* Fully managed deployment service
* Can deploy to both AWS and your on-premises servers
* Completely automates software deployments
* Eliminates error-prone manual operations
  
### AWS CodePipeline

* Fully managed continuous delivery service
* Helps you automate release pipelines
* Can automate the build, test and deploy phases
* Integrates with both CodeCommit and Github
