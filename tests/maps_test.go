/**
  @author: Zero
  @date: 2023/3/31 16:24:58
  @desc:

**/

package tests

import (
	"github.com/zlx2019/toys"
	"testing"
)

func TestForEach(t *testing.T) {
	m := map[string]any{"name": "张三", "age": 18, "address": "广州"}
	toys.ForEachMap(m, func(s string, a any) {

	})
}
