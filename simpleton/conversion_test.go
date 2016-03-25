package simpleton

import(
	"testing"
)

func TestBytesToString(t *testing.T) {
	in := []byte("acdc")
	const out = "acdc"
	if x := BytesToString(in); x != out {
		t.Errorf("BytesToString(%v) = %v, want %v", in, x, out)
	}
}

func BenchmarkBytesToString(b *testing.B) {
	in := []byte("acdc")
	for i := 0; i < b.N; i++ {
		BytesToString(in)
	}
}
