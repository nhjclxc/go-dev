package main

import (
	"fmt"
	"go-dev/src/test9_often_package/often_pkg_06_cache/test_anonymous_embedding/anonymous_user"
)

func main() {

	// 测试匿名嵌入的封装效果，参考：github.com/patrickmn/go-cache的Cache与cache的关系
	user := anonymous_user.NewUser("张三", 18)

	fmt.Printf("user = %#v, %#V \n", user, user)

	user.SetName("里斯")
	fmt.Printf("user = %#v, %#V \n", user, user)

}
