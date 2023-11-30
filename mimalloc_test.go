package mimalloc_test

import (
	"testing"

	"github.com/mxmauro/go-mimalloc"
)

// -----------------------------------------------------------------------------

func TestAllocator(t *testing.T) {
	ptr := mimalloc.Malloc(1000)
	t.Logf("Allocated pointer at @ %v", ptr)
	mimalloc.Free(ptr)
}

func TestArenaAllocator(t *testing.T) {
	arena := mimalloc.ArenaCreate()
	ptr := mimalloc.ArenaMalloc(arena, 1000)
	t.Logf("Allocated pointer at @ %v", ptr)
	mimalloc.ArenaFree(arena, ptr)
	mimalloc.ArenaDestroy(arena)
}
