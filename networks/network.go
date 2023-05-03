/**
  @author: Zero
  @date: 2023/5/3 08:08:18
  @desc: 网络相关函数库

**/

package networks

import (
	"github.com/zlx2019/toys/system"
	"net"
	"os/exec"
	"time"
)

// IsTelnet 检验一个主机端口是否可以被连接
func IsTelnet(host, port string) bool {
	adder := host + ":" + port
	// 尝试建立连接,超时时间为3s
	conn, err := net.DialTimeout("tcp", adder, time.Second*3)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// IsPing 检验一个主机地址是否可以Ping通
func IsPing(host string) bool {
	cmd := exec.Command("ping", host, "-c", "4", "-W", "6")
	if system.IsWindows() {
		cmd = exec.Command("ping", host, "-n", "4", "-w", "6")
	}
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}
