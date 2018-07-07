package simpleton

import(
	"bytes"
	"testing"
)

func TestLength(t *testing.T) {
	in, out := "acdc", 4
	if x := Length(in); x != out {
		t.Errorf("Length(%v) == %v, want %v", in, x, out)
	}

	in, out = "世界世界", 4
	if x := Length(in); x != out {
		t.Errorf("Length(%v) == %v, want %v", in, x, out)
	}
}

func TestGetChars(t *testing.T) {
	in, outFrom, outTo, out := "acdc", 2, 3, "d"
	if x := GetChars(in, outFrom, outTo); x != out {
		t.Errorf("GetChars(%v, %v, %v) = %v, want %v", in, outFrom, outTo, x, out)
	}

	outTo, out = 5, "dc"
	if x := GetChars(in, outFrom, outTo); x != out {
		t.Errorf("GetChars('%s', %v, %v) = '%s', want '%s'", in, outFrom, outTo, x, out)
	}

	outTo, out = -1, "d"
	if x := GetChars(in, outFrom, outTo); x != out {
		t.Errorf("GetChars('%s', %v, %v) = '%s', want '%s'", in, outFrom, outTo, x, out)
	}

	in, outFrom, outTo, out = "ac世界dc世界", 2, 3, "世"
	if x := GetChars(in, outFrom, outTo); x != out {
		t.Errorf("GetChars(%v, %v, %v) = %v, want %v", []int32(in), outFrom, outTo, []int32(x), []int32(out))
	}


	in, outFrom, outTo, out = "世界世界", 2, 3, "世"
	if x := GetChars(in, outFrom, outTo); x != out {
		t.Errorf("GetChars(%v, %v, %v) = %v, want %v", []int32(in), outFrom, outTo, []int32(x), []int32(out))
	}

	outTo, out = 5, "世界"
	if x := GetChars(in, outFrom, outTo); x != out {
		t.Errorf("GetChars('%s', %v, %v) = '%s', want '%s'", in, outFrom, outTo, x, out)
	}

	outTo, out = -1, "世"
	if x := GetChars(in, outFrom, outTo); x != out {
		t.Errorf("GetChars('%s', %v, %v) = '%s', want '%s'", in, outFrom, outTo, x, out)
	}

	outTo, out = 0, eMPTY_STRING
	if x := GetChars(in, outFrom, outTo); x != out {
		t.Errorf("GetChars('%s', %v, %v) = '%s', want '%s'", in, outFrom, outTo, x, out)
	}

	in = eMPTY_STRING
	if x := GetChars(in, outFrom, outTo); x != out {
		t.Errorf("GetChars('%s', %v, %v) = '%s', want '%s'", in, outFrom, outTo, x, out)
	}
}


func TestGetLastChar(t *testing.T) {
	in, out := "acdc", "c"
	if x := GetLastChar(in); x != out {
		t.Errorf("GetLastChar(%s) = %s, want %s", in, x, out)
	}

	in, out = "世界世界", "界"
	if x := GetLastChar(in); x != out {
		t.Errorf("GetLastChar(%s) = %s, want %s", in, x, out)
	}

	in, out = eMPTY_STRING, eMPTY_STRING
	if x := GetLastChar(in); x != out {
		t.Errorf("GetLastChar('%s') = '%s', want '%s'", in, x, out)
	}
}


func TestGetLastChars(t *testing.T) {
	in, num, out := "acdc", 2, "dc"
	if x := GetLastChars(in, num); x != out {
		t.Errorf("GetLastRunes(%v, %v) = %s, want %s", in, num, x, out)
	}

	in, num, out = "世界世界", 2, "世界"
	if x := GetLastChars(in, num); x != out {
		t.Errorf("GetLastRunes(%v, %v) = %v, want %v", []int32(in), num, []int32(x), []int32(out))
	}

	in, out = "世", "世"
	if x := GetLastChars(in, num); x != out {
		t.Errorf("GetLastRunes(%v, %v) = %v, want %v", []int32(in), num, []int32(x), []int32(out))
	}

	in, out = eMPTY_STRING, eMPTY_STRING
	if x := GetLastChars(in, num); x != out {
		t.Errorf("GetLastRunes(%v, %v) = %v, want %v", []int32(in), num, []int32(x), []int32(out))
	}
}


