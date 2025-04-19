package main

import (
	"log"
	"net"
	"strconv"
	"time"
)

func main() {

	// 1、打开一个对话
	dial, errDial := net.Dial("tcp", `localhost:8090`)
	if errDial != nil {
		log.Fatalln("client.errDial.err: ", errDial.Error())
		return
	}

	// 2、设置超时时间
	timeout := time.Now().Add(5 * time.Second)
	errSetDeadline := dial.SetDeadline(timeout)
	if errSetDeadline != nil {
		log.Fatalln("client.errSetDeadline.err: ", errSetDeadline.Error())
		return
	}

	// 3、发送、接收消息
	var exitChan chan bool = make(chan bool)
	var waitChan chan bool = make(chan bool)

	// 开一个协程发数据
	go func() {
		for i := 0; i < 10; i++ {
			_, errWrite := dial.Write([]byte("Ping" + strconv.Itoa(i)))
			if errWrite != nil {
				log.Fatalln("client.errWrite.err: ", errWrite.Error())
				return
			}
			waitChan <- true
		}

		exitChan <- true
	}()

	dial.Write([]byte("CLOSE_SERVER"))
	// 开一个协程不断读取数据
	go func() {
		for true {
			<-waitChan
			buf := make([]byte, 512)
			read, errRead := dial.Read(buf)
			if errRead != nil {
				log.Fatalln("client.errRead.err: ", errRead.Error())
				return
			}
			log.Println("client.read: size = ", read, ", read = ", string(buf))
		}
	}()

	<-exitChan
	// 最后关闭连接
	errClose := dial.Close()
	if errClose != nil {
		log.Fatalln("client.errClose.err: ", errClose.Error())
		return
	}

}
