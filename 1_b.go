package main

import "fmt"

func main() {
   var ch = make(chan string)
	go func() {
		ch <- "啊啊啊啊"
		fmt.Println("下山的路又堵起了")
	}()
    <- ch
}
