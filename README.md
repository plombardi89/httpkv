# httpkv

Exposes a map data structure with a simple HTTP interface. Written for the exceedingly (rare) situations where 
[Redis](https://redis.io) just feels like overkill and you do not need persistence.

# API

## Put

Create or update an existing item.

## Get

Get an entry from the map by key. If the key does not exist then a `HTTP 404 - Not Found` is returned.

## Delete

Delete an entry from the map by key. Delete always succeeds and returns a `HTTP 204 - No Content`.

## Export

Return the entire map data structure serialized as JSON.

# License

Apache 2.0. Please read [LICENSE](LICENSE) for details.