func TestGetFirstChar(t *testing.T) {
	in, out := "acdc", "a"
	if x := GetFirstChar(in); x != out {
		t.Errorf("GetFirstChar(%s) = %s, want %s", in, x, out)
	}

	in, out = "世界世界", "世"
	if x := GetFirstChar(in); x != out {
		t.Errorf("GetFirstChar(%s) = %s, want %s", in, x, out)
	}

	in, out = "世", "世"
	if x := GetFirstChar(in); x != out {
		t.Errorf("GetFirstChar(%s) = %s, want %s", in, x, out)
	}

	in, out = eMPTY_STRING, eMPTY_STRING
	if x := GetFirstChar(in); x != out {
		t.Errorf("GetFirstChar(%s) = %s, want %s", in, x, out)
	}
}

func TestGetFirstChars(t *testing.T) {
	in, num, out := "acdc", 2, "ac"
	if x := GetFirstChars(in, num); x != out {
		t.Errorf("GetFirstChars(%s, %d) = %s, want %s", in, num, x, out)
	}

	in, out = "a", "a"
	if x := GetFirstChars(in, num); x != out {
		t.Errorf("GetLastChars(%s, %d) = %s, want %s", in, num, x, out)
	}

	in, out = "世界世界", "世界"
	if x := GetFirstChars(in, num); x != out {
		t.Errorf("GetLastChars(%s, %d) = %s, want %s", in, num, x, out)
	}

	in, out = "世", "世"
	if x := GetFirstChars(in, num); x != out {
		t.Errorf("GetLastChars(%s, %d) = %s, want %s", in, num, x, out)
	}

	in, out = eMPTY_STRING, eMPTY_STRING
	if x := GetFirstChars(in, num); x != out {
		t.Errorf("GetLastChars(%s, %d) = %s, want %s", in, num, x, out)
	}
}


func TestFirstCharsEqual(t *testing.T) {
	in, in2, out := "acdc", "ac", true
	if x := FirstCharsEqual(in, in2); x != out {
		t.Errorf("FirstCharsEqual(%v, %v) = %v, want %v", in, in2, x, out)
	}

	in2, out = "dc", false
	if x := FirstCharsEqual(in, in2); x != out {
		t.Errorf("FirstCharsEqual(%v, %v) = %v, want %v", in, in2, x, out)
	}

	in, in2, out = "世界世界", "世界", true
	if x := FirstCharsEqual(in, in2); x != out {
		t.Errorf("FirstCharsEqual(%v, %v) = %v, want %v", []int32(in), []int32(in2), x, out)
	}

	in, in2, out = eMPTY_STRING, eMPTY_STRING, true
	if x := FirstCharsEqual(in, in2); x != out {
		t.Errorf("FirstCharsEqual(%v, %v) = %v, want %v", in, in2, x, out)
	}
}

func TestLastCharsEqual(t *testing.T) {
	in, in2, out := "acdc", "dc", true
	if x := LastCharsEqual(in, in2); x != out {
		t.Errorf("LastCharsEqual(%s, %s) = %v, want %v", in, in2, x, out)
	}

	in2, out = "ac", false
	if x := LastCharsEqual(in, in2); x != out {
		t.Errorf("LastCharsEqual(%v, %v) = %v, want %v", in, in2, x, out)
	}

	in, in2, out = "世界世界", "世界", true
	if x := FirstCharsEqual(in, in2); x != out {
		t.Errorf("LastCharsEqual(%v, %v) = %v, want %v", []int32(in), []int32(in2), x, out)
	}

	in2, out = "界世", false
	if x := LastCharsEqual(in, in2); x != out {
		t.Errorf("LastCharsEqual(%v, %v) = %v, want %v", []int32(in), []int32(in2), x, out)
	}

	in, in2, out = eMPTY_STRING, eMPTY_STRING, true
	if x := LastCharsEqual(in, in2); x != out {
		t.Errorf("LastCharsEqual(%v, %v) = %v, want %v", in, in2, x, out)
	}
}

