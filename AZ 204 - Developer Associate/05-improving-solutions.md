# Optmizing Solutions

## Content Delivery Networks

Serves content to outside the region where resource lives.

* Audio, video, css, static content are stored and not fetched from main server
* Lower latency
* Streaming video, more reliable for IoT
* Good for access from far away

This is done through caching techniques on regional CDNs

* CDN Profiles
  * Register to CDNs
    * `Get-AzResourceProvider -ProviderNamespace Microsoft.Cdn`
    * `Register-AzResourceProvider -ProviderNamespace Microsoft.Cdn`
  * Create
    * Unique name
    * Choose Subscription
    * Choose pricing tier (mostly same price but a **Verizon** plan which is $$$$$)
    * Name it

### Caching Rules for CDN

Configuration after it is created

* Global caching
  * 1 per endpoint
  * affects all requests
  * will override http cache-directive headers
* Custom caching rules
  * assigned to paths and extensions
  * processed in order and override global

### Query string caching

Happens in the web request and consists of key-value pair.

* Modes:
  * IgnoreQueryString: (WEIRD NAME)
    * Passed initial string and caches the object
    * New Requests are ignored until the cache expires
    * Default settings
  * BypassCaching
    * Requests with query string are not cached and will pull from source
  * Cache Every Unique URL
    * Every unique URL cached

## Caching

### Cache for Redis

In-memory caching for handling sessions, application cache, data storage, job and message queueing.

Transactions and so on.

* Configuration
  * Install redis SDK
  * Add to configuration (host and key)

### FrontDoor

Holds on to data 'chuncks' until the full object has been retrieved or the session is **closed**.
Objects are then cached and re-used
Finds the best route and way to upload or download. (can be used in conjunction with the other ones).

Finds the optimal route to the resource requested (Source or CDN)
Utilizes query string caching
* Ignore or Cache Every
  * Similar to CDN in this case

Purges are done with TTL
* Single Path Purge
  * Purge individual files based on their full path and extension
* Wildcard Purge
  * Purge everything from the endpoint below the wildcard
* Root Domain Purge
  * Full purge

All of these are defined in request headers
