package main

import (
	"bufio"
	"fmt"
	"os"
)

// 控制台输入
//func main() {
//	args := ""
//	fmt.Printf("请输入:")
//	fmt.Scanln(&args)
//	counts := make(map[string]int)
//	input := bufio.NewScanner(strings.NewReader(args))
//	for input.Scan(){
//		counts[input.Text()]++
//	}
//	for line, n := range counts{
//		fmt.Printf("%d\t%s\n", line, n)
//		if n>1 {
//			fmt.Printf("%d\t%s\n", line, n)
//		}
//	}
//}

// 多次输入重复的行个数和文本
func main(){
	counts := make(map[string]int)
	//filePath := "./json_file/dup.txt"
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "")
	} else {
		for _, arg := range files{
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
		}
		for line, n := range counts{
			//fmt.Printf("%d\t%s\n", n, line)
			if n>1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}

func countLines(f *os.File, counts map[string]int, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()+"\t"+arg]++   // 等价 line := input.Text()   counts[line] = counts[line] + 1
	}
}
