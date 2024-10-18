# lrucache

## LRU Cache
An in memory cache implementation that expires the least recently used items, and limits cache size by a maximum number of items.

## API (ruby example code; implemented in Golang)
```bash
# A cache object can be instantiated in memory. It requires the max number of records as an argument:
cache = Cache.new(max_size: 100)

#Converts the cache into a hash:
cache.to_h

# An object may be written to a string cache key:
cache.write("key", value)

# That object may be retrieved by a key, or nil is returned if it is not found:
cache.read("key")

# A cached value may be deleted by key:
cache.delete("key")

# All values may be deleted:
cache.clear

# The number of records can be fetched at any time:
cache.count
```

## Unit Tests
```bash
go test ./...
```
