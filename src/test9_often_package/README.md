# go常用的包

[常用包和第三方包介绍](https://golangguide.top/golang/%E5%B8%B8%E7%94%A8%E5%8C%85%E5%A4%A7%E5%85%A8.html)


## 基础功能
 - fmt：格式化输入输出（Println、Printf 等）
 - os：操作系统相关（文件、环境变量等）
 - io/ioutil：I/O 读写
 - bufio
 - strconv：字符串和基本数据类型转换
 - time：时间和日期处理
 - math/rand：数学计算、随机数
 - unicode：Unicode 处理
 - log：日志记录
 - sort：排序
 - data_convert：包含多种数据格式之间的转化：对象、json、map、xml等数据格式之间的转化【encoding/json】
 - path/filepath：路径处理
 - archive/zip / compress/gzip：文件压缩解压
 - regexp：正则表达式
 - reflect/unsafe
 - error
 - go-num，类似 Python 的 numpy, [go-num](https://github.com/gonum/gonum)
 - builtin：go的基础数据结构，[builtin](https://pkg.go.dev/builtin)
 - Uuid：这个包提供一个纯Go实现的UUID，[github.com/gofrs/uuid](https://github.com/gofrs/uuid.git)
 - "math/big"：精密计算和 big 包
 - z
 - 


## 并发编程
 - sync：同步原语（Mutex、WaitGroup、Once）
 - sync/atomic：原子操作
 - context：上下文管理（WithCancel、WithTimeout）
 - runtime：Go 运行时信息（Goroutine 数量、垃圾回收）
 - 


## 网络编程
 - net：TCP/UDP 网络编程
 - net/http：HTTP 服务器与客户端
 - net/url：URL 解析
 - net/rpc：RPC 远程调用
 - crypto / crypto/*：加密相关（md5、sha256 等）
 - mime/multipart：文件上传处理
 - 


## 开源包
1. [samber/lo](https://github.com/samber/lo)：类似于 Java 中的 Stream API
2. [go-funk](https://github.com/thoas/go-funk)：类似于 Java 中的 Stream API，允许你在 Go 中进行更简洁的链式操作
3. [aws-lambda-go](https://github.com/aws/aws-lambda-go)：提供类似 Java Stream 的 map 和 filter 等功能
 - z


## 基础功能
 - z


