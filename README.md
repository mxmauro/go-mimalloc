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
	// Do whatever
	mimalloc.Free(ptr)
}
```

## LICENSES

Underlying `mimalloc` library: [MIT](/mimalloc_LICENSE.txt)

Wrapper: [MIT](/LICENSE.txt)
