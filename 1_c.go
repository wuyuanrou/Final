package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main(){
	wg.Add(9)
	go Sentence("啊这")
	go Sentence("第二个")
	go Sentence("第三个")
	go Sentence("第四个")
	go Sentence("第五个")
	go Sentence("第六个")
	go Sentence("第七个")
	go Sentence("第八个")
	go Sentence("第九个")
	Sentence("第十个")
	wg.Wait()
	fmt.Println("子协程打印完啦")
}

func Sentence(s string){
	fmt.Println(s)
	wg.Done()
}