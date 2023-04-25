/**
  @author: Zero
  @date: 2023/4/25 16:45:50
  @desc: 网络通信数据包工具接口,用于对数据进行封包与解包。
**/

package packer

import (
	"github.com/zlx2019/toys/network/protocol"
	"io"
)

const (
	// HeaderByteSize 数据头部所占字节数
	HeaderByteSize = 8
	// IDByteSize 数据ID所占字节数
	IDByteSize = 8
)

// IPacker 数据包处理接口
type IPacker interface {
	// Pack 数据打包
	// 将一个Message数据包序列化为字节数组
	Pack(message *protocol.Message) ([]byte, error)

	// UnPack 数据解包
	// 将字节数组反序列化为Message数据包
	UnPack(bytes []byte) (*protocol.Message, error)

	// UnPackFromReader 数据解包
	// 从一个可读字节流中读取数据,并且反序列化为Message数据包
	UnPackFromReader(reader io.Reader) (*protocol.Message, error)
}
