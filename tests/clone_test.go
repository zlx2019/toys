/**
  @author: Zero
  @date: 2023/3/27 19:04:36
  @desc:

**/

package tests

import (
	"github.com/stretchr/testify/assert"
	"github.com/zlx2019/toys"
	"testing"
)

type Comm struct {
	Address string
}

type User struct {
	Name     string
	Age      int
	Locked   bool
	Hobby    []string
	C        *Comm
	Username string `copier:"Alias"` //指定别名,通过该别名拷贝
	Ignore   string //明确指定该字段忽略掉
}

type Person struct {
	Name     string
	Age      int
	Locked   bool
	Hobby    []string
	C        *Comm
	Nickname string `copier:"Alias"` //指定别名,通过该别名拷贝
}

func TestCopy(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	u := User{
		Name:     "小张",
		Age:      18,
		Locked:   false,
		Hobby:    []string{"打球", "rap"},
		C:        &Comm{Address: "广州"},
		Username: "root",
		Ignore:   "要忽略的字段",
	}
	p := &Person{}
	toys.CopyProperties(&p, &u)
	is.True(u.Name == p.Name)

	p2, _ := toys.CopyPropertiesTo[Person](&u)
	is.True(u.Name == p2.Name)
}
