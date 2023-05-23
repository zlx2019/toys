/**
  @author: Zero
  @date: 2023/3/31 16:55:10
  @desc:

**/

package valida

import (
	"encoding/json"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	alphaMatcher         *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z]+$`)
	letterRegexMatcher   *regexp.Regexp = regexp.MustCompile(`[a-zA-Z]`)
	intStrMatcher        *regexp.Regexp = regexp.MustCompile(`^[\+-]?\d+$`)
	urlMatcher           *regexp.Regexp = regexp.MustCompile(`^((ftp|http|https?):\/\/)?(\S+(:\S*)?@)?((([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))|(([a-zA-Z0-9]+([-\.][a-zA-Z0-9]+)*)|((www\.)?))?(([a-z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-z\x{00a1}-\x{ffff}]{2,}))?))(:(\d{1,5}))?((\/|\?|#)[^\s]*)?$`)
	dnsMatcher           *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`)
	emailMatcher         *regexp.Regexp = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	chineseMobileMatcher *regexp.Regexp = regexp.MustCompile(`^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$`)
	chineseIdMatcher     *regexp.Regexp = regexp.MustCompile(`^[1-9]\d{5}(18|19|20|21|22)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	chineseMatcher       *regexp.Regexp = regexp.MustCompile("[\u4e00-\u9fa5]")
	chinesePhoneMatcher  *regexp.Regexp = regexp.MustCompile(`\d{3}-\d{8}|\d{4}-\d{7}|\d{4}-\d{8}`)
	creditCardMatcher    *regexp.Regexp = regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|(222[1-9]|22[3-9][0-9]|2[3-6][0-9]{2}|27[01][0-9]|2720)[0-9]{12}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11}|6[27][0-9]{14})$`)
	base64Matcher        *regexp.Regexp = regexp.MustCompile(`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`)
)

// IsAlpha 字符串是否只包含英文字母
func IsAlpha(str string) bool {
	return alphaMatcher.MatchString(str)
}

// IsAllUpper 字符串是否全是大写英文字母
func IsAllUpper(str string) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return str != ""
}

// IsAllLower 字符串是否全是小写英文字母
func IsAllLower(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return str != ""
}

// ContainUpper 验证字符串是否包含至少一个英文大写字母
func ContainUpper(str string) bool {
	for _, r := range str {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// ContainLower 验证字符串是否包含至少一个英文小写字母。
func ContainLower(str string) bool {
	for _, r := range str {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// ContainLetter 字符串是否至少包含一个英文字母
func ContainLetter(str string) bool {
	return letterRegexMatcher.MatchString(str)
}

// IsJSON 检查字符串是否为有效的JSON。
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsNumberStr 检查字符串是否可以转换为数字。
func IsNumberStr(s string) bool {
	return IsIntStr(s) || IsFloatStr(s)
}

// IsFloatStr check if the string can convert to a float.
// Play: https://go.dev/play/p/LOYwS_Oyl7U
func IsFloatStr(str string) bool {
	_, e := strconv.ParseFloat(str, 64)
	return e == nil
}

// IsIntStr check if the string can convert to a integer.
// Play: https://go.dev/play/p/jQRtFv-a0Rk
func IsIntStr(str string) bool {
	return intStrMatcher.MatchString(str)
}

// IsIp 检查字符串是否为 IP 地址。
func IsIp(ipstr string) bool {
	ip := net.ParseIP(ipstr)
	return ip != nil
}

// IsIpV4 检查字符串是否为 IPv4 地址。
func IsIpV4(ipstr string) bool {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return false
	}
	return strings.Contains(ipstr, ".")
}

// IsIpV6 检查字符串是否为 IPv6 地址。
func IsIpV6(ipstr string) bool {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return false
	}
	return strings.Contains(ipstr, ":")
}

// IsPort 检查字符串是否为有效的网络端口。
func IsPort(str string) bool {
	if i, err := strconv.ParseInt(str, 10, 64); err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

// IsUrl 检查字符串是否为 URL。
func IsUrl(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}

	return urlMatcher.MatchString(str)
}

// IsDns check if the string is dns.
// Play: https://go.dev/play/p/jlYApVLLGTZ
func IsDns(dns string) bool {
	return dnsMatcher.MatchString(dns)
}

// IsEmail 是否是电子邮箱
func IsEmail(email string) bool {
	return emailMatcher.MatchString(email)
}

// IsChineseMobile 验证字符串是否是中国手机号码
func IsChineseMobile(mobileNum string) bool {
	return chineseMobileMatcher.MatchString(mobileNum)
}

// IsChineseIdNum check if the string is chinese id card.
// Play: https://go.dev/play/p/d8EWhl2UGDF
func IsChineseIdNum(id string) bool {
	return chineseIdMatcher.MatchString(id)
}

// ContainChinese 字符串是否包含中文字符
func ContainChinese(s string) bool {
	return chineseMatcher.MatchString(s)
}

// IsChinesePhone check if the string is chinese phone number.
// Valid chinese phone is xxx-xxxxxxxx or xxxx-xxxxxxx.
// Play: https://go.dev/play/p/RUD_-7YZJ3I
func IsChinesePhone(phone string) bool {
	return chinesePhoneMatcher.MatchString(phone)
}

// IsCreditCard check if the string is credit card.
// Play: https://go.dev/play/p/sNwwL6B0-v4
func IsCreditCard(creditCart string) bool {
	return creditCardMatcher.MatchString(creditCart)
}

// IsBase64 字符串是否是base64编码
func IsBase64(base64 string) bool {
	return base64Matcher.MatchString(base64)
}

// IsEmptyString check if the string is empty.
// Play: https://go.dev/play/p/dpzgUjFnBCX
func IsEmptyString(str string) bool {
	return len(str) == 0
}

// IsRegexMatch check if the string match the regexp.
// Play: https://go.dev/play/p/z_XeZo_litG
func IsRegexMatch(str, regex string) bool {
	reg := regexp.MustCompile(regex)
	return reg.MatchString(str)
}

// IsStrongPassword check if the string is strong password, if len(password) is less than the length param, return false
// Strong password: alpha(lower+upper) + number + special chars(!@#$%^&*()?><).
// Play: https://go.dev/play/p/QHdVcSQ3uDg
func IsStrongPassword(password string, length int) bool {
	if len(password) < length {
		return false
	}
	var num, lower, upper, special bool
	for _, r := range password {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsLower(r):
			lower = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}

	return num && lower && upper && special
}

// IsWeakPassword 检查字符串是否为弱密码
// Weak password: 只有字母或只有数字或字母+数字。
func IsWeakPassword(password string) bool {
	var num, letter, special bool
	for _, r := range password {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsLetter(r):
			letter = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}

	return (num || letter) && !special
}

// IsZeroValue 判断一个值是否为该类型的默认零值
func IsZeroValue(value any) bool {
	if value == nil {
		return true
	}
	// 获取值的信息
	rv := reflect.ValueOf(value)
	// 如果这个值为指针类型,则获取指针的值
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if !rv.IsValid() {
		return true
	}
	switch rv.Kind() {
	case reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Ptr, reflect.Chan, reflect.Func, reflect.Interface, reflect.Slice, reflect.Map:
		return rv.IsNil()
	}
	return reflect.DeepEqual(rv.Interface(), reflect.Zero(rv.Type()).Interface())
}

// IsGBK 检查数据编码是否为 GBK
func IsGBK(data []byte) bool {
	i := 0
	for i < len(data) {
		if data[i] <= 0xff {
			i++
			continue
		} else {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}

	return true
}

// IsUTF8 判断一个 []byte| string | rune 是否是合法的utf-8编码格式
func IsUTF8[T []byte | string | rune](value T) bool {
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return utf8.ValidString(v.String())
	case reflect.Int32:
		// rune
		// int64 to rune
		r := rune(v.Int())
		return utf8.ValidRune(r)
	case reflect.Slice:
		// []byte
		// 获取切片的元素类型
		elemKind := v.Type().Elem().Kind()
		// 判断该切片的元素是否是byte(uint8)类型
		if elemKind == reflect.Uint8 {
			return utf8.Valid(v.Bytes())
		}
	}
	return false
}

// IsFullUTF8 判断该字节序列是否包含完整的UTF-8编码的字符
// 通常需要先校验utf8是否合法,再调用该方法进一步校验
func IsFullUTF8(value []byte) bool {
	return utf8.FullRune(value)
}
