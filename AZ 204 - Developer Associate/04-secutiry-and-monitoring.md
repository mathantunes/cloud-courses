# Security

Exam -> What to do for authentication and authorization using Azure.

Authentication -> Verify identity of user or service
Authorization -> What this service or user can access, once it's validated

## Authentication and Authorization

### OAuth2

Industry standard for authentication (HTTP)

User attempts to open web app -> Redirect to Azure AD -> Gets token -> Passed to Web App

**Must register app in Azure AD**
* Enterprise application section on Azure AD
* Choose predefined registration from library
* Returns:
  * Application ID -> identifies app
  * Redirect URI -> direct responses back to service
  * Scenario URI -> links for specific (terms of service, gtc, gdpr...)

* Token is refreshed once in a while

#### Roles

Resource Owner
* Request service from app
* Owns data, assigns access to resources

Resource Server
* Service -> Where data relies
* Trusts authorization service to give access

Authorization Server
* Identity provider
* Handles auth

### Shared access signatures (SAS)

These are URIs that grant secure access to Azure Storage resources.

One can provide a SAS URI to someone who needs access to specific resources

* Levels of access
  * Account level -> root level access
  * Service level -> delegates access to only one service within an account: Blob, Queue, Table..
  * User level -> only **BLOB** using **Azure AD**, can be used to access specific or blob containers

**SNAPSHOT NOT SUPORTED**

* App will "craft" the URIs dynamically

### Azure Active Directory

### Role based access controls (RBAC)

`Coarse grain access`

In essence, read, write or both to resources inside of a storage account

**RBAC uses role assignments to give permissions to Security Principals**

A **Security Principal** is an object representing user, group, service principal or managed identity that exists within **Azure AD**

**Examples of roles to Security Principal:**
* Storage Blob Data Owner -> Full access, modify ACLs
* Storage Blob Contributor -> Read, Write, Delete access -> Modifies permissions on what it owns
* Storage Blob Reader -> Read allowed resources

#### ACLs

Access Control Lists -> used to apply `Finer grain access to Storage accounts`

Access ACLs:
  * Provides access to an object. Files and directories both have it
Default ACLs:
  * Parent template for access to child **objects and directories**. Files **DO NOT** have default ACLs.

Permissions:
  * Read (R) -> Can read file / Needs Execute to read from directory (list)
  * Write (W) -> Can write or append to file / Need write and execute to create child items in directory
  * Execute (X) -> N/A / Required to traverse the child items of a directory 

## Centralizing Security -> App Configuration

Centralize configuration, settings and security for applications to use.

Scale, settings across all instances in one place.

Can point to anything: Webapps, functions, k8, vms...

Series of Key/Value pairs to replace configuration files
Point applications to use a certain key on app configuration

* Use `:` or `/` for hierarchical keys (`magic:earth:color`)
  * Although there is no parsing done by app config
  * Case sensitive

* Values: unicode strings
* Labels: used for variant keys (**Why???**)
  * This means we can have 2 keys `a:b` and then label it `c` and `d` to differenciate

### Connecting app to App Configuration

1. Connect application to app configuration **store**
1. Configure app to use the key/values

**PRO TIP: Feature flags -> to enable/disable features from one place**

### Security

256bit AES encrypted by Azure

Can use access keys for authentication and authorization
* Generated when created app configuration -> Read/Write or ReadOnly keys
* Rotate on console

Can leverage Azure AD to create **RBAC**
* Owner -> Administrative
* Data Reader -> Can only view
* Contributor -> Read Write
* Reader -> Also a reader?

### Tiers

Free Tier -> Pretty much for test
* 1 resource allowed
* 10 MB of storage
* 1000 requests per day
* No SLA (**LOL**)

Standard Tier -> Prod tier
* Unlimited resources
* 1 GB storage
* 20k requests/h
* 99.9% availability SLA
* 1,2 USD / day

## Azure Key Vault

Store secrets, certificates and keys for use by applications and users.

Can be auditted

* Securely access
* Access keys, SAS tokens
* Assign access to whom needs it

Can be used in conjunction with **App Configuration**
* Secure app configuration values -> Use KeyVaultAPI

**CAN SOFT DELETE RESOURCES -> Retained for 90 days**

### API Interaction

* Azure CLI
  * `az keyvault create --name "necro" --resource-group "magic" --location "schweiz"`
  * `az keyvault secret set --vault-name "necro" --name "secret" --value "super"`
  * access by `https://.vault.azure.net/secrets/{NAME}`
    * Can configure access permissions
  * `az keyvault secret show --name "necro" --vault-name "necro" --query "value"`
* Powershell
* Rest
* Resource Manager
* .NET
* Other SDK are available but not covered

### Logical Vault

General Standard

* Secret management
* Key management
* Certificate management

### Hardware security module

Only good for keys

## Graphs and Gremlins

Microsoft graphapi is a RESTful web API that allows to access resources on Office 365, Windows 10 user data and some other weird shit

* Exam focus on **Azure AD**
  * Includes applications users are connected to
  * What do they do
  * Wants to create analytics reports and predictions

**USES OData to query LOOOOOOOL**

## Managed Identities

Allow developers to create accounts in Azure AD for when users may need to authenticate for some token.

Instead of storing a hash somewhere, use azure managed account specific to the application to gain access.

Create with `Azure CLI`
* System
  * `az appconfig identity assign --name "service" --resource-group "one"`
* User
  * `az appconfig identity assign --name "user" --resource-group "one" --identities /subscription/1 Microsoft.ManagedIdentity/userAssignIdentities/user1`

```C#
AzureServiceTokenProvider.GetAccessTokenAsync()
```

### System Assigned Identity

* Identity that is part of the configuration store. 
* If you remove the configuration store, this identity goes with it.
* There can only be **ONE** system-assigned identity

* Created as part of a resource (VMs, App Service and so on)
* Shared life cycle with the application -> App down, managed identity is also down
* Cannot be shared
* Used for:
  * Workloads that are contained within a single Azure resource
  * Workloads for which you need independent identities

### User-Assigned Identity

* Independent resources
  * Users or AD Groups
* Configuration stores can have multiple
* Independent life cycle (must cleanup by yourself)
* Used for:
  * Workloads that runs in multiple resources and can share a single identity
  * Shared storage amongst VMs or so

# Monitoring

## Azure Monitor

Consumes data from multiple resources
* Applications, OS, Resources, Subscriptions, Tenants (AD)
* Custom data from **REST Calls**

Can integrate and consume data from Azure onitor in many ways
* Insights, PowerBI and some others
* Alerts
  * Define scope (VM or so)
  * Conditions -> Performance / Level
  * Actions -> Notifying, log or some other stuff

### Log Data

Multiple properties, may vary in size.

Traces performance events, errors and so on

Used in log analytics
* Kusto Query Language is used in this (KQL)

### Metric Data

Numeric values that show some aspect of a resource at a point in time
Consumes very little storage

## Application Insights

**Application instrumenting**

Powerful tool to help with cloud application data.

Insights on individual applications.

* Request rates, response times, failure rates
* Dependency rates
* Exceptions
* Page views, load performance,
* AJAX Calls
* User and session counts
* Performance counters
* Host diagnostics
* Diagnostic traces
* Custom events

Can be setup @ **Runtime** when using Web App Services or IIS on VM or on-prem

* Can add custom event tracking
  * Use instrumentationkey (endpoint)

### Features

* Users / Sessions / Events
  * Most interesting pages
  * User location
  * Browser, Device
* Funnels
  * Check progression on purchasing
* Cohorts
  * Set of users, sessions or events that have common
* Impact
  * Load time and other properties influence in convertion
  * Location and convertion
* Retention
  * User return to app
  * Frequency of using a certain feature
* User Flows
  * Visualize how users navigate
  * What do they click
  * Where do they churn
  * Where do they repeat actions

### Availability Alerts

* URL Ping Test
  * Simple ping test that can be configured from console
* Multi-Step webtest (ONLY ON **VS ENTERPRISE**)
  * Recording of webrequests in an ordered sequence to test complex scenarios
* Custom Tracking
  * Requires calls to `TrackAvailability` method on code
