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

/*
// #define _OW_SOURCE 
func CryptGenSalt(prefix string, count int, input string, size int) string {
	cPrefix := C.CString(prefix)
	cInput := C.CString(input)
	cSize := C.int(size)
	cCount := C.ulong(count)
	//data := C.struct_crypt_data{}
	out := C.GoString(C.crypt_gensalt(cPrefix, cCount, cInput, cSize))
	//			&data, cOutputSize ))
	C.free(unsafe.Pointer(cPrefix))
	C.free(unsafe.Pointer(cSalt))
	C.free(unsafe.Pointer(cInput))
	C.free(unsafe.Pointer(cSize))
	C.free(unsafe.Pointer(cOutputSize))
	return out
}*/

/*
type Crypt struct {
	salter
}

func (crypt *Crypt) Hash(key string) string {

}

func (crypt *Crypt) SetSalter(to SalterInterface) {
	crypt.salter = to
}

func (crypt *Crypt) GetSalter() SalterInterface {
	if crypt.salter == nil {
		crypt.salter = NewCryptSalter()
	}

	return crypt.salter
}
*/
