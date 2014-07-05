package libgosimpleton

import(
	"bytes"
	"strings"
)

const(
	eMPTY_STRING = ""
)

func Length(r string) int {
	return len([]int32(r))
}

func GetChars(r string, i, j int) string {
	s := []int32(r)
	if len(s) > 0 {
		if j < 0 {
			j = len(s)+j
		}

		if i >= 0 && j > 0 && j < len(s) {
			return string(s[i:j])
		} else if j > len(s) {
			return string(s[i:])
		}
	}
	return eMPTY_STRING
}


func GetLastChar(r string) string {
	if len(r) > 0 {
		s := []int32(r)
		return string(s[len(s)-1:len(s)])
	}
	return eMPTY_STRING
}

func GetLastChars(r string, n int) string {
	s := []int32(r)
	if len(s)-n >= 0 {
		return string(s[len(s)-n:len(s)])
	} else if len(s) > 0 {
		return string(s)
	}
	return eMPTY_STRING
}


func GetFirstChar(r string) string {
	if len(r) > 0 {
		s := []int32(r)
		return string(s[0:1])
	}
	return eMPTY_STRING
}

func GetFirstChars(r string, n int) string {
	s := []int32(r)
	if len(s) > n {
		return string(s[0:n])
	}
	return r
}


func FirstCharsEqual(s, p string) bool {
	return CharsEqual(s, p, 0)
}

func LastCharsEqual(s, p string) bool {
	return CharsEqual(s, p, len(p))
}

func CharsEqual(r, rp string, start int) bool {
	s := []int32(r)
	p := []int32(rp)
	if len(s) >= len(p)+start {
		slice := s[start:start+len(p)]
		return strings.EqualFold(string(slice), rp)
	}
	return false
}


func InsertBeforeFirstChar(s, n string) string {
	b, nn := []byte(s), []byte(n)
	return BytesToString(InsertBeforeFirstByte(b, nn))
}

func InsertAfterLastChar(s, n string) string {
	b, nn := []byte(s), []byte(n)
	return BytesToString(InsertAfterLastByte(b, nn))
}

func InsertString(s, n string, i int) string {
	if len(s) == 0 {
		return InsertBeforeFirstChar(s, n)
	}

	firstChars := GetFirstChars(s, i)
	b, nn := []byte(s), []byte(n)
	ii := bytes.Index(b, []byte(firstChars)) + len(firstChars)
	return BytesToString(InsertBytes(b, nn, ii))
}


func RemoveFirstChars(r string, n int) string {
	s := []int32(r)
	if len(s) > n {
		return string(s[n:len(s)])
	}
	return eMPTY_STRING
}

func RemoveLastChars(r string, n int) string {
	s := []int32(r)
	if len(s) > n {
		return string(s[0:len(s)-n])
	}
	return eMPTY_STRING
}

func RemoveChars(s string, i, j int) string {
	b := []byte(s)
	return BytesToString(RemoveBytes(b, i, j))
}


func RemoveFirstCharsByString(s, n string) string {
	return RemoveFirstChars(s, Length(n))
}

func RemoveLastCharsByString(s, n string) string {
	return RemoveLastChars(s, Length(n))
}

func RemoveCharsByString(s string, i int, n string) string {
	return RemoveChars(s, i, Length(n))
}


func GetLongestString(s ...string) string {
	_, index := GetLongestStringLength(s...)
	if index > -1 {
		return s[index]
	}
	return eMPTY_STRING
}

func GetShortestString(s ...string) string {
	_, index := GetShortestStringLength(s...)
	if index > -1 {
		return s[index]
	}
	return eMPTY_STRING
}

func GetLongestStringLength(s ...string) (max, index int) {
	index, max = -1, -1
	for i, v := range s {
		if Length(v) > max {
			max = Length(v)
			index = i
		}
	}
	return
}

func GetShortestStringLength(s ...string) (min, index int) {
	index, min = -1, -1
	for i, v := range s {
		if Length(v) < min || min == -1 {
			min = Length(v)
			index = i
		}
	}
	return
}
