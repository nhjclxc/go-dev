# 基础功能
支持使用`dig @127.0.0.1 local.com A`命令来指定使用本地权威dns服务器查询local.com域名的ip，代码看[main.go](main.go)

# 增强功能
 - [enhance01.go](enhance01%2Fenhance01.go)
    - 在[main.go](main.go)基础上监听 TCP 和 UDP（兼容大包）【保证服务器能接收请求，UDP 用于常规查询，TCP 用于大包和区分】
 - [enhance02.go](enhance02%2Fenhance02.go)
   - 在[enhance01.go](enhance01%2Fenhance01.go)基础上支持处理更多 Qtype【扩展本地权威解析，处理 AAAA / CNAME / MX / TXT 等类型】
 - [enhance03.go](enhance03%2Fenhance03.go)
   - 在[enhance02.go](enhance02%2Fenhance02.go)基础上 支持多 Question 【一个报文可以同时请求多个域名，提高客户端兼容性】
 - [enhance05.go](enhance01%2Fenhance05.go)
   - 在[enhance03.go](enhance03%2Fenhance03.go)的基础上 支持内存缓存（LRU / TTL 过期）【对重复查询提高性能，也减少上游递归查询压力，提高性能】
 - [enhance06.go](enhance01%2Fenhance06.go)
   - 在[enhance05.go](enhance05%2Fenhance05.go)的基础上 支持通配符，例如 *.local.com【增加灵活性】
 - [enhance07.go](enhance01%2Fenhance07.go)
   - 在[enhance06.go](enhance01%2Fenhance06.go)的基础上 支持 递归查询 / 上游转发【如果本地没有答案，向上游 DNS 查询，类似 dnsmasq；依赖缓存和 Question 处理】
 - [enhance08.go](enhance08%2Fenhance08.go)
   - 在[enhance06.go](enhance01%2Fenhance06.go)的基础上 支持从配置文件热加载域名 → IP 【方便动态更新域名记录，依赖本地权威解析和缓存】
 - 添加 EDNS0 / DNSSEC【高级功能，增加兼容性和安全性，可在前面基础上叠加】



## 其他常见增强功能（可选）

1. **日志和监控**

   * 记录每次查询、响应延迟、上游查询失败次数等，便于运维。

2. **访问控制 / 白名单 / 黑名单**

   * 限制哪些客户端可以访问 DNS，防止滥用。

3. **限流 / QPS 控制**

   * 防止被 DDoS 攻击，保护服务器稳定性。

4. **统计报表**

   * 查询频率、最常访问域名等，用于性能分析。

5. **缓存持久化**

   * 将缓存记录写入磁盘，服务重启后保留热数据。

6. **支持 IPv6 上游查询**

   * 上游服务器也可以是 IPv6，提高兼容性。

7. **可扩展插件机制**

   * 允许自定义拦截、重写、转发规则，例如像 CoreDNS 插件架构。




# 知识点

## dns查询类型
| 类型    | 描述      |
| ----- | ------- |
| A     | IPv4 地址 |
| AAAA  | IPv6 地址 |
| CNAME | 别名记录    |
| MX    | 邮件交换记录  |
| TXT   | 文本记录    |
| SRV   | 服务记录    |

## QType类型
| Qtype | 类型值 | 含义 / 返回数据                                 |
| ----- | --- | ----------------------------------------- |
| A     | 1   | IPv4 地址，例如 `127.0.0.1`                    |
| AAAA  | 28  | IPv6 地址，例如 `::1`                          |
| CNAME | 5   | 别名记录（Canonical Name），表示域名的别名指向另一个域名       |
| NS    | 2   | Name Server，指定域名的权威 DNS 服务器               |
| MX    | 15  | 邮件交换记录，指定邮件服务器                            |
| TXT   | 16  | 文本记录，可存储任意文本信息，例如 SPF 或验证信息               |
| SOA   | 6   | Start of Authority，域名的起始授权记录，包括主 DNS、序列号等 |
| PTR   | 12  | 指针记录，用于反向解析 IP → 域名                       |
| SRV   | 33  | 服务定位记录，例如 `_sip._tcp.example.com`         |
| ANY   | 255 | 查询该域名的所有记录（客户端要求返回所有类型）                   |


## DNS解析流程图
```
┌───────────────────────┐
│       客户端请求        │
│   dig / 浏览器 / App    │
│   域名 + Qtype          │
└───────────┬───────────┘
            │
            ▼
┌───────────────────────┐
│ ServeDNS(w, r *dns.Msg) │
│ - r.Question[0] => qname │
│ - r.Question[0] => qtype │
└───────────┬───────────┘
            │
            ▼
  ┌─────────────────────┐
  │ 检查 domainMap      │
  │ 自定义域名？         │
  └───────┬─────────────┘
          │Yes
          ▼
 ┌─────────────────────┐
 │ 本地权威解析         │
 │ customDNSQuery()     │
 │ - Qtype == A         │
 │     → dns.A          │
 │ - Qtype == AAAA      │
 │     → dns.AAAA       │
 │ - Qtype == CNAME     │
 │     → dns.CNAME      │
 │ - Qtype == ANY       │
 │     → 返回所有 RR    │
 └───────┬─────────────┘
          │
          ▼
   ┌───────────────┐
   │ 返回客户端响应 │
   │ w.WriteMsg(msg)│
   └───────────────┘

          │No
          ▼
 ┌─────────────────────────┐
 │ 上游递归查询             │
 │ upstreamDNSQuery()       │
 │ - 新建 Msg               │
 │ - SetQuestion(qname,qtype)│
 │ - RecursionDesired=true   │
 │ - 循环上游服务器查询      │
 └───────────┬─────────────┘
             │
             ▼
 ┌─────────────────────────┐
 │ 上游返回响应             │
 │ 修正 Msg.Id = r.Id       │
 │ 保留 Answer / Authority / Additional │
 └───────────┬─────────────┘
             │
             ▼
      ┌───────────────┐
      │ 返回客户端响应 │
      │ w.WriteMsg()  │
      └───────────────┘
```
