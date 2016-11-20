// +build !go1.7

package iotool

import(
	"strconv"
)

// version for go 1.6 -  will just return string of GroupId
func (info *FileInfo) Group() (group string) {
	group = "???"; groupId := info.GroupId()
	group = strconv.Itoa(int(groupId))
	return
}
