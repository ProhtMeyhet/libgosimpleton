package simpleton

import(
	"reflect"
)

/*
* CompareAndSwap compares values from right
* to left and swaps them if they are  not equal
* or t1 is nil
*
* Arrays and Slices must have same length
* 
* Structs are only compareable if
* Struct1.Field(i).Type() == Struct2.Field(i).Type()
*/
func CompareAndSwap(values ...interface{}) {
	reflects := make([]reflect.Value,len(values),len(values))
	for k, v := range(values) {
		reflects[k] = reflect.ValueOf(v)
	}

	for i := len(reflects); i > 1; i-- {
		compareAndSwap(reflects[i-2], reflects[i-1])
	}
}

func compareAndSwap(t1, t2 reflect.Value) {
	if t1.Type() != t2.Type() {
		return
	}

	// TODO
	// add recursion detection

	switch t1.Kind() {
	case reflect.Slice:
		if t1.IsNil() {
			if t1.CanSet() {
				t1.Set(t2)
			}
			return
		}
		fallthrough
	case reflect.Array:
		if t1.Len() != t2.Len() {
			return
		}
		for i := 0; i < t1.Len(); i++ {
			compareAndSwap(t1.Index(i), t2.Index(i))
		}
	case reflect.Struct:
		for i, n := 0, t1.NumField(); i < n; i++ {
			compareAndSwap(t1.Field(i), t2.Field(i))
		}
	case reflect.Map:
		if t1.IsNil() {
			if t1.CanSet() {
				t1.Set(t2)
			}
			return
		}
		for _, k := range t1.MapKeys() {
			compareAndSwap(t1.MapIndex(k), t2.MapIndex(k))
		}
	case reflect.Ptr:
		fallthrough
	case reflect.Interface:
		if t1.IsNil() {
			t1.Set(t2)
			return
		}
		compareAndSwap(t1.Elem(), t2.Elem())
	default:
		if t1.Interface() != t2.Interface() {
			if t1.CanSet() {
				t1.Set(t2)
			}
		}
	}
}
