/**
  @author: Zero
  @date: 2023/3/31 15:43:49
  @desc:

**/

package slices

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type member struct {
	Hobby []string
}

func TestRemove(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ints := Remove([]int{1, 2, 3, 4, 5}, 2)
	is.True(EqualSlice(ints, []int{1, 2, 4, 5}))

	ints = RemoveByElement(ints, 4)
	is.True(EqualSlice(ints, []int{1, 2, 5}))
}

func TestFlatMap(t *testing.T) {
	t.Parallel()
	cxk := member{Hobby: []string{"打球", "唱", "跳", "rap"}}
	lrj := member{Hobby: []string{"乒乓球", "足球", "篮球"}}
	slice := []member{cxk, lrj}
	hobbyList := FlatMap(slice, func(m member) []string {
		return m.Hobby
	})
	fmt.Println(hobbyList) // [打球 唱 跳 rap 乒乓球 足球 篮球]
}

func TestReduce(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	sum := Reduce(numbers, 0, func(prev int, item int) int {
		return prev + item
	})
	is.Equal(sum, 55)
}