func TestCharsEqual(t *testing.T) {
	in, in2, start, out := "acdc", "cd", 1, true
	if x := CharsEqual(in, in2, start); !x {
		t.Errorf("CharsEqual(%s, %s, %v) = %v, want %v", in, in2, start, x, out)
	}

	in, in2, start, out = "世界世界", "界世", 1, true
	if x := CharsEqual(in, in2, start); !x {
		t.Errorf("CharsEqual(%v, %v, %v) = %v, want %v", []int32(in), []int32(in2), start, x, out)
	}

	in2, out = "世界", false
	if x := CharsEqual(in, in2, start); x {
		t.Errorf("CharsEqual(%v, %v, %v) = %v, want %v", []int32(in), []int32(in2), start, x, out)
	}
}


func TestRemoveFirstChars(t *testing.T) {
	in, num, out := "acdc", 2, "dc"
	if x := RemoveFirstChars(in, num); x != out {
		t.Errorf("RemoveFirstChars(%s, %d) = %s, want %s", in, num, x, out)
	}

	in, out = "世界世界", "世界"
	if x := RemoveFirstChars(in, num); x != out {
		t.Errorf("RemoveFirstChars(%s, %d) = %s, want %s", in, num, x, out)
	}
}

func TestRemoveLastChars(t *testing.T) {
	in, num, out := "acdc", 2, "ac"
	if x := RemoveLastChars(in, num); x != out {
		t.Errorf("RemoveLastChars(%s, %d) = %s, want %s", in, num, x, out)
	}

	in, out = "世界世界", "世界"
	if x := RemoveFirstChars(in, num); x != out {
		t.Errorf("RemoveFirstChars(%s, %d) = %s, want %s", in, num, x, out)
	}
}

func TestInsertBeforeFirstChar(t *testing.T) {
	in, insert, out := "acdc", "greatest: ", "greatest: acdc"
	if x := InsertBeforeFirstChar(in, insert); x != out {
		t.Errorf("InsertBeforeFirstChar(%s, %s) = %s, want %s", in, insert, x, out)
	}

	in, out = eMPTY_STRING, insert
	if x := InsertBeforeFirstChar(in, insert); x != out {
		t.Errorf("InsertBeforeFirstChar(%s, %s) = %s, want %s", in, insert, x, out)
	}

	in, insert, out = "世界世界", "/世界/", "/世界/世界世界"
	if x := InsertBeforeFirstChar(in, insert); x != out {
		t.Errorf("InsertBeforeFirstChar(%s, %s) = %s, want %s", in, insert, x, out)
	}

	in, insert, out = eMPTY_STRING, eMPTY_STRING, eMPTY_STRING
	if x := InsertBeforeFirstChar(in, insert); x != out {
		t.Errorf("InsertBeforeFirstChar(%s, %s) = %s, want %s", in, insert, x, out)
	}
}


func TestInsertAfterLastChar(t *testing.T) {
        in, insert, out := "acdc", " are the greatest", "acdc are the greatest"
        if x := InsertAfterLastChar(in, insert); x != out {
                t.Errorf("InsertAfterLastChar(%s, %s) = %s, want %s", in, insert, x, out)
        }

	in, out = eMPTY_STRING, insert
	if x := InsertAfterLastChar(in, insert); x != out {
		t.Errorf("InsertAfterLastChar(%s, %s) = %s, want %s", in, insert, x, out)
	}

	in, insert, out = "世界世界", "世界", "世界世界世界"
	if x := InsertAfterLastChar(in, insert); x != out {
		t.Errorf("InsertAfterLastChar(%v, %v) = %v, want %v", []int32(in), []int32(insert), []int32(x), []int32(out))
	}

	in, insert, out = eMPTY_STRING, eMPTY_STRING, eMPTY_STRING
	if x := InsertAfterLastChar(in, insert); x != out {
		t.Errorf("InsertAfterLastChar(%s, %s) = %s, want %s", in, insert, x, out)
	}
}


