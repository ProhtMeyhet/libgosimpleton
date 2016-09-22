package iotool

import(
	"os"
	osuser "os/user"
	"strconv"
	"syscall"
)

// wrapper around os.FileInfo for easier access to os.FileInfo.Sys()
type FileInfo struct {
	os.FileInfo

	path string
}

// fresh and shiny
func NewFileInfo(apath string, from os.FileInfo) (info *FileInfo) {
	info = &FileInfo{ from, apath }
	return
}

// the given path unaltered
func (info *FileInfo) Path() string {
	return info.path
}

// number of hard links
func (info *FileInfo) HardLinks() (hardLinks uint64) {
	sys := info.Sys(); if sys != nil {
		// 32 bit linux Nlink is 32bit
		// 64 bit linux Nlink is 64bit
		hardLinks = uint64(sys.(*syscall.Stat_t).Nlink)
	}; return
}

// blocksize for file system I/O
func (info *FileInfo) BlockSize() (blockSize int64) {
	sys := info.Sys(); if sys != nil {
		// 32 bit linux Blksize is 32bit
		// 64 bit linux Blksize is 64bit
		blockSize = int64(sys.(*syscall.Stat_t).Blksize)
	}; return
}

// number of 512B blocks allocated
func (info *FileInfo) Blocks() (blocks int64) {
	sys := info.Sys(); if sys != nil {
		blocks = sys.(*syscall.Stat_t).Blocks
	}; return
}

// inode number
func (info *FileInfo) Inode() (inode uint64) {
	sys := info.Sys(); if sys != nil {
		inode = sys.(*syscall.Stat_t).Ino
	}; return
}

// ID of device containing file
func (info *FileInfo) DeviceId() (deviceId uint64) {
	sys := info.Sys(); if sys != nil {
		deviceId = sys.(*syscall.Stat_t).Dev
	}; return
}

// user name of owner
func (info *FileInfo) Owner() (owner string) {
	owner = "?"; userId := info.UserId(); if userId >= 0 {
		userInfo, e := osuser.LookupId(strconv.Itoa(userId)); if e == nil {
			owner = userInfo.Username
		}
	}; return
}

// complete username of owner (eg: Username = gdm, Name = Gnome Display Manager)
func (info *FileInfo) Username() (user string) {
	user = "?"; userId := info.UserId(); if userId >= 0 {
		userInfo, e := osuser.LookupId(strconv.Itoa(userId)); if e == nil {
			user = userInfo.Name
		}
	}; return
}

// user ID of owner
func (info *FileInfo) UserId() (userId int) {
	sys := info.Sys(); if sys != nil {
		userId = int(sys.(*syscall.Stat_t).Uid)
	}; return
}

// group name of owner
func (info *FileInfo) Group() (group string) {
	group = "???"; groupId := info.GroupId(); if groupId >= 0 {
	//FIXME in go 1.7 there will be user.LookupGroup
	//	groupInfo, e := osuser.LookupGroup(userId); if e == nil {
	//		group = groupInfo.Name
			group = strconv.Itoa(int(groupId))
	//	}
	}; return
}

// group id of owner
func (info *FileInfo) GroupId() (groupId int) {
	sys := info.Sys(); if sys != nil {
		groupId = int(sys.(*syscall.Stat_t).Gid)
	}; return
}

// is file not a directory and exeCUTEable
func (info *FileInfo) IsExecuteable() bool {
	return !info.IsDir() && info.Mode() & 0111 == 0111
}

// are two files the same
func (info *FileInfo) Same(compare os.FileInfo) bool {
	// os.SameFile checks only against it's own underlying structure
	if fileInfo, ok := compare.(*FileInfo); ok {
		compare = fileInfo.FileInfo
	}

	return os.SameFile(info.FileInfo, compare)
}
