/**
  @author: Zero
  @date: 2023/3/27 14:37:36
  @desc:

**/

package system

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWaitStopSignal(t *testing.T) {
	result, fail, _ := Command("ls")
	fmt.Println(result)
	fmt.Println(fail)
}

func TestCommand(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	ok := CommandOk("ls", "-a")
	is.True(ok)

	lines, ok := CommandLines("ls")
	is.True(ok)
	is.Equal(len(lines), 2)
}
