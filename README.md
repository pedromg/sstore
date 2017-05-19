sstore
======

Simple Go key:value file storage

Open and store key:value specs, per line.
Useful for in file definitions.

## Use

To use it, just __import__ the package in the import section:

```go
import (
	"github.com/pedromg/sstore"
)
```

Then initialize a struct like:

```go
var my_store sstore.SStore
```

To __create__ a file:

```go
my_store.Name = "filename" // can have path and extension. I use .sstore
my_store.Descr = "This is the first line, unusable but descriptive."
my_store.Data = map[string]string {
	"key1":"value1",
	"key2":"value2",
}
err := sstore.CreateFile(my_store)
```

That's it. 

To __read__ a file the struct is populated by the data, not the first line.

```go
var err error
my_store, err = sstore.ReadFile("filename")
my_store.Data["key1"] // will return value1
```

## Test
```
$ go test sstore_test.go
```

### License
Copyright (c) pedro mg 2014 (pedro.mota@gmail.com)
All code is licensed under the [MIT License](http://www.opensource.org/licenses/mit-license.php)
