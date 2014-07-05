package libgosimpleton

import(
	"bytes"
	"testing"
)


func TestGetBytes(t *testing.T) {
	in, outFrom, outTo, out := []byte("acdc"), 2, 3, []byte("d")
	if x := GetBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
		t.Errorf("GetBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
	}

	outTo, out = 5, []byte("dc")
	if x := GetBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
		t.Errorf("GetBytes('%s', %v, %v) = '%s', want '%s'", in, outFrom, outTo, x, out)
	}

	outTo, out = -1, []byte("d")
	if x := GetBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
		t.Errorf("GetBytes('%s', %v, %v) = '%s', want '%s'", in, outFrom, outTo, x, out)
	}

	outTo, out = 0, eMPTY_BYTE
	if x := GetBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
		t.Errorf("GetBytes('%s', %v, %v) = '%s', want '%s'", in, outFrom, outTo, x, out)
	}

	in = eMPTY_BYTE
	if x := GetBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
		t.Errorf("GetBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, in)
	}
}

func TestGetLastByte(t *testing.T) {
	in, out := []byte("acdc"), []byte("c")
	if x := GetLastByte(in); !bytes.Equal(x, out) {
		t.Errorf("GetLastByte(%s) = %s, want %s", in, x, out)
	}
}

func TestGetLastBytes(t *testing.T) {
	in, num, out := []byte("acdc"), 2, []byte("dc")
	if x := GetLastBytes(in, num); !bytes.Equal(x, out) {
		t.Errorf("GetLastBytes(%v, %v) = %v, want %v", in, num, x, out)
	}
}


func TestGetFirstByte(t *testing.T) {
	in, out := []byte("acdc"), []byte("a")
	if x := GetFirstByte(in); !bytes.Equal(x, out) {
		t.Errorf("GetFirstByte(%v) = %v, want %v", in, x, out)
	}
}

func TestGetFirstBytes(t *testing.T) {
	in, num, out := []byte("acdc"), 2, []byte("ac")
	if x := GetFirstBytes(in, num); !bytes.Equal(x, out) {
		t.Errorf("GetFirstBytes(%v, %v) = %v, want %v", in, num, x, out)
	}
}


func TestFirstBytesEqual(t *testing.T) {
	in, in2, out := []byte("acdc"), []byte("ac"), true
	if x := FirstBytesEqual(in, in2); !x {
		t.Errorf("FirstBytesEqual(%v, %v) = %v, want %v", in, in2, x, out)
	}
}

func TestLastBytesEqual(t *testing.T) {
	in, in2, out := []byte("acdc"), []byte("dc"), true
	if x := LastBytesEqual(in, in2); !x {
		t.Errorf("LastBytesEqual(%s, %s) = %b, want %b", in, in2, x, out)
	}
}

func TestBytesEqual(t *testing.T) {
	in, in2, start, out := []byte("acdc"), []byte("cd"), 1, true
	if x := BytesEqual(in, in2, start); !x {
		t.Errorf("LastBytesEqual(%s, %s) = %b, want %b", in, in2, x, out)
	}
}


func TestInsertBeforeFirstByte(t *testing.T) {
	in, insert, out := []byte("acdc"), []byte("greatest: "), []byte("greatest: acdc")
	if x := InsertBeforeFirstByte(in, insert); !bytes.Equal(x, out) {
		t.Errorf("InsertBeforeFirstByte(%s, %s) = %s, want %s", in, insert, x, out)
	}

	in, out = eMPTY_BYTE, insert
	if x := InsertBeforeFirstByte(in, insert); !bytes.Equal(x, out) {
		t.Errorf("InsertBeforeFirstByte(%s, %s) = %s, want %s", in, insert, x, out)
	}
}

func TestInsertAfterLastByte(t *testing.T) {
        in, insert, out := []byte("acdc"), []byte(" are the greatest"), []byte("acdc are the greatest")
        if x := InsertAfterLastByte(in, insert); !bytes.Equal(x, out) {
                t.Errorf("InsertAfterLastByte(%s, %s) = %s, want %s", in, insert, x, out)
        }

	in, out = eMPTY_BYTE, insert
	if x := InsertAfterLastByte(in, insert); !bytes.Equal(x, out) {
		t.Errorf("InsertAfterLastByte(%s, %s) = %s, want %s", in, insert, x, out)
	}
}


func TestInsertBytes(t *testing.T) {
	in, insert, pos, out := []byte("acdc"), []byte("//"), 2, []byte("ac//dc")
	if x := InsertBytes(in, insert, pos); !bytes.Equal(x, out) {
		t.Errorf("InsertBytes(%s, %s, %v) = %s, want %s", in, insert, pos, x, out)
	}

	pos, out = 0, []byte("//acdc")
	if x := InsertBytes(in, insert, pos); !bytes.Equal(x, out) {
		t.Errorf("InsertBytes(%s, %s, %v) = %s, want %s", in, insert, pos, x, out)
	}

	pos, out = len(in), []byte("acdc//")
        if x := InsertBytes(in, insert, pos); !bytes.Equal(x, out) {
                t.Errorf("InsertBytes(%s, %s, %v) = %s, want %s", in, insert, pos, x, out)
        }

	in, out = eMPTY_BYTE, insert
	if x := InsertBytes(in, insert, pos); !bytes.Equal(x, out) {
		t.Errorf("InsertBytes(%s, %s, %v) = %s, want %s", in, insert, pos, x, out)
	}
}


