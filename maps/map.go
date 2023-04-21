/**
  @author: Zero
  @date: 2023/3/30 14:44:20
  @desc: Map结构函数库

**/

package maps

// Keys 获取一个Map所有的Key
func Keys[M ~map[K]V, K comparable, V any](maps M) []K {
	keys := make([]K, len(maps))
	for k := range maps {
		keys = append(keys, k)
	}
	return keys
}

// Values 获取一个Map所有的Value
func Values[M ~map[K]V, K comparable, V any](maps M) []V {
	values := make([]V, len(maps))
	for _, v := range maps {
		values = append(values, v)
	}
	return values
}

// ForEach 遍历map的元素,通过action函数进行处理.
func ForEach[M ~map[K]V, K comparable, V any](maps M, action func(K, V)) {
	for k, v := range maps {
		action(k, v)
	}
}

// Map 遍历map的元素,通过action函数进行处理,并返回处理后的结果.
func Map[M ~map[K]V, K comparable, V any](maps M, action func(K, V) V) []V {
	values := make([]V, len(maps))
	for k, v := range maps {
		values = append(values, action(k, v))
	}
	return values
}
