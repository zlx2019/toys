/**
  @author: Zero
  @date: 2023/4/28 18:17:41
  @desc: 配置文件加载器

**/

package configs

import (
	"github.com/spf13/viper"
	"log"
)

// Loader 配置文件加载器,基于Viper
// 支持 Json、Yaml、Toml、Properties等格式
type Loader struct {
	Viper *viper.Viper
}

// NewJsonLoader 创建一个Json文件加载器
func NewJsonLoader(fileName string, filePath ...string) (*Loader, error) {
	return NewLoader(fileName, "json", filePath...)
}

// NewYamlLoader 创建一个Yaml文件加载器
func NewYamlLoader(fileName string, filePath ...string) (*Loader, error) {
	return NewLoader(fileName, "yaml", filePath...)
}

// NewLoader 创建加载器
// fileName 文件名
// fileType 文件类型 json、yaml、toml...
// filePath 文件路径
func NewLoader(fileName string, fileType string, filePath ...string) (*Loader, error) {
	v := viper.New()
	// 设置文件名
	v.SetConfigName(fileName)
	// 设置文件类型
	v.SetConfigType(fileType)
	// 设置多个文件搜索路径
	for _, path := range filePath {
		v.AddConfigPath(path)
	}
	// 查找并且读取文件
	if err := v.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Println("无法找到配置文件!")
		case viper.ConfigParseError:
			log.Println("配置文件格式不正确,无法解析!")
		}
		return nil, err
	}
	return &Loader{Viper: v}, nil
}

// Load 将配置文件内容映射到一个值中
func (loader *Loader) Load(value any) error {
	return loader.Viper.Unmarshal(value)
}

// GetString 根据一个Key获取一个String类型的值
func (loader *Loader) GetString(key string) string {
	return loader.Viper.GetString(key)
}

// GetInt 根据一个Key获取一个int类型的值
func (loader *Loader) GetInt(key string) int {
	return loader.Viper.GetInt(key)
}

// GetBool 根据Key获取一个bool类型的值
func (loader *Loader) GetBool(key string) bool {
	return loader.Viper.GetBool(key)
}

// GetFloat 根据Key获取一个float64类型的值
func (loader *Loader) GetFloat(key string) float64 {
	return loader.Viper.GetFloat64(key)
}

// GetList 根据Key获取一个string类型的切片
func (loader *Loader) GetList(key string) []string {
	return loader.Viper.GetStringSlice(key)
}

// GetMap 根据Key获取一个Map结构配置
func (loader *Loader) GetMap(key string) map[string]interface{} {
	return loader.Viper.GetStringMap(key)
}
