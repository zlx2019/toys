/**
  @author: Zero
  @date: 2023/3/27 20:45:56
  @desc:

**/

package tests

import (
	"fmt"
	"github.com/zlx2019/toys"
	"testing"
)

func TestAnyToBytes(t *testing.T) {
	bytes, _ := toys.ToBytes(111111.12)
	fmt.Println(bytes)
}
