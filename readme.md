---
   rock dns plugin
---

# 配置
```lua
    local dns = rock.dns{
        nameserver = "114.114.114.114:53",
        timeout = 5,
    }
```

# 使用
```go
    //set
    obj := ud.(dns.Dns)
    r , size := obj.Query("www.baidu.com")
    //
```