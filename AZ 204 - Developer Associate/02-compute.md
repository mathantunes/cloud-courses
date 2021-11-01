# Azure compute services

## Virtual machines

Server on Azure infrastructure

Scalable and flexible, but requires more responsibility

**Many different sizes of VMs**
**Many different OS -> MS or Linux**

* Azure console -> Too easy.
* CLI -> Scripted deployment

### Design considerations

* Name of application resources
* Location
* Size of VM
* Max number of replicas
* OS
* VM start configuration
* Related resources

**SLA** of 99.9% if using 1 premium storage for all disks
**SLA** of 99.95% with 2+ VMS within an availability set

### Sizes

* A series -> Entry level -> Test runs and small servers
* GP D Series
  * *B, DSv3, Dasv4, Dav4, Av2, DC, DCv2, Ddsv4*
  * balanced CPU-to-Memory ratio **Low to medium traffic / small to medium DBs**
    * Da -> AMD Epyc processor
    * D -> Intel Xeon
    * Dc -> Security -> Hardware encryption
* Compute F Series
  * *F, Fs, Fsv2, FX* -> High CPU-to-Memory ratio
  * **Good for medium traffic web servers, network appliances, batches and servers**
* Memory E Series
  * *Esv3, Ev3, Easv4, Edsv4, Edv4, Mv2, MDSv2, Dv2*
  * Low CPU-to-Memory ratio (Meaning more memory) **Relational DBs, Caches, In-Memory analytics**
* Storage L Series
  * *Lsv2*
  * NVMe storage
  * High disk throughput and IO **Big Data, SQL, NoSQL, DWH**
* GPU
  * Graphics card integrated
  * *NC, NCv2, NCv3, NCasT4_v3, ND, NDv2, NV, NVv3, NVv4*
  * **Heavy graphic render, video editing, model training, deep learning**
* High performance compute
  * *HB, HBv2, HBv3, HC, H*
  * **Fastest CPU with optional RDMA network interface**

### PowerShell

Az.Compute
* New -> deploy new systems or scale sets from scratch
* Get -> Gather information
* Start-AzVM / Stop-AzVM -> up or down a VM
  * STOP DOES NOT **DEALLOCATE**

```powershell
# Creates azure resource group
New-AzResourceGroup `
   -ResourceGroupName "myResourceGroupVM" `
   -Location "EastUS"

$cred = Get-Credential

# Creates azure VM
New-AzVm `
    -ResourceGroupName "myResourceGroupVM" `
    -Name "myVM" `
    -Location "EastUS" `
    -VirtualNetworkName "myVnet" `
    -SubnetName "mySubnet" `
    -SecurityGroupName "myNetworkSecurityGroup" `
    -PublicIpAddressName "myPublicIpAddress" `
    -Credential $cred

# Start VM
Start-AzVM `
   -ResourceGroupName "myResourceGroupVM"  `
   -Name $vm.name

# Stop VM
Stop-AzVM `
    -ResourceGroupName "myResourceGroupVM" `
    -Name "myVM" -Force

# Get size
Get-AzVMSize -Location "EastUS"

# Resize VM
$vm = Get-AzVM `
   -ResourceGroupName "myResourceGroupVM"  `
   -VMName "myVM"
$vm.HardwareProfile.VmSize = "Standard_DS3_v2"
Update-AzVM `
   -VM $vm `
   -ResourceGroupName "myResourceGroupVM"

# Connect remotely
Get-AzPublicIpAddress `
    -ResourceGroupName "myResourceGroupVM"  | Select IpAddress
mstsc /v:<publicIpAddress>
```

### CLI

```sh
# Create resource group
az group create --name myResourceGroupVM --location eastus
# Create VM
az vm create \
    --resource-group myResourceGroupVM \
    --name myVM \
    --image UbuntuLTS \
    --admin-username azureuser \
    --generate-ssh-keys
# Output of create VM
# {
#   "fqdns": "",
#   "id": "/subscriptions/d5b9d4b7-6fc1-0000-0000-000000000000/resourceGroups/myResourceGroupVM/providers/Microsoft.Compute/virtualMachines/myVM",
#   "location": "eastus",
#   "macAddress": "00-0D-3A-23-9A-49",
#   "powerState": "VM running",
#   "privateIpAddress": "10.0.0.4",
#   "publicIpAddress": "52.174.34.95",
#   "resourceGroup": "myResourceGroupVM"
# }

ssh azureuser@$publicIpAddress

# View current size
az vm show --resource-group myResourceGroupVM --name myVM --query hardwareProfile.vmSize
# List size options
az vm list-vm-resize-options --resource-group myResourceGroupVM --name myVM --query [].name
# Resize
az vm resize --resource-group myResourceGroupVM --name myVM --size Standard_DS4_v2
# Deallocate
az vm deallocate --resource-group myResourceGroupVM --name myVM
# Start
az vm start --resource-group myResourceGroupVM --name myVM
```

### ARM Templates

Pre-configured scripts to deploy single machines to whole environments
**JSON script**

Can be deployed via multiple methods:
* Portal
* CLI
* PowerShell
* Rest API
* Github
* Cloud Shell

### LAB

Setup VM using Powershell

<img src=./assets/lab_diagram_Hands-On1.png/>

```powershell
Get-AzResourceGroup #171-bb356f84-creating-azure-virtual-machines-using

