/**
  @author: Zero
  @date: 2023/5/30 11:48:48
  @desc: Go程序相关函数库

**/

package system

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

// GetCurrentProcessID 获取当前Go程序的进程ID
func GetCurrentProcessID() string {
	return strconv.Itoa(os.Getpid())
}

// GetCurrentGoroutineID 获取当前协程ID
func GetCurrentGoroutineID() string {
	buf := make([]byte, 128)
	// 读取当前栈信息
	buf = buf[:runtime.Stack(buf, false)]
	stackInfo := string(buf)
	// 去除其他信息,获取 `goroutine 协程ID`
	idStr := strings.Split(stackInfo, "[running]")[0]
	// 再去除`goroutine`
	id := strings.Split(idStr, "goroutine")[1]
	return strings.TrimSpace(id)
}
