package user

import(
	osuser "os/user"
	"strconv"
)

// wrapper around os.User for linux
type LinuxUser struct {
	*osuser.User

	id,
	groupId uint32
}

// get current user
func Current() (user *LinuxUser, e error) {
	osUser, e := osuser.Current(); if e != nil { return }
	user = &LinuxUser{ osUser, 0, 0 }
	return user, user.parse()
}

// for more future parsing
func (user *LinuxUser) parse() (e error) {
	e = user.parseId(); if e != nil { return }
	return user.parseGid()
}

// parse id
func (user *LinuxUser) parseId() (e error) {
	userId, e := strconv.ParseUint(user.Uid, 10, 32); if e != nil { return }
	user.id = uint32(userId); return
}

// parse group id
func (user *LinuxUser) parseGid() (e error) {
	groupId, e := strconv.ParseUint(user.Gid, 10, 32); if e != nil { return }
	user.groupId = uint32(groupId); return
}

// getter
func (user *LinuxUser) Id() uint32 {
	return user.id
}

// getter
func (user *LinuxUser) GroupId() uint32 {
	return user.groupId
}
