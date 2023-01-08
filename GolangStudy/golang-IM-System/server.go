package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
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
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message

		// 将消息发送给全部的在线User
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

// BroadCast 广播消息的方法
func (s *Server) BroadCast(user *User, msg string) {
	s.Message <- fmt.Sprintf("[%s]%s:%s", user.Addr, user.Name, msg)
}

func (s *Server) Handler(conn net.Conn) {
	//...当前链接的业务
	log.Println("Successful link establishment...")

	user := NewUser(conn, s)

	user.Online()
	// 监听用户是否活跃的channel
	isLive := make(chan bool)

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
	s.BroadCast(user, "Already online")

	// 当前handle阻塞
	for {
		select {
		case <-isLive:
			// 当前用户是活跃的，应该重置定时器
			// 不做任何事情，为了激活select，更新下面的定时器

		case <-time.After(time.Second * 10):
			// 已经超时
			// 将当前的User强制的关闭
			user.SendMsg("You've been forced offline")

			// 销毁用户的资源
			close(user.C)

			// 关闭连接
			err := conn.Close()
			if err != nil {
				return
			}
			// 退出当前的Handler
			return // runtime.Goexit()
		}

	}
}

// Start 启动服务器的窗口
func (s *Server) Start() {
	// socket listen
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		log.Println("net.Lister err:", err)
	}
	// close listen socket
	defer listen.Close()

	// 启动监听Message的goroutine
	go s.ListenMessage()

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
