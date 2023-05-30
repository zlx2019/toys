/**
  @author: Zero
  @date: 2023/5/29 21:58:30
  @desc:

**/

package randoms

import (
	"fmt"
	"testing"
)

func TestRandomUUID(t *testing.T) {
	fmt.Println(RandomUUID())
	fmt.Println(RandomUUID())
	fmt.Println(RandomUUID())
	fmt.Println(RandomString(15))
	fmt.Println(RandomString(15))
	fmt.Println(RandomString(15))
	fmt.Println(RandomString(15))
}
