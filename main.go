package main

import (
	"fmt"
	"github.com/liuzhenchuan/algorithm/num"
)

func main() {
	// 统计素数
	n := num.Eratosthenes(100)
	fmt.Println(n)
}
