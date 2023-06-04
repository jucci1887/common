/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	"github.com/jucci1887/common"
	"testing"
)

func TestStringStart(t *testing.T) {
	Test.Start("String")
}

func TestSub(t *testing.T) {
	dateFormat := "2006-01-02"
	msg := "2006"
	result := common.String.Sub(dateFormat, 0, 4)

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestSubLen(t *testing.T) {
	dateFormat := "2006-01-02"
	msg := "2006"
	result := common.String.SubLen(dateFormat, 0, 4)

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestReverse(t *testing.T) {
	str := "abcdef"
	msg := "fedcba"
	result := common.String.Reverse(str)

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestStringEnd(t *testing.T) {
	Test.End("String")
}
