/**
  @author: Zero
  @date: 2023/3/27 14:46:03
  @desc:

**/

package tests

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/zlx2019/toys"
	"testing"
)

func TestTry(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ok := toys.Try(func() error {
		return nil
	})
	is.True(ok)

	ok = toys.Try(func() error {
		return errors.New("error")
	})
	is.False(ok)
}

type MyErr struct {
}

func (e *MyErr) Error() string {
	return ""
}

func TestErrAs(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	_, ok := toys.ErrAs[*MyErr](fmt.Errorf("Hello"))
	is.False(ok)

	_, ok = toys.ErrAs[*MyErr](&MyErr{})
	is.True(ok)

	_, ok = toys.ErrAs[*MyErr](nil)
	is.False(ok)
}
