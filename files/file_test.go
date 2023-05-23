/**
  @author: Zero
  @date: 2023/5/3 11:33:09
  @desc: 单元测试

**/

package files

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试下载网络文件
func TestDownloadFile(t *testing.T) {
	err := DownloadFile("Default4.jpg", "https://limestart-assets.retiehe.com/wallpapers/Default4.jpg")
	assert.New(t).NoError(err)
}

func TestFolderFileNames(t *testing.T) {
	names, _ := FolderFileNames("./")
	fmt.Println(len(names))
	fmt.Println(names)
}
