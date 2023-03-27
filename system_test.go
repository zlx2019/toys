/**
  @author: Zero
  @date: 2023/3/27 14:37:36
  @desc:

**/

package toys

import (
	"fmt"
	"testing"
)

func TestWaitStopSignal(t *testing.T) {
	WaitStopSignal(func() {
		fmt.Println("回调函数...")
	})
}
