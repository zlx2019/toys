/**
  @author: Zero
  @date: 2023/3/27 20:10:12
  @desc: 类型转换相关函数库

**/

package toys

import (
	"bytes"
	"encoding/binary"
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

// ToJson Any 转换为 String
func ToJson(value any) (string, error) {
	result, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(result), nil
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

//  数字类型转 []byte
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
