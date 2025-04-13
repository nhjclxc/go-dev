package main

import "fmt"

// 14.3 协程的同步：关闭通道-测试阻塞的通道
func main11() {
	/*
		通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。只有发送者需要关闭通道，接收者永远不会需要。
	*/

	var strChan chan string = make(chan string)

	var flagChan chan bool = make(chan bool)

	// sned
	go func(in chan<- string) {
		defer close(in)

		in <- "qaz1"
		in <- "qaz2"
		in <- "qaz3"
		in <- "qaz4"
		in <- "qaz5"

	}(strChan)

	// receiver
	go func(out <-chan string) {
		//for true {
		//	if val, ok := <-out; ok {
		//		fmt.Println("receiver: ", val)
		//	} else {
		//		break
		//	}
		//}

		for val := range out {
			fmt.Println("receiver: ", val)
		}

		flagChan <- true
	}(strChan)

	<-flagChan

}
