/**
  @author: Zero
  @date: 2023/4/25 21:28:15
  @desc:

**/

package packer

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zlx2019/toys/network/protocol"
	"testing"
)

func TestNormalPacker_Pack(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	m := &protocol.Message{
		ID:      100001,
		Payload: []byte("Hello,~你好"),
	}
	packer := NewNormalPacker()
	// 封包
	pack, err := packer.Pack(m)
	is.NoError(err)
	fmt.Println(pack)

	// 解包
	message, err := packer.UnPack(pack)
	is.NoError(err)
	is.Equal(m.ID, message.ID)
	is.Equal(string(m.Payload), string(message.Payload))
	fmt.Println(string(m.Payload))
	fmt.Println(string(message.Payload))
}
