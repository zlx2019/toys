/**
  @author: Zero
  @date: 2023/4/12 13:42:22
  @desc: 真假运算函数库

**/

package cond

import "reflect"

// Bool 判断一个值为`真`|`假`
//
// 如果是一个类型的零值均为`假`
// 如果是slices或者map,length大于0时为真，否则返回false
// 如果出入类型参数含有Bool方法, 会调用该方法并返回
// 如果传入类型参数有IsZero方法, 返回IsZero方法返回值的取反
func Bool[T any](value T) bool {
	// 断言类型
	switch v := any(value).(type) {
	case interface{ Bool() bool }:
		// 如果该值有Bool()方法,直接返回该方法结果
		return v.Bool()
	case interface{ IsZero() bool }:
		// 如果该值有IsZero()方法,直接返回该方法结果
		return v.IsZero()
	}
	return isValueZero(&value)
}

// And 并且运算
// 当两个值都为`真`才返回True,反之False。
func And[A, B any](a A, b B) bool {
	return Bool(a) && Bool(b)
}

// Or 或运算
// 当有其中一个值为`真`就返回True。
func Or[A, B any](a A, b B) bool {
	return Bool(a) || Bool(b)
}

// Xor 异或运算
// 两个值相等则返回False,不相等则返回True
func Xor[A, B any](a A, b B) bool {
	vala := Bool(a)
	valb := Bool(b)
	return (vala || valb) && vala != valb
}

// Nor 或非运算
// 两个值都为`假`则返回True
func Nor[A, B any](a A, b B) bool {
	return !(Bool(a) || Bool(b))
}

// XNor 如果a和b都是真的或a和b均是假的，则返回true。
func XNor[A, B any](a A, b B) bool {
	valA := Bool(a)
	valB := Bool(b)
	return (valA && valB) || (!valA && !valB)
}

// 判断一个值是否为零值
func isValueZero(value any) bool {
	// 断言value.Kind()
	switch v := reflect.ValueOf(value).Elem(); v.Kind() {
	case reflect.Map, reflect.Slice:
		return v.Len() != 0
	default:
		return !v.IsZero()
	}
}
