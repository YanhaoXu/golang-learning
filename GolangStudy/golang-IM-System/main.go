package main

import "log"

func main() {
	log.Println("golang-IM-System Start...")
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
