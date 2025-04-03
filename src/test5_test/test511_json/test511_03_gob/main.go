package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

/*
Gob 是 Go 自己的以二进制形式序列化和反序列化程序数据的格式；

Gob 通常用于远程方法调用（RPCs，参见 15.9 节的 rpc 包）参数和结果的传输，以及应用程序和机器之间的数据传输。
Gob 特定地用于纯 Go 的环境中，例如，两个用 Go 写的服务之间的通信。这样的话服务可以被实现得更加高效和优化。
Gob 不是可外部定义，语言无关的编码方式。因此它的首选格式是二进制，而不是像 JSON 和 XML 那样的文本格式。
Gob 并不是一种不同于 Go 的语言，而是在编码和解码过程中用到了 Go 的反射。
*/

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main() {

	var network bytes.Buffer

	// 模拟这个为发送者
	enc := gob.NewEncoder(&network)

	pp := P{0, 1, 2, "zhansgan 里斯"}
	err1 := enc.Encode(pp)
	fmt.Println(err1)

	// 模拟这个为接收者
	dec := gob.NewDecoder(&network)

	pp2 := P{}
	fmt.Println(pp2)
	err2 := dec.Decode(&pp2)
	fmt.Println(err2)
	fmt.Println(pp2)

}
