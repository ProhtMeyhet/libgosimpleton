package iotool

import(
	"io"
	"os"
)

func WriteTo(helper *FileHelper, path string, data []byte) (e error) {
	helper = WriteOnly().Copy(helper).ToggleCreate()
	handler, e := Open(helper, path); if e != nil { return }
	written, e := handler.Write(data); if e == nil && written < len(data) { e = io.ErrShortWrite }
	if e1 := handler.Close(); e1 == nil { e = e1 }
	return
}

func WriteFile(helper *FileHelper, path string, data []byte) (e error) {
	return WriteTo(helper, path, data)
}

//TODO simplee for now mirror os functions
func IsExist(e error) bool {
	return os.IsExist(e)
}

//TODO simplee for now mirror os functions
func IsNotExist(e error) bool {
	return os.IsNotExist(e)
}
