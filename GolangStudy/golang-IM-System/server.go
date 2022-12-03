package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的Channel
	Message chan string
}

// NewServer 创建一个Server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// ListenMessager 监听Message广播消息channel的goroutine，一但有消息就发送给全部的在线User
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		// 将消息发送给全部的在线User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := fmt.Sprintf("[%s]%s:%s", user.Addr, user.Name, msg)
	this.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	log.Println("链接建立成功...")
}

// Start 启动服务器的窗口
func (s *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		log.Println("net.Lister err:", err)
	}
	// close listen socket
	defer listen.Close()

	for {
		//accept
		conn, err := listen.Accept()
		if err != nil {
			log.Println("listener accept err:", err)
			continue
		}

		// do handler
		go s.Handler(conn)
	}

}