func TestInsertString(t *testing.T) {
	in, insert, pos, out := "acdc", "//", 2, "ac//dc"
	if x := InsertString(in, insert, pos); x != out {
		t.Errorf("InsertString(%s, %s, %v) = %s, want %s", in, insert, pos, x, out)
	}

	t.Log(bytes.Index([]byte(in), []byte(GetFirstChars(in, pos))))

	pos, out = 0, "//acdc"
	if x := InsertString(in, insert, pos); x != out {
		t.Errorf("InsertString(%s, %s, %v) = %s, want %s", in, insert, pos, x, out)
	}

	pos, out = len(in), "acdc//"
        if x := InsertString(in, insert, pos); x != out {
                t.Errorf("InsertString(%s, %s, %v) = %s, want %s", in, insert, pos, x, out)
        }

	in, out = eMPTY_STRING, insert
	if x := InsertString(in, insert, pos); x != out {
		t.Errorf("InsertString(%s, %s, %v) = %s, want %s", in, insert, pos, x, out)
	}

	in, insert, pos, out = "世界世界", "/世界/", 2, "世界/世界/世界"
	if x := InsertString(in, insert, pos); x != out {
		t.Errorf("InsertString(%v, %v, %v) = %v, want %v", []int32(in), []int32(insert), pos, []int32(x), []int32(out))
	}
}

func TestRemoveFirstCharsByString(t *testing.T) {
	in, in2, out := "acdc", "zz", "dc"
	if x := RemoveFirstCharsByString(in, in2); x != out {
		t.Errorf("RemoveFirstCharsByString(%s, %s) = %s, want %s", in, in2, x, out)
	}

	in, in2, out = "世界世界", "世界", "世界"
	if x := RemoveFirstCharsByString(in, in2); x != out {
		t.Errorf("RemoveFirstCharsByString(%v, %v) = %v, want %v", in, in2, x, out)
	}
}

func TestRemoveLastCharsByString(t *testing.T) {
	in, in2, out := "acdc", "zz", "ac"
	if x := RemoveLastCharsByString(in, in2); x != out {
		t.Errorf("RemoveLastCharsByString(%s) = %s, want %s", in, x, out)
	}
}



func TestGetLongestStringLength(t *testing.T) {
	in := GetStringSlice()
	out, outIndex := 8, 3
	if x, index := GetLongestStringLength(in...); x != out || index != outIndex {
		t.Errorf("GetLongestStringLength(%v) = %v, want %v or index %v != %v", in, x, out, index, outIndex)
	}
}

func TestGetShortestStringLength(t *testing.T) {
	in := GetStringSlice()
	out, outIndex := 2, 4
	if x, index := GetShortestStringLength(in...); x != out || index != outIndex {
		t.Errorf("GetShortestStringLength(%v) = %v, want %v or index %v != %v", in, x, out, index, outIndex)
	}
}

func TestGetLongestString(t *testing.T) {
	in := GetStringSlice()
	out := "greatest"
	if x := GetLongestString(in...); x != out {
		t.Errorf("GetLongestString(%v) = %v, want %v", in, x, out)
	}
}

func TestGetShortestString(t *testing.T) {
	in := GetStringSlice()
	out := "in"
	if x := GetShortestString(in...); x != out {
		t.Errorf("GetShortestString(%v) = %v, want %v", in, x, out)
	}
}








func GetStringSlice() []string {
	return []string {
		"acdc",
		"are",
		"the",
		"greatest",
		"in",
		"the",
		"world",
	}
}
