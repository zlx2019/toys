/**
  @author: Zero
  @date: 2023/4/25 20:01:30
  @desc: 数据包体

**/

package protocol

// Message 传输的数据包结构体
type Message struct {
	// 数据包ID 唯一标识。所占8个字节
	ID uint64
	// 具体的数据载体。所占字节非固定
	Payload []byte
}
