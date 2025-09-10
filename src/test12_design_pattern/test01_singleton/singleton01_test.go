package test01_singleton

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Singleton struct {
	Value int
	CTime time.Time
}

var singletonObj *Singleton
var singletonLock sync.Mutex

func NewSingleton() *Singleton {
	singletonLock.Lock()
	defer singletonLock.Unlock()
	if singletonObj == nil {
		t := time.Now()
		fmt.Printf("第一次创建对象：%v \n", t)
		singletonObj = &Singleton{
			Value: t.Second(),
			CTime: t,
		}
	}
	return singletonObj
}

func Test01(t *testing.T) {
	s := NewSingleton()
	fmt.Printf("val: %#v \n", s)
	s1 := NewSingleton()
	fmt.Printf("val: %#v \n", s1)
	s2 := NewSingleton()
	fmt.Printf("val: %#v \n", s2)
}

func Test02(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			s := NewSingleton()
			fmt.Printf("val: %#v \n", s)
		}()
	}

	time.Sleep(3 * time.Second)
}
