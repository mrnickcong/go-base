package main

import (
	"fmt"
	"sync"
	"time"
)

/**
多线程通过共享内存来通信
*/
var syncMap = sync.Map{}

func readAndWriteGlobeMemory() {
	if value, exists := syncMap.Load("myKey"); exists {
		fmt.Println("value:", value)
	} else {
		syncMap.Store("myKey", "myValue")
	}
}
func main() {
	go readAndWriteGlobeMemory()
	go readAndWriteGlobeMemory()
	go readAndWriteGlobeMemory()
	go readAndWriteGlobeMemory()

	//休眠1秒钟
	time.Sleep(time.Second)
}
