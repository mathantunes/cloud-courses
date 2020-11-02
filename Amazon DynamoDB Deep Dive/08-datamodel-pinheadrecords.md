# Version 1 Data Model Overview by Mark Richman

[Version 1 Repo](https://github.com/linuxacademy/content-dynamodb-deepdive/tree/master/webapp-v1)

This version still utilizes many Scan operations without any Secondary indexes

## Album

Images are now on S3

```json
{
    "id": 1,
    "album_art": "/path/to/image/1.jpg",
    "artist_id": 51,
    "format": "12 Vinyl",
    "price": 48.33,
    "title": "Title",
    "year": 1998,
}
```

## Artist

```json
{
    "id": 51,
    "name": "Edu Happy",
}
```

## Track

```json
{
    "id": 123,
    "album_id": 1, // Album Id Refference
    "name": "Nice Track",
    "number": 2,
    "length": 270066,
}
```

## User

```json
{
    "email": "email@email.com",
    "password": "encrypted",
}
```