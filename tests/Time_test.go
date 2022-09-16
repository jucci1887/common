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
	"time"
)

func TestTimeStart(t *testing.T) {
	Test.Start("Time")
}

func TestGetTodayDate(t *testing.T) {
	dateFormat := "2006-01-02"
	msg := time.Now().Format(dateFormat)
	result := common.Time.Now().GetTodayDate()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestToStandard(t *testing.T) {
	timeFormat := "2006-01-02 15:04:05"
	msg := time.Now().Format(timeFormat)
	result := common.Time.Now().ToStandard()

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestToTimestamp(t *testing.T) {
	timeFormat := "2006-01-02 15:04:05"
	standard := time.Now().Format(timeFormat)
	toUnixTime, _ := time.Parse(timeFormat, standard)
	msg := toUnixTime.Unix()
	result := common.Time.Now().ToTimestamp()

	if result == msg {
		Test.T(t).Logs(fmt.Sprint(msg)).Ok(result)
	} else {
		Test.T(t).Logs(fmt.Sprint(msg)).No(result)
	}
}

func TestGetTodayDateString(t *testing.T) {
	dateFormat := "20060102"
	msg := time.Now().Format(dateFormat)
	result := common.Time.Now().GetTodayDateString(dateFormat)

	if result == msg {
		Test.T(t).Logs(msg).Ok(result)
	} else {
		Test.T(t).Logs(msg).No(result)
	}
}

func TestTimeEnd(t *testing.T) {
	Test.End("Time")
}
