package main

import (
	"fmt"
)

var(
	s [100][10000]int  //（是否有）楼梯数组
    num [100][10000]int //房间内号码数组
	m int //共几层
	n int//每层几间
	p int//第几层
	r int//第几间
	count = 0//密码
	t int//第几个有楼梯的房间
)


func main(){
	fmt.Scanf("%d %d", &m, &n)
	Situ(m, n)
}

func Situ(m int, n int){
	var  q int
	for p=1; p<=m; p++ {
		for q = 0; q <= n; q++ {
			fmt.Scanln("%d %d", &s[p][q], &num[p][q])
		}
	}//每个房间的情况
	fmt.Scanln("%d", &r)
    for p=1; p<=m; p++ {
		Init(r)
	}//每层楼找房间
	fmt.Printf("密码是:%d", count)
}

func Init(r int)int {
	//一开始的房间是第r间，第t间有楼梯的房子
	var newr int
	count += num[p][r] //算密码

	if s[p][r] == 1 {
		t = 1
	} else {
		t = 0
	} //判断该层第一间是不是有楼梯的

	for ; t <= num[p][r]; r++ {
		if r > n {
			r = 0
		} //从头开始找

		if t == num[p][r] {
			newr = r
			break
		} //返回新房间的序号

		if num[p][r] == 1 {
			t++
		} else {
			continue
		}
	} //找该层第几间有楼梯的房间
	return newr
}

//思路：
//用循环把每个房间的情况存入数组，再用Init函数(入参为进入该层第一间房间的编号)寻找，每找一个房间则r++,若r值大于该层房间数则把r清零
//若房间有楼梯则t++,直到t等于第一间房间里的num，退出循环，并返回一个新的房间编号（r），并且楼层数p自增，再执行Init函数