/**
  @author: Zero
  @date: 2023/3/27 20:24:38
  @desc: 自定义扩展类型

**/

package toys

// Integer 整型
type Integer interface {
	Signed | Unsigned
}

// Signed 无符号整型
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned 有符号整型
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
