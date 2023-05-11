/**
  @author: Zero
  @date: 2023/4/21 15:26:58
  @desc: Json序列化与反序列化

**/

package converts

import (
	"errors"
	"github.com/json-iterator/go"
	"strings"
)

const JsonTag = "json"

var coder jsoniter.API

// 初始化 自定义JSON序列化器配置
func init() {
	coder = jsoniter.Config{
		// 控制 JSON 格式化时每个缩进级别的空格数,默认为0
		IndentionStep: 0,
		//指定是否按字典序对 Map 的键进行排序。默认为 false
		SortMapKeys: false,
		// 指定是否对 HTML 特殊字符进行转义。
		EscapeHTML: true,
		// ：指定是否将数字解码为 Go 内置的 Number 类型而不是 float64。
		UseNumber: true,
		//指定在结构体字段中查找的 JSON 标签的名称。
		TagKey: JsonTag,
		// 指定是否只解析具有 json(TagKey) 标记的结构体字段。
		// 如果为false,则解析所有字段，包括未标记的字段,默认使用变量名作为json字段名
		OnlyTaggedField: false,
		// 指定是否验证 json.RawMessage 是否有效。
		ValidateJsonRawMessage: true,
		// 指定是否将结构体字段视为简单字符串类型，即不进行解码，而是直接将其作为字符串返回
		ObjectFieldMustBeSimpleString: false,
		// 指定是否区分字段名大小写。
		CaseSensitive: false,
	}.Froze()

}

// AnyToJson 将任意类型转换为JSON字符串
func AnyToJson(value any) (string, error) {
	return coder.MarshalToString(value)
}

// AnyToJsonBytes 将任意类型转换为JSON字节数组
func AnyToJsonBytes(value any) ([]byte, error) {
	return coder.Marshal(value)
}

// ReadJson 解析JSON字符串,并且映射到一个结构体中
// T 表示要映射到的结构体类型
func ReadJson[T any](json string) (T, error) {
	var target T
	err := coder.UnmarshalFromString(json, &target)
	if err != nil {
		return target, err
	}
	return target, nil
}

// ReadJsonToAny  解析JSON字符串,并且映射到一个结构体中
// 需要指定要映射的结构体指针
func ReadJsonToAny(json string, value any) error {
	return coder.UnmarshalFromString(json, value)
}

// ReadJsonBytes 解析JSON字节数组,并且映射到一个结构体中
// T 表示要映射到的结构体类型
func ReadJsonBytes[T any](bytes []byte) (T, error) {
	var target T
	err := coder.Unmarshal(bytes, &target)
	if err != nil {
		return target, err
	}
	return target, nil
}

// ReadJsonBytesToAny 解析JSON字节数组,并且映射到一个结构体中
// 需要指定要映射的结构体指针
func ReadJsonBytesToAny(bytes []byte, value any) error {
	return coder.Unmarshal(bytes, value)
}

// GetJsonNodeToString 在一个JSON字符串中,通过字段名的路径获取对应的值
func GetJsonNodeToString(json string, keyPath string) (string, error) {
	if len(keyPath) == 0 {
		return "", errors.New("keyPath is empty")
	}
	// 解析path,根据.分割
	keys := nodePathSplit(keyPath)
	// 获取节点
	node := coder.Get([]byte(json), keys...)
	// 判断节点是否存在
	if node.ValueType() == jsoniter.InvalidValue {
		return "", errors.New("json Node Not Found")
	}
	return node.ToString(), nil
}

// GetJsonNode 根据字段名的路径获取对应的值,使用泛型动态获取任意类型的值
func GetJsonNode[T any](json string, keyPath string, action func(jsoniter.Any) T) (T, error) {
	var result T
	if len(keyPath) == 0 {
		return result, errors.New("keyPath is empty")
	}
	// 解析path,根据.分割
	keys := nodePathSplit(keyPath)
	// 获取节点
	node := coder.Get([]byte(json), keys...)
	// 判断节点是否存在
	if node.ValueType() == jsoniter.InvalidValue {
		return result, errors.New("json Node Not Found")
	}
	return action(node), nil
}

func GetJsonNodeFromPaths(json string, paths ...string) (string, error) {
	return GetJsonNodeToString(json, strings.Join(paths, "."))
}

// 将一个字符串根据`.`分割成切片,并且将切片中的每个元素转换为interface{}
// 例如: "a.b.c" -> []interface{}{"a","b","c"}
func nodePathSplit(keyPath string) []interface{} {
	if !strings.Contains(keyPath, ".") {
		return []interface{}{keyPath}
	}
	paths := strings.Split(keyPath, ".")
	var keys = make([]interface{}, len(paths))
	for i, path := range paths {
		keys[i] = path
	}
	return keys
}
