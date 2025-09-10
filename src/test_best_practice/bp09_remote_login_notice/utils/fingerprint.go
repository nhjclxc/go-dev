package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

// Fingerprinter 可配置的指纹生成器
type Fingerprinter struct {
	secret            []byte   // 用于 HMAC 的密钥，推荐传入
	headers           []string // 要包含的 header 列表（原始名字）
	cookies           []string // 要包含的 cookie 名称
	includeIP         bool     // 是否包含 c.ClientIP()
	includeProxyChain bool     // 是否包含 X-Forwarded-For 的代理链
	includeQuery      bool     // 是否包含有序的 query 参数
	extraHeader       string   // 额外包含某个 header 的值，例如前端 JS 指纹 'X-Client-FP'
	truncate          int      // 输出 hex 截断长度 (0 表示不截断)
	debug             bool     // debug 模式：会把 canonical string 写到 context (key: fingerprint_raw)
}

// Option 函数式配置选项
type Option func(*Fingerprinter)

// 构造器，secret 可为空（若为空将使用 sha256，但不推荐）
func NewFingerprinter(secret string, opts ...Option) *Fingerprinter {
	f := &Fingerprinter{
		secret:            []byte(secret),
		headers:           []string{"User-Agent", "Accept", "Accept-Language", "Accept-Encoding", "Referer"},
		cookies:           nil,
		includeIP:         true,
		includeProxyChain: false,
		includeQuery:      false,
		extraHeader:       "",
		truncate:          0,
		debug:             false,
	}
	for _, o := range opts {
		o(f)
	}
	return f
}

// option helpers
func WithCookies(c []string) Option       { return func(f *Fingerprinter) { f.cookies = c } }
func WithIncludeIP(v bool) Option         { return func(f *Fingerprinter) { f.includeIP = v } }
func WithIncludeProxyChain(v bool) Option { return func(f *Fingerprinter) { f.includeProxyChain = v } }
func WithIncludeQuery(v bool) Option      { return func(f *Fingerprinter) { f.includeQuery = v } }
func WithExtraHeader(name string) Option  { return func(f *Fingerprinter) { f.extraHeader = name } }
func WithTruncate(n int) Option           { return func(f *Fingerprinter) { f.truncate = n } }
func WithDebug(v bool) Option             { return func(f *Fingerprinter) { f.debug = v } }

// GetProxyChain2 : 返回代理链（基于 X-Forwarded-For / X-Real-Ip / RemoteAddr）
func GetProxyChain2(c *gin.Context) []string {
	forwardedFor := c.GetHeader("X-Forwarded-For")
	if forwardedFor != "" {
		parts := strings.Split(forwardedFor, ",")
		for i := range parts {
			parts[i] = strings.TrimSpace(parts[i])
		}
		return parts
	}
	xreal := strings.TrimSpace(c.GetHeader("X-Real-Ip"))
	if xreal != "" {
		return []string{xreal}
	}
	// fallback to RemoteAddr
	if addr := c.Request.RemoteAddr; addr != "" {
		if ip, _, err := net.SplitHostPort(addr); err == nil {
			return []string{ip}
		}
		return []string{addr}
	}
	return []string{""}
}

// buildCanonicalString: 将选定字段按确定顺序拼接成原始字符串（保证稳定性）
func (f *Fingerprinter) buildCanonicalString(c *gin.Context) string {
	parts := make([]string, 0)

	// headers（按字典序以稳定顺序）
	hdrs := make([]string, len(f.headers))
	copy(hdrs, f.headers)
	sort.Strings(hdrs)
	for _, h := range hdrs {
		v := strings.TrimSpace(c.GetHeader(h))
		parts = append(parts, strings.ToLower(h)+":"+v)
	}

	// cookies（按字典序）
	if len(f.cookies) > 0 {
		cks := make([]string, len(f.cookies))
		copy(cks, f.cookies)
		sort.Strings(cks)
		for _, name := range cks {
			val, err := c.Cookie(name)
			if err != nil {
				val = ""
			}
			parts = append(parts, "cookie:"+name+"="+val)
		}
	}

	// IP（ClientIP）
	if f.includeIP {
		ip := strings.TrimSpace(c.ClientIP())
		parts = append(parts, "ip:"+ip)
	}

	// 代理链
	if f.includeProxyChain {
		chain := GetProxyChain(c)
		parts = append(parts, "proxy_chain:"+strings.Join(chain, ","))
	}

	// query（按 key 排序，且 values 排序）
	if f.includeQuery {
		q := c.Request.URL.Query()
		keys := make([]string, 0, len(q))
		for k := range q {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			vals := q[k]
			sort.Strings(vals)
			parts = append(parts, "q:"+k+"="+strings.Join(vals, ","))
		}
	}

	// 额外客户端 header（例如前端 JS 计算并通过 X-Client-FP 发送）
	if f.extraHeader != "" {
		parts = append(parts, strings.ToLower(f.extraHeader)+":"+strings.TrimSpace(c.GetHeader(f.extraHeader)))
	}

	// 合并
	return strings.Join(parts, "|")
}

// Fingerprint 生成最终的 hex 字符串（HMAC-SHA256，如果 secret 为空则用 SHA256）
func (f *Fingerprinter) Fingerprint(c *gin.Context) string {
	raw := f.buildCanonicalString(c)

	var sum []byte
	if len(f.secret) > 0 {
		mac := hmac.New(sha256.New, f.secret)
		mac.Write([]byte(raw))
		sum = mac.Sum(nil)
	} else {
		s := sha256.Sum256([]byte(raw))
		sum = s[:]
	}

	hexs := hex.EncodeToString(sum)
	if f.truncate > 0 && f.truncate < len(hexs) {
		hexs = hexs[:f.truncate]
	}

	if f.debug {
		// 把原始字符串暴露到 context 方便调试（生产慎用）
		c.Set("fingerprint_raw", raw)
	}
	return hexs
}

// Middleware 返回一个 Gin 中间件：计算指纹并写入 context（key）或响应 header
// ctxKey: 如果非空，会在 c.Set(ctxKey, fingerprint)
// setHeader: 如果为 true，会在响应 header 中写入 X-Server-Fingerprint
func (f *Fingerprinter) Middleware(ctxKey string, setHeader bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		fp := f.Fingerprint(c)
		if ctxKey != "" {
			c.Set(ctxKey, fp)
		}
		if setHeader {
			c.Writer.Header().Set("X-Server-Fingerprint", fp)
		}
		c.Next()
	}
}

// GetFingerprintFromContext 从 context 读取
func GetFingerprintFromContext(c *gin.Context, key string) (string, bool) {
	v, ok := c.Get(key)
	if !ok {
		return "", false
	}
	s, ok := v.(string)
	return s, ok
}
