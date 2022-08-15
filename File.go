/*
 Author: Kernel.Huang
 Mail: kernelman79@gmail.com
 File: FileService
 Date: 3/21/21 12:50 AM
*/
package common

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/kavanahuang/log"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const MAX_ULIMIT = 65535

type fileNode struct {
	Name     string
	Size     int64
	Modified time.Time
	Children []*fileNode
}

type FileServices struct {
	name    string
	path    string
	Perm    os.FileMode
	DirPerm os.FileMode
	osFile  *os.File
}

var Files = new(FileServices)

func init() {
	Files.Perm = 0644
	Files.DirPerm = 0755
}

func (file *FileServices) FileInfo(filename string) os.FileInfo {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Logs.Error("Get file info error: ", err)
	}

	return fileInfo
}

/**
 * 检查文件存活时间
 */
func (file *FileServices) CheckFileActive(filename string) int64 {
	return file.FileInfo(filename).ModTime().Unix()
}

/**
 * 获取文件大小
 */
func (file *FileServices) GetFileSize(filename string) int64 {
	return file.FileInfo(filename).Size()
}

func (file *FileServices) PutFile(filename string) *FileServices {
	file.name = filename
	return file
}

/**
 * 读取文件内容
 */
func (file *FileServices) GetFile(filename string) []byte {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Logs.Error("Get file: "+filename+" error: ", err)
	}

	return buf
}

/**
 * 写入字符串内容
 * 示例： Files.PutFile(filePath).FromString("text")
 */
func (file *FileServices) FromString(str string) bool {
	return file.BufIoPut(file.name, str)
}

/**
 * []byte写入方式
 * 示例： Files.PutFile(filePath).FromByte([]byte("text")
 */
func (file *FileServices) FromByte(content []byte) bool {
	err := ioutil.WriteFile(file.name, content, file.Perm)
	if err != nil {
		log.Logs.Error("Write file: "+file.name+" error: ", err)
		return false
	}

	return true
}

/**
 * []byte写入方式
 */
func (file *FileServices) IoPut(filename string, content []byte) bool {
	err := ioutil.WriteFile(filename, content, file.Perm)
	if err != nil {
		log.Logs.Error("Io put file: "+filename+" error: ", err)
		return false
	}

	return true
}

/**
 * BufIO写入模式
 */
func (file *FileServices) BufIoPut(filename string, content string) bool {
	create := file.Create(filename)
	putFile := bufio.NewWriter(create)

	_, err := putFile.WriteString(content)
	if err != nil {
		log.Logs.Error("Write data error: ", err)
		return false
	}

	err = putFile.Flush()
	if err != nil {
		log.Logs.Error("Push data error: ", err)
		return false
	}

	defer func() { _ = create.Close() }()
	return true
}

/**
 * 追加写入模式
 *	open := Files.AppendWrite(filePath)
 *	defer func() {_ = open.Close()}()
 *	write := bufio.NewWriter(open)
 *	_, _ := write.WriteString("text")
 *	_ = write.Flush()
 */
func (file *FileServices) AppendWrite(filename string) *os.File {
	open, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, file.Perm)
	if err != nil {
		log.Logs.Error("Open file error: ", err)
	}

	return open
}

/**
 * 追加写入模式
 */
func (file *FileServices) AddWrite(filename string) *FileServices {
	open, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, file.Perm)
	if err != nil {
		log.Logs.Error("Open file error: ", err)
	}

	file.osFile = open
	return file
}

/**
 * 单次追加写入文本内容
 * 示例： Files.AddWrite(filePath).AddString("text")
 */
func (file *FileServices) AddString(content string) bool {
	if file.osFile == nil {
		log.Logs.Error("Add write error: the file.osFile value is nil")
	}

	defer func() { _ = file.osFile.Close() }()

	// 使用缓存写入
	write := bufio.NewWriter(file.osFile)
	_, err := write.WriteString(content)
	if err != nil {
		log.Logs.Error("Add string content error: ", err)
		return false
	}

	// 将缓存中的内容写入磁盘
	err = write.Flush()
	if err != nil {
		log.Logs.Error("Flush add string content error: ", err)
		return false
	}

	return true
}

/**
 * 多次追加写入文本内容
 * 示例：
 * write := Files.AddWrite(filePath)
 * write.MultiAddString("text1")
 * write.MultiAddString("text2")
 * write.CloseFile()
 */
func (file *FileServices) MultiAddString(content string) bool {
	if file.osFile == nil {
		log.Logs.Error("Add write error: the file.osFile value is nil")
	}

	// 使用缓存写入
	write := bufio.NewWriter(file.osFile)
	_, err := write.WriteString(content)
	if err != nil {
		log.Logs.Error("Add string content error: ", err)
		return false
	}

	// 将缓存中的内容写入磁盘
	err = write.Flush()
	if err != nil {
		log.Logs.Error("Flush add string content error: ", err)
		return false
	}

	return true
}

/**
 * 关闭文件句柄，使其无法用于I/O
 * 示例：
 * write := Files.AddWrite(filePath)
 * write.MultiAddString("text")
 * write.CloseFile()
 */
