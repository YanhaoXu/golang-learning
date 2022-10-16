package main

import (
	"gofile/read"
	"gofile/write"
	"log"
	"os"
)

func init() {
	// 将日志输出到标准输出
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetOutput(os.Stdout)
}

func main() {
	log.Println("main")

	// readDemo()
	writeDemo()
}

func readDemo() {
	// 读取文件 基本使用
	log.Println("读取文件 基本使用")
	read.BasicRead()

	// bufio读取文件
	log.Println("bufio读取文件")
	read.BufioRead()

	// ioutil.ReadFile读取整个文件
	log.Println("ioutil.ReadFile读取整个文件")
	read.IoutilRead()
}

func writeDemo() {
	// 写入文件 基本使用
	log.Println("写入文件 基本使用")
	write.BasicWrite()
}
