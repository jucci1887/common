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
	"strconv"
	"testing"
)

func TestFormatStart(t *testing.T) {
	Test.Start("Format")
}

func TestIntToString(t *testing.T) {
	msg := "10"
	result := common.Format.FromInt(10).ToString()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestStringToInt(t *testing.T) {
	msg := 10
	result := common.Format.FromString("10").ToInt()

	if result == msg {
		Test.T(t).Logs(strconv.Itoa(msg)).Ok(result)
	} else {
		Test.T(t).Logs(strconv.Itoa(msg)).No(result)
	}
}

func TestFormatEnd(t *testing.T) {
	Test.End("Format")
}
