/**
  @author: Zero
  @date: 2023/3/27 20:10:12
  @desc: 类型转换相关函数库

**/

package converts

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
)

// StringToBool String 转换为 Bool
func StringToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

// ToBytes Any 转换为 []byte
func ToBytes(value any) ([]byte, error) {
	// 通过反射获取到value的值
	v := reflect.ValueOf(value)
	// 断言类型
	switch value.(type) {
	case bool:
		return strconv.AppendBool([]byte{}, v.Bool()), nil
	case string:
		return []byte(v.String()), nil
	case []byte:
		return v.Bytes(), nil
	case int, int8, int16, int32, int64:
		return numeralToBytes(binary.Write, binary.BigEndian, v.Int())
	case uint, uint8, uint16, uint32, uint64:
		return numeralToBytes(binary.Write, binary.BigEndian, v.Uint())
	case float32:
		return floatToBytes(v.Float())
	case float64:
		return floatToBytes(v.Float())
	default:
		newValue, err := json.Marshal(value)
		return newValue, err
	}
}

// ToString Any 转换为 String
func ToString(value any) string {
	if value == nil {
		return ""
	}
	switch val := value.(type) {
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case int:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case string:
		return val
	case []byte:
		return string(val)
	default:
		b, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		return string(b)
	}
}

// ToInt Any 转换为 Int
func ToInt(value any) (int64, error) {
	v := reflect.ValueOf(value)
	var result int64
	err := fmt.Errorf("ToInt: invalid value type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = v.Int()
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = int64(v.Uint())
		return result, nil
	case float32, float64:
		result = int64(v.Float())
		return result, nil
	case string:
		result, err = strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			result = 0
		}
		return result, err
	default:
		return result, err
	}
}

// SliceToMap 切片 转换为 Map
// T 切片元素类型
// K Map-Key类型
// V Map-Value类型
// operator 实现K和V,分别使用T的哪些属性。
func SliceToMap[T any, K comparable, V any](slice []T, operator func(T) (K, V)) map[K]V {
	maps := make(map[K]V, len(slice))
	for _, item := range slice {
		k, v := operator(item)
		maps[k] = v
	}
	return maps
}

// MapToSlice Map 转换为 切片
func MapToSlice[T any, K comparable, V any](maps map[K]V, operator func(K, V) T) []T {
	slice := make([]T, 0, len(maps))
	for k, v := range maps {
		slice = append(slice, operator(k, v))
	}
	return slice
}

// ToReadOnlyChannel 将一个切片异步转换成一个只读通道
func ToReadOnlyChannel[T any](list []T) <-chan T {
	channel := make(chan T, len(list))
	go func() {
		for _, item := range list {
			channel <- item
		}
		close(channel)
	}()
	return channel
}

// EncoderBytes 将对象序列化为[]byte
func EncoderBytes(value any) ([]byte, error) {
	// 创建一个缓冲区
	buffer := bytes.NewBuffer(nil)
	// 创建一个编码器
	encoder := gob.NewEncoder(buffer)
	// 进行编码
	if err := encoder.Encode(value); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// DecoderBytes 将[]byte反序列化 到一个对象中
func DecoderBytes(values []byte, target any) error {
	// 创建一个缓冲区
	buffer := bytes.NewBuffer(values)
	// 创建一个解码器
	decoder := gob.NewDecoder(buffer)
	// 解码
	return decoder.Decode(target)
}

// 数字类型转 []byte
func numeralToBytes(writeFunc func(writer io.Writer, order binary.ByteOrder, data interface{}) error, byteOrder binary.ByteOrder, value interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	err := writeFunc(buf, byteOrder, value)
	return buf.Bytes(), err
}

// float类型转 []byte
func floatToBytes(value float64) ([]byte, error) {
	bits := math.Float64bits(value)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)
	return bytes, nil
}
