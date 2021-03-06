package afGoface

import "net"

type IConnection interface {

	//启动链接 让当前的链接准备工作

	Start()

	//停止链接

	Stop()
	//获取当前链接绑定的socket conn
	GetTCPConnection() *net.TCPConn

	//获取当前链接米快的Id
	GetConnID() uint32

	//获取远程客户端的TCP状态

	GetRemoteAddr() net.Addr
	//发送数据，讲数据发给远程的客户端

	SendMsg(msgId uint32, data []byte) error

	//设置连接属性
	SetProperty(key string, value interface{})

	//移除链接属性
	DelProperty(key string)

	//获取连接属性

	GetProperty(key string) (interface{}, error)
}

//定义一个处理链接业务的方法

type HandleFunc func(*net.TCPConn, []byte, int) error
