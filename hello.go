package main

import (
	"fmt"
	"time"
)

func main() {
	var args string
	fmt.Printf("请输入数值：")
	fmt.Scanln(&args)
	start := time.Now()
	for i, arg := range args {
		fmt.Printf("%d\t%s\n", i, arg)

	}
	fmt.Printf("%.10fs", time.Since(start).Seconds())
}