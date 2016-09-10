package iotool

import(
	"io"
	"os"
)

// alias for WriteFile to mirror io.ioutil behaviour
func WriteTo(helper *FileHelper, path string, data []byte) (e error) {
	return WriteFile(helper, path, data)
}

// write data to a file
// if you want to append, use FileHelper.ToggleAppend()
// if you want it truncated, use FileHelper.ToggleTruncate()
func WriteFile(helper *FileHelper, path string, data []byte) (e error) {
	helper = WriteOnly().Copy(helper).ToggleCreate()
	handler, e := Open(helper, path); if e != nil { return }
	written, e := handler.Write(data); if e == nil && written < len(data) { e = io.ErrShortWrite }
	if e1 := handler.Close(); e1 == nil { e = e1 }
	return
}

// FIXME needs rethinking
func Remove(handler FileInterface) error {
	return os.Remove(handler.Name())
}

func IsDirectoryE(e error) bool {
	if e == nil { return false }
	message := e.Error()
	return message == IsDirectoryError.Error() || len(message) >= 14 && message[len(message)-14:] == "is a directory"
}

//TODO simplee for now mirror os functions
func IsExist(e error) bool {
	return os.IsExist(e)
}

//TODO simplee for now mirror os functions
func IsNotExist(e error) bool {
	return os.IsNotExist(e)
}
