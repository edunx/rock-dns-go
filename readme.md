---
   rock dns plugin
---

# 配置
```lua
    local dns = rock.dns{
        nameserver = "114.114.114.114:53",
        timeout = 5,
        type_name = "CNAME",
    }
    -- 测试代码
   dns.query( "www.baidu.com." )

 
```

# 使用
```golang
    //set
    obj := ud.(dns.Dns)
    r , size := obj.Query("www.baidu.com")
    //
```