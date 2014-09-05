package libgosimpleton

import(
	"bytes"
)

var eMPTY_BYTE []byte

func GetBytes(b []byte, i, j int) []byte {
	if j < 0 {
		j = len(b)+j
	}

	if j < i {
		return eMPTY_BYTE
	}

	if len(b) > 0 {
		if i >= 0 && j > 0 && j < len(b) {
			return b[i:j]
		} else if j > len(b) {
			return b[i:]
		}
	}

	return b
}

func GetLastByte(b []byte) []byte {
	if len(b) > 0 {
		return b[len(b)-1:len(b)]
	}

	return b
}

func GetLastBytes(b []byte, n int) []byte {
	if len(b)-n >= 0 {
		return b[len(b)-n:len(b)]
	} else if len(b) > 0 {
		return b
	}

	return b
}


func GetFirstByte(b []byte) []byte {
	if len(b) > 0 {
		return b[0:1]
	}
	return b
}

func GetFirstBytes(b []byte, n int) []byte {
	if len(b) > n {
		return b[0:n]
	}
	return b
}


func FirstBytesEqual(b, p []byte) bool {
	return BytesEqual(b, p, 0)
}

func LastBytesEqual(b, p []byte) bool {
	return BytesEqual(b, p, len(p))
}

func BytesEqual(b, p []byte, start int) bool {
	if len(b) >= len(p)+start {
		return bytes.Equal(b[start:start+len(p)], p)
	}
	return false
}


func InsertBeforeFirstByte(b, n []byte) []byte {
	pr := make([]byte, len(b)+len(n))
	copy(pr, n[:])
	copy(pr[len(n):], b)
	return pr
}

func InsertAfterLastByte(b, n []byte) []byte {
        pr := make([]byte, len(b)+len(n))
        copy(pr, b[:])
        copy(pr[len(b):], n)
        return pr
}

func InsertBytes(b, n []byte, i int) []byte {
	if i >= len(b) {
		return InsertAfterLastByte(b, n)
	} else if i < 1 {
		return InsertBeforeFirstByte(b, n)
	}

	pr := make([]byte, len(b)+len(n))
	copy(pr, b[:i])
	copy(pr[i:], n)
	copy(pr[i+len(n):], b[i:])
	return pr
}


func RemoveFirstBytes(b []byte, n int) []byte {
	if len(b) > n {
		return b[n:len(b)]
	}
	return eMPTY_BYTE
}

func RemoveLastBytes(b []byte, n int) []byte {
	if len(b) > n {
		return b[0:len(b)-n]
	}
	return eMPTY_BYTE
}

func RemoveBytes(b []byte, i, j int) []byte {
	if j > len(b)+i {
		j = len(b)
	} else if j < 1 {
		j = len(b) + j
	}

	if i < 0 {
		i = len(b) + i
	}

	if len(b) > i && i >= 0 && j > i {
		pr := make([]byte,len(b)+i-j)
		copy(pr, b[0:i])
		copy(pr[i:], b[j:])
		return pr
	}
	return eMPTY_BYTE
}

func RemoveFirstBytesByBytes(b, n []byte) []byte {
	return RemoveFirstBytes(b, len(n))
}

func RemoveLastBytesByBytes(b, n []byte) []byte {
	return RemoveLastBytes(b, len(n))
}

func RemoveBytesByBytes(b []byte, i int, n []byte) []byte {
	return RemoveBytes(b, i, len(n))
}


func GetLongestByte(b ...[]byte) []byte {
	_, index := GetLongestByteLength(b...)
	if index > -1 {
		return b[index]
	}
	return eMPTY_BYTE
}

func GetShortestByte(b ...[]byte) []byte {
	_, index := GetShortestByteLength(b...)
	if index > -1 {
		return b[index]
	}
	return eMPTY_BYTE
}

func GetLongestByteLength(b ...[]byte) (max, index int) {
	index, max = -1, -1
	for i, v := range b {
		if len(v) > max {
			max = len(v)
			index = i
		}
	}
	return
}

func GetShortestByteLength(b ...[]byte) (min, index int) {
	index, min = -1, -1
	for i, v := range b {
		if len(v) < min || min == -1 {
			min = len(v)
			index = i
		}
	}
	return
}

// converts []byte slice to string slice
func ByteSliceToStringSlice(bc ...[]byte) (s []string) {
	s = make([]string, len(bc))
	for key, b := range bc {
		s[key] = string(b)
	}

	return
}

// get the sum of bytes in a non-nested []byte slice
func GetSumOfBytes(of ...[]byte) uint64 {
	return getSumOfBytes(false, of...)
}

// get the sum of bytes in a non-nested byte slice
// trims first and last element
func GetSumOfBytesTrimOuter(of ...[]byte) uint64 {
	return getSumOfBytes(true, of...)
}

// get the sum of bytes in a non-nested []byte slice
// if trim == true, trim first and last element
func getSumOfBytes(trim bool, of ...[]byte) (sum uint64) {
	for key, s := range of {
		if trim && (key == 0 || key == len(of) - 1) {
			s = bytes.TrimSpace(s)
		}

		sum += uint64(len(s))
	}

	return
}
