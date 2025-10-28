https://www.liwenzhou.com/posts/Go/gRPC/#c-9-2



```go
func(ctx context.Context, method string, req, reply interface{}, cc *ClientConn, invoker UnaryInvoker, opts ...CallOption) error
```

一元拦截器的实现通常可以分为三个部分: 调用 RPC 方法之前（预处理）、调用 RPC 方法（RPC调用）和调用 RPC 方法之后（调用后）。

预处理：用户可以通过检查传入的参数(如 RPC 上下文、方法字符串、要发送的请求和 CallOptions 配置)来获得有关当前 RPC 调用的信息。
RPC调用：预处理完成后，可以通过执行invoker执行 RPC 调用。
调用后：一旦调用者返回应答和错误，用户就可以对 RPC 调用进行后处理。通常，它是关于处理返回的响应和错误的。 若要在 ClientConn 上安装一元拦截器，请使用DialOptionWithUnaryInterceptor的DialOption配置 Dial 。




