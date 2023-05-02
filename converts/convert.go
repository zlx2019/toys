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
	"fmt"
	"github.com/zlx2019/toys"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

// StringToBool String 转换为 Bool
func StringToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}

// ToBytes Any任意类型 转换为 []byte
func ToBytes(value any) ([]byte, error) {
	// 通过反射获取到value的值
	v := reflect.ValueOf(value)
	// 断言类型
	switch value.(type) {
	case bool:
		return strconv.AppendBool([]byte{}, v.Bool()), nil
	case string:
		return StringToBytes(v.String()), nil
	case []byte:
		return v.Bytes(), nil
	// Int类型处理
	case int, int64:
		// int 强转为int64处理
		return IntegerToBytes(v.Int())
	case int8:
		return IntegerToBytes(int8(v.Int()))
	case int16:
		return IntegerToBytes(int16(v.Int()))
	case int32:
		return IntegerToBytes(int32(v.Int()))
	// uint类型处理
	case uint, uint64:
		// uint直接强转uint64处理
		return IntegerToBytes(v.Uint())
	case uint8:
		return IntegerToBytes(uint8(v.Uint()))
	case uint16:
		return IntegerToBytes(uint16(v.Uint()))
	case uint32:
		return IntegerToBytes(uint32(v.Uint()))
	case float32, float64:
		return FloatToBytes(v.Float())
	default:
		newValue, err := AnyToJsonBytes(value)
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
		return BytesToString(val)
	default:
		str, err := AnyToJson(val)
		if err != nil {
			return ""
		}
		return str
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

// ToBool 判断一个值是否为零值。
// 将任意值转换为Bool(值为`类型零值`则为false)
// 指针类型 == nil为false，反之true
// int~ | uint~ | float~ | complex等类型 !=0 则为true,反之false
func ToBool(value any) bool {
	if value == nil {
		return false
	}
	v := reflect.ValueOf(value)
	// 如果是指针类型
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return true
		}
		//不为Nil则获取指针所指向的具体值的类型
		v = v.Elem()
	}
	// 断言值类型
	switch v.Kind() {
	case reflect.Bool:
		return v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Int() != 0
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0
	case reflect.Complex64, reflect.Complex128:
		return v.Complex() != 0
	case reflect.String:
		return v.String() != ""
	case reflect.Slice, reflect.Array, reflect.Map:
		return v.Len() != 0
	case reflect.Interface:
		itf := v.Interface()
		switch val := itf.(type) {
		case interface{ IsZero() bool }:
			return val.IsZero()
		default:
			return val != nil
		}
	default:
		return false
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

// StringToBytes 通过unsafe将String序列化为[]byte。性能高
func StringToBytes(value string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct {
		string
		Cap int
	}{value, len(value)}))
}

// BytesToString 通过unsafe将[]byte反序列化为String。性能非常高
func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

// IntegerToBytes Int~ | uint~ 类型转为 []byte
func IntegerToBytes[T toys.Integer](value T) ([]byte, error) {
	// 通过反射获取value的类型
	valueType := reflect.TypeOf(value).Kind()
	switch valueType {
	case reflect.Int:
		// 由于int是不固定字节大小类型,所以导致无法写入缓冲区
		// 统一将int类型转为int64固定大小字节类型再转换
		return intAndUintToBytes(int64(value))
	case reflect.Uint:
		// 由于uint是不固定字节大小类型,所以导致无法写入缓冲区
		// 统一将int类型转为uint64固定大小字节类型再转换
		return intAndUintToBytes(uint64(value))
	default:
		// 其余类型直接转换即可
		return intAndUintToBytes(value)
	}
}

// BytesToInteger []byte类型转为Int~类型
func BytesToInteger[T toys.Integer](value []byte) (T, error) {
	// 将字节数组读取到缓冲区
	buffer := bytes.NewBuffer(value)
	// 创建返回值
	var result T
	// 将数据写入返回值
	if err := binary.Read(buffer, binary.BigEndian, &result); err != nil {
		return result, err
	}
	return result, nil
}

// intAndUintToBytes int~ | uint~ 转为 []byte类型
func intAndUintToBytes[T toys.Integer](value T) ([]byte, error) {
	// 根据类型的字节大小,并且创建缓冲区
	buffer := bytes.NewBuffer(make([]byte, unsafe.Sizeof(value)))
	// 将数据写入到缓冲区
	if err := binary.Write(buffer, binary.BigEndian, value); err != nil {
		return nil, err
	}
	// 返回缓冲区数据字节
	return buffer.Bytes(), nil
}

// FloatToBytes float64类型转 []byte
func FloatToBytes(value float64) ([]byte, error) {
	bits := math.Float64bits(value)
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, bits)
	return bytes, nil
}

// BytesToFloat []byte 转float64
func BytesToFloat(value []byte) float64 {
	bits := binary.LittleEndian.Uint64(value)
	return math.Float64frombits(bits)
}
