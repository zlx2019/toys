/**
  @author: Zero
  @date: 2023/5/3 10:37:43
  @desc:

**/

package networks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPing(t *testing.T) {
	ass := assert.New(t)
	ping := IsPing("www.baidu.com")
	ass.True(ping)

	ping = IsPing("www.xawdw.xax")
	ass.False(ping)
}

func TestIsTelnet(t *testing.T) {
	telnet := IsTelnet("127.0.0.1", "8080")
	assert.New(t).True(telnet)
}
