/**
  @author: Zero
  @date: 2023/4/29 21:59:40
  @desc:

**/

package configs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadJson(t *testing.T) {
	loader, err := NewLoader("d", "../")
	if err != nil {
		panic(err)
	}
	get := loader.Viper.GetString("package.name")
	fmt.Println(get)
}

// 映射实体
type mapping struct {
	Name   string         `json:"name" yaml:"name"`
	Age    int            `json:"age" yaml:"age"`
	Locked bool           `json:"locked" yaml:"locked"`
	Hobby  []string       `json:"hobby" yaml:"hobby"`
	Info   map[string]any `json:"info" yaml:"info"`
}

func TestLoadJson(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	loader, err := NewJsonLoader("a.json", "./config/")
	is.NoError(err)

	name := loader.GetString("name")
	is.Equal(name, "张三")

	m := &mapping{}
	err = loader.Load(&m)
	is.NoError(err)
	is.Equal(m.Name, name)
	fmt.Println(m)
}

func TestLoadYaml(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	loader, err := NewYamlLoader("b.yml", "./config/")
	is.NoError(err)
	age := loader.GetInt("age")
	is.Equal(age, 18)

	m := &mapping{}
	err = loader.Load(m)
	is.NoError(err)
	is.Equal(m.Name, "张三")
	fmt.Println(m)
}
