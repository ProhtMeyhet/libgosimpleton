package libgocredentials

import(
	"unsafe"
)

// #cgo LDFLAGS: -lcrypt -Wall
// #define _GNU_SOURCE
// #define _XOPEN_SOURCE
// #include <crypt.h>
// #include <stdlib.h>
import "C"

// wrapper around crypt_r
// see man crypt(3) for help
func CCrypt(key, salt string) string {
	cKey := C.CString(key)
	cSalt := C.CString(salt)
	data := C.struct_crypt_data{}
	out := C.GoString(C.crypt_r(cKey, cSalt, &data))
	C.free(unsafe.Pointer(cKey))
	C.free(unsafe.Pointer(cSalt))
	return out
}
