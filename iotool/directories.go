package iotool

import(
	"os"
)

func ListDirectory(path string) (names []string, e error) {
	fd, e := os.Open(path); if e != nil { return }; defer fd.Close()
	return fd.Readdirnames(-1)
}
