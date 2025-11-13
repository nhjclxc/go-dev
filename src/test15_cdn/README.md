# test15_cdn



```shell


.cdn
├── cdn-gslb-service           // 全局负载均衡：根据用户地理位置、延迟、健康状态选择最优节点
├── cdn-slb-service            // 节点内部负载均衡：将请求分发到节点内的 Edge Cache Server
├── cdn-origin-service         // 源站服务：提供原始内容和回源接口给 ECS
├── cdn-ecs-service            // Edge Cache Server：边缘缓存文件，提供下载，异步回源和缓存管理
├── cdn-log-service            // 日志收集：访问日志、回源日志、下载日志的采集和存储
├── cdn-monitor-service        // 监控与告警：监控节点健康、流量、缓存命中率，并触发告警
├── cdn-pay-service            // 计费与支付：用户计费、流量计费、套餐管理
├── cdn-config-service         // 配置管理：动态下发节点信息、缓存策略、回源策略、域名配置
├── cdn-cache-manager-service  // 缓存管理：缓存清理、过期文件回收、缓存预热（Preheat）
├── cdn-auth-service           // 安全鉴权：防盗链、请求签名验证、访问权限控制
├── cdn-traffic-control-service // 流量控制：限速、限流、带宽配额管理、QoS
├── cdn-backup-service         // 数据备份与同步：文件冗余、节点间同步、高可用保障
├── cdn-alert-service          // 告警通知：节点宕机、异常流量、缓存命中率低等告警通知
├── cdn-api-gateway            // API 网关：提供运维后台和业务系统访问接口
├── cdn-metrics-service        // 指标收集与聚合：统计各服务性能指标，支持报表和调度决策
├── cdn-preload-service        // 内容预热：提前将热点文件推送到边缘节点，提升访问速度

```

```shell
各个模块有单独的go mod可以独立打包部署。整个cdn项目又有go work来整合方便本地开发代码相互使用


微服务模块对外提供代码的方式
    方式 A：直接作为库模块被引用：模块内部有一个 pkg/ 或 internal/ 目录放公共函数、结构体、接口。其他模块直接引用
    方式 B：通过接口 + 依赖注入解耦：先在通用模块（cdn-common）里面写接口（PaymentProcessor），然后在目标模块（cdn-pay-service）里面写具体实现（AlipayPayService、WeChatPayService），最后在要使用的模块（cdn-api-gateway）注入具体实现
    方式 C：微服务之间通过 gRPC/HTTP 调用：在common里面定义proto，在pay里面实现服务端，在api里面实现客户端

```


```shell

用户请求
    │
    ▼
1️⃣ 权威 DNS 查询 → cdn-gslb-service
    - GSLB 根据用户地理位置、延迟、节点健康状态选择最优节点
    - 返回节点 IP 给用户
    │
    ▼
2️⃣ 用户请求到节点 → cdn-slb-service
    - 节点内部 SLB 将请求分发到 ECS（Edge Cache Server）
    - 分发策略：轮询/最少连接/缓存命中率等
    │
    ▼
3️⃣ cdn-ecs-service
    - 检查本地缓存是否命中
        ├─ 命中：直接返回文件给用户
        └─ 未命中：
            │
            ▼
4️⃣ 回源 → cdn-origin-service
    - ECS 向源站请求原始文件
    - 回源文件后，ECS 异步缓存

```


| 模块                              | 触发时机                | 作用                      |
| ------------------------------- | ------------------- | ----------------------- |
| **cdn-config-service**          | ECS 或 SLB 启动/动态策略下发 | 提供节点信息、缓存策略、回源策略、域名配置   |
| **cdn-cache-manager-service**   | ECS 缓存管理            | 清理过期缓存、缓存预热             |
| **cdn-auth-service**            | ECS / SLB 处理请求前     | 验证请求合法性、防盗链、签名鉴权        |
| **cdn-traffic-control-service** | SLB / ECS 分发请求      | 限速、限流、带宽控制              |
| **cdn-log-service**             | 请求处理时               | 收集访问日志、回源日志、下载日志        |
| **cdn-monitor-service**         | 整个流程                | 监控节点健康、缓存命中率、流量，触发告警    |
| **cdn-metrics-service**         | 请求或定时汇报             | 聚合性能指标，辅助调度和报表          |
| **cdn-alert-service**           | 监控发现异常时             | 节点宕机、缓存命中率低、异常流量告警      |
| **cdn-preload-service**         | ECS 缓存管理            | 可在热点文件未被请求前预热缓存         |
| **cdn-api-gateway**             | 外部运维或业务后台访问         | 查询状态、管理配置、触发缓存清理等       |
| **cdn-pay-service**             | 用户付费相关              | 流量计费、套餐控制，可在请求访问计费逻辑时调用 |
| **cdn-backup-service**          | ECS 缓存更新 / 源站同步     | 异步处理文件冗余和节点同步           |

