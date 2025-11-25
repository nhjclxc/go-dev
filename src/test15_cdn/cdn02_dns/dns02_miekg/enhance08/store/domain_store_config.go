package store

import (
	"context"
	"fmt"
	"github.com/miekg/dns"
	"github.com/spf13/viper"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// 动态配置文件实现自定义域名的动态解析

// DomainStoreConfig 想
type DomainStoreConfig struct {
	ctx       context.Context
	filepath  string
	reload    time.Duration
	config    *Config
	domainMap map[string][]dns.RR // domain → RR 列表
	mu        sync.Mutex          // 用于热更新控制
}

func NewDomainStoreConfig(ctx context.Context, filepath string, reload time.Duration) *DomainStoreConfig {
	if filepath == "" {
		log.Fatal("配置文件路径为空！！！")
		return nil
	}
	dsc := &DomainStoreConfig{
		ctx:       ctx,
		filepath:  filepath,
		reload:    reload,
		config:    &Config{},
		domainMap: make(map[string][]dns.RR),
	}
	err := dsc.Load()
	// 开启热更新
	go dsc.ReReload()

	if err != nil {
		fmt.Println("配置加载错误：", err)
	}
	return dsc
}

func (d *DomainStoreConfig) GetRecords(qname string) ([]dns.RR, bool) {
	_, rrRecords, found := d.matchDomain(qname)
	_ = rrRecords
	if found {
		// 自定义域名 customDnsQuery
		// ...
	} else {
		// 其他域名 upstreamQuery
		// ...
	}

	return nil, false
}

// matchDomain *匹配域名
// matchedDomain, rrRecords, found
func (d *DomainStoreConfig) matchDomain(qname string) (string, []dns.RR, bool) {
	// 1、域名完全匹配
	rrRecords, found := d.domainMap[qname]
	if found {
		return qname, rrRecords, true
	}

	// 2、*匹配
	for domain, rrs := range d.domainMap {
		if strings.HasPrefix(domain, "*") {
			// *.local.com. *.example.com. ...
			// 转化为：local.com.  example.com.
			baseDomain := domain[2:]
			if strings.HasSuffix(qname, baseDomain) {
				return domain, rrs, true
			}
		}
	}

	return "", nil, false
}

func (d *DomainStoreConfig) Load() error {
	d.mu.Lock()
	defer d.mu.Unlock()

	v := viper.New()            // 创建viper对象
	v.SetConfigFile(d.filepath) // 设置配置文件地址
	v.SetConfigType("yaml")

	// 读取配置
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	//// 🔁 热更新配置（这是一个可选配置）
	//d.v.WatchConfig()
	//d.v.OnConfigChange(func(e fsnotify.Event) {
	//	// 配置更新后回进入这个回调函数
	//	fmt.Println("config file changed:", e.Name)
	//	// 配置更新后，将 Viper 中已加载的配置数据（比如从 config.yaml 中）解析并填充到结构体 Global.GlobalConfig 中。
	//	if err = d.v.Unmarshal(&d.config); err != nil {
	//		fmt.Println(err)
	//	}
	//})
	// 将 Viper 中已加载的配置数据（比如从 config.yaml 中）解析并填充到结构体 global.GVA_CONFIG 中。
	if err = v.Unmarshal(&d.config); err != nil {
		log.Fatal(fmt.Errorf("fatal error unmarshal config: %w", err))
		return err
	}

	// 构造map[domain] -> []dns.RR
	d.buildRR()

	return nil
}

// ReReload 执行热加载
func (d *DomainStoreConfig) ReReload() {

	if d.reload == 0 {
		fmt.Println("not set")
	}

	ticker := time.NewTicker(d.reload)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Println("重新加载配置文件")
			d.Load()
			fmt.Println(d.config.Upstream.Servers)
		case <-d.ctx.Done():
			fmt.Println("程序关闭，退出配置热更新！")
			return
		}
	}

}

func (d *DomainStoreConfig) buildRR() {

	// 构建 RR Map
	newRRMap := make(map[string][]dns.RR)

	for _, rec := range d.config.Records {
		domain := dns.Fqdn(rec.Domain) // 自动补全结尾的 点.

		for _, ans := range rec.Answers {

			var rr dns.RR

			switch ans.Qtype {
			case "A":
				rr = &dns.A{
					Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: rec.Ttl},
					A:   net.ParseIP(ans.Value),
				}

			case "AAAA":
				rr = &dns.AAAA{
					Hdr:  dns.RR_Header{Name: domain, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: rec.Ttl},
					AAAA: net.ParseIP(ans.Value),
				}

			case "CNAME":
				rr = &dns.CNAME{
					Hdr:    dns.RR_Header{Name: domain, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: rec.Ttl},
					Target: dns.Fqdn(ans.Value),
				}

			case "TXT":
				rr = &dns.TXT{
					Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: rec.Ttl},
					Txt: []string{ans.Value},
				}

			default:
				log.Println("未知 Qtype:", ans.Qtype)
				continue
			}

			newRRMap[domain] = append(newRRMap[domain], rr)
		}
	}

	d.domainMap = newRRMap
}
