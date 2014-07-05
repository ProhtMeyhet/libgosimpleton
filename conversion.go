package libgosimpleton

import(
	"reflect"
	"unsafe"
)

// Quick copy 
func BytesToString(b []byte) string {
        bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
        stringHeader := reflect.StringHeader{bytesHeader.Data, bytesHeader.Len}
        return *(*string)(unsafe.Pointer(&stringHeader))
}
