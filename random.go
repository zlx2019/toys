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
	NUMBERS = "0123456789"
)

func init() {
	rand.Seed(time.Now().Unix())
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
	rand.Seed(time.Now().Unix())
	for index := range runes {
		runes[index] = charset[rand.Intn(maxLength)]
	}
	return string(runes)
}

// RandomMobileCode 生成一个6位数手机验证码
func RandomMobileCode() string {
	return RandomNumberCode(6)
}

// RandomNumberCode 随机生成指定长度的数字字符
func RandomNumberCode(length int) string {
	numLength := len(NUMBERS)
	// 设置一个随机种子
	rand.Seed(time.Now().UnixNano())
	var code bytes.Buffer
	//var code strings.Builder
	for i := 0; i < length; i++ {
		// 每次随机从0-9获取一个字符,追加到code
		code.WriteByte(NUMBERS[rand.Intn(numLength)])
	}
	return code.String()
}
