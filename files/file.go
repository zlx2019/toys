/**
  @author: Zero
  @date: 2023/5/3 10:43:40
  @desc: 文件系统相关函数库

**/

package files

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

// Exist 查看一个文件是否存在
func Exist(file string) bool {
	_, err := os.Stat(file)
	// 存在
	if err == nil {
		return true
	}
	// 不存在异常
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// Create 创建一个文件
func Create(filePath string) bool {
	if Exist(filePath) {
		// 已存在,无需创建
		return true
	}
	file, err := os.Create(filePath)
	if err != nil {
		return false
	}
	defer file.Close()
	return true
}

// CreateFolder 创建一个目录
func CreateFolder(folderPath string) error {
	// 777
	return os.MkdirAll(path.Dir(folderPath), os.ModePerm)
}

// Remove 删除一个文件
func Remove(path string) error {
	return os.Remove(path)
}

// IsFolder 是否是一个文件目录
func IsFolder(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// Copy 拷贝一个文件
func Copy(source, target string) error {
	// 打开源文件,只读模式
	srcFile, err := os.OpenFile(source, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	//创建目标文件
	tarFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer tarFile.Close()
	//定义一个缓冲区,循环读写
	buf := make([]byte, 1024*4)
	for {
		// 读取源文件
		readCount, err := srcFile.Read(buf)
		// 错误处理 EOF表示已经没有可读的数据
		if err != nil {
			if err == io.EOF {
				// 已经读取完毕,正常退出
				return nil
			}
			return err
		}
		// 将读取到的数据,写入新文件
		_, err = tarFile.Write(buf[:readCount])
		if err != nil {
			return err
		}
	}
}

// ClearFile 清空一个文件内容
func ClearFile(filePath string) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString("")
	return err
}

// ReadFileToString 读取文件内容,转换为string
func ReadFileToString(filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// ReadFileLines 读取文件内容,转换为[]string,按行获取
func ReadFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	lines := make([]string, 0)
	// 创建读取器,将文件内容写入读取器
	readBuf := bufio.NewReader(file)
	for {
		// 按`\n`换行符进行读取
		line, _, err := readBuf.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}

// FolderFileNames 获取一个目录内所有文件的名字
func FolderFileNames(path string) ([]string, error) {
	if !Exist(path) {
		return []string{}, nil
	}
	files, err := os.ReadDir(path)
	fileNum := len(files)
	if err != nil || fileNum == 0 {
		return []string{}, nil
	}
	names := make([]string, fileNum)
	for _, file := range files {
		if !file.IsDir() {
			file.Type().IsRegular()
			names = append(names, file.Name())
		}
	}
	return names, nil
}

// DownloadFile 下载网络资源文件
// filePath 下载到的本地文件路径
// url 要下载的资源文件
func DownloadFile(filePath, url string) error {
	// 发起HTTP请求
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	// 创建文件
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	// 将数据拷贝到文件中
	_, err = io.Copy(file, response.Body)
	return err
}

// UploadFile 将本地文件,上传至网络某个文件传输接口(HTTP Post)
// filePath 要上传的文件路径
// uploadUrl 上传地址
func UploadFile(filePath, uploadUrl string) (bool, error) {
	if !Exist(filePath) {
		// 文件不存在
		return false, errors.New("file not exist")
	}
	// 获取文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	// 创建文件缓冲区
	fileBuf := &bytes.Buffer{}
	// 创建FormFile表单文件对象
	writer := multipart.NewWriter(fileBuf)
	formFile, err := writer.CreateFormFile("file", fileInfo.Name())
	if err != nil {
		return false, err
	}
	// 打开要上传的文件
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()
	// 将file文件内容拷贝到FormFile中
	_, err = io.Copy(formFile, file)
	if err != nil {
		return false, err
	}
	contentType := writer.FormDataContentType()
	writer.Close()
	// 上传文件
	_, err = http.Post(uploadUrl, contentType, fileBuf)
	if err != nil {
		return false, err
	}
	return true, nil
}
