/**
  @author: Zero
  @date: 2023/3/27 14:41:58
  @desc: 错误处理相关函数库

**/

package toys

import (
	"errors"
	"fmt"
)

// Try 执行一个函数,如果函数内出错,则返回false,反之返回true
func Try(callback func() error) (ok bool) {
	ok = true
	defer func() {
		if err := recover(); err != nil {
			ok = false
			return
		}
	}()
	err := callback()
	if err != nil {
		ok = false
	}
	return
}

// Assert 断言,如果ok为false,则返回一个error
func Assert(ok bool, format string, args ...any) error {
	if !ok {
		return fmt.Errorf(fmt.Sprintf(format, args))
	}
	return nil
}

// ErrAs 判断一个error实例 是否为一个error的类型
func ErrAs[T error](err error) (T, bool) {
	var t T
	ok := errors.As(err, &t)
	return t, ok
}
