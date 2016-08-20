package iotool

import(
	"os"
	"testing"
)

// actually a compiler test. if it doesn't blow up, it's good
func TestIfCounterfeit(t *testing.T) {
	var f FileInterface

	f = NewFakeStdFile(os.Stdin, os.Stdout)

	f.Fd()
}
