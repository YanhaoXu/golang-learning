package main

import (
	"fmt"
	"log"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// NewServer 创建一个Server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
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
