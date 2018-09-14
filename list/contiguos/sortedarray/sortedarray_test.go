package sortedarray_test

import (
	"testing"

	"github.com/rodolfo-r/datastructures/list/contiguos/array"
)

func TestNew(t *testing.T) {
	array.New()
}

func TestInsert(t *testing.T) {
	arr := array.New()
	for n := 0; n < 12; n++ {
		arr.Insert(n)
	}
}

func TestSearch(t *testing.T) {
	arr := array.New()
	for n := 11; n >= 0; n-- {
		arr.Insert(n)
	}

	for n := 0; n < 12; n++ {
		if i, ok := arr.Search(n); !ok {
			t.Fatalf("expected to find %v", n)
		} else if index := 11 - n; i != index {
			t.Fatalf("have index: %v, want %v", i, index)
		}
	}
}

func TestGet(t *testing.T) {
	arr := array.New()
	for n := 11; n >= 0; n-- {
		arr.Insert(n)
	}

	for i := 0; i < 12; i++ {
		if n, ok := arr.Get(i); !ok {
			t.Fatalf("expected successful get at index %v", n)
		} else if num := 11 - i; num != n {
			t.Fatalf("have num: %v, want %v", n, num)
		}
	}
}
func TestDelete(t *testing.T) {
	arr := array.New()
	for n := 0; n < 12; n++ {
		arr.Insert(n)
	}

	for a := 0; a < 12; a++ {
		num, _ := arr.Get(0)
		if ok := arr.Delete(0); !ok {
			t.Fatal("expected successful delete")
		}

		if num, ok := arr.Search(num); ok {
			t.Fatalf("expected unsuccessful search of %v", num)
		}
	}

	if ok := arr.Delete(100); ok {
		t.Fatal("expected unsuccessful delete")
	}

	if ok := arr.Delete(-1); ok {
		t.Fatal("expected unsuccessful delete")
	}
}
