package main

import (
	"fmt"
	"io"
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

// ListenMessage 监听Message广播消息channel的goroutine，一但有消息就发送给全部的在线User
func (this *Server) ListenMessage() {
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

// BroadCast 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	this.Message <- fmt.Sprintf("[%s]%s:%s", user.Addr, user.Name, msg)
}

func (this *Server) Handler(conn net.Conn) {
	//...当前链接的业务
	log.Println("Successful link establishment...")

	user := NewUser(conn, this)

	user.Online()

	// 接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				log.Println("Conn Read err:", err)
				return
			}

			// 提取用户的消息（去除'\n'）
			msg := string(buf[:n-1])

			// 用户针对msg进行消息处理
			user.DoMessage(msg)
		}
	}()

	// 广播当前用户上线的消息
	this.BroadCast(user, "Already online")

	select {}
}

// Start 启动服务器的窗口
func (this *Server) Start() {
	// socket listen
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		log.Println("net.Lister err:", err)
	}
	// close listen socket
	defer listen.Close()

	// 启动监听Message的goroutine
	go this.ListenMessage()

	for {
		//accept
		conn, err := listen.Accept()
		if err != nil {
			log.Println("listener accept err:", err)
			continue
		}

		// do handler
		go this.Handler(conn)
	}

}
