/**
  @author: Zero
  @date: 2023/3/27 14:46:03
  @desc:

**/

package toys

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTry(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	ok := Try(func() error {
		return nil
	})
	is.True(ok)

	ok = Try(func() error {
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
	_, ok := ErrAs[*MyErr](fmt.Errorf("Hello"))
	is.False(ok)

	_, ok = ErrAs[*MyErr](&MyErr{})
	is.True(ok)

	_, ok = ErrAs[*MyErr](nil)
	is.False(ok)
}
