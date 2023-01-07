package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// NewUser 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// Online 用户的上线业务
func (this *User) Online() {
	// 用户上线，将用户加入到onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "user already online.")
}

// Offline 用户的下线业务
func (this *User) Offline() {
	// 用户下线，将用户从onlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前用户下线消息
	this.server.BroadCast(this, "user already offline.")
}

// SendMsg 给当前User对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// DoMessage 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户都有哪些

		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			this.SendMsg(fmt.Sprintf("[%s]%s:online...\n", user.Addr, user.Name))
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式：rename|张三
		newName := strings.Split(msg, "|")[1]
		// 判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("The current user name is in use\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("You have updated the user name:" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式：to|张三|消息内容

		// 1 获取对方的用户名 得到对方的User对象
		remoteName := strings.Split(msg, "|")[1]

		if remoteName == "" {
			this.SendMsg("The message format is not correct, please use the \"| to|zhangsan|hello \" format.\n")
			return
		}

		// 2 根据用户名 得到对方User对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("The user name does not exist\n")
			return
		}

		// 3 获取消息内容，通过对方的User对象将消息内容发送过去
		connect := strings.Split(msg, "|")[2]
		if connect == "" {
			this.SendMsg("No message, please resend\n")
			return
		}
		remoteUser.SendMsg(this.Name + "say to you:" + connect)

	} else {
		this.server.BroadCast(this, msg)
	}
}

// ListenMessage 监听当前的User channel的方法，一但有消息，就直接发送给对方的客户端
func (this User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
