/**
  @author: Zero
  @date: 2023/3/31 16:24:58
  @desc:

**/

package maps

import (
	"testing"
)

func TestForEach(t *testing.T) {
	m := map[string]any{"name": "张三", "age": 18, "address": "广州"}
	ForEachMap(m, func(s string, a any) {

	})
}
