package common

import (
	"encoding/hex"
	uuid "github.com/iris-contrib/go.uuid"
	"strings"
)

type uid struct{}

var UUID = new(uid)

func (d *uid) UUIDToSqlserverId(u uuid.UUID) string {
	return strings.ToUpper(d.UuidToSqlserverId(u))
}

func (d *uid) UuidToSqlserverId(u uuid.UUID) string {
	buf := make([]byte, 36)

	hex.Encode(buf[0:2], u[3:4])
	hex.Encode(buf[2:4], u[2:3])
	hex.Encode(buf[4:6], u[1:2])
	hex.Encode(buf[6:8], u[0:1])

	buf[8] = '-'
	hex.Encode(buf[9:11], u[5:6])
	hex.Encode(buf[11:13], u[4:5])

	buf[13] = '-'
	hex.Encode(buf[14:16], u[7:8])
	hex.Encode(buf[16:18], u[6:7])

	buf[18] = '-'
	hex.Encode(buf[19:23], u[8:10])

	buf[23] = '-'
	hex.Encode(buf[24:], u[10:])

	return string(buf)
}
