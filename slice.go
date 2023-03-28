/**
  @author: Zero
  @date: 2023/3/27 15:39:05
  @desc: 切片相关操作函数库

**/

package toys

// Contains 切片中是否包含某一个元素
// value  要匹配的元素
func Contains[T comparable](slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// ContainsBy 切片中是否包含满足谓词条件的元素
// predicate 谓词条件函数
func ContainsBy[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}

// ContainsSub 切片中是否包含另一个子切片(所有元素)
func ContainsSub[T comparable](slice, subSlice []T) bool {
	for _, item := range subSlice {
		if !Contains(slice, item) {
			// 有一个不存在则返回false
			return false
		}
	}
	return true
}

// Filter 从一个切片中过滤出符合条件的元素
// predicate 过滤谓词条件函数
func Filter[T any](slice []T, predicate func(T) bool) []T {
	list := make([]T, 0)
	for _, item := range slice {
		if predicate(item) {
			list = append(list, item)
		}
	}
	return list
}

// Chunk 将一个切片,按数量分为多个切片
func Chunk[T any](slice []T, size int) [][]T {
	result := [][]T{}
	if len(slice) == 0 || size <= 0 {
		return result
	}
	for _, item := range slice {
		l := len(result)
		if l == 0 || len(result[l-1]) == size {
			result = append(result, []T{})
			l++
		}
		result[l-1] = append(result[l-1], item)
	}
	return result
}

// Concat 将多个切片合成一个切片
func Concat[T any](slice []T, slices ...[]T) []T {
	result := append([]T{}, slice...)
	for _, v := range slices {
		result = append(result, v...)
	}
	return result
}

// EqualSlice 判断两个切片是否完全相同,元素数量、顺序、值都相等。
func EqualSlice[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// MatchAll 只有切片中的所有元素都匹配谓词条件才会返回true
func MatchAll[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if !predicate(item) {
			return false
		}
	}
	return true
}

// MatchNone 只有切片中的所有元素都不匹配谓词条件才会返回true
func MatchNone[T any](slice []T, predicate func(T) bool) bool {
	l := 0
	for _, v := range slice {
		if !predicate(v) {
			l++
		}
	}
	return l == len(slice)
}

// MatchAny 只要切片中有一个元素可以匹配谓词条件就会返回true
func MatchAny[T any](slice []T, predicate func(T) bool) bool {
	for _, item := range slice {
		if predicate(item) {
			return true
		}
	}
	return false
}
