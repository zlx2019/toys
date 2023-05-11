/**
  @author: Zero
  @date: 2023/5/9 20:02:52
  @desc: json格式序列化反序列化函数库(基于sonic)

**/

package jsons

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/zlx2019/toys/texts"
	"strings"
)

// sonic序列化处理器
var coder sonic.API

func init() {
	// 使用sonic最快效率配置
	//sonicCoder = sonic.ConfigFastest
	// 使用sonic默认配置(没有任何配置,宗旨以效率为主)
	//sonicCoder = sonic.ConfigDefault
	// 使用sonic标准配置(为了兼容json标准库)
	//sonicCoder = sonic.ConfigStd
	// 手动配置sonic配置
	coder = sonic.Config{
		// 序列化Json之前是否根据Key进行排序(会比较损耗性能)
		SortMapKeys: false,
		// 无需验证
		CompactMarshaler: true,
		// 所有的空数组和对象都转义为`[]`和`{}`,而不是null
		NoNullSliceOrMap: true,
		// 反序列化时,将整数映射为int64,而不是float64
		UseInt64: true,
		// 反序列化时,将数字都映射为数字类型,而不是float64
		UseNumber: false,
		// 反序列化时,如果有不匹配的字段是否返回错误
		DisallowUnknownFields: false,
		// 指定编码时,通过复制而不是引用字符串值
		CopyString: true,
		// 当 JSON 的字符串值中未转义控件字符 （\u0000-\u001f） 时，解码器将返回错误。
		ValidateString: true,
	}.Froze()
}

// Marshal 将一个值对象,序列化为Json字符串
func Marshal(value any) (string, error) {
	return coder.MarshalToString(value)
}

// MarshalBytes 将一个值对象,序列化为Json字节切片
func MarshalBytes(value any) ([]byte, error) {
	return coder.Marshal(value)
}

// Unmarshal 通过指定的泛型创建指定的值对象,将一个Json字符串反序列化写入到值对象中,并且返回
func Unmarshal[V any](json string) (*V, error) {
	var value V
	if err := coder.UnmarshalFromString(json, &value); err != nil {
		return nil, err
	}
	return &value, nil
}

// UnmarshalToValue 将一个Json字符串,反序列化写入到一个值对象中
func UnmarshalToValue(json string, value any) error {
	return coder.UnmarshalFromString(json, value)
}

// UnmarshalBytes 通过指定的泛型创建指定的值对象,将一个Json字节切片反序列化写入到值对象中,并且返回
func UnmarshalBytes[V any](bytes []byte) (*V, error) {
	var value V
	if err := coder.Unmarshal(bytes, &value); err != nil {
		return nil, err
	}
	return &value, nil
}

// UnmarshalBytesToValue 将一个Json字节切片,反序列化写入到一个值对象
func UnmarshalBytesToValue(bytes []byte, value any) error {
	return coder.Unmarshal(bytes, value)
}

// Valida 校验一个Json编码的字节格式是否有效
func Valida(bytes []byte) bool {
	return coder.Valid(bytes)
}

// GetNode 从一个Json字符串中,根据Path获取一个任意类型的节点值(通过闭包)
func GetNode[V any](json string, keyPath string, action func(ast.Node) (V, error)) (V, error) {
	var temp V
	if texts.IsEmpty(keyPath) {
		return temp, errors.New("json nodePath empty")
	}
	//将路径根据.分割,并且组装
	paths := pathSplit(keyPath)
	// 获取node
	node, err := sonic.GetFromString(json, paths...)
	if err != nil {
		return temp, err
	}
	// 通过泛型+闭包获取具体的返回值
	return action(node)
}

// GetString 从一个Json字符串中,根据path获取一个String类型的节点值
func GetString(json string, keyPath string) (string, error) {
	return GetNode[string](json, keyPath, func(node ast.Node) (string, error) {
		return node.String()
	})
}

// 将一个字符串根据`.`分割成切片,并且将切片中的每个元素转换为interface{}
// 例如: "a.b.c" -> []interface{}{"a","b","c"}
func pathSplit(keyPath string) []interface{} {
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
