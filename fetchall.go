// fetchall 并发获取URL并报告它们得时间和大小
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:]{
		go fetch(url, ch)  // 启动一个goroutine
	}
	for range os.Args[1:]{
		fmt.Println(<-ch) // 从通道ch接受
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // 发送到通道ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %d   %s", secs, nbytes, url)
}


// 找一个产生大量数据得网站，连续两次运行fetchall, 看报告得时间是否会有大得变化， 调查缓存情况。每一次获取得内容
// 一样吗？修改fetchall将内容输出到文件


// 使用更长的参数列表来尝试