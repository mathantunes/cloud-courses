# Database Model V3 by Mark Richman

## Schema

### Key Schema:

* PK: Type
* SK: ID (Entity ID)
  
#### Data Attributes

* album_art
* year
* format
* price
* name_title
* artist_id
* sku

### LSI

* PK: Type
* SK: name_tile (**can be artist_name, album_name, track_name**)

### GSIs

* name_title
* year
* price
* format

**It requires new queries for album data, track data**

Through use of caching, it may be quite fast, but more requests are made to fill in pages.