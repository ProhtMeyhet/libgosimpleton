libgosimpleton
==============

libgosimpleton - a minimalistic library for simpletons like me that don't
wanna reserve memory in their minds for stuff like:

```go
copy(a[i:], a[j:])
for k, n := len(a)-j+i, len(a); k < n; k ++ {
       a[k] = s.empty
}
a = a[:len(a)-j+i]
```

SLICES
------
run MAKE_SLICES.sh to make slices.

LICENCE
-------
see LICENCE
