package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		}
	}
}
