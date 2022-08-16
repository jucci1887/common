/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"bufio"
	"github.com/kavanahuang/common"
	"os"
	"testing"
)

func TestFileStart(t *testing.T) {
	Test.Start("File")
}

func TestPerm(t *testing.T) {
	var assert os.FileMode

	assert = 0644
	msg := "FilePerm: "
	result := common.Files.Perm

	if result == assert {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestFileInfo(t *testing.T) {
	msg := "-rw-r-----"
	filePath := "/var/log/syslog"
	result := common.Files.FileInfo(filePath)

	fileMode := result.Mode().String()
	if fileMode == msg {
		Test.T(t).Logs(msg).Ok(fileMode)
	} else {
		Test.T(t).Logs(msg).No(fileMode)
	}
}

func TestCheckFileActive(t *testing.T) {
	var assert int64
	assert = 0

	msg := "CheckFileActive: "
	filePath := "/var/log/syslog"
	result := common.Files.CheckFileActive(filePath)
	if result != assert {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetFileSize(t *testing.T) {
	var assert int64
	assert = 1068
	msg := "GetFileSize: "

	filePath := "../LICENSE"
	result := common.Files.GetFileSize(filePath)
	if result == assert {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestIsExists(t *testing.T) {
	msg := "Is Exists: File_test.go"
	filePath := "File_test.go"
	result := common.Files.IsExists(filePath)
	if result == nil {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestFromString(t *testing.T) {
	assert := "go"
	msg := "FromString: "

	filePath := "/tmp/file_test"
	result := common.Files.PutFile(filePath).FromString("go")
	if result {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}

	getFile := common.Files.GetFile(filePath)
	toString := string(getFile)

	if toString == assert {
		Test.T(t).Logs(msg).Ok(toString)
	} else {
		Test.T(t).Logs(msg).No(toString)
	}
}

func TestFromByte(t *testing.T) {
	assert := "byte"
	msg := "FromByte: "

	filePath := "/tmp/file_byte"
	result := common.Files.PutFile(filePath).FromByte([]byte("byte"))
	if result {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}

	getFile := common.Files.GetFile(filePath)
	toString := string(getFile)

	if toString == assert {
		Test.T(t).Logs(msg).Ok(toString)
	} else {
		Test.T(t).Logs(msg).No(toString)
	}
}

func TestIoPut(t *testing.T) {
	assert := "IoPut"
	msg := "IoPut: "

	filePath := "/tmp/file_ioput"
	result := common.Files.IoPut(filePath, []byte("IoPut"))
	if result {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}

	getFile := common.Files.GetFile(filePath)
	toString := string(getFile)

	if toString == assert {
		Test.T(t).Logs(msg).Ok(toString)
	} else {
		Test.T(t).Logs(msg).No(toString)
	}
}

func TestBufIoPut(t *testing.T) {
	assert := "BufIoPut"
	msg := "BufIoPut: "

	filePath := "/tmp/file_bufioput"
	result := common.Files.BufIoPut(filePath, "BufIoPut")
	if result {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}

	getFile := common.Files.GetFile(filePath)
	toString := string(getFile)

	if toString == assert {
		Test.T(t).Logs(msg).Ok(toString)
	} else {
		Test.T(t).Logs(msg).No(toString)
	}
}

func TestAppendWrite(t *testing.T) {
	assert := "BufIoPutAppendWrite"
	msg := "AppendWrite: "
	appendStr := "AppendWrite"

	filePath := "/tmp/file_bufioput"
	open := common.Files.AppendWrite(filePath)
	defer func() { _ = open.Close() }()

	write := bufio.NewWriter(open)
	_, err := write.WriteString(appendStr)
	if err != nil {
		Test.T(t).Logs(msg).No(err)
	}

	err = write.Flush()
	if err != nil {
		Test.T(t).Logs(msg).No(err)
	} else {
		Test.T(t).Logs(msg).Ok(appendStr)
	}

	getFile := common.Files.GetFile(filePath)
	toString := string(getFile)

	if toString == assert {
		Test.T(t).Logs(assert).Ok(toString)
	} else {
		Test.T(t).Logs(assert).No(toString)
	}
}

func TestAddString(t *testing.T) {
	assert := "BufIoPutAppendWriteAddString"
	msg := "AddString: "
	appendStr := "AddString"

	filePath := "/tmp/file_bufioput"
	addWrite := common.Files.AddWrite(filePath)
	result := addWrite.AddString(appendStr)

	if result {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}

	getFile := common.Files.GetFile(filePath)
	toString := string(getFile)

	if toString == assert {
		Test.T(t).Logs(assert).Ok(toString)
	} else {
		Test.T(t).Logs(assert).No(toString)
	}
}

func TestMultiAddString(t *testing.T) {
	assert := "BufIoPutAppendWriteAddStringMultiAddString"
	msg := "MultiAddString: "
	appendStr := "MultiAddString"

	filePath := "/tmp/file_bufioput"
	addWrite := common.Files.AddWrite(filePath)
	result := addWrite.MultiAddString(appendStr)

	if result {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}

	getFile := common.Files.GetFile(filePath)
	toString := string(getFile)

	if toString == assert {
		Test.T(t).Logs(assert).Ok(toString)
	} else {
		Test.T(t).Logs(assert).No(toString)
	}

	assert = "BufIoPutAppendWriteAddStringMultiAddStringMultiAddString"
	msg = "MultiAddString: "
	appendStr = "MultiAddString"

	result = addWrite.MultiAddString(appendStr)
	addWrite.CloseFile()

	if result {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}

	getFile = common.Files.GetFile(filePath)
	toString = string(getFile)

	if toString == assert {
		Test.T(t).Logs(assert).Ok(toString)
	} else {
		Test.T(t).Logs(assert).No(toString)
	}
}

func TestFileEnd(t *testing.T) {
	Test.End("File")
}
