/**
  @author: Zero
  @date: 2023/3/27 14:08:42
  @desc: 系统相关函数库

**/

package system

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// WaitSignal 阻塞等待系统信号
func WaitSignal(fn func(sig os.Signal) bool) {
	// 用来接收操作系统信号的通道
	signalChannel := make(chan os.Signal, 1)
	// 可接受的信号
	allowed := []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGPIPE}
	// 如果监听到其中一个信号后,将信号发送到信号通道中
	signal.Notify(signalChannel, allowed...)
	// 阻塞等待信号
	for signal := range signalChannel {
		if !fn(signal) {
			close(signalChannel)
			break
		}
	}
}

// WaitStopSignal 阻塞等待系统关闭信号
// release 回调函数,可以用来执行释放一些资源操作
func WaitStopSignal(release func()) {
	// 创建一个接收操作系统信号的通道
	exitChannel := make(chan os.Signal)
	// 这里表示如果接收到了SIGINT或者SIGTERM系统信号,则会把信号向exit通道发送.
	// syscall.SIGINT: 		用户发送INTR字符,例如在终端执行(Ctrl+C) 触发 kill -2 pid然后进程结束
	// syscall.SIGTERM: 	结束程序(可以被捕获、阻塞或忽略)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞,直到接收到两种信号其中一种...
	<-exitChannel
	// 执行回调函数
	if release != nil {
		release()
	}
}

// GetEnv 根据Key获取环境变量
func GetEnv(key string) string {
	return os.Getenv(key)
}

// SetEnv 设置环境变量
func SetEnv(key, value string) error {
	return os.Setenv(key, value)
}

// DelEnv 根据Key删除环境变量
func DelEnv(key string) error {
	return os.Unsetenv(key)
}

// GetSystem 操作系统
func GetSystem() string {
	return runtime.GOOS
}

// IsMac 是否是Mac操作系统
func IsMac() bool {
	return GetSystem() == "darwin"
}

// IsWindows 是否是Windows系统
func IsWindows() bool {
	return GetSystem() == "windows"
}

// IsLinux 是否是Linux系统
func IsLinux() bool {
	return GetSystem() == "linux"
}
