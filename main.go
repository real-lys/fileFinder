package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var matches int = 0

func main() {
	var path string        //搜索路径
	if len(os.Args) == 1 { //如果用户没有指定路径，则默认为当前目录
		path = "./"
	} else {
		path = os.Args[1] // 读取命令行参数，path为用户指定的路径
	}

	if !strings.HasSuffix(path, "/") { //如果路径后没有“/”，则自动加上
		path += "/"
	}
	fmt.Print("input the filename: ")

	inputReader := bufio.NewReader(os.Stdin)
	query, _ := inputReader.ReadString('\n') //读取用户输入
	query = strings.TrimSuffix(query, "\n")

	var start time.Time = time.Now() //记录开始时间
	finder(path, query)              //搜索文件
	fmt.Println("the number of ", query, " is ", matches)
	fmt.Println(time.Since(start)) //总共用时
}

func finder(path string, query string) {
	files, err := os.ReadDir(path)
	if err == nil {
		for _, file := range files {
			filename := file.Name()
			if filename == query {
				matches++ //搜索到的文件数量+1
				var wg sync.WaitGroup
				wg.Add(1)
				go func() {
					fmt.Println(path + filename)
					wg.Done()
				}()
				wg.Wait()
			}
			if file.IsDir() {
				finder(path+filename+"/", query) //如果file为目录，则递归搜索目录
			}
		}
	}
}
