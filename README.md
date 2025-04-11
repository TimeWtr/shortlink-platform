# shortlink-platform
基于Go实现的高可用、高性能短链接生成平台


### 1. 单条短码生成流程
```mermaid

sequenceDiagram
    participant Client
    participant API
    participant Bloomfilter
    participant Cache
    participant DB

    Client->>API: post/rpc (提交URL)
    API->>Cache: 查询短码池中是否有可用
    alt 短码池中有可用短码
        Cache-->>API: 返回可用短码
    else 短码池中无可用短码
        loop 循环获取3次
           API ->> Cache: 从短码池获取可用短码
           alt 短码池中有可用短码
               Cache->>API: 返回可用短码，跳出循环
           else 短码池中无可用短码
               API ->> API: 检查循环次数是否达到3次
               alt 未达到3次
                 API ->> API: 循环计数+1
                 API ->> Cache: 尝试获取可用短码
               else 达到3次
                 API -->> Client: 生成短码失败
               end
           end
        end
    else
      API->>DB: 将短码与原始URL的映射关系写入数据库
      API->>Cache: 将短码与原始URL的映射关系缓存
      API -->>Client: 生成成功
    end
```

### 2. 短码池预生成流程
#### 2.1 哈希+解决冲突算法
```mermaid
sequenceDiagram
    participant Scheduler
    participant Instance
    participant Pool[Redis List]
    participant Bloom filter
    
    Instance ->> Scheduler: 抢占定时任务
    Scheduler -->> Instance: 抢占成功
    Instance ->> Pool[Redis List]: 查询池中的预生成短码数量是否小于阈值(100000条)
    alt 数量满足
        Pool[Redis List] -->> Instance: 结束定时任务
    else 数量不够, 本地预生成足够数量的短码
        loop
            Instance ->> Instance: 检查生成的数量是否满足
            alt 数量满足
                note right of Instance: 达到阈值，退出循环
            else 不满足
               Instance ->> Instance: 生成一条新的短码
               Instance ->> Bloom filter: 查询短码是否存在
               alt 存在
                  Instance ->> Instance: 加入随机盐值
                  note right of Instance: 重复此步骤直至哈希短码不再重复
               else 不存在
                  Instance ->> Instance: 本地缓存
                  alt 本批次达到1000条短码
                      note right of Instance: Pipeline写入/Lua脚本写入
                      Instance ->> Pool[Redis List]: 批量写入到短码池[并发安全]
                      Instance ->> Bloom filter: 批量新增到过滤器[并发安全]
                  end
               end
           end
        end
    end
    Instance -->> Scheduler: 结束抢占到的定时任务，等待下次执行
```
#### 2.2 递增序列+Base62算法
```mermaid
sequenceDiagram
    participant Scheduler
    participant Instance
    participant Pool[Redis List]
    participant DB
    
    Instance ->> Scheduler: 抢占定时任务
    Scheduler -->> Instance: 抢占成功
    Instance ->> Pool[Redis List]: 查询池中的预生成短码数量是否小于阈值(100000条)
    alt 数量足够
        note right of Instance: 结束定时任务
    else 数量小于阈值
        Instance ->> DB: 查询上次定时任务生成的最新递增ID
        loop
            alt 数量满足
                note right of Instance: 达到阈值，退出循环
            else
                Instance ->> Instance: 生成新的递增ID
                Instance ->> Instance: Base62计算，获取新的短码
                Instance ->> Instance: 本地缓存
                alt 本批次达到1000条短码
                    note right of Instance: Pipeline写入/Lua脚本写入
                    Instance ->> Pool[Redis List]: 批量写入预生成短码到短码池，并更新总数[并发安全]
                    Instance ->> DB: 更新最新的递增ID到数据库[并发安全]
                end
            end
        end
    end
    
    Instance -->> Scheduler: 结束抢占到的定时任务，等待下次执行
```