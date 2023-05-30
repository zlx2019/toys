package prints

import "testing"

func TestGetASCIIColor(t *testing.T) {
	SetDefaultTextStyle(High)
	Println("哈哈哈")
	PrintlnRed("哈哈哈")
	PrintlnBlack("哈哈哈")
	PrintlnBlue("哈哈哈")
	PrintlnCyan("哈哈哈")
	PrintlnGreen("哈哈哈")
	PrintlnPurple("哈哈哈")
	PrintlnYellow("哈哈哈")
	PrintlnWhite("哈哈哈")
}
