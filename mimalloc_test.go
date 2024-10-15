package mimalloc_test

import (
	"testing"

	"github.com/mxmauro/go-mimalloc"
)

// -----------------------------------------------------------------------------

func TestAllocator(t *testing.T) {
	ptr := mimalloc.Malloc(1000)
	if ptr == nil {
		t.Fatalf("insufficient memory")
	}
	defer mimalloc.Free(ptr)

	t.Logf("Allocated pointer at @ %v", ptr)
}

func TestArenaAllocator(t *testing.T) {
	arena := mimalloc.ArenaCreate()
	if arena == nil {
		t.Fatalf("insufficient memory")
	}
	defer mimalloc.ArenaDestroy(arena)

	ptr := mimalloc.ArenaMalloc(arena, 1000)
	if ptr == nil {
		t.Fatalf("insufficient memory")
	}
	defer mimalloc.ArenaFree(arena, ptr)

	t.Logf("Allocated pointer at @ %v", ptr)
}
