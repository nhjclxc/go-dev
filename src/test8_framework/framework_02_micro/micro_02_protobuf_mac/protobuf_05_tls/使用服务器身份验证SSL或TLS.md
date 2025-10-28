https://www.liwenzhou.com/posts/Go/gRPC/#c-8-3

gRPC 内置支持 SSL/TLS，可以通过 SSL/TLS 证书建立安全连接，对传输的数据进行加密处理。

这里我们演示如何使用自签名证书进行server端加密。


# 生成证书
## 生成私钥

执行下面的命令生成私钥文件——server.key。[server.key](server.key)
```shell
openssl ecparam -genkey -name secp384r1 -out server.key
```

vim server.cnf[server.cnf](server.cnf)
```cnf
[ req ]
default_bits       = 4096
default_md		= sha256
distinguished_name = req_distinguished_name
req_extensions     = req_ext

[ req_distinguished_name ]
countryName                 = Country Name (2 letter code)
countryName_default         = CN
stateOrProvinceName         = State or Province Name (full name)
stateOrProvinceName_default = BEIJING
localityName                = Locality Name (eg, city)
localityName_default        = BEIJING
organizationName            = Organization Name (eg, company)
organizationName_default    = DEV
commonName                  = Common Name (e.g. server FQDN or YOUR name)
commonName_max              = 64
commonName_default          = github.com/nhjclxc

[ req_ext ]
subjectAltName = @alt_names

[alt_names]
DNS.1   = localhost
DNS.2   = github.com/nhjclxc
IP      = 127.0.0.1
```

之后执行以下命令即可生成[server.crt](server.crt)
```
openssl req -nodes -new -x509 -sha256 -days 3650 -config server.cnf -extensions 'req_ext' -key server.key -out server.crt
```

之后server端加载证书server.cert和秘钥server.key
Server端使用credentials.NewServerTLSFromFile函数分别加载证书server.cert和秘钥server.key
```go
creds, _ := credentials.NewServerTLSFromFile("./server.crt", "./server.key")
s := grpc.NewServer(grpc.Creds(creds))
lis, _ := net.Listen("tcp", "127.0.0.1:8972")
// error handling omitted
s.Serve(lis)
```

客户端使用以下代码进行认证
```go
creds, _ := credentials.NewClientTLSFromFile("./server.crt", "")
conn, _ := grpc.NewClient("127.0.0.1:8972", grpc.WithTransportCredentials(creds))
// error handling omitted
client := pb.NewGreeterClient(conn)
// ...
```