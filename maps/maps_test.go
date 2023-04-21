/**
  @author: Zero
  @date: 2023/3/31 16:24:58
  @desc:

**/

package maps

import (
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	m := map[string]any{"name": "张三", "age": 18, "address": "广州"}
	ForEach(m, func(key string, value any) {
		fmt.Println(key, value)
	})
}