func TestRemoveFirstBytes(t *testing.T) {
	in, num, out := []byte("acdc"), 2, []byte("dc")
	if x := RemoveFirstBytes(in, num); !bytes.Equal(x, out) {
		t.Errorf("RemoveFirstBytes(%s, %d) = %s, want %s", in, num, x, out)
	}
}

func TestRemoveLastByties(t *testing.T) {
	in, num, out := []byte("acdc"), 2, []byte("ac")
	if x := RemoveLastBytes(in, num); !bytes.Equal(x, out) {
		t.Errorf("RemoveLastBytes(%v, %v) = %v, want %v", in, num, x, out)
	}
}

func TestRemoveBytes(t *testing.T) {
        in, outFrom, outTo, out := []byte("acdc"), 2, 3, []byte("acc")
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }

	outFrom, outTo = -2, -1
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }

	out, outFrom, outTo = []byte("acc"), 2, -1
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }

	out, outFrom, outTo = []byte("ac"), 1, -1
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }

	out, outFrom, outTo = []byte("acd"), 3, 0
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }

	out, outFrom, outTo = []byte("c"), 0, -1 //getlastbyte
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }


	// empty & out of index check
	out = eMPTY_BYTE

	outFrom, outTo = -10, 8
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }
	outFrom, outTo = 0, 8
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }

	outFrom, outTo = 0, 0
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }

	outFrom, outTo = 7, 8
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }

	in = eMPTY_BYTE
        if x := RemoveBytes(in, outFrom, outTo); !bytes.Equal(x, out) {
                t.Errorf("RemoveBytes(%s, %v, %v) = %s, want %s", in, outFrom, outTo, x, out)
        }
}


func TestRemoveFirstBytesByBytes(t *testing.T) {
	in, in2, out := []byte("acdc"), []byte("zz"), []byte("dc")
	if x := RemoveFirstBytesByBytes(in, in2); !bytes.Equal(x, out) {
		t.Errorf("RemoveFirstBytesByBytes(%v, %v) = %v, want %v", in, in2, x, out)
	}

	in, out = eMPTY_BYTE, eMPTY_BYTE
        if x := RemoveFirstBytesByBytes(in, in2); !bytes.Equal(x, out) {
                t.Errorf("RemoveFirstBytesByBytes(%v, %v) = %v, want %v", in, in2, x, out)
        }
}

func TestRemoveLastBytesByBytes(t *testing.T) {
	in, in2, out := []byte("acdc"), []byte("zz"), []byte("ac")
	if x := RemoveLastBytesByBytes(in, in2); !bytes.Equal(x, out) {
		t.Errorf("RemoveLastBytesByBytes(%v) = %v, want %v", in, x, out)
	}

	in, out = eMPTY_BYTE, eMPTY_BYTE
        if x := RemoveLastBytesByBytes(in, in2); !bytes.Equal(x, out) {
		t.Errorf("RemoveLastBytesByBytes(%v) = %v, want %v", in, x, out)
	}
}



func TestGetLongestByteLength(t *testing.T) {
	in := GetByteSlice()
	out, outIndex := 8, 3
	if x, index := GetLongestByteLength(in...); x != out || index != outIndex {
		t.Errorf("GetLongestByteLength(%v) = %v, want %v or index %v != %v", in, x, out, index, outIndex)
	}
}

func TestGetShortestByteLength(t *testing.T) {
	in := GetByteSlice()
	out, outIndex := 2, 4
	if x, index := GetShortestByteLength(in...); x != out || index != outIndex {
		t.Errorf("GetShortestByteLength(%v) = %v, want %v or index %v != %v", in, x, out, index, outIndex)
	}
}

func TestGetLongestByte(t *testing.T) {
	in := GetByteSlice()
	out := []byte("greatest")
	if x := GetLongestByte(in...); !bytes.Equal(x, out) {
		t.Errorf("GetLongestByte(%v) = %v, want %v", in, x, out)
	}
}

func TestGetShortestByte(t *testing.T) {
	in := GetByteSlice()
	out := []byte("in")
	if x := GetShortestByte(in...); !bytes.Equal(x, out) {
		t.Errorf("GetShortestByte(%v) = %v, want %v", in, x, out)
	}

	in2, out := []byte("greatest"), []byte("in")
	if x := GetShortestByte(in2, out); !bytes.Equal(x, out) {
		t.Errorf("GetShortestByte(%v) = %v, want %v", in, x, out)
	}

	if x := GetShortestByte(out, in2); !bytes.Equal(x, out) {
		t.Errorf("GetShortestByte(%v) = %v, want %v", in, x, out)
	}
}







func GetByteSlice() [][]byte {
	return [][]byte {
		[]byte("acdc"),
		[]byte("are"),
		[]byte("the"),
		[]byte("greatest"),
		[]byte("in"),
		[]byte("the"),
		[]byte("world"),
	}
}
