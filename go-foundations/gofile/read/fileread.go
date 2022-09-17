package read

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// ioutil.ReadFile读取整个文件
func IoutilRead() {

	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		log.Println("open file failed!, err:", err)
		return
	}

	log.Println(string(content))
}

// bufio读取文件
func BufioRead() {
	file, err := os.Open("./testFile.txt")
	if err != nil {
		log.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			if len(line) != 0 {
				log.Println(line)
			}
			log.Println("文件读完了")
			break
		}

		if err != nil {
			log.Println("read file failed, err:", err)
			return
		}

		log.Println(line)
	}

}

// 读取文件 基本使用
func BasicRead() {
	// 只读方式打开当前目录下的main。go文件
	file, err := os.Open("./main.go")
	if err != nil {
		log.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	defer file.Close()

	// 循环读取文件
	var content []byte
	// 使用Read方法读取数据
	var tmp = make([]byte, 128)
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			log.Println("文件读完了")
			break
		}

		if err != nil {
			log.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	log.Println(string(content))
}
