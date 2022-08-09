/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"github.com/kavanahuang/common"
	"testing"
)

func TestStringStart(t *testing.T) {
	Test.Start("System")
}

func TestGetCurrentDir(t *testing.T) {
	msg := "/home/kernel/project/go/common/tests"
	result := common.GetCurrentDir()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetCurrentFilename(t *testing.T) {
	msg := "/home/kernel/project/go/common/tests/t.test"
	result := common.GetCurrentFilename()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestCurrentAndAbsPath(t *testing.T) {
	msg := "/home/kernel/project/go/common/tests/t.test"
	result := common.CurrentAndAbsPath()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestSetCurrentPath(t *testing.T) {
	msg := "./t.test"
	result := common.SetCurrentPath()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetAbsPath(t *testing.T) {
	msg := "/home/kernel/project/go/common/tests/t.test"
	path := common.SetCurrentPath()
	result := common.GetAbsPath(path)

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetConfigDir(t *testing.T) {
	msg := "/home/kernel/project/go/common/config"
	result := common.GetConfigDir()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetLastPath(t *testing.T) {
	msg := "/home/kernel/project/go/common"
	path := common.GetCurrentDir()
	result := common.GetLastPath(path)

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestGetRootPath(t *testing.T) {
	msg := "/home/kernel/project/go/common"
	result := common.GetRootPath()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestParseEnvVar(t *testing.T) {
	msg := "/home/kernel"
	path := "${HOME}"
	result := common.ParseEnvVar(path)

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestCallBrowser(t *testing.T) {
	url := "http://www.bing.com"
	msg := "http://www.bing.com"
	err := common.CallBrowser(url)

	if err != nil {
		Test.T(t).Logs(msg).No(err)
	} else {
		Test.T(t).Logs(msg).Ok(err)
	}
}

func TestGetOS(t *testing.T) {
	msg := "linux"
	result := common.GetOS()

	if result != msg {
		Test.T(t).Logs(msg).No(result)
	} else {
		Test.T(t).Logs(msg).Ok(result)
	}
}

func TestGetCustomConfigPath(t *testing.T) {
	msg := "/home/kernel/project/go/common/config/config.toml"
	result := common.GetCustomConfigPath("config", "config.toml")

	if result != msg {
		Test.T(t).Logs(msg).No(result)
	} else {
		Test.T(t).Logs(msg).Ok(result)
	}
}

func TestStringEnd(t *testing.T) {
	Test.End("System")
}
