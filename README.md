# 简介
***
go语言实现简单分布式缓存goCache，参考groupcache

实现功能如下：

* LRU缓存淘汰策略
* 单机并发缓存
* HTTP服务端
* 一致性哈希
* 分布式结点
* 防止缓存击穿

# Quick Start
```
$ git clone git@github.com:wjh791072385/goCache.git
$ cd goCache
$ ./run.sh
```

# 缓存数据处理流程
***
数据处理流程：

（1）先检查本机是否有缓存，如果无进行第二步

（2）通过一致性哈希算法获取peer结点，检查peer结点是否有当前缓存值，如果有则返回。如果没有则进行第三步

（3）从数据库中加载缓存值，利用LRU算法更新缓存

# 代码模块
***
```
├── README.md
├── go.mod
├── goCache         
│   ├── byteview.go     //只读数据结构，存储缓存值
│   ├── cache.go        //内部控制并发            
│   ├── consistentHash
│   │   └── consistentHash.go   //一致性哈希
│   ├── goCache.go      //外部交互，主流程控制
│   ├── goCache_test.go 
│   ├── http.go         //http交互模块
│   ├── http_test.go
│   ├── lru
│   │   ├── lru.go   //实现LRU淘汰策略   
│   │   └── lru_test.go
│   ├── peers.go        //peer结点接口相关     
│   └── singleflight
│       └── singleflight.go     //防止缓存击穿策略
├── main.go     //主函数进行测试
└── run.sh      //脚本测试
```






