package iotool

import(
	"testing"

	"os"
)

func TestOpenFlags(t *testing.T) {
	helper := ReadOnly()

	if helper.OpenFlags() != os.O_RDONLY {
		t.Errorf("expected %v (ReadOnly), got %v", os.O_RDONLY, helper.OpenFlags())
	}

	helper = helper.ToggleCreate()
	if helper.OpenFlags() != os.O_RDONLY|os.O_CREATE || !helper.HasCreate() {
		t.Errorf("expected %v (ReadOnly|Create), got %v", os.O_RDONLY|os.O_CREATE, helper.OpenFlags())
	}

	helper = WriteOnly()
	if helper.OpenFlags() != os.O_WRONLY {
		t.Errorf("expected %v (WriteOnly), got %v", os.O_WRONLY, helper.OpenFlags())
	}

	helper = helper.ToggleCreate()
	if helper.OpenFlags() != os.O_WRONLY|os.O_CREATE || !helper.HasCreate() {
		t.Errorf("expected %v (ReadOnly|Create), got %v", os.O_WRONLY|os.O_CREATE, helper.OpenFlags())
	}

	helper = helper.ToggleCreate().ToggleAppend()
	if helper.OpenFlags() != os.O_WRONLY|os.O_APPEND || !helper.HasAppend() || helper.HasCreate() {
		t.Errorf("expected %v (ReadOnly|Append), got %v", os.O_WRONLY|os.O_CREATE|os.O_APPEND, helper.OpenFlags())
	}

	helper = ReadAndWrite()
	if helper.OpenFlags() != os.O_RDWR {
		t.Errorf("expected %v (ReadWrite), got %v", os.O_RDWR, helper.OpenFlags())
	}

	// TODO more testing
}
