package main

import (
	"fmt"
	myUser "go-dev/src/test5_test/test57_package/user"
)

func main() {

	//user1 := &user.User{}
	user1 := &myUser.User{}
	fmt.Println(user1)
	user1.ID = 666
	user1.Name = "张三"
	fmt.Println(user1)

}
