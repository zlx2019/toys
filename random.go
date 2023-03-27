/**
  @author: Zero
  @date: 2023/3/27 15:26:06
  @desc: 随机生成相关函数库

**/

package toys

import (
	"bytes"
	"math/rand"
	"time"
)

const (
	Numeral      = "0123456789"
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Letters      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
}

// RandomRune 从一个rune切片中,随机挑选n个rune,生成一个string
func RandomRune(length int, charset []rune) string {
	if length <= 0 {
		panic("[RandomStr] The length cannot be less than 0")
	}
	if len(charset) == 0 {
		panic("[RandomStr] charset must not empty")
	}
	runes := make([]rune, length)
	maxLength := len(charset)
	for index := range runes {
		runes[index] = charset[rand.Intn(maxLength)]
	}
	return string(runes)
}

// RandomMobileCode 生成一个6位数手机验证码
func RandomMobileCode() string {
	return RandomNumeral(6)
}

// RandomInt 在一个范围区间,随机获取一个值
func RandomInt(min, max int) int {
	if min == max {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return rand.Intn(max-min) + min
}

// RandomString 随机生成指定长度在字符串(只包含大小写字母)
func RandomString(length int) string {
	return RandomText(Letters, length)
}

// RandomUpper 随机生成指定长度在字符串(只包含大写字母)
func RandomUpper(length int) string {
	return RandomText(UpperLetters, length)
}

// RandomLower 随机生成指定长度在字符串(只包含小写字母)
func RandomLower(length int) string {
	return RandomText(Letters, length)
}

// RandomNumeral 随机生成指定长度的字符串(只包含数字)
func RandomNumeral(length int) string {
	return RandomText(Numeral, length)
}

// RandomNumeralAndString 随机生成指定长度的字符串(包含字母和数字)
func RandomNumeralAndString(length int) string {
	return RandomText(Numeral+Letters, length)
}

// RandomText 随机生成一个指定长度在字符串,根据指定的文本内容
func RandomText(text string, length int) string {
	// 获取文本的长度,在这个长度范围内随机挑选
	textLen := len(text)
	// 字符串拼接
	var result bytes.Buffer
	for i := 0; i < length; i++ {
		result.WriteByte(text[rand.Intn(textLen)])
	}
	return result.String()
}
