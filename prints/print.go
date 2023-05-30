/**
  @author: Zero
  @date: 2023/5/28 08:58:25
  @desc: 终端输出工具库
**/

package prints

import (
	"fmt"
)

// DefaultTextStyle 默认的文本输出样式
var DefaultTextStyle = Default

// Println 输出无颜色文本
func Println(message string) {
	PrintlnColor(message, 0)
}

// PrintlnBlue 输出蓝色文本
func PrintlnBlue(message string) {
	PrintlnColor(message, Blue)
}

// PrintlnYellow 输出黄色文本
func PrintlnYellow(message string) {
	PrintlnColor(message, Yellow)
}

// PrintlnGreen 输出绿色文本
func PrintlnGreen(message string) {
	PrintlnColor(message, Green)
}

// PrintlnCyan 输出青色文本
func PrintlnCyan(message string) {
	PrintlnColor(message, Cyan)
}

// PrintlnRed 输出红色文本
func PrintlnRed(message string) {
	PrintlnColor(message, Red)
}

// PrintlnBlack 输出黑色文本
func PrintlnBlack(message string) {
	PrintlnColor(message, Black)
}

// PrintlnPurple 输出紫色文本
func PrintlnPurple(message string) {
	PrintlnColor(message, Purple)
}

// PrintlnWhite 输出白色文本
func PrintlnWhite(message string) {
	PrintlnColor(message, White)
}

// PrintlnColor 输出带有颜色符号的消息
func PrintlnColor(message string, textColor int) {
	fmt.Println(GetColor(message, textColor))
}

// SprintfColor 获取带有颜色符号的消息
func SprintfColor(message string, textColor int) string {
	return GetColor(message, textColor)
}

// SetDefaultTextStyle 设置默认的字体显示样式
func SetDefaultTextStyle(textStyle int) {
	DefaultTextStyle = textStyle
}

// GetColor 拼接输出内容和样式符号
func GetColor(message string, textColo int) string {
	return fmt.Sprintf("%s%s", getASCIIColor(DefaultTextStyle, 0, textColo), message)
}

// getASCIIColorDefault 获取重置默认样式符号
func getASCIIColorDefault() string {
	return fmt.Sprintf("%c[0m", 0x1B)
}

// getASCIIColor 根据指定的样式,获取ASCII文字符
// 格式: 0x1B[文字样式;背景颜色;前景颜色m
// 示例: 0x1B[1;40;31m
// 解释: 0x1B[高亮;黑色背景m
func getASCIIColor(textStyle, backgroundColor, textColo int) string {
	if textStyle == 0 {
		return fmt.Sprintf("%c[%d;%dm", 0x1B, backgroundColor, textColo)
	} else if backgroundColor == 0 {
		return fmt.Sprintf("%c[%d;%dm", 0x1B, textStyle, textColo)
	} else {
		return fmt.Sprintf("%c[%d;%d;%dm", 0x1B, textStyle, backgroundColor, textColo)
	}
}
