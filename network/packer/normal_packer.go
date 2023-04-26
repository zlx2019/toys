/**
  @author: Zero
  @date: 2023/4/25 20:23:42
  @desc: 基础的IPacker封包与解包工具

**/

package packer

import (
	"encoding/binary"
	"github.com/zlx2019/toys/network/protocol"
	"io"
	"net"
	"time"
)

// NormalPacker 基础数据包处理实现
// IPacker 默认实现
type NormalPacker struct {
	// 字节顺序处理器
	ByteOrder binary.ByteOrder
}

// NewNormalPacker 创建一个数据包工具
func NewNormalPacker() *NormalPacker {
	return &NormalPacker{ByteOrder: binary.BigEndian}
}

// Pack 数据打包
// 将一个Message数据包序列化为字节数组
// 数据包序列化后的结构:
//
//	主要分为三块结构: 		  [{数据头}|{数据ID}|{数据体}]
//	其结构内容分布:   		  [{数据载体字节长度}|{数据包ID}|{数据载体}]
//	分别所占字节:           [8字节|8字节|不固定大小]
func (packer NormalPacker) Pack(message *protocol.Message) ([]byte, error) {
	//读取 数据载体所占字节数量
	//计算数据包的总大小 (8 + 8 + 数据载体字节大小)
	totalSize := HeaderByteSize + IDByteSize + len(message.Payload)
	// 创建数据包的字节数组
	packs := make([]byte, totalSize)
	// 设置数据包头区域: 索引0-8位置,为数据包所占字节总大小
	packer.ByteOrder.PutUint64(packs[:HeaderLocation], uint64(totalSize))
	// 设置数据包ID区域: 索引8-16位置内容为Message的ID
	packer.ByteOrder.PutUint64(packs[HeaderLocation:IDLocation], message.ID)
	// 设置数据载体: 将数据载体设置到数据包的索引16(8+8)以后的空余位置
	copy(packs[IDLocation:], message.Payload)
	return packs, nil
}

// UnPack 数据解包
// 将字节数组反序列化为Message数据包
func (packer NormalPacker) UnPack(bytes []byte) (*protocol.Message, error) {
	// 读取数据包头(数据包总大小)[0:8]
	totalSize := packer.ByteOrder.Uint64(bytes[:HeaderByteSize])
	// 读取数据包ID [8:16]
	id := packer.ByteOrder.Uint64(bytes[HeaderByteSize : HeaderByteSize+IDByteSize])
	// 计算数据载体字节大小(数据包总大小 - 数据头大小 - 数据ID大小)
	payloadSize := totalSize - HeaderByteSize - IDByteSize
	// 读取数据载体,这样哪怕发生粘包,也不会读取多余的数据[16:16 + payloadSize]
	payload := bytes[HeaderByteSize+IDByteSize : HeaderByteSize+IDByteSize+payloadSize]
	return &protocol.Message{ID: id, Payload: payload}, nil
}

// UnPackReader 数据解包 (性能更高,推荐使用)
// 从一个可读字节流中读取数据,并且反序列化为Message数据包
func (packer NormalPacker) UnPackReader(reader io.Reader) (*protocol.Message, error) {
	// 创建可容纳数据包头部和数据包ID大小的字节数组
	buf := make([]byte, HeaderByteSize+IDByteSize)
	// 读取数据包头和数据包ID[0:16]
	// 将reader中的字节数据写入到buf中,直到把buf写满(8+8字节)
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return nil, err
	}
	// 读取数据包头(数据包总大小)[0:8]
	totalSize := packer.ByteOrder.Uint64(buf[:HeaderLocation])
	// 读取数据包ID [8:16]
	id := packer.ByteOrder.Uint64(buf[HeaderLocation:IDLocation])
	// 计算数据载体所占字节大小(数据包总大小 - 数据头 - 数据ID)
	payloadSize := totalSize - HeaderByteSize - IDByteSize
	// 创建固定大小的数据载体字节数组
	payloadBuf := make([]byte, payloadSize)
	// 读取数据载体,由于[]byte大小是已固定的,所以即使产生粘包也不会读取额外的数据内容
	_, err = io.ReadFull(reader, payloadBuf)
	if err != nil {
		return nil, err
	}
	return &protocol.Message{ID: id, Payload: payloadBuf}, nil
}

// UnPackConn 数据解包
// 从一个TCP连接中读取数据,并且反序列化为Message数据包
func (packer NormalPacker) UnPackConn(conn net.Conn, timeout time.Duration) (*protocol.Message, error) {
	// 设置读取该连接数据的超时时间
	err := conn.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return nil, err
	}
	return packer.UnPackReader(conn)
}
