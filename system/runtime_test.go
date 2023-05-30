/**
  @author: Zero
  @date: 2023/5/30 11:54:02
  @desc:

**/

package system

import (
	"fmt"
	"testing"
)

func TestGetCurrentGoroutineID(t *testing.T) {
	fmt.Println(GetCurrentProcessID())
	fmt.Println(GetCurrentGoroutineID())
}
