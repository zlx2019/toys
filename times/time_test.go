/**
  @author: Zero
  @date: 2023/3/28 15:41:39
  @desc:

**/

package times

import (
	"fmt"
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

	fmt.Println(Format(EndOfYear(time.Now())))
}
