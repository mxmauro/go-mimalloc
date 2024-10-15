# go-mimalloc

A Golang wrapper of the [mimalloc](https://github.com/microsoft/mimalloc) general-purpose allocator library.

## Requeriments

This library makes use of `C` code. To compile it, `CGO` must be enabled by adding `CGO_ENABLED=1` environment variable
and `gcc` compiler must be installed and accessible through the system PATH.

## Usage

```golang
package main

import (
	"github.com/mxmauro/go-mimalloc"
)

func main() {
	ptr := mimalloc.Malloc(1000)
	if ptr == nil {
		panic("insufficient memory")
    }
	defer mimalloc.Free(ptr)

	// Do whatever
	
}
```

## LICENSE

Like the original mimalloc code, this wrapper also uses the [MIT](/LICENSE) license.
