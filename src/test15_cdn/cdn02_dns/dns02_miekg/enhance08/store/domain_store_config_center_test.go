package store

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestEtcdPut(t *testing.T) {
	// /cdn-dns/records/local.com/name       => "local.com"
	key := "/cdn-dns/records/local.com/name"
	value := "local.com"
	putResp, err := EtcdClient.Put(context.Background(), key, value)
	_ = err
	fmt.Println("操作成功：", putResp.Header.Revision)

}

func TestEtcdGet(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	key := "/cdn-dns/records/local.com/name"
	getResponse, err := EtcdClient.Get(ctx, key)
	if err != nil {
		fmt.Println("etcd put err", err)
		return
	}
	fmt.Println("getResponse.Header.ClusterId", getResponse.Header.ClusterId)
	fmt.Println("getResponse.Header.MemberId", getResponse.Header.MemberId)
	fmt.Println("getResponse.Header.String()", getResponse.Header.String())
	fmt.Println("getResponse.Header.Size()", getResponse.Header.Size())
	fmt.Println("getResponse.Header.Revision ", getResponse.Header.Revision)
	fmt.Println(getResponse.More)

	// getResponse.Count > 0 表示获取到了数据，并且遍历getResponse.Kvs可以获取到相应的数据
	fmt.Println("getResponse.Count 判断操作是否成功", getResponse.Count)
	for k, kv := range getResponse.Kvs {
		fmt.Println("k = ", k, "kv = ", kv, "kv byte value", string(kv.Value))
	}
	fmt.Println(getResponse.More)

}

func TestDomainStoreConfigCenter(t *testing.T) {

	ctx, cancle := context.WithCancel(context.Background())
	_ = cancle

	dscc := NewDomainStoreConfigCenter(ctx, EtcdClient)
	fmt.Println(dscc.domainMap)
	// etcdctl put foo bar
	// etcdctl get foo
	// etcdctl del foo
	// etcdctl put "/cdn-dns/records/local.com/name" "local.com"
	// etcdctl put "/cdn-dns/records/local.com/a/0" "127.0.0.1"
	// etcdctl get "/cdn-dns/records/local.com/name"
	// etcdctl del "/cdn-dns/records/local.com/name"

	//go func() {
	//	time.Sleep(5 * time.Second)
	//	cancle()
	//}()

	select {}
}

func Test22(t *testing.T) {
	eventBuffer := make([]string, 0)
	fmt.Println(eventBuffer)
	eventBuffer = append(eventBuffer, "a", "b", "c")
	fmt.Println(eventBuffer)
	eventBuffer = eventBuffer[:0]
	fmt.Println(eventBuffer)
}
