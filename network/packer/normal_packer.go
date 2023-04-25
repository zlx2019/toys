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
)

// NormalPacker 基础数据包处理实现
// IPacker 默认实现
type NormalPacker struct {
	// 字节顺序处理器
	ByteOrder binary.ByteOrder
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
	//计算数据包的总大小 8 + 8 + 数据载体
	totalSize := HeaderByteSize + IDByteSize + len(message.Payload)
	// 创建数据包的字节数组
	packs := make([]byte, totalSize)
	// 设置数据包头区域: 索引0-8位置,为数据包所占字节总大小
	packer.ByteOrder.PutUint64(packs[:HeaderByteSize], uint64(totalSize))
	// 设置数据包ID区域: 索引8-16位置内容为Message的ID
	packer.ByteOrder.PutUint64(packs[HeaderByteSize:HeaderByteSize+IDByteSize], message.ID)
	// 设置数据载体: 将数据载体设置到数据包的索引16(8+8)以后的空余位置
	copy(packs[HeaderByteSize+IDByteSize:], message.Payload)
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

func (packer NormalPacker) UnPackFromReader(reader io.Reader) (*protocol.Message, error) {
	//TODO implement me
	panic("implement me")
}

// NewNormalPacker 创建一个数据包工具
func NewNormalPacker() *NormalPacker {
	return &NormalPacker{ByteOrder: binary.BigEndian}
}
