package libgosimpleton

import(
	"fmt"
)

/*
* come on, i'm not gonna remeber that!
* that looks creepy!
* https://code.google.com/p/go-wiki/wiki/SliceTricks
**/

// ----------delete everything below this line --------------- //
var emptyString string

func HasKeyString(a []string, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceStringEqual(a, to []string) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceStringExactlyEqual(a, to []string) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyString(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceStringCut(a []string, i, j int) []string {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyString
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceStringDelete(a []string, i int) []string {
	if i > 0 {
		return SliceStringCut(a, i-1, i)
	} else {
		return SliceStringCut(a, i, i+1)
	}
}

func SliceStringInsert(a []string, what string, i int) []string {
	a = append(a, emptyString)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceStringPop(a []string) (string, []string) {
	// pop is empty
	if len(a) < 1 {
		return emptyString, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceStringPush(a []string, x string) []string {
	a = append(a, x)
	return a
}
var emptyInt int

func HasKeyInt(a []int, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceIntEqual(a, to []int) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceIntExactlyEqual(a, to []int) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyInt(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceIntCut(a []int, i, j int) []int {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyInt
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceIntDelete(a []int, i int) []int {
	if i > 0 {
		return SliceIntCut(a, i-1, i)
	} else {
		return SliceIntCut(a, i, i+1)
	}
}

func SliceIntInsert(a []int, what int, i int) []int {
	a = append(a, emptyInt)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceIntPop(a []int) (int, []int) {
	// pop is empty
	if len(a) < 1 {
		return emptyInt, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceIntPush(a []int, x int) []int {
	a = append(a, x)
	return a
}
var emptyUint uint

func HasKeyUint(a []uint, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceUintEqual(a, to []uint) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceUintExactlyEqual(a, to []uint) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyUint(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceUintCut(a []uint, i, j int) []uint {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyUint
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceUintDelete(a []uint, i int) []uint {
	if i > 0 {
		return SliceUintCut(a, i-1, i)
	} else {
		return SliceUintCut(a, i, i+1)
	}
}

func SliceUintInsert(a []uint, what uint, i int) []uint {
	a = append(a, emptyUint)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceUintPop(a []uint) (uint, []uint) {
	// pop is empty
	if len(a) < 1 {
		return emptyUint, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceUintPush(a []uint, x uint) []uint {
	a = append(a, x)
	return a
}
var emptyUint8 uint8

func HasKeyUint8(a []uint8, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceUint8Equal(a, to []uint8) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceUint8ExactlyEqual(a, to []uint8) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyUint8(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceUint8Cut(a []uint8, i, j int) []uint8 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyUint8
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceUint8Delete(a []uint8, i int) []uint8 {
	if i > 0 {
		return SliceUint8Cut(a, i-1, i)
	} else {
		return SliceUint8Cut(a, i, i+1)
	}
}

func SliceUint8Insert(a []uint8, what uint8, i int) []uint8 {
	a = append(a, emptyUint8)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceUint8Pop(a []uint8) (uint8, []uint8) {
	// pop is empty
	if len(a) < 1 {
		return emptyUint8, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceUint8Push(a []uint8, x uint8) []uint8 {
	a = append(a, x)
	return a
}
var emptyUint16 uint16

func HasKeyUint16(a []uint16, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceUint16Equal(a, to []uint16) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceUint16ExactlyEqual(a, to []uint16) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyUint16(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceUint16Cut(a []uint16, i, j int) []uint16 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyUint16
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceUint16Delete(a []uint16, i int) []uint16 {
	if i > 0 {
		return SliceUint16Cut(a, i-1, i)
	} else {
		return SliceUint16Cut(a, i, i+1)
	}
}

func SliceUint16Insert(a []uint16, what uint16, i int) []uint16 {
	a = append(a, emptyUint16)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceUint16Pop(a []uint16) (uint16, []uint16) {
	// pop is empty
	if len(a) < 1 {
		return emptyUint16, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceUint16Push(a []uint16, x uint16) []uint16 {
	a = append(a, x)
	return a
}
var emptyUint32 uint32

func HasKeyUint32(a []uint32, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceUint32Equal(a, to []uint32) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceUint32ExactlyEqual(a, to []uint32) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyUint32(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceUint32Cut(a []uint32, i, j int) []uint32 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyUint32
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceUint32Delete(a []uint32, i int) []uint32 {
	if i > 0 {
		return SliceUint32Cut(a, i-1, i)
	} else {
		return SliceUint32Cut(a, i, i+1)
	}
}

func SliceUint32Insert(a []uint32, what uint32, i int) []uint32 {
	a = append(a, emptyUint32)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceUint32Pop(a []uint32) (uint32, []uint32) {
	// pop is empty
	if len(a) < 1 {
		return emptyUint32, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceUint32Push(a []uint32, x uint32) []uint32 {
	a = append(a, x)
	return a
}
var emptyUint64 uint64

func HasKeyUint64(a []uint64, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceUint64Equal(a, to []uint64) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceUint64ExactlyEqual(a, to []uint64) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyUint64(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceUint64Cut(a []uint64, i, j int) []uint64 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyUint64
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceUint64Delete(a []uint64, i int) []uint64 {
	if i > 0 {
		return SliceUint64Cut(a, i-1, i)
	} else {
		return SliceUint64Cut(a, i, i+1)
	}
}

func SliceUint64Insert(a []uint64, what uint64, i int) []uint64 {
	a = append(a, emptyUint64)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceUint64Pop(a []uint64) (uint64, []uint64) {
	// pop is empty
	if len(a) < 1 {
		return emptyUint64, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceUint64Push(a []uint64, x uint64) []uint64 {
	a = append(a, x)
	return a
}
var emptyInt8 int8

func HasKeyInt8(a []int8, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceInt8Equal(a, to []int8) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceInt8ExactlyEqual(a, to []int8) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyInt8(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceInt8Cut(a []int8, i, j int) []int8 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyInt8
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceInt8Delete(a []int8, i int) []int8 {
	if i > 0 {
		return SliceInt8Cut(a, i-1, i)
	} else {
		return SliceInt8Cut(a, i, i+1)
	}
}

func SliceInt8Insert(a []int8, what int8, i int) []int8 {
	a = append(a, emptyInt8)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceInt8Pop(a []int8) (int8, []int8) {
	// pop is empty
	if len(a) < 1 {
		return emptyInt8, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceInt8Push(a []int8, x int8) []int8 {
	a = append(a, x)
	return a
}
var emptyInt16 int16

func HasKeyInt16(a []int16, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceInt16Equal(a, to []int16) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceInt16ExactlyEqual(a, to []int16) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyInt16(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceInt16Cut(a []int16, i, j int) []int16 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyInt16
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceInt16Delete(a []int16, i int) []int16 {
	if i > 0 {
		return SliceInt16Cut(a, i-1, i)
	} else {
		return SliceInt16Cut(a, i, i+1)
	}
}

func SliceInt16Insert(a []int16, what int16, i int) []int16 {
	a = append(a, emptyInt16)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceInt16Pop(a []int16) (int16, []int16) {
	// pop is empty
	if len(a) < 1 {
		return emptyInt16, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceInt16Push(a []int16, x int16) []int16 {
	a = append(a, x)
	return a
}
var emptyInt32 int32

func HasKeyInt32(a []int32, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceInt32Equal(a, to []int32) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceInt32ExactlyEqual(a, to []int32) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyInt32(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceInt32Cut(a []int32, i, j int) []int32 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyInt32
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceInt32Delete(a []int32, i int) []int32 {
	if i > 0 {
		return SliceInt32Cut(a, i-1, i)
	} else {
		return SliceInt32Cut(a, i, i+1)
	}
}

func SliceInt32Insert(a []int32, what int32, i int) []int32 {
	a = append(a, emptyInt32)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceInt32Pop(a []int32) (int32, []int32) {
	// pop is empty
	if len(a) < 1 {
		return emptyInt32, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceInt32Push(a []int32, x int32) []int32 {
	a = append(a, x)
	return a
}
var emptyInt64 int64

func HasKeyInt64(a []int64, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceInt64Equal(a, to []int64) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceInt64ExactlyEqual(a, to []int64) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyInt64(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceInt64Cut(a []int64, i, j int) []int64 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyInt64
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceInt64Delete(a []int64, i int) []int64 {
	if i > 0 {
		return SliceInt64Cut(a, i-1, i)
	} else {
		return SliceInt64Cut(a, i, i+1)
	}
}

func SliceInt64Insert(a []int64, what int64, i int) []int64 {
	a = append(a, emptyInt64)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceInt64Pop(a []int64) (int64, []int64) {
	// pop is empty
	if len(a) < 1 {
		return emptyInt64, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceInt64Push(a []int64, x int64) []int64 {
	a = append(a, x)
	return a
}
var emptyFloat32 float32

func HasKeyFloat32(a []float32, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceFloat32Equal(a, to []float32) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceFloat32ExactlyEqual(a, to []float32) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyFloat32(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceFloat32Cut(a []float32, i, j int) []float32 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyFloat32
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceFloat32Delete(a []float32, i int) []float32 {
	if i > 0 {
		return SliceFloat32Cut(a, i-1, i)
	} else {
		return SliceFloat32Cut(a, i, i+1)
	}
}

func SliceFloat32Insert(a []float32, what float32, i int) []float32 {
	a = append(a, emptyFloat32)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceFloat32Pop(a []float32) (float32, []float32) {
	// pop is empty
	if len(a) < 1 {
		return emptyFloat32, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceFloat32Push(a []float32, x float32) []float32 {
	a = append(a, x)
	return a
}
var emptyFloat64 float64

func HasKeyFloat64(a []float64, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceFloat64Equal(a, to []float64) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceFloat64ExactlyEqual(a, to []float64) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyFloat64(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceFloat64Cut(a []float64, i, j int) []float64 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyFloat64
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceFloat64Delete(a []float64, i int) []float64 {
	if i > 0 {
		return SliceFloat64Cut(a, i-1, i)
	} else {
		return SliceFloat64Cut(a, i, i+1)
	}
}

func SliceFloat64Insert(a []float64, what float64, i int) []float64 {
	a = append(a, emptyFloat64)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceFloat64Pop(a []float64) (float64, []float64) {
	// pop is empty
	if len(a) < 1 {
		return emptyFloat64, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceFloat64Push(a []float64, x float64) []float64 {
	a = append(a, x)
	return a
}
var emptyComplex64 complex64

func HasKeyComplex64(a []complex64, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceComplex64Equal(a, to []complex64) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceComplex64ExactlyEqual(a, to []complex64) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyComplex64(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceComplex64Cut(a []complex64, i, j int) []complex64 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyComplex64
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceComplex64Delete(a []complex64, i int) []complex64 {
	if i > 0 {
		return SliceComplex64Cut(a, i-1, i)
	} else {
		return SliceComplex64Cut(a, i, i+1)
	}
}

func SliceComplex64Insert(a []complex64, what complex64, i int) []complex64 {
	a = append(a, emptyComplex64)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceComplex64Pop(a []complex64) (complex64, []complex64) {
	// pop is empty
	if len(a) < 1 {
		return emptyComplex64, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceComplex64Push(a []complex64, x complex64) []complex64 {
	a = append(a, x)
	return a
}
var emptyComplex128 complex128

func HasKeyComplex128(a []complex128, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceComplex128Equal(a, to []complex128) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceComplex128ExactlyEqual(a, to []complex128) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyComplex128(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceComplex128Cut(a []complex128, i, j int) []complex128 {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyComplex128
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceComplex128Delete(a []complex128, i int) []complex128 {
	if i > 0 {
		return SliceComplex128Cut(a, i-1, i)
	} else {
		return SliceComplex128Cut(a, i, i+1)
	}
}

func SliceComplex128Insert(a []complex128, what complex128, i int) []complex128 {
	a = append(a, emptyComplex128)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceComplex128Pop(a []complex128) (complex128, []complex128) {
	// pop is empty
	if len(a) < 1 {
		return emptyComplex128, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceComplex128Push(a []complex128, x complex128) []complex128 {
	a = append(a, x)
	return a
}
var emptyByte byte

func HasKeyByte(a []byte, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceByteEqual(a, to []byte) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceByteExactlyEqual(a, to []byte) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyByte(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceByteCut(a []byte, i, j int) []byte {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyByte
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceByteDelete(a []byte, i int) []byte {
	if i > 0 {
		return SliceByteCut(a, i-1, i)
	} else {
		return SliceByteCut(a, i, i+1)
	}
}

func SliceByteInsert(a []byte, what byte, i int) []byte {
	a = append(a, emptyByte)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceBytePop(a []byte) (byte, []byte) {
	// pop is empty
	if len(a) < 1 {
		return emptyByte, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceBytePush(a []byte, x byte) []byte {
	a = append(a, x)
	return a
}
var emptyRune rune

func HasKeyRune(a []rune, k int) (b bool) {
	defer func() {
		if p := recover(); p != nil {
			b = false
		}
	}()

	b = true

	_ = a[k]

	return
}

func SliceRuneEqual(a, to []rune) bool {
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", to)
}

func SliceRuneExactlyEqual(a, to []rune) bool {
	if len(a) != len(to) {
		return false
	}

	for k, v := range a {
		if HasKeyRune(to, k) {
			if v != to[k] {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func SliceRuneCut(a []rune, i, j int) []rune {
	if j < 0 {
		j = len(a) + j
	}

	if j <= i {
		return a
	}

	// consider the slice cut
	if len(a) < i {
		return a
	}

	if len(a) < j {
		j = len(a)
	}

	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k ++ {
	    a[k] = emptyRune
	}
	a = a[:len(a)-j+i]
	return a
}

func SliceRuneDelete(a []rune, i int) []rune {
	if i > 0 {
		return SliceRuneCut(a, i-1, i)
	} else {
		return SliceRuneCut(a, i, i+1)
	}
}

func SliceRuneInsert(a []rune, what rune, i int) []rune {
	a = append(a, emptyRune)
	copy(a[i+1:], a[i:])
	a[i] = what
	return a
}

func SliceRunePop(a []rune) (rune, []rune) {
	// pop is empty
	if len(a) < 1 {
		return emptyRune, a
	}

	x := a[len(a)-1]
	a = a[:len(a)-1]
	return x, a
}

func SliceRunePush(a []rune, x rune) []rune {
	a = append(a, x)
	return a
}
