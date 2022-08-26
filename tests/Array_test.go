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

func TestArrayStart(t *testing.T) {
	Test.Start("Array")
}

func TestReverseStringArray(t *testing.T) {
	arr := []string{"0", "1", "2"}
	msg := []string{"2", "1", "0"}
	result := common.Array.ReverseStringArray(arr)

	if common.Array.EqStringSlice(result, msg) {
		Test.T(t).Logs(fmt.Sprint(msg)).Ok(result)
	} else {
		Test.T(t).Logs(fmt.Sprint(msg)).No(result)
	}
}

func TestEqStringSlice(t *testing.T) {
	arr := []string{"0", "1", "2"}
	target := []string{"2", "1", "0"}
	result := common.Array.EqStringSlice(arr, target)

	if !result {
		Test.T(t).Logs(fmt.Sprint(false)).Ok(result)
	} else {
		Test.T(t).Logs(fmt.Sprint(false)).No(result)
	}

	arr = []string{"0", "1", "2"}
	target = []string{"0", "1", "2"}
	result = common.Array.EqStringSlice(arr, target)

	if result {
		Test.T(t).Logs(fmt.Sprint(true)).Ok(result)
	} else {
		Test.T(t).Logs(fmt.Sprint(true)).No(result)
	}
}

func TestArrayEnd(t *testing.T) {
	Test.End("Array")
}
