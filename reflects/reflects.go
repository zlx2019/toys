/**
  @author: Zero
  @date: 2023/5/1 22:14:59
  @desc: 反射工具包

**/

package main

import (
	"reflect"
)

// GetTags 获取一个结构体中所有字段的指定tag
// 以字段名为Key,tag值为Value返回一个Map
func GetTags(data any, tag string) map[string]string {
	typeInfo := reflect.TypeOf(data)
	if typeInfo.Kind() == reflect.Pointer {
		typeInfo = typeInfo.Elem()
	}
	count := typeInfo.NumField()
	maps := make(map[string]string, count)
	for i := 0; i < count; i++ {
		field := typeInfo.Field(i)
		if tagValue, ok := field.Tag.Lookup(tag); ok {
			maps[field.Name] = tagValue
		}
	}
	return maps
}
