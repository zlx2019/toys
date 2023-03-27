/**
  @author: Zero
  @date: 2023/3/27 15:13:33
  @desc: 字符串相关函数库

**/

package toys

import "unicode/utf8"

// RuneCount 统计一个字符串的字符数量(包含中文)
func RuneCount(content string) int {
	return utf8.RuneCountInString(content)
}

// ChunkString 将一个字符串切片 按数量分割成多份
func ChunkString[T ~string](str T, size int) []T {
	if size <= 0 {
		panic("[ChunkString] The size cannot be less than 0")
	}
	if len(str) == 0 {
		return []T{""}
	}
	if size >= len(str) {
		return []T{str}
	}
	var chunks []T = make([]T, 0, ((len(str)-1)/size)+1)
	currentLen := 0
	currentStart := 0
	for i := range str {
		if currentLen == size {
			chunks = append(chunks, str[currentStart:i])
			currentLen = 0
			currentStart = i
		}
		currentLen++
	}
	chunks = append(chunks, str[currentStart:])
	return chunks
}
