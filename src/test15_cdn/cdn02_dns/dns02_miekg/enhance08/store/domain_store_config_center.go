package store

import (
	"context"
	"fmt"
	"github.com/miekg/dns"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 使用配置中心（Etcd）实现自定义域名的动态解析
// Etcd：强一致性、快速下发、 可靠性

/*
使用etcd下发配置文件，必须使用前缀树下发，如下：
| Key                              | Value                 |
| -------------------------------- | --------------------- |
| `/cdn-dns/records/local.com/ttl` | `"300"`               |
| `/cdn-dns/records/local.com/a`   | `["1.1.1.1"]`         |
| `/cdn-dns/records/*.local.com/a` | `["127.222.222.222"]` |


records:
  - domain: "local.com"
    ttl: 300
    answers:
      - qtype: "A"
        value: "127.0.0.1"
      - qtype: "A"
        value: "127.0.0.2"
      - qtype: "A"
        value: "127.0.0.3"
      - qtype: "AAAA"
        value: "::1"
      - qtype: "TXT"
        value: "hello=world"
      - qtype: "CNAME"
        value: "target.local.com."

转化为前缀树
/cdn-dns/records/local.com/name       => "local.com"
/cdn-dns/records/local.com/ttl        => "300"
/cdn-dns/records/local.com/a/0        => "127.0.0.1"
/cdn-dns/records/local.com/a/1        => "127.0.0.2"
/cdn-dns/records/local.com/a/2        => "127.0.0.3"
/cdn-dns/records/local.com/aaaa/0     => "::1"
/cdn-dns/records/local.com/txt/0      => "hello=world"
/cdn-dns/records/local.com/cname/0    => "target.local.com."


...

*/

type DomainStoreConfigCenter struct {
	ctx        context.Context
	etcdClient *clientv3.Client
	// key: domain
	// value: map[qtype/index]RR 例如 "a/0", "a/1"
	domainMap           map[string]map[string]dns.RR
	ttlMap              map[string]uint32 // domain → ttl
	nameMap             map[string]string // domain → name
	mu                  sync.Mutex        // 用于热更新控制
	keyPrefix           string
	initConfigReversion int64 // etcd的版本控制（类似mysql里面的乐观锁）
}

var defaultTtl uint32 = 300

// 不必关系 GetRecords 的实现
func (d *DomainStoreConfigCenter) GetRecords(qname string) ([]dns.RR, bool) {
	return nil, false
}

// 不必关系 Reload 的实现
func (d *DomainStoreConfigCenter) Reload() error {
	return nil
}

func NewDomainStoreConfigCenter(ctx context.Context, etcdClient *clientv3.Client) *DomainStoreConfigCenter {
	keyPrefix := "/cdn-dns/records/"
	dscc := DomainStoreConfigCenter{
		ctx:        ctx,
		etcdClient: etcdClient,
		keyPrefix:  keyPrefix,
		domainMap:  make(map[string]map[string]dns.RR),
		ttlMap:     make(map[string]uint32),
		nameMap:    make(map[string]string),
	}

	getResp, err := dscc.etcdClient.Get(dscc.ctx, dscc.keyPrefix, clientv3.WithPrefix())
	if err != nil {
		fmt.Println("etcd get err", err)
		return nil
	}
	dscc.initConfigReversion = getResp.Header.Revision
	dscc.buildRrMap(getResp)

	// 监听key变化
	dscc.Watch()

	return &dscc
}

func (d *DomainStoreConfigCenter) getDomainTtl(domain string) uint32 {
	ttl, found := d.ttlMap[domain]
	if found {
		return ttl
	}
	key := d.keyPrefix + domain + "/ttl"
	ttlStr, err := d.getEtcd(key)
	if err != nil {
		return defaultTtl
	}
	atoi, err := strconv.Atoi(ttlStr)
	if err != nil {
		return defaultTtl
	}
	d.ttlMap[domain] = uint32(atoi)
	return uint32(atoi)
}

func (d *DomainStoreConfigCenter) getDomainName(domain string) string {
	name, found := d.nameMap[domain]
	if found {
		return name
	}
	key := d.keyPrefix + domain + "/name"
	nameStr, err := d.getEtcd(key)
	if err != nil {
		return domain
	}
	d.nameMap[domain] = nameStr
	return nameStr
}

func (d *DomainStoreConfigCenter) getEtcd(key string) (string, error) {
	resp, err := d.etcdClient.Get(d.ctx, key)
	if err != nil {
		fmt.Println(" GetDomainTtl err", err)
		return "", nil
	}
	var value []byte
	for _, kv := range resp.Kvs {
		value = kv.Value
	}
	return string(value), err
}
func (d *DomainStoreConfigCenter) Watch() {
	go func() {
		fmt.Println("开始监听配置中心...")
		// 创建watch
		// clientv3.WithRev(s.initConfigReversion+1) 确保在初始化到watch这段时间内配置没有被更新过
		watchChan := d.etcdClient.Watch(d.ctx, d.keyPrefix, clientv3.WithPrefix(), clientv3.WithRev(d.initConfigReversion+1))
		fmt.Println("listening config change ...")

		for {
			select {
			case <-d.ctx.Done():
				fmt.Println("配置监听已被取消，退出 Watch")
				return
			case watchResponse, ok := <-watchChan:
				if !ok {
					fmt.Println("watchChan 已关闭，退出 Watch")
					return
				}
				for _, event := range watchResponse.Events {
					fmt.Printf("配置中心操作类型：%d, 操作配置：%s \n", event.Type, string(event.Kv.Key))
					switch event.Type {
					case clientv3.EventTypePut:
						// 新增或更新配置
						d.doBuildRrMap(clientv3.EventTypePut, string(event.Kv.Key), string(event.Kv.Value))
					case clientv3.EventTypeDelete:
						// 删除配置
						d.doBuildRrMap(clientv3.EventTypeDelete, string(event.Kv.Key), string(event.Kv.Value))
					}
				}
			}
		}
	}()
}
func (d *DomainStoreConfigCenter) buildRrMap(resp *clientv3.GetResponse) {
	d.mu.Lock()
	defer d.mu.Unlock()

	fmt.Println("resp.Count 判断操作是否成功", resp.Count)
	for _, kv := range resp.Kvs {
		d.doBuildRrMap(clientv3.EventTypePut, string(kv.Key), string(kv.Value))
	}
}

func (d *DomainStoreConfigCenter) doBuildRrMap(optType mvccpb.Event_EventType, key, value string) {
	/*
		前缀树
		/cdn-dns/records/local.com/name       => "local.com"
		/cdn-dns/records/local.com/ttl        => "300"
		/cdn-dns/records/local.com/a/0        => "127.0.0.1"
		/cdn-dns/records/local.com/a/1        => "127.0.0.2"
		/cdn-dns/records/local.com/a/2        => "127.0.0.3"
		/cdn-dns/records/local.com/aaaa/0     => "::1"
		/cdn-dns/records/local.com/txt/0      => "hello=world"
		/cdn-dns/records/local.com/cname/0    => "target.local.com."
	*/

	// key = /cdn-dns/records/local.com/name -> keySuffix = local.com/name
	// key = /cdn-dns/records/local.com/a/0  -> keySuffix = local.com/a/0
	keySuffix := strings.TrimPrefix(key, d.keyPrefix)
	domain := strings.Split(keySuffix, "/")[0]
	// keySuffix = local.com/name -> domainSuffix = name
	// keySuffix = local.com/a/0  -> domainSuffix = a/0
	domainSuffix := strings.TrimPrefix(keySuffix, domain+"/")

	// 处理删除
	if optType == clientv3.EventTypeDelete {
		if strings.Contains(domainSuffix, "/") {
			if domainMap, ok := d.domainMap[domain]; ok {
				delete(domainMap, domainSuffix)
				if len(domainMap) == 0 {
					delete(d.domainMap, domain)
				}
			}
		} else if domainSuffix == "name" {
			delete(d.nameMap, domain)
		} else if domainSuffix == "ttl" {
			delete(d.ttlMap, domain)
		}
		//return
	} else if optType == clientv3.EventTypePut {
		// 处理新增或修改
		if strings.Contains(domainSuffix, "/") {
			qtypeStr := strings.Split(domainSuffix, "/")[0]
			name := d.getDomainName(domain)
			ttl := d.getDomainTtl(domain)

			// 如果没有，则创建一个map来存储
			if _, ok := d.domainMap[domain]; !ok {
				d.domainMap[domain] = make(map[string]dns.RR)
			}
			var rr dns.RR

			switch qtypeStr {
			case "a":
				rr = &dns.A{
					Hdr: dns.RR_Header{Name: name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: ttl},
					A:   net.ParseIP(value),
				}
			case "aaaa":
				rr = &dns.AAAA{
					Hdr:  dns.RR_Header{Name: name, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: ttl},
					AAAA: net.ParseIP(value),
				}
			case "cname":
				rr = &dns.CNAME{
					Hdr:    dns.RR_Header{Name: name, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: ttl},
					Target: dns.Fqdn(value),
				}
			case "txt":
				rr = &dns.TXT{
					Hdr: dns.RR_Header{Name: name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: ttl},
					Txt: []string{value},
				}
			}

			// 往对应的map上放数据，新增或更新
			d.domainMap[domain][domainSuffix] = rr
		} else {
			if "name" == domainSuffix {
				d.nameMap[domain] = value
			} else if "ttl" == domainSuffix {
				ttl, err := strconv.Atoi(value)
				if err != nil {
					ttl = int(defaultTtl)
				}
				d.ttlMap[domain] = uint32(ttl)
			}
		}
	}

	fmt.Println("构建 rrMap...d.domainMap = ", d.domainMap)
	fmt.Println("构建 rrMap...d.nameMap = ", d.nameMap)
	fmt.Println("构建 rrMap...d.ttlMap = ", d.ttlMap)
	fmt.Println("------------------------------------------")
}

// Watch2 收集一个批量数据然后进行数据更新到内存的map中
func (d *DomainStoreConfigCenter) Watch2() {
	go func() {
		fmt.Println("开始监听配置中心...")
		watchChan := d.etcdClient.Watch(d.ctx, d.keyPrefix, clientv3.WithPrefix(), clientv3.WithRev(d.initConfigReversion+1))

		// 事件缓冲队列
		var eventBuffer []*clientv3.Event
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case watchResp, ok := <-watchChan:
				if !ok {
					fmt.Println("watchChan closed")
					return
				}
				eventBuffer = append(eventBuffer, watchResp.Events...)
			case <-ticker.C:
				if len(eventBuffer) == 0 {
					continue
				}
				// 批量处理
				d.processEventBatch(eventBuffer)
				// eventBuffer[:0] 是 Go 中 清空 slice 的标准做法。它会把 slice 的长度重置为 0，但 不会释放底层数组的内存，容量 (cap) 仍然保持不变。
				eventBuffer = eventBuffer[:0]
			case <-d.ctx.Done():
				fmt.Println("退出配置监听")
				return
			}
		}
	}()
}

// 批量处理函数
func (d *DomainStoreConfigCenter) processEventBatch(events []*clientv3.Event) {
	d.mu.Lock()
	defer d.mu.Unlock()

	// 按 domain 分组，批量更新
	for _, event := range events {
		key := string(event.Kv.Key)
		value := string(event.Kv.Value)
		d.doBuildRrMap(event.Type, key, value)
	}
}
