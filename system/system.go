/**
  @author: Zero
  @date: 2023/3/27 14:08:42
  @desc: 系统相关函数库

**/

package system

import (
	"bytes"
	"github.com/zlx2019/toys/slices"
	"github.com/zlx2019/toys/valida"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"unicode/utf8"
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

// CommandOk 执行shell终端命令,只返回成功或者失败
func CommandOk(commands ...string) bool {
	command := strings.Join(commands, " ")
	_, _, err := Command(command)
	if err != nil {
		return false
	}
	return true
}

// CommandLines 执行shell终端命令,如果成功则将结果转换为字符串行切片
func CommandLines(commands ...string) ([]string, bool) {
	command := strings.Join(commands, " ")
	successResult, _, err := Command(command)
	if err != nil {
		return nil, false
	}
	lines := slices.Filter(strings.Split(successResult, "\n"), func(item string) bool {
		return item != ""
	})
	return lines, true
}

// Command 执行shell终端命令
// successResult 命令执行成功结果
// failResult 命令执行失败结果
func Command(command string) (successResult, failResult string, err error) {
	// 成功结果内容
	var successBuf bytes.Buffer
	// 失败结果内容
	var failBuf bytes.Buffer
	// 创建终端
	cmd := exec.Command("/bin/bash", "-c", command)
	if IsWindows() {
		cmd = exec.Command("powershell.exe", command)
	}
	// 将终端标准输出和错误输出 写入缓冲区
	cmd.Stdout = &successBuf
	cmd.Stderr = &failBuf
	// 执行命令
	err = cmd.Run()
	if err != nil {
		// 返回错误内容
		if utf8.Valid(failBuf.Bytes()) {
			failResult = byteToString(failBuf.Bytes(), "UTF8")
		} else if valida.IsGBK(failBuf.Bytes()) {
			failResult = byteToString(failBuf.Bytes(), "GBK")
		}
		return
	}
	datas := successBuf.Bytes()
	if utf8.Valid(datas) {
		successResult = byteToString(datas, "UTF8")
	} else if valida.IsGBK(datas) {
		successResult = byteToString(datas, "GBK")
	}
	return
}

// 根据指定的编码格式,将字节切片转为string。
func byteToString(data []byte, charset string) string {
	var result string
	switch charset {
	case "GBK":
		decodeBytes, _ := simplifiedchinese.GBK.NewDecoder().Bytes(data)
		result = string(decodeBytes)
	case "GB18030":
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(data)
		result = string(decodeBytes)
	case "UTF8":
		fallthrough
	default:
		result = string(data)
	}

	return result
}