new-azvm -ResourceGroup "171-bb356f84-creating-azure-virtual-machines-using" -Name "TestVM" -VirtualNetworkName "MyVNet" -SubnetName "MySubnet" -SecurityGroupName "MySecurityGroup" -PublicIpAddressName "MyPubIP" -OpenPorts 80,3389
# It then prompts for user and password -> very strict password rules

Get-AzPublicIpAddress #13.67.223.249

Remote connect to it
```

**Further labs**

* [Linux setup](https://docs.microsoft.com/en-us/learn/modules/create-linux-virtual-machine-in-azure/)

### Security

Azure has a *Shared Responsibility Management*

<img src=./assets/shared.png/>

* Protect network traffic
* Identify users -> IAM (Use least privilege principle)
* Encrypt in motion and at rest

**You have to patch the OS of VMs**
* Can do it manually
* Can do it with integrated tools

**Patch programs and services installed on the VM**
* Microsoft does not have anything to help with this

* Scaling -> Clicks or automated
  * VM Scale sets
* Backups
  * Azure backup
  * Understand process to reduce downtime

**Navigating Azure Portal**
* VM
  * Networking settings
    * Check effective rules
  * Connect
    * To remotely connect
  * Disks
    * Check disk configuration  
    * Encryption
  * Size settings
    * Scale if needed
  * Availability + Scaling
    * Update Scale sets
  * Backups
    * Restore
    * Take adhoc backup
      * Retain it until
  * Disaster Recovery
    * Replica server in another region
    * Configure and control secondary env
  * Update management
    * Setup so Azure patches the OS for us
    * Track updates with analytics

### Virtual Networks

When creating a VM, we create a Virtual Network or use an existing one
-> This is what allows VM to be accessed by Public internet
It is possible to setup during VM setup:
* Network Interfaces
* IP Addresses
* Virtual Network and subnets 
* Network security groups
* Load balancers

**Subnet** -> Range of IPs in Virtual Network
**Network Interface** -> Interconnection between VM and VN
**Network Security Group** -> ACL to allow or deny traffic

### Automate ARM Templates

Preconfigured scripts
* JSON files
* Deploy via Console, CLI, Github

**Required fields**
* $schema -> Version of template (tailored per deploy method)
* contentVerion -> Control by the user
* resources -> VMs, Containers, Security groups.....

**Optional fields**
* apiProfile -> Azure Stack mostly (helps consistency between different locations)
* parameters -> Builds rules and static responses 
* variables 
* functions -> expressions to be reused
* outputs -> returns values from deployed resources

### LAB TODO

## Containers

**Good for testing LOL**

Azure Container Instances service

* Create image
  * CLI, PS or Console
  * `az container create` `New-AzContainerGroup` (on PS, default port 80)
* Deploy a container to ACI
* Maintain

Can use docker hub images
Can provide DNS label to get a public API and DNS ***.azurewebsites.com

### Container groups

All resources that are pertinent to the container (Libraries, dependencies and service...)

### Container operation
??

### Container Registry

Manages all images stored in Azure

* How to publish to AzCR
  * 3 Tiers (MICROSOFT) -> Basic, Standard and Premium
    * Basic -> dev tier -> low throughput
    * Standard -> More storage + more throughput (PROD)
    * Premium -> High performance + georeplication + content trust + private link compatibility
    * `az acr create --resource-group a --name b --sku basic` -> NAME IS UNIQUE IN AZURE
      * From output -> get the **loginServer** -> push with this URL
    * `az acr login --name b` `docker tag thisdockerfile loginServer/image:v1` `docker push image:v1`
    * `az acr repository list --name b --output table`
    * Get access keys -> username and password
      * Create container instance -> use user and pass
  * Lock against new updates
* Secure registry
* Run from registry

Will use docker CLI (hehe)
`docker build github.com/{BRANCH}/{FILE}.git#{BRANCH}:{DIRECTORY}`

## App services

Deploy apps and APIs without needing to manage underlying host

Logging options
* Application
  * Capture log info produced by the application
* Web Server Diagnostics
  * Detailed error logging, failed request tracing and web server logs

### Web app services

