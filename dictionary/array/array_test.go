package array_test

import (
	"log"
	"testing"

	"github.com/techmexdev/algos/dictionary/array"
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
		i, err := arr.Search(n)
		if err != nil {
			t.Fatalf("have error: %s, want index %v", err, n)
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
		num, err := arr.Get(n)
		if err != nil {
			t.Fatalf("have error: %s, want num %v", err, num)
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
		err := arr.Delete(0)
		if err != nil {
			t.Fatalf("have error: %s, want nil", err)
		}

		num, err = arr.Search(num)
		if err == nil {
			log.Fatalf("want error, have num %v and nil error", num)
		}
	}

	err := arr.Delete(100)
	if err == nil {
		t.Fatalf("have error nil, want error")
	}

	err = arr.Delete(-1)
	if err == nil {
		t.Fatalf("have error nil, want error")
	}
}
