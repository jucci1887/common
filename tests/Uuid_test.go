/**
 * Created by IntelliJ IDEA.
 * User: kernel
 * Mail: kernelman79@gmail.com
 * Date: 2017/8/22
 * Time: 01:38
 */

package tests

import (
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/jucci1887/common"
	"testing"
)

func TestUuidStart(t *testing.T) {
	Test.Start("Uuid")
}

func TestUuidToSqlserverId(t *testing.T) {
	sqlserverId := "44DB726E-9A41-48BE-90A2-FDAB3878AA19"
	goUuid := "6e72db44-419a-be48-90a2-fdab3878aa19"
	toUuid, _ := uuid.FromString(goUuid)
	result := common.UUID.UUIDToSqlserverId(toUuid)

	if result == sqlserverId {
		Test.T(t).Logs(sqlserverId).Ok(result)
	} else {
		Test.T(t).Logs(sqlserverId).No(result)
	}
}

func TestUuidEnd(t *testing.T) {
	Test.End("Uuid")
}
