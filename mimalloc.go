//go:build cgo

package mimalloc

//go:generate go run ./internal/gen.go

// #include "mimalloc_amalgamation.h"
import "C"
import (
	"unsafe"
)

// -----------------------------------------------------------------------------

// Malloc allocates unmanaged memory of the given size in bytes.
func Malloc(size uintptr) unsafe.Pointer {
	ptr := C.mi_malloc(C.size_t(size))
	return unsafe.Pointer(ptr)
}

// Calloc allocates a contiguous block of unmanaged memory. Total amount in bytes is count multiplied by size.
func Calloc(count uintptr, size uintptr) unsafe.Pointer {
	ptr := C.mi_calloc(C.size_t(count), C.size_t(size))
	return unsafe.Pointer(ptr)
}

// Realloc changes the size of a previously allocated block of memory. Returns the new location on success.
func Realloc(ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
	newptr := C.mi_realloc(ptr, C.size_t(size))
	return unsafe.Pointer(newptr)
}

// Free deallocates a block of memory.
func Free(ptr unsafe.Pointer) {
	C.mi_free(ptr)
}

// MSize returns the size of an allocated block of unmanaged memory.
func MSize(ptr unsafe.Pointer) uintptr {
	return uintptr(C.mi_usable_size(ptr))
}

// -----------------------------------------------------------------------------

// ArenaCreate creates a new arena. An arena groups a set of memory allocations.
func ArenaCreate() unsafe.Pointer {
	ptr := C.mi_heap_new()
	return unsafe.Pointer(ptr)
}

// ArenaDestroy destroys the given arena. All allocated memory of this arena will be also freed.
func ArenaDestroy(arena unsafe.Pointer) {
	C.mi_heap_destroy((*C.mi_heap_t)(arena))
}

// ArenaMalloc allocates memory like Malloc but in the given arena.
func ArenaMalloc(arena unsafe.Pointer, size uintptr) unsafe.Pointer {
	ptr := C.mi_heap_malloc((*C.mi_heap_t)(arena), C.size_t(size))
	return unsafe.Pointer(ptr)
}

// ArenaCalloc allocates memory like Calloc but in the given arena.
func ArenaCalloc(arena unsafe.Pointer, count uintptr, size uintptr) unsafe.Pointer {
	ptr := C.mi_heap_calloc((*C.mi_heap_t)(arena), C.size_t(count), C.size_t(size))
	return unsafe.Pointer(ptr)
}

// ArenaRealloc re-allocates memory like Realloc but in the given arena.
func ArenaRealloc(arena unsafe.Pointer, ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
	newptr := C.mi_heap_realloc((*C.mi_heap_t)(arena), ptr, C.size_t(size))
	return unsafe.Pointer(newptr)
}

// ArenaFree deallocates a block of memory in the given arena.
func ArenaFree(_ unsafe.Pointer, ptr unsafe.Pointer) {
	// NOTE: See https://github.com/microsoft/mimalloc/issues/598
	C.mi_free(ptr)
}

// ArenaMSize returns the size of an allocated block of unmanaged memory.
func ArenaMSize(_ unsafe.Pointer, ptr unsafe.Pointer) uintptr {
	return uintptr(C.mi_usable_size(ptr))
}
