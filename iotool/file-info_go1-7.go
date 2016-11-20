// +build go1.7

package iotool

import(
	osuser "os/user"
	"strconv"
)

// group name of owner
func (info *FileInfo) Group() (group string) {
	group = "???"; groupId := info.GroupId(); if groupId >= 0 {
		stringGroupId := strconv.Itoa(int(groupId))
		groupInfo, e := osuser.LookupGroupId(stringGroupId); if e == nil {
			group = groupInfo.Name
		}
	}; return
}
