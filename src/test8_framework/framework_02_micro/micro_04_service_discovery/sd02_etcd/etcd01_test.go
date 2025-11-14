package sd02_etcd

import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

// etcd的基本操作

func Test10(t *testing.T) {
	_ = etcdClient
	defer etcdClient.Close()
}

// etcd是一个类似于redis的key-value数据库，一个key只能对应一个value，但是可以有key前缀来对应多个key
// key可以是一个确定的字符串，也可以是一个key的前缀，还可以是一个范围的key

func Test11(t *testing.T) {
	//etcd put

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer etcdClient.Close()

	putResponse, err := etcdClient.Put(ctx, "etcd-put", "etcd-put-value11")
	if err != nil {
		fmt.Println("etcd put err", err)

		if errors.Is(err, context.Canceled) {
			// ctx is canceled by another routine
		} else if errors.Is(err, context.DeadlineExceeded) {
			// ctx is attached with a deadline and it exceeded
		} else if errors.Is(err, rpctypes.ErrEmptyKey) {
			// client-side error: key is not provided
		} else if ev, ok := status.FromError(err); ok {
			code := ev.Code()
			if code == codes.DeadlineExceeded {
				// server-side context might have timed-out first (due to clock skew)
				// while original client-side context is not timed-out yet
			}
		} else {
			// bad cluster endpoints, which are not etcd servers
		}
		return
	}
	fmt.Println("putResponse.Header.ClusterId", putResponse.Header.ClusterId)
	fmt.Println("putResponse.Header.MemberId", putResponse.Header.MemberId)
	fmt.Println("putResponse.Header.String()", putResponse.Header.String())
	fmt.Println("putResponse.Header.Size()", putResponse.Header.Size())

	// putResponse.Header.Revision > 0表示put操作成功
	fmt.Println("putResponse.Header.Revision 判断操作是否成功", putResponse.Header.Revision)

}

func Test12(t *testing.T) {
	//etcd get

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer etcdClient.Close()

	getResponse, err := etcdClient.Get(ctx, "etcd-put")
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

func Test13(t *testing.T) {
	// etcd delete

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer etcdClient.Close()

	deleteResponse, err := etcdClient.Delete(ctx, "etcd-put")
	if err != nil {
		return
	}

	// deleteResponse.Deleted > 0，表示操作成功，数值即为删除key的数量，若=0则表示删除失败。
	// 当key是一个前缀或是一个范围时deleteResponse.Deleted可能会大于1
	fmt.Println("deleteResponse.Deleted 判断是否操作成功", deleteResponse.Deleted)
}

func Test151(t *testing.T) {
	//如何表示一个key前缀？

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	defer etcdClient.Close()

	etcdClient.Put(ctx, "put-key-a", "put-key-a-value")
	etcdClient.Put(ctx, "put-key-b", "put-key-b-value")
	etcdClient.Put(ctx, "put-key-c", "put-key-c-value")
	etcdClient.Put(ctx, "put-key-c-c", "put-key-c-c-value")
	etcdClient.Put(ctx, "put-key-c-c-c", "put-key-c-c-c-value")
	etcdClient.Put(ctx, "put-key-c-i-a", "put-key-c-i-a-value")

	// 获取一个前缀批量的数据
	// 使用clientv3.WithPrefix()表示要操作某一个前缀的key，具体操作有前面的方法体现（如Get，Delete，Watch等）
	// 只要前缀是put-key-，后面无论任何都可以匹配
	// 前缀查询本质上是：[key, key + next_possible_byte)
	getResp, err := etcdClient.Get(ctx, "put-key-", clientv3.WithPrefix())
	if err != nil {
		fmt.Println("etcd get err", err)
		return
	}
	fmt.Println("count", getResp.Count)
	for key, keyValue := range getResp.Kvs {
		fmt.Println("key", key, "keyValue", string(keyValue.Value))
	}

}

func Test152(t *testing.T) {
	//如何表示一个范围的key？

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	defer etcdClient.Close()

	etcdClient.Put(ctx, "a", "a-value")
	etcdClient.Put(ctx, "b", "b-value")
	etcdClient.Put(ctx, "c", "c-value")
	etcdClient.Put(ctx, "h", "h-value")
	etcdClient.Put(ctx, "x", "x-value")
	etcdClient.Put(ctx, "y", "y-value")

	// key的范围匹配
	// 使用clientv3.WithRange()来表示此次要操作一个范围的key，其中的参数是endKey，而此时前面操作函数的key即为startKey
	//etcd 范围查询是：[startKey, endKey)
	getResp, err := etcdClient.Get(ctx, "a", clientv3.WithRange("z"))
	if err != nil {
		fmt.Println("etcd get err", err)
		return
	}
	fmt.Println("getResp.Count", getResp.Count)
	for key, keyValue := range getResp.Kvs {
		fmt.Println("key", key, "keyValue", string(keyValue.Value))
	}

}
