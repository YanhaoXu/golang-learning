package write

import (
	"log"
	"os"
)

func BasicWrite() {
	file, err := os.OpenFile("D:/DevOps/testData/testfile/111.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Println("open file failed, err:", err)
		return
	}

	defer file.Close()

	str := "hello world"
	file.Write([]byte(str))
	file.WriteString("hello golang")
}
