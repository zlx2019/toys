/**
  @author: Zero
  @date: 2023/3/30 14:44:20
  @desc: Map结构函数库

**/

package toys

// Keys 获取一个Map所有的Key
func Keys[M ~map[K]V, K comparable, V any](value M) []K {
	keys := make([]K, len(value))
	for k := range value {
		keys = append(keys, k)
	}
	return keys
}

// Values 获取一个Map所有的Value
func Values[M ~map[K]V, K comparable, V any](value M) []V {
	values := make([]V, len(value))
	for _, v := range value {
		values = append(values, v)
	}
	return values
}
