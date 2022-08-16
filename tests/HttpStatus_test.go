/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"fmt"
	"github.com/kavanahuang/common"
	"testing"
)

type statusCode struct {
	code int
	msg  string
	data interface{}
}

func TestHttpStatusStart(t *testing.T) {
	Test.Start("HttpStatus")
}

func TestHttpStatus(t *testing.T) {
	msg := "Test http status "
	assert := 200
	result := common.OK.Code
	if result == assert {
		Test.T(t).Logs(msg + "code: " + fmt.Sprint(assert)).Ok(result)
	} else {
		Test.T(t).Logs(msg + "code: " + fmt.Sprint(assert)).No(result)
	}

	assertMsg := "OK"
	resultMsg := common.OK.Msg
	if assertMsg == resultMsg {
		Test.T(t).Logs(msg + "Msg: " + assertMsg).Ok(resultMsg)
	} else {
		Test.T(t).Logs(msg + "Msg: " + assertMsg).No(resultMsg)
	}
}

func TestHttpStatusEnd(t *testing.T) {
	Test.End("HttpStatus")
}
