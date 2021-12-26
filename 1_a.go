package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()                     //mu.Lock()应该在执行语句前，才能保证同时只有一个协程访问该语句
	}()
        mu.Unlock()
}
