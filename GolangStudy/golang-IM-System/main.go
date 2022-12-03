package main

import "log"

func main() {
	log.Println("golang-IM-System Start...")
	server := NewServer("127.0.0.1", 8888)
	log.Println(server)
	//log.Println(*server)
	log.Println(&server)
	log.Println(&*server)
	log.Println(*&server)
	server.Start()
}
