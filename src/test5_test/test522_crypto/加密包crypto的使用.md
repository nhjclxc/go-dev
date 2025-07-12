加密包crypto的使用


src/test8_framework/framework_08_safe还有

Go 的 `crypto` 标准库提供了多种哈希与加密算法，`import "crypto/sha1"` 和 `import "crypto/md5"` 属于 **哈希函数（Hash Functions）** 模块的一部分，主要用于摘要算法。下面是 Go 官方 `crypto` 相关库中提供的常见加密/哈希算法模块：

---

## 🔐 一、哈希函数模块（不可逆摘要算法）

适用于生成摘要、签名校验等：

| 包名              | 描述                                                       |
| --------------- | -------------------------------------------------------- |
| `crypto/md5`    | MD5 哈希函数，128位输出（16字节）                                    |
| `crypto/sha1`   | SHA-1 哈希函数，160位输出（20字节）                                  |
| `crypto/sha256` | SHA-256、SHA-224 等（SHA-2家族）                               |
| `crypto/sha512` | SHA-512、SHA-384、SHA-512/224/256                          |
| `hash/adler32`  | Adler-32 校验和，适用于快速校验                                     |
| `hash/crc32`    | CRC-32 校验码                                               |
| `hash/crc64`    | CRC-64 校验码                                               |
| `crypto/sha3`   | SHA-3（Keccak）家族算法，包括 SHA3-256、SHA3-512、Shake128、Shake256 |

---

## 🔐 二、对称加密算法（可逆）

适用于加密解密（需要密钥）：

| 包名              | 描述                          |
| --------------- | --------------------------- |
| `crypto/aes`    | AES 对称加密算法（推荐）              |
| `crypto/cipher` | 加密块封装接口，例如 AES 的 CBC、CFB 模式 |
| `crypto/des`    | DES、TripleDES（不推荐，已过时）      |

---

## 🔐 三、非对称加密算法（公私钥机制）

适用于签名、加解密：

| 包名                | 描述                    |
| ----------------- | --------------------- |
| `crypto/rsa`      | RSA 算法                |
| `crypto/ecdsa`    | ECDSA 椭圆曲线签名算法        |
| `crypto/elliptic` | 提供椭圆曲线实现（供 ECDSA 使用）  |
| `crypto/ed25519`  | Ed25519 快速签名算法（现代、推荐） |
| `crypto/dsa`      | DSA 签名算法（老旧，不推荐）      |

---

## 🔐 四、其他安全相关包

| 包名              | 描述                                       |
| --------------- | ---------------------------------------- |
| `crypto/hmac`   | 基于哈希的消息认证码（用于 HMAC-SHA256 等）             |
| `crypto/rand`   | 安全的随机数生成器                                |
| `crypto/tls`    | TLS 加密协议实现（HTTPS）                        |
| `crypto/x509`   | X.509 证书解析、验证                            |
| `crypto/pbkdf2` | 密码派生函数（在 `golang.org/x/crypto/pbkdf2` 中） |
| `crypto/subtle` | 常量时间比较函数，避免时间侧信道攻击                       |

---

## 📌 示例（SHA256 与 HMAC-SHA256）

```go
// sha256
import (
    "crypto/sha256"
    "fmt"
)

func main() {
    data := []byte("hello world")
    hash := sha256.Sum256(data)
    fmt.Printf("%x\n", hash)
}
```

```go
// HMAC-SHA256
import (
    "crypto/hmac"
    "crypto/sha256"
    "fmt"
)

func main() {
    key := []byte("my-secret-key")
    message := []byte("hello world")
    mac := hmac.New(sha256.New, key)
    mac.Write(message)
    expectedMAC := mac.Sum(nil)
    fmt.Printf("%x\n", expectedMAC)
}
```

---

如需推荐使用的加密方式：

* **哈希：** 推荐使用 `sha256` 或 `sha3`
* **对称加密：** 推荐 `AES` + `crypto/cipher`
* **非对称加密：** 推荐 `ed25519` 或 `rsa`
* **签名与认证：** 使用 `HMAC` 或 `ed25519`

需要我帮你封装一套常用的加密工具类吗？比如支持：MD5、SHA256、HMAC、AES 等？
