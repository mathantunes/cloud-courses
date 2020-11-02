# Databases 101 by Mark Richman

Storing information

## Relational Scheme

Relate by some sort of business agnostic ID to access related data.

* Table -> Collection of related data
* Row -> Collection of related values of columns
* Column -> Units of data with a datatype
* Schema -> Relationships between tables
* Key -> column for rapid and sorted access
    * Candidate Key -> Unique identification
    * Primary Key -> Unique identifier
    * Foreign Key -> Identifies a value which is a Primary key at another table
* Query -> Data retrieval statement

## Relational DBs and SQL

### ACID Properties

A sequence of database operations that satisfies the **ACID** is called a **transaction**

* Atomicity either all occur or nothing occurs
* Consistency state of the database before and after is consistent (no integrity error)
* Isolation any changes made by a transaction is isolated from the rest of the system until the transaction is commited
* Durability data is saved by the system such that even in the event of a failure and system restart, the data is in the correct state.

### Scaling

#### Replication (too much load)

Replicate on read replicas.

Writes go to master database that replicates to replicated read replicas.

* **Writes are still not scaled**

* **Replication lag**

#### Sharding (too much data)

Instead of having one master, have many masters

**PROS**
* Can Store larges data sets
* Can handle more load than a single master node

**CONS**
* Queries become more complex (range can hit all nodes)
* Joins are difficult (tables are no longer all on the same nome)

### CAP Theorem

**It is impossible to provide more than 2 out of these 3**

* Consistency - All clients see the same data at the same time
* Availability - The system continues to operate even in the presence of node failures.
* Partition Tolerance - The system continues to operate in spite of network failures.

**Combinations**

* CA - Data is consistent between all nodes, as long as all nodes are online
* CP - Data is consistent between all nodes and maintains partition tolerance by becoming unavailable when a node goes down.
* AP - Nodes remain online even if they can't communicate with each other and will re-sync data one the partition is resolved.

**Distributed systems must choose the partition tolerance P; as it need to tolerate packet loss over the network**, so the choice is between consistency (C) and Availability (A).

### Introduction to NoSQL

**Optimized for compute**

* Denormalized/hierarchical
* Instantiated views
* Scale horizontally
* Built for OLTP at scale

Scale through Sharding and Clustering

#### No SQL Types

* Key-Value -> Simple method to store data
* Document-oriented -> JSON-like data
* Column-oriented -> each column of data in sequential blocks on disk (analytics) 
* Graph -> store relationships. Data is stored as nodes and edges. very fast.

#### ACID vs. BASE

* Basic availability: The database appears to work most of the time.
* Soft state: Stores do not have to be write-consistent, nor do different replicas have to be mutually consistent all the time
* Eventual consistency: Stores exhibit consistency at some later point (lazily at read time)

**Base** 

* Weak consistency
* Availability first
* Best effort
* Optimistic
* Faster and easir schema evolution

