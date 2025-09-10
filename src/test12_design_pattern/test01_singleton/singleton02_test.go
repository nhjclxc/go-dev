package test01_singleton

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Singleton2 struct {
	Value int
	CTime time.Time
}

var singleton2Obj *Singleton2
var once sync.Once

func NewSingleton2() *Singleton {
	once.Do(func() {
		t := time.Now()
		fmt.Printf("第一次创建对象：%v \n", t)
		singletonObj = &Singleton{
			Value: t.Second(),
			CTime: t,
		}
	})
	return singletonObj
}

func Test012(t *testing.T) {
	s := NewSingleton2()
	fmt.Printf("val: %#v \n", s)
	s1 := NewSingleton2()
	fmt.Printf("val: %#v \n", s1)
	s2 := NewSingleton2()
	fmt.Printf("val: %#v \n", s2)
}

func Test022(t *testing.T) {
	for i := 0; i < 5; i++ {
		go func() {
			s := NewSingleton2()
			fmt.Printf("val: %#v \n", s)
		}()
	}

	time.Sleep(3 * time.Second)
}
