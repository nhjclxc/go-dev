package store

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestDomainStoreConfig(t *testing.T) {
	filepath := "domain_store.yaml"
	//filepath = "cdn02_dns/dns02_miekg/enhance08/store/domain_store.yaml"
	//filepath = "test15_cdn/cdn02_dns/dns02_miekg/enhance08/store/domain_store.yaml"
	//filepath = "/Users/lxc20250729/lxc/code/go-dev/src/test15_cdn/cdn02_dns/dns02_miekg/enhance08/store/domain_store.yaml"

	ctx, cancle := context.WithCancel(context.Background())
	_ = cancle

	domainStoreConfig := NewDomainStoreConfig(ctx, filepath, 3*time.Second)

	//fmt.Println(domainStoreConfig)
	//fmt.Println(domainStoreConfig.config)
	fmt.Println(domainStoreConfig.filepath)
	//fmt.Println(domainStoreConfig.config.Upstream.Servers)

	//count := 0
	//for {
	//	select {
	//	default:
	//		time.Sleep(3 * time.Second)
	//		count++
	//		if count == 15 {
	//			return
	//		}
	//	}
	//}

}
