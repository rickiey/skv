# SKV

simple KV Storage

## supprt 

+ File (json format) Default
+ redis
+ ...

### File

default path and file

```go

var defaultFile string = "skv-tmp-default-file.json"
var defaultPath string = "/tmp/simpleKV-default-tmp-dir/"
```

or set custom path and file 

```go
SetFsRepo(filepathm, filename)
```

### Redis

set Default SKV : redis

```go
SetRedisKv("127.0.0.1:6379", "redis-password")
```