libgosimpleton
==============

libgosimpleton is a library for functionality i commonly use.

old description: a minimalistic library for simpletons like me that don't
wanna reserve memory in their minds for stuff like:

```go
copy(a[i:], a[j:])
for k, n := len(a)-j+i, len(a); k < n; k ++ {
       a[k] = s.empty
}
a = a[:len(a)-j+i]
```

PARALLEL
--------
Do some work in parallel:

```go
work, count := parallel.NewWorkManual(16), 0
work.Do(func() {
	work.Lock()
	count++
	work.Unlock()
})

println(count) // prints 16
```

Open a file, read in one thread do your function in another thread:
```go
work := OpenFileDoWork(helper, path, func(buffered *NamedBuffer) {
	into := make([]byte, 512)
	for {
		println("reading " + buffered.Name())
		read, e := buffered.Read(into); if e != nil { /* handle errors and EOF */ }
 		/* do work */
	}
})

work.Wait()
```
hash files, print to Stdout

```go
OpenFilesFromListDoWork(helper, func(buffered *iotool.NamedBuffer) {
	hasher := sha256.New()
	io.Copy(hasher, buffered)
	println(hex.EncodeToString(hasher.Sum(nil)))	
}, path1, path2, paths...).Wait()
```

IOTOOL
------
Functions for io. Open files with a helper capable of doing file advice. This also makes opening files clearer compared to os.Open in the standard library.

```go
helper := iotool.ReadOnly().FileAdviceReadSequential()
handler, e := iotool.Open(helper, "filename")
```

SLICES
------
run MAKE_SLICES.sh to make slices.


CREDENTIALS
-----------
credentials library (and credentialstool) for editing user and password infrastructure. currently unix (/etc/shadow) and sql is implemented (not fully in credentialstool).


LICENCE
-------
see LICENCE
