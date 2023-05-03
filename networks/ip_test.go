/**
  @author: Zero
  @date: 2023/5/3 13:16:35
  @desc:

**/

package networks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInternalIP(t *testing.T) {
	ip := GetInternalIP()
	assert.New(t).Equal(ip, "192.168.0.101")
}

func TestGetPublicIP(t *testing.T) {
	ip := GetPublicIP()
	assert.New(t).Equal(ip, "103.156.242.60")
}

func TestGetIPAddressName(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	// 获取本机公网IP地址
	name, err := GetLocalAddressName()
	is.NoError(err)
	is.Equal(name, "香港")

	// 根据IP 获取地址名称
	name, err = GetIPAddressName("66.150.130.201")
	is.NoError(err)
	is.Equal(name, "美国")
}