* Create
  * Supported languages (C#, Node, PHP, Java, python, Ruby and custom containers)
  * Azure CLI or PS
    * `az webapp up --sku F1 --name abc --os-type linux`
  * SKU -> Service tier
    * Free -> Shit
    * Dev/Test -> Supports custom domain but still shit (SHARED COMPUTE TYPE) **NO LINUX ALLOWED ON SHARED PATHETIC**
    * Dedicated Dev/Test (WTF) -> almost the price of production but no autoscaling (DEDICATED COMPUTE)
    * Production -> 50GB disk (meh)
    * Performance and scale optimized -> 250GB disk
    * High-performance, sec & isolation -> 1TB and ISOLATED COMPUTE
* Enable diagnostics
  * Monitoring -> Log -> Configure what, where and how long
    * Storage account is needed if longer than 72h
    * Can use FTP analytics or built in azure analytics
* Deploy code
  * Have to select method on Deployment Center (WTF)
  * Approval workflows with Deployment slots
  * Manually
    * FTP, Git, Dropbox
  * CI/CD
    * Azure DevOps, Github, private git
* Configure SLL, API, Connection Strings
  * Settings -> Configuration
  * Custom domains -> Add -> Validate
  * TLS SSL Settings -> Setup certificates (private and public or **BUY at azure**)
  * Can Scale up or out
* Auto Scale rules
??

### Microservice

### Orchestration

## Azure functions

True serverless. event driven processes

* Can be run as part of the web app services (WHY?)
* Independent serverless functions (AS THEY ARE)

* Linux and Windows (with nuances obviously)

* How to create
  * Pick runtime, some other config and create
  * Pick plan (AGAIN)
    * Consumption
      * Serverless operation
      * Pay per user
      * Scale automatically
    * Premium
      * Serverless + more performance
      * Keep machines warm
      * VNet con
    * App Service Plan
      * Not serverless -> uses nodes on service plan
      * Use custom image
      * Always on plan if enabled
  * Monitor with insights
* Configure triggers
  * In Azure we have Trigger -> Input Binding -> Output binding
  * Trigger runs function (hook, schedule...)

### Somehow amazing binding examples

[REF1](https://docs.microsoft.com/en-us/azure/azure-functions/functions-triggers-bindings?tabs=csharp#bindings-code-examples)
[REF](https://docs.microsoft.com/en-us/azure/azure-functions/functions-bindings-cosmosdb-v2-input?tabs=javascript#example)

```json
{
  "bindings": [
    {
      "authLevel": "anonymous",
      "name": "req",
      "type": "httpTrigger",
      "direction": "in",
      "methods": [
        "get",
        "post"
      ],
      "route":"todoitems/{partitionKeyValue}/{id}"
    },
    {
      "name": "$return",
      "type": "http",
      "direction": "out"
    },
    {
      "type": "cosmosDB",
      "name": "toDoItem",
      "databaseName": "ToDoItems",
      "collectionName": "Items",
      "connectionStringSetting": "CosmosDBConnection",
      "direction": "in",
      "Id": "{id}",
      "PartitionKey": "{partitionKeyValue}"
    }
  ],
  "disabled": false
}
```

```js
module.exports = function (context, req, toDoItem) {
    context.log('JavaScript queue trigger function processed work item');
    if (!toDoItem)
    {
        context.log("ToDo item not found");
    }
    else
    {
        context.log("Found ToDo item, Description=" + toDoItem.Description);
    }

    context.done();
};
```
  
### Azure Durable Functions

**Leverage the Durable Task Framework (WTF)**

Extension to azure functions -> Stateful workflows

Coordination, function chaining
fn1 -> fn2 -> fn3
Fan-out / Fan-in

**Don't support as many runtimes**

### Custom handlers

Small web servers that retrieve events from functions host. **That's if the language supports HTTP primitives**
* Can use go or rust
* or Deno js environment

Trigger -> Binding -> Functions Host -> HTTP CALLS to CUSTOM HANDLER -> GETS PAYLOAD BACK -> OUTPUT BINDING -> Target

**Requires**
* host.json file on root
  * address of webserver
* local.settings.json on root
  * application settings -> env
* function.json
  * just like azure functions within each function folder
  * must be in the folder with same function name as the function
* runs a webserver with script

## Overview

* VMs
  * part of IaaS
  * deployment -> console, cli, ps
  * responsibility (SSM) -> patching and so on
  * IaaS availability -> backups and scale sets
* ARM Templates
  * tooling for IaaC
  * json script
    * schema, contentVersion, resources
* Container 
  * part of PaaS
  * can use docker apis
  * Instances
  * Registries
  * Container groups
* Azure app services
  * Web app services
    * tiers (lol)
* Azure functions
  * Weird plans
  * Triggers, input and output bindings
  * Usage of app services infrastructure
  * Durable functions
* Orchestration -> easy

### Links

[Pricing for WebAppServices](https://azure.microsoft.com/en-us/pricing/details/app-service/windows/)
[ARM](https://docs.microsoft.com/en-us/azure/azure-resource-manager/templates/overview)
[Virtual Network](https://docs.microsoft.com/en-us/azure/virtual-network/network-overview)
[Containers](https://docs.microsoft.com/en-us/azure/container-instances/container-instances-overview)