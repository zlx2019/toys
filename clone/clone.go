/**
  @author: Zero
  @date: 2023/3/27 18:51:22
  @desc: 值拷贝相关函数库  参考 https://github.com/jinzhu/copier
**/

package clone

import "github.com/jinzhu/copier"

// 拷贝选项
var defaultOption = copier.Option{
	// 忽略为零值的字段
	IgnoreEmpty: true,
	// 深拷贝
	DeepCopy: true,
}

// CopyPropertiesTo 结构体深拷贝
// 如果字段名不同,通过`copier:"Alias"` 指定为同一个tag,别名必须大写开头
func CopyPropertiesTo(target, source any) error {
	return copier.CopyWithOption(target, source, defaultOption)
}

// CopyProperties 通过拷贝一个旧对象,生成一个新的对象
func CopyProperties[T any](source any) (T, error) {
	var target T
	if err := CopyPropertiesTo(&target, source); err != nil {
		return target, err
	}
	return target, nil
}
