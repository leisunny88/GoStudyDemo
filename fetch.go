// fetch 输出从URL获取得内容
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main_() {
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: readng%s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// 函数io.Copy(dst,src),从src读，并且写入dst。使用它代替 ioutil.ReadAll来复制响应内容到os.Stdout, 这样
// 不需要装下整个响应数据流得缓冲区。确保检查io.Copy返回的错误结果

func main() {
	for _, url := range os.Args[1:]{
		resp, err := http.Get(url)
		if err != nil{
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(2)
		}
		out, err := os.Create("/json_file/out.txt")
		wt := bufio.NewWriter(out)
		b, err := io.Copy(wt, resp.Body)
		for i := 0; i<4; i++{
			fmt.Println("%s", i)
		}
		fmt.Println("write" , b)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading%s: %v\n", url, err)
			os.Exit(1)
		}
		wt.Flush()
		fmt.Printf("%s", b)
	}
}
// 修改fetch程序添加一个http://前缀（假如该URL参数缺失协议前缀）可能会用到strings.HasPrefix
func urlGet(url string) {
	resp, err := http.Get(url)
	if err != nil{
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(2)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil{
		fmt.Fprintf(os.Stderr, "fetch: readng%s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

func mainBL() {
	for _, url := range os.Args[1:]{
		if strings.HasPrefix(url, "http://"){  //strings.HasPrefix  判断是否以什么为前缀，返回bool类型
			urlGet(url)
		}else {
			url := "http://" + url
			urlGet(url)
		}
	}
}

// 修改fetch来输出Http得状态码， 可以在resp.Status中找到它

//func main() {
//	for _, url := range os.Args[1:]{
//
//	}
//}

