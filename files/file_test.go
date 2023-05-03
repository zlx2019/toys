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
	err := DownloadFile("new.png", "https://zlx2019.github.io/img/logo/Vertx-logo.png")
	assert.New(t).NoError(err)
}

func TestFolderFileNames(t *testing.T) {
	names, _ := FolderFileNames("./")
	fmt.Println(len(names))
	fmt.Println(names)
}
