/**
  @author: Zero
  @date: 2023/3/28 15:41:39
  @desc:

**/

package tests

import (
	"fmt"
	"github.com/zlx2019/toys"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	//fmt.Println(Now())
	//fmt.Println(NowDate())
	//fmt.Println(NowTime())
	//
	//fmt.Println(ParseDateTime(Now(),"2006-01-02 15:04:05"))
	//fmt.Println(Format(time.Now()))

	fmt.Println(toys.Format(toys.EndOfYear(time.Now())))
}
