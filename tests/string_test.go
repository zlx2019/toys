/**
  @author: Zero
  @date: 2023/3/27 13:10:06
  @desc:

**/

package tests

import (
	"fmt"
	"github.com/zlx2019/toys"
	"testing"
)

func TestString(t *testing.T) {
	str := toys.RandomRune(20, []rune("哈哈哈哈你好xxxsaas"))
	fmt.Println(str)
}
