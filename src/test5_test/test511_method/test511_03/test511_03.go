package main

// 方法集
// https://topgoer.com/方法/方法集.html
func main0() {
	/*
		Golang方法集 ：每个类型都有与之关联的方法集，这会影响到接口实现规则。
			• 类型 T 方法集包含全部 receiver T 方法。
			• 类型 *T 方法集包含全部 receiver T + *T 方法。
			• 如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。
			• 如类型 S 包含匿名字段 *T，则 S 和 *S 方法集包含 T + *T 方法。
			• 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。
	*/

}
