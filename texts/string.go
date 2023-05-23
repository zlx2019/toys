/**
  @author: Zero
  @date: 2023/3/27 15:13:33
  @desc: 字符串相关函数库

**/

package texts

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// IsBlank 字符串是否为空串
func IsBlank(str string) bool {
	return IsEmpty(str) || len(strings.TrimSpace(str)) == 0
}

// NotBlank 字符串是否为非空串
func NotBlank(str string) bool {
	return !IsBlank(str)
}

// NotEmpty 字符串是否非空
func NotEmpty(str string) bool {
	return !IsEmpty(str)
}

// IsEmpty 字符串是否为空
func IsEmpty(str string) bool {
	if str == "" {
		return true
	}
	return false
}

// RuneCount 统计一个字符串的字符数量(包含中文)
func RuneCount(content string) int {
	return utf8.RuneCountInString(content)
}

// Format 字符串格式化
func Format(template string, values ...any) string {
	return fmt.Sprintf(template, values...)
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

// ComputeStrSimilarity 计算两个字符串的相似度
func ComputeStrSimilarity(s1 string, s2 string) float64 {
	ed := editDistance(s1, s2)
	lenMax := float64(max(len(s1), len(s2)))
	val := 1 - float64(ed)/lenMax
	// 保留两位小数
	similarity, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
	return similarity
}

func editDistance(s1 string, s2 string) int {
	m := len(s1)
	n := len(s2)

	// 创建一个二维数组来存储编辑距离的值
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 初始化第一行和第一列
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// 计算编辑距离
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}

	// 返回编辑距离
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 辅助函数，用于返回三个数中的最小值
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
