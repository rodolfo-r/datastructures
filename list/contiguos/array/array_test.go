package array_test

import (
	"testing"

	"github.com/techmexdev/datastructures/list/contiguos/array"
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
	for n := 0; n < 12; n++ {
		arr.Insert(n)
	}

	for n := 0; n < 12; n++ {
		if i, ok := arr.Search(n); !ok {
			t.Fatalf("expected to find %v", n)
		} else if i != n {
			t.Fatalf("have index: %v, want %v", i, n)
		}
	}
}

func TestGet(t *testing.T) {
	arr := array.New()
	for n := 0; n < 12; n++ {
		arr.Insert(n)
	}

	for n := 0; n < 12; n++ {
		if num, ok := arr.Get(n); !ok {
			t.Fatalf("expected successful get at index %v", n)
		} else if n != num {
			t.Fatalf("have num: %v, want %v", num, n)
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
