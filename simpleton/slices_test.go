package simpleton

import(
	"testing"
)

//https://code.google.com/p/go-wiki/wiki/SliceTricks

func TestSliceStringHasKey(t *testing.T) {
	in := []string{ "a", "c", "d", "c" }
	in2 := 0
	out := true

	if x := HasKeyString(in, in2); x != out {
		t.Errorf("HasKeyString(%v, %v) = %v, want %v", in, in2, x, out)
	}

	in2 = -1
	out = false
	if x := HasKeyString(in, in2); x != out {
		t.Errorf("HasKeyString(%v, %v) = %v, want %v", in, in2, x, out)
	}

	in2 = 4
	if x := HasKeyString(in, in2); x != out {
		t.Errorf("HasKeyString(%v, %v) = %v, want %v", in, in2, x, out)
	}
}

func TestSliceStringEqual(t *testing.T) {
	in := []string{ "a", "c", "d", "c" }
	in2 := []string{ "A", "C", "D", "C" }
	out := false

	if x := SliceStringEqual(in, in2); x != out {
		t.Errorf("SliceStringEqual(%v, %v) = %v, want %v", in, in2, x, out)
	}

	in2 = []string{ "acdc" }
	if x := SliceStringEqual(in, in2); x != out {
		t.Errorf("SliceStringEqual(%v, %v) = %v, want %v", in, in2, x, out)
	}

	in = []string{ "A", "C", "D", "C" }
	if x := SliceStringEqual(in, in2); x != out {
		t.Errorf("SliceStringEqual(%v, %v) = %v, want %v", in, in2, x, out)
	}

	in2 = []string{ "A", "C", "D", "C" }
	out = true
	if x := SliceStringEqual(in, in2); x != out {
		t.Errorf("SliceStringEqual(%v, %v) = %v, want %v", in, in2, x, out)
	}
}

func TestSliceStringCut(t *testing.T) {
	in := []string{ "a", "c", "d", "c" }
	out := []string{ "c", "d", "c" }
	i, j := 0, 1

	if x := SliceStringCut(in, i, j); !SliceStringEqual(x, out) {
		t.Errorf("SliceStringCut(%v, %v, %v) = %v, want %v", in, i, j, x, out)
	}

	in = []string{ "a", "c", "d", "c" }
	out = []string{ "a", "c", "d", "c" }
	j = 0
	if x := SliceStringCut(in, i, j); !SliceStringEqual(x, out) {
		t.Errorf("SliceStringCut(%v, %v, %v) = %v, want %v", in, i, j, x, out)
	}

	in = []string{ "a", "c", "d", "c" }
	out = []string{ "a", "c" }
	i, j = 1, -1
	if x := SliceStringCut(in, i, j); !SliceStringEqual(x, out) {
		t.Errorf("SliceStringCut(%v, %v, %v) = %v, want %v", in, i, j, x, out)
	}
}

func TestSliceStringDelete(t *testing.T) {
	in := []string{ "a", "c", "d", "c" }
	out := []string{ "c", "d", "c" }
	i := 0

	if x := SliceStringDelete(in, i); !SliceStringEqual(x, out) {
		t.Errorf("SliceStringDelete(%v, %v) = %v, want %v", in, i, x, out)
	}
}

func TestSliceStringInsert(t *testing.T) {
	in := []string{ "a", "c", "d", "c" }
	insert := "/"
	out := []string{ "a", "c", "/", "d", "c" }
	i := 2

	if x := SliceStringInsert(in, insert, i); !SliceStringEqual(x, out) {
		t.Errorf("SliceStringInsert(%v, %v, %v) = %v, want %v", in, insert, i, x, out)
	}
}

func TestSliceStringPop(t *testing.T) {
	in := []string{ "a", "c", "d", "c" }
	out1, out2 := []string{ "a", "c", "d" }, "c"

	if x, y := SliceStringPop(in); !SliceStringEqual(y, out1) || x != out2 {
		t.Errorf("SliceStringPop(%v) = %v, %v want %v, %v", in, x, y, out2, out1)
	}

	in = []string{ }
	out1, out2 = []string{ }, ""

	if x, y := SliceStringPop(in); !SliceStringEqual(y, out1) || x != out2 {
		t.Errorf("SliceStringPop(%v) = %v, %v want %v, %v", in, x, y, out2, out1)
	}
}

func TestSliceStringPush(t *testing.T) {
	in, in2 := []string{ "a", "c", "d" }, "c"
	out := []string{ "a", "c", "d", "c" }

	if x := SliceStringPush(in, in2); !SliceStringEqual(x, out) {
		t.Errorf("SliceString.Push(%v, %v) = %v want %v", in, in2, x, out)
	}

	in, in2 = []string{ }, ""
	out = []string{ "" }

	if x := SliceStringPush(in, in2); !SliceStringEqual(x, out) {
		t.Errorf("SliceString.Push(%v, %v) = %v want %v", in, in2, x, out)
	}
}

