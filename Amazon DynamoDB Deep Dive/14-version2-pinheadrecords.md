# Database Model V2 by Mark Richman

[V2 Ref GitHub](https://github.com/linuxacademy/content-dynamodb-deepdive/blob/master/webapp-v2/webapp/models.py)

**I noticed the weird LSIs not being used (Artist Name and Title and Artist Name and Year)**

This version of the Pinehead Records data model features:

* Better table structure (single hierarchical)
* Local and Global secondary indexes
* User accounts in DynamoDB

## Pinehead Records

### Key Schema:

* PK (artist_name)
* SK (id) **album_id**

Represents a one to many relationship between artist and album

### Data Attributes:

* album_art -> **S3 Path**
* year
* format
* price
* title
* track_info -> JSON array for name, number and length

### Local Indexes:

#### Artist Name and Year

Primary key is still artist_name but sorting key is now year to return all data attributes

This allows for searching an artist and all its records on a certain year.

#### Artist Name and Title

Primary key is still artist_name but sorting key is now title to return all data attributes

This allows for searching an artist and a certain title.

### Global Indexes:

#### Title Index

Title data attribute used as primary key to return all data attributes

This allows for **searching a title from different artists**.

#### Year Index

Year data attribute used as primary key to return all data attributes

This allows for **searching all titles released on that year**.

#### Price Index

Price data attribute used as primary key to return all data attributes

#### Format Index

Format data attribute used as primary key to return all data attributes

