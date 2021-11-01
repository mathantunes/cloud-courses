# Azure Commands

* Create new container registry
  * `az`
  * `New-AzContainerRegistry`
* Show all the images in an Azure Container Registry -> in a table
  * `az acr repository list --name acgRegistry --output table`
  * `New-Az`
* Create a managed identity and ensure it was built as part of the scale set
  * `az vmss identity assign --resource-group AzureDale --name ElasticMagic`
  * `New-Az`
* Create a new Azure Key Vault access policy for secrets
  * `az keyvault set-policy --name chamberOfSecrets --object-id --secret-permissions`
  * `New-Az`
* Create a new message queue called llamaqueue in namespace ACGServiceBusNS
  * `az`
  * `New-AzServiceBusQueue -ResourceGroupName acg204RG -NamespaceName ACGServiceBusNS -Name llamaqueue`
* Create an Event Hubs namespace called az204hubns
  * `az eventhubs namespace create --name az204hubns --resource-group acg204RG -l 'East US'`
  * `New-Az`
* Create a new API Management
  * `az apim api create --service-name llamadrama -g az204RG --api-id LlamaDrama --path '/llama' --display-name 'Llama Drama'`
  * `New-Az`
* Create a new Azure Key Vault access policy for secrets
  * `az`
  * `New-Az`
* Create a new Azure Key Vault access policy for secrets
  * `az`
  * `New-Az`
* Create a new Azure Key Vault access policy for secrets
  * `az`
  * `New-Az`
* Create a new Azure Key Vault access policy for secrets
  * `az`
  * `New-Az`
* Create a new Azure Key Vault access policy for secrets
  * `az`
  * `New-Az`

# Azure Logic App JSON

# Test tips

* Secure and Lot of memory app service -> ISOLATED
* Authenticate but without sending credentials to API -> (7) -> Utilize managed identity
* Ensure App Service traffic is secure (8) -> Install SLL Certificate on the App Service to encrypt all traffic
* Azure AD API Permission (10) -> Delegated Permission
* Login with Google, Microsoft and so on (19) -> Azure AD B2C
* Enforce HTTPS on Static Website and CDN (20)
* Service access to secrets only (33) -> Register app with AD, register app with Key Vault, Associate a certificate with azure AD web app, make app use the certificate to authenticate to Key Vault
* Immutable blob storage policy -> Done on blob storage level
* C# BlobContainerService manipulates the entities
* Notification Hub -> Retrieve PNS Handle, STORE PNS Handle, Send Notification to PNS, Send to Device



`az appservice plan create --name $webappname --resource-group rgroup --sku FREE`
`az webapp create --name $webappname --resource-group rgroup --plan $webappname`
`az webapp deployment source config --name $webappname --resource-group rgroup -- repo-url $gitrepo`

* Storage blob trigger might have delays of up to 10 minutes LOL

* `group`, `app service plan`, `webapp`, `webapp deployment slot`, `webapp deployment source`, 

* Azure Functions from QUEUE
  * Default retry 5
  * Default batch of 16 messages
  * Track availability (performance) with a time triggered function LOL
* Run script on Azure Web App devops
  * WEBSITE_RUN_FROM_PACKAGE setting in host.json
  * .deployment file in root
* Web App for Containers
  * Persist files on **/home** by settings **WEBSITES_ENABLE_APP_SERVICES_STORAGE** to **true**

Blob LEASE????
* Like a semaphore but with a shit name
* Can be acquired for some time or indefinitely

CosmosDB
* Multi region `--locations 'southcentralus=0 eastus=1 westus=2`
* Change Feed PROCESSOR
  * Monitored container -> has the data from which the change feed is generated
  * Lease container -> acts as a state storage and coordinates processing across multiple workers
  * Host -> Application instance that uses the change feed processor to listen for changes
  * Delegate -> Code that defines what to do with each batch
  * Min throughput is 10% of configured max throughput
  * Ranged query -> goes through different partitions
  * 

Kubernetes
* INGRESS CONTROLLER
  * Provides reverse proxy
  * configurable traffic routing
  * TLS

Secure Logic App with **INTEGRATION SERVICE ENVIRONMENT**

Service Bus Filters
* Boolean Filters
  * True -> Get all arriving messages
  * False -> Get none
  * They are both derived from SQL Filter LOL
* SQL Filters
  * Complex querying
* Correlation Filters
  * Set of conditions that are matched against one or more arriving message or property

Retrieve and update user profile info stored in **AD** with **Microsoft Graph API or Azure API Management**

* SAS Revoking
  * Revoke delegation key
  * Delete stored access policy

* Cache Coding
  * `IDatabase cache = Connection.GetDatabase()`
  * `cache.KeyDelete("Key")`


* Secure VNET Azure Functions
  * Create azure function on premium plan
  * Create system assigned managed identity
  * Create key vault policy