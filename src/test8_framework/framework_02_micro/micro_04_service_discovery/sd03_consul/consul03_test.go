package sd03_consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"testing"
	"time"
)

// consul实现分布式锁
func Test31(t *testing.T) {

	// 1、创建session会话
	name := "web"

	// 给name对应的key上锁，
	//session存活时间是10s，
	//session到期时的操作api.SessionBehaviorRelease
	//当某个节点释放锁时，Consul 会等待 LockDelay 再允许其他节点获取锁
	sessionEntry := api.SessionEntry{
		Name:      name,
		TTL:       "10s",
		Behavior:  api.SessionBehaviorRelease,
		LockDelay: 0,
	}
	sessionId, meta, err := ConsulClient.Session().Create(&sessionEntry, nil)
	if err != nil {
		fmt.Println("创建session失败！", err)
		return
	}
	fmt.Println(meta.RequestTime)

	// 2、申请锁
	key := "session-key"
	kvPair := api.KVPair{
		Key:     key,
		Value:   []byte("session-value"),
		Session: sessionId,
	}
	acquire, _, err := ConsulClient.KV().Acquire(&kvPair, nil)
	if err != nil {
		fmt.Println("上锁失败！", err)
		return
	}
	if acquire {
		fmt.Println("上锁成功")

		// 3s后释放锁
		time.Sleep(time.Second * 3)

		kvPair2 := api.KVPair{
			Key:     key,
			Session: sessionId,
		}
		release, _, err := ConsulClient.KV().Release(&kvPair2, nil)
		if err != nil {
			fmt.Println("释放锁失败！", err)
			return
		}
		fmt.Println("释放锁结果：", release)

	} else {
		fmt.Println("上锁失败")
	}

}

// 封装成方法
// sessionName会话名称
// string：返回会话id
func CreateSession(sessionName string) (string, error) {
	sessionEntry := api.SessionEntry{
		Name:      sessionName,
		LockDelay: 0,
		Behavior:  api.SessionBehaviorRelease,
		TTL:       "10s",
	}
	create, _, err := ConsulClient.Session().Create(&sessionEntry, nil)
	return create, err
}

// 申请锁
func Acquire(sessionId string) {

}
