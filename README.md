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


SIMPLEE
-------
errors with sprintf format capability:

```go
e := simplee.New("Apple is %s")
return e.Format("crap")
```


CREDENTIALS
-----------
credentials library and credentialstool for editing user and password infrastructure. currently unix (/etc/shadow) and sql is implemented.


LICENCE
-------
see LICENCE
