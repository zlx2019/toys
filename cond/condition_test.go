/**
  @author: Zero
  @date: 2023/4/12 14:10:50
  @desc:

**/

package cond

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type element1 struct {
}

func (e *element1) Bool() bool {
	return true
}

type element2 struct {
}

func (e *element2) Bool() bool {
	return false
}

func TestBool(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	is.True(Bool(true))
	is.False(Bool(false))

	is.True(Bool(1))
	is.False(Bool(0))

	is.True(Bool(" "))
	is.False(Bool(""))

	is.True(Bool(&element1{}))
	is.False(Bool(&element2{}))
}

func TestXor(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	is.True(Xor(0, 1))
	is.True(Xor(1, 0))

	is.False(Xor(1, 1))
	is.False(Xor(0, 0))
}

func TestNor(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	is.True(Nor(0, 0))
	is.False(Nor(1, 1))
	is.False(Nor(0, 1))
	is.False(Nor(1, 0))
}
