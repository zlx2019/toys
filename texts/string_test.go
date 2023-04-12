/**
  @author: Zero
  @date: 2023/3/27 13:10:06
  @desc:

**/

package texts

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zlx2019/toys/randoms"
	"testing"
)

func TestString(t *testing.T) {
	str := randoms.RandomRune(20, []rune("哈哈哈哈你好xxxsaas"))
	fmt.Println(str)
}

func TestEditDistance(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	s1 := "12345"
	s2 := "12345"
	similarity := ComputeStrSimilarity(s1, s2)
	is.Equal(similarity, float64(1))
	s3 := "12345"
	s4 := "1234"
	is.Equal(ComputeStrSimilarity(s3, s4), float64(0.8))
}