func (file *FileServices) CloseFile() bool {
	if file.osFile == nil {
		log.Logs.Error("Close file error: the file.osFile is nil")
	}

	defer func() { _ = file.osFile.Close() }()
	return true
}

/**
 * 覆盖写入模式
 * 示例：
 *	open := Files.AppendWrite(filePath)
 *	defer func() {_ = open.Close()}()
 *	write := bufio.NewWriter(open)
 *	_, _ := write.WriteString("text")
 *	_ = write.Flush()
 */
func (file *FileServices) Overwrite(filename string) *os.File {
	open, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, file.Perm)
	if err != nil {
		log.Logs.Error("Open file error: ", err)
	}

	return open
}

/**
 * Perm传参写入模式
 * 示例：
 * FileService.PutFile(path).Write()
 * FileService.PutFile(path).Write(0640)
 */
func (file *FileServices) Write(perm ...os.FileMode) *os.File {
	if len(perm) != 0 {
		file.Perm = perm[0]
	}

	open, err := os.OpenFile(file.name, os.O_CREATE|os.O_EXCL|os.O_WRONLY, file.Perm)
	if err != nil {
		log.Logs.Error("Open file error: ", err)
	}

	return open
}

/**
 * 使用默认Perm:0666模式创建文件
 */
func (file *FileServices) Create(filename string) *os.File {
	open, err := os.Create(filename)
	if err != nil {
		log.Logs.Error("Created file: "+filename+" error: ", err)
	}
	return open
}

/**
 * 检查文件或目录是否存在,并返回错误
 */
func (file *FileServices) IsExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return err
	}
	return nil
}

/**
 * 检查文件或目录是否存在
 */
func (file *FileServices) PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		if err != nil {
			log.Logs.Error(path+" not found: ", err)
		}
	}
	return false
}

/**
 * 路径赋值
 */
func (file *FileServices) NewPath(path string) *FileServices {
	file.path = path
	return file
}

/**
 * 新建一个目录
 */
func (file *FileServices) NewFolder(perm ...os.FileMode) bool {
	if len(perm) != 0 {
		file.DirPerm = perm[0]
	}

	err := os.Mkdir(file.path, file.DirPerm)
	if err != nil {
		log.Logs.Error("Created dir: "+file.path+" error: ", err)
		return false
	}

	return true
}

/**
 * 新建目录: path的所经过目录路径不存在则创建, 否则忽略
 */
func (file *FileServices) NewNonFolder(perm ...os.FileMode) bool {
	if len(perm) != 0 {
		file.DirPerm = perm[0]
	}

	err := os.MkdirAll(file.path, file.DirPerm)
	if err != nil {
		log.Logs.Error("Created dir: "+file.path+" error: ", err)
		return false
	}

	return true
}

/**
 * 新建一个目录
 */
func (file *FileServices) NewDir(path string, perm ...os.FileMode) bool {
	if len(perm) != 0 {
		file.DirPerm = perm[0]
	}

	err := os.Mkdir(path, file.DirPerm)
	if err != nil {
		log.Logs.Error("New dir: "+path+" error: ", err)
		return false
	}

	return true
}

/**
 * 创建目录: path的所经过目录路径不存在则创建, 否则忽略
 */
func (file *FileServices) CreateDir(path string, perm ...os.FileMode) bool {
	if len(perm) != 0 {
		file.DirPerm = perm[0]
	}

	err := os.MkdirAll(path, file.DirPerm)
	if err != nil {
		log.Logs.Error("Created dir: "+path+" error: ", err)
		return false
	}

	return true
}

/**
 * 获取文件树列表
 */
func (file *FileServices) ShowFileTree() *fileNode {
	node := &fileNode{}
	if info, err := os.Stat(file.path); err == nil {
		if err := file.RecursionIterateDir(file.path, info, node, new(uint)); err != nil {
			log.Logs.Error("Get File list error: ", err)
		}
	}

	return node
}

/**
 * 递归遍历目录
 */
func (file *FileServices) RecursionIterateDir(path string, info os.FileInfo, node *fileNode, number *uint) error {
	if (!info.IsDir() && !info.Mode().IsRegular()) || strings.HasPrefix(info.Name(), ".") {
		return errors.New("ERROR: Non-regular file")
	}

	*number++
	// 限制遍历的文件数量
	if (*number) > MAX_ULIMIT {
		return errors.New("ERROR: Over file limit")
	}

	node.Name = info.Name()
	node.Size = info.Size()
	node.Modified = info.ModTime()
	if !info.IsDir() {
		return nil
	}

	children, err := ioutil.ReadDir(path)
	if err != nil {
		return fmt.Errorf("ERROR: Failed to list files: %w", err)
	}

	node.Size = 0
	for _, i := range children {
		c := &fileNode{}
		p := filepath.Join(path, i.Name())
		if err := file.RecursionIterateDir(p, i, c, number); err != nil {
			continue
		}

		node.Size += c.Size
		node.Children = append(node.Children, c)
	}

	return nil
}
