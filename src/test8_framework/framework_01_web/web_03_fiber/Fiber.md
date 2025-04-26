### Fiber

[Fiber开源代码](https://github.com/gofiber/fiber)
[官方文档](https://docs.gofiber.io/)
[官方文档-中文](https://docs.fiber.org.cn/)
[官方的学习示例代码](https://github.com/gofiber/recipes)


安装：`go get github.com/gofiber/fiber/v2`

#### 学习资料

- https://geekdaxue.co/read/fiber/welcome
- https://learnku.com/docs/gofiber/2.x
-

### 强大的中间件支持
https://docs.fiber.org.cn/category/-middleware/#google_vignette

以下为包含在Fiber框架中的中间件列表.

中间件 描述
basicauth 基本身份验证中间件提供 HTTP 基本身份验证。 它为有效凭证调用下一个处理程序，为丢失或无效凭证调用 401 Unauthorized
cache 用于拦截和缓存响应
compress Fiber 的压缩中间件，默认支持 deflate，gzip 和 brotli
cors 使用各种选项启用跨源资源共享(CORS)
csrf 保护来自 CSRF 的漏洞
encryptcookie 加密 cookie 值的加密中间件
envvar 通过提供可选配置来公开环境变量
etag 让缓存更加高效并且节省带宽, 让 web 服务在响应内容未变更的情况下不再需要重发送整个响应体
expvar 通过其 HTTP 服务器运行时间提供 JSON 格式的暴露变体
favicon 如果提供了文件路径，则忽略日志中的图标或从内存中服务
filesystem Fiber 文件系统中间件，特别感谢 Alireza Salary
limiter 用于 Fiber 的限速中间件。 用于限制对公共 api 或对端点的重复请求，如密码重置
logger HTTP 请求/响应日志
monitor 用于报告服务器指标，受 Express-status-monitor 启发
pprof 特别感谢 Matthew Lee (@mthli)
proxy 允许您将请求proxy到多个服务器
recover Recover 中间件将可以堆栈链中的任何位置将 panic 恢复，并将处理集中到 ErrorHandler
requestid 为每个请求添加一个 requestid.
session Session 中间件. 注意: 此中间件使用了我们的存储包.
skip Skip 中间件会在判断条为 true 时忽略此次请求
timeout 添加请求的最大时间，如果超时则发送给ErrorHandler 进行处理.
adaptor net/http 处理程序与 Fiber 请求处理程序之间的转换器，特别感谢 @arsmn！
helmet 通过设置各种 HTTP 头帮助保护您的应用程序
keyauth Key auth 中间件提供基于密钥的身份验证
redirect 用于重定向请求的中间件
rewrite Rewrite 中间件根据提供的规则重写URL路径。它有助于向后兼容或者创建更清晰、更具描述性的链接
以下为外部托管的中间件列表，由 Fiber团队 维护。

中间件 描述
jwt JWT 返回一个 JSON Web Token(JWT) 身份验证中间件
storage 包含实现 Storage 接口的数据库驱动，它的设计旨在配合 fiber 的其他中间件来进行使用
template 该中间件包含 8 个模板引擎，可与 Fiber v1.10.x Go 1.13或更高版本一起使用
websocket 基于 Fasthttp WebSocket for Fiber 实现，支持使用 Locals ！