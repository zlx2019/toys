/**
  @author: Zero
  @date: 2023/4/12 15:49:27
  @desc: clone 基准测试
	go test -bench=. clone_benchmark_test.go clone_test.go clone.go


**/

package clone

import "testing"

func BenchmarkCopyProperties(b *testing.B) {
	u := User{
		Name:     "小张",
		Age:      18,
		Locked:   false,
		Hobby:    []string{"打球", "rap"},
		C:        &Comm{Address: "广州"},
		Username: "root",
		Ignore:   "要忽略的字段",
	}
	for i := 0; i < b.N; i++ {
		CopyProperties[User](&u)
	}
}

func BenchmarkCopyPropertiesTo(b *testing.B) {
	u := User{
		Name:     "小张",
		Age:      18,
		Locked:   false,
		Hobby:    []string{"打球", "rap"},
		C:        &Comm{Address: "广州"},
		Username: "root",
		Ignore:   "要忽略的字段",
	}
	var u2 User
	for i := 0; i < b.N; i++ {
		CopyPropertiesTo(&u2, u)
	}
}
