package libgosimpleton

import(
	"testing"
	"reflect"
)


func TestCompareAndSwap(t *testing.T) {
	standard :=	one { A: "asdf", B: "uiop" }
	garbage :=	one { A: "jkl;", B: "qwer" }
	wanted :=	one { A: "jkl;", B: "qwer" }

	CompareAndSwap(garbage, standard)

	if !reflect.DeepEqual(garbage, wanted) {
		t.Errorf("CompareAndSwap failed! garbage and wanted not equal!")
	}

	standard =	one { A: "asdf", B: "uiop" }
	garbage =	one { A: "jkl;", B: "qwer" }
	wanted =	one { A: "asdf", B: "uiop" }

	CompareAndSwap(standard, garbage)

	if !reflect.DeepEqual(standard, wanted) {
		t.Errorf("CompareAndSwap failed! standard and wanted not equal!")
	}

	empty :=	one{}
	standard =	one { A: "asdf", B: "uiop" }
	wanted =	one { A: "asdf", B: "uiop" }

	CompareAndSwap(empty, standard)

	if !reflect.DeepEqual(standard, wanted) {
		t.Errorf("CompareAndSwap failed! standard and wanted not equal!")
	}

	standard =	one { C: []string{ "a", "c" } }
	garbage =	one { C: []string{ "d", "c" } }
	wanted =	one { C: []string{ "a", "c" } }

	CompareAndSwap(garbage, standard)

	if !reflect.DeepEqual(standard, wanted) {
		t.Errorf("CompareAndSwap failed! standard and wanted not equal!")
	}

}


type one struct {
	A string
	B string
	C []string
}
