package libgosimpleton

import(
	"testing"

	"reflect"
)

func Assert(t *testing.T, value1, value2 interface{}) {
	if !reflect.DeepEqual(value1, value2) {
		t.Fatalf("expected %v, got %v\n", value1, value2)
	}
}
