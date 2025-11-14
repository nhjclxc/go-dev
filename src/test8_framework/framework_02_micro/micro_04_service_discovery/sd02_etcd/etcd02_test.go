package sd02_etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"strconv"
	"sync"
	"testing"
	"time"
)

// etcd 的watch操作

func TestWatch21(t *testing.T) {
	// watch操作
	// Watch = 监听一个 key 或前缀，一旦变化（Put/Delete）立即收到事件通知。
	// etcd 的 Watch 是通过 Raft 日志驱动的，所以它具有 强一致性（顺序一致的事件序列）。
	//
	//适用场景：
	//	服务注册/发现（监控健康 Key 是否消失），详细内容看：etcd02_watch01_service_register
	//	配置动态下发（监听配置 key） ，详细内容看：etcd02_watch02_dynamic_config
	//	分布式锁变化监听（与 etcd concurrency 配合），详细内容看：etcd02_watch03_lock
	//	大规模调度系统的事件监听（Kubernetes 就 heavily 依赖 watch）
	//
	// Watch重点：
	// 1）Watch 底层来自 Raft Log Index（Revision）每一次 etcd 写操作都会生成一个自增的 Revision：
	// 2）Watch 分为两类： 1普通 Watch：从当前 revision 继续看；2历史 Watch（带 Revision）：从指定 revision（历史）开始看
	// 3）Watch 的三层结构（非常重要）：Watcher → WatchableKV → Raft
	//		Wtach的流程：Raft Log 被 commit → WatchableKV 触发匹配 Watch → 推给客户端
	// 4）Watch 会保持长连接（gRPC streaming）：Watch 是一个 双向 gRPC 流，所以：无需频繁轮询，极低延迟，事件实时推送，断线自动重连

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer etcdClient.Close()
	wg := sync.WaitGroup{}

	watchKey := "watch-key3"
	watchChan := etcdClient.Watch(ctx, watchKey)

	wg.Add(1)
	go func() {
		// 监听协程
		for watchResponse := range watchChan {
			for index, event := range watchResponse.Events {
				// event.Type：PUT 或 DELETE
				fmt.Println("index", index, "event", event.PrevKv, event.Type, event.Kv)
				fmt.Println("create", event.IsCreate(), "modify", event.IsModify())

				if watchKey+"-value-2" == string(event.Kv.Value) {
					wg.Done()
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		// 写操作协程, 模拟写操作
		for i := range 3 {
			time.Sleep(time.Duration(i) * time.Second)
			etcdClient.Put(ctx, watchKey, watchKey+"-value-"+strconv.Itoa(i))
		}
		wg.Done()
	}()

	fmt.Println("etcd watching ...")
	wg.Wait()
	fmt.Println("etcd done")

}

func TestWatch22(t *testing.T) {
	//	使用etcd的watch实现：服务注册/发现（监控健康 Key 是否消失）

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer etcdClient.Close()
	wg := sync.WaitGroup{}

	watchKey := "micro-service-" // micro-service-1,micro-service-2,micro-service-3 ...
	watchChan := etcdClient.Watch(ctx, watchKey, clientv3.WithPrefix())

	wg.Add(1)
	go func() {
		// 监听服务是否在线
		for watchResponse := range watchChan {
			for _, event := range watchResponse.Events {
				if event.Type == mvccpb.PUT {
					// 微服务上线
					fmt.Printf("[Service] %s is alive \n", event.Kv.Key)
				} else {
					//} else if event.Type == mvccpb.DELETE {
					fmt.Printf("[Service] %s is offline \n", event.Kv.Key)
				}
			}
		}

	}()

	// 模拟三个微服务
	for i := range 3 {
		wg.Add(1)
		go func(sleepTime int) {
			// 每个微服务模拟3次上线下线操作
			microService := watchKey + strconv.Itoa(sleepTime)
			for i2 := range 6 {
				time.Sleep(time.Duration(sleepTime) * time.Second)
				if i2%2 == 0 {
					etcdClient.Put(ctx, microService, microService)
					fmt.Printf("[Client] %s online \n", microService)
				} else {
					etcdClient.Delete(ctx, microService)
					fmt.Printf("[Client] %s offline \n", microService)
				}
			}
			wg.Done()
		}(i + 1)
	}

	fmt.Println("etcd watching ...")
	wg.Wait()
	fmt.Println("etcd done.")

}

func TestWatch23(t *testing.T) {
	// Lease（租约）
	//租约是 etcd 的 TTL（Time-To-Live）机制
	//创建租约时可以指定过期时间，例如 5 秒
	//租约过期后，绑定在该租约上的 key 会 自动删除
	//作用：把 key 和租约关联，实现 自动过期 / 心跳机制

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer etcdClient.Close()

	// ======================
	// 4️⃣ 消费者 Watch 服务前缀
	// ======================
	watchChan := etcdClient.Watch(ctx, "/services/user-service/", clientv3.WithPrefix(), clientv3.WithPrevKV())

	go func() {
		for watchResp := range watchChan {
			for _, ev := range watchResp.Events {
				switch ev.Type {
				case clientv3.EventTypePut:
					fmt.Println("服务上线:", string(ev.Kv.Key), string(ev.Kv.Value))
				case clientv3.EventTypeDelete:
					fmt.Println("服务下线:", string(ev.Kv.Key))
				}
			}
		}
	}()

	// 模拟三个微服务
	for i := range 1 {
		go func(sleepTime int) {

			serviceKey := "/services/user-service/instance-" + strconv.Itoa(sleepTime)
			serviceValue := `{"ip":"10.0.0.1","port":3000}` + strconv.Itoa(sleepTime)

			// 1️⃣ 创建租约（TTL 5秒）
			leaseResp, _ := etcdClient.Grant(context.Background(), 3)

			// 2️⃣ Put key + lease，注册服务
			// clientv3.WithLease（租约）
			etcdClient.Put(ctx, serviceKey, serviceValue, clientv3.WithLease(leaseResp.ID))

			// 3️⃣ 自动续租，即向服务端发送心跳数据
			ch, _ := etcdClient.KeepAlive(context.Background(), leaseResp.ID)
			go func() {
				for ka := range ch {
					fmt.Println("续租成功，leaseID:", ka.ID)
				}
			}()

			time.Sleep(time.Duration(4) * time.Second)
			// 服务下线
			//etcdClient.Delete(ctx, serviceKey)
			// 停止续租
			etcdClient.Revoke(ctx, leaseResp.ID)

		}(i + 1)
	}

	// 模拟服务运行
	select {}
}
