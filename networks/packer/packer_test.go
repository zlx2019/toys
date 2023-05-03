/**
  @author: Zero
  @date: 2023/4/25 21:28:15
  @desc: 单元测试

**/

package packer

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zlx2019/toys/networks/protocol"
	"net"
	"testing"
	"time"
)

// 测试[]byte To *Message
func TestNormalPacker_Pack(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	m := &protocol.Message{
		ID:      100001,
		Payload: []byte("Hello,~你好"),
	}
	packer := NewNormalPacker()
	// 封包
	pack, err := packer.Pack(m)
	is.NoError(err)
	fmt.Println(pack)

	// 解包
	message, err := packer.UnPack(pack)
	is.NoError(err)
	is.Equal(m.ID, message.ID)
	is.Equal(string(m.Payload), string(message.Payload))
	fmt.Println(string(m.Payload))
	fmt.Println(string(message.Payload))
}

// 测试Reader To *Message
func TestNormalPacker_UnPackReader(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	m := &protocol.Message{
		ID:      100001,
		Payload: []byte("Hello,~你好"),
	}
	packer := NewNormalPacker()
	// 封包
	pack, err := packer.Pack(m)
	is.NoError(err)
	fmt.Println(pack)

	// 解包
	reader := bytes.NewReader(pack)
	message, err := packer.UnPackReader(reader)
	is.NoError(err)
	is.Equal(m.ID, message.ID)
	is.Equal(string(m.Payload), string(message.Payload))
	fmt.Println(string(m.Payload))
	fmt.Println(string(message.Payload))
}

// 测试UnPackConn TCP客户端
func TestNormalPacker_UnPackTCPClient(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	conn, _ := net.Dial("tcp", ":9091")
	defer conn.Close()
	pack, _ := NewNormalPacker().Pack(&protocol.Message{ID: 100, Payload: []byte("你好")})
	_, err := conn.Write(pack)
	is.NoError(err)
}

// 测试UnPackConn TCP服务端
func TestNormalPacker_UnPackTCPServer(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":9091")
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	defer listener.Close()
	conn, _ := listener.Accept()
	message, err := NewNormalPacker().UnPackConn(conn, time.Second*3)
	is.NoError(err)
	is.Equal(message.ID, uint64(100))
	is.Equal(string(message.Payload), "你好")
	fmt.Println(message.ID)
	fmt.Println(string(message.Payload))
}

// 测试UnPackConn UDP客户端
func TestNormalPacker_UnPackUDPClient(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	conn, _ := net.Dial("udp", ":9092")
	defer conn.Close()
	pack, _ := NewNormalPacker().Pack(&protocol.Message{ID: 200, Payload: []byte("Hello,")})
	_, err := conn.Write(pack)
	is.NoError(err)
}

// 测试UnPackConn UDP服务端
func TestNormalPacker_UnPackConnUDPServer(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	udpAddr, _ := net.ResolveUDPAddr("udp", ":9092")
	listen, _ := net.ListenUDP("udp", udpAddr)
	defer listen.Close()
	buf := make([]byte, 1024*1024)
	// 读取数据 1mb
	_, err := listen.Read(buf)
	is.NoError(err)
	// 解包
	message, err := NewNormalPacker().UnPack(buf)
	is.NoError(err)
	is.Equal(message.ID, uint64(200))
	is.Equal(string(message.Payload), "Hello,")
}
