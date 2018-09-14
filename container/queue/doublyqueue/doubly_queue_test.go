package doublyqueue_test

import (
	"testing"

	"github.com/rodolfo-r/datastructures/container/queue/doublyqueue"
)

func TestEnqueue(t *testing.T) {
	q := doublyqueue.New()
	for n := 1; n <= 5; n++ {
		q.Enqueue(n)
	}
}

func TestDequeue(t *testing.T) {
	q := doublyqueue.New()
	for n := 1; n <= 5; n++ {
		q.Enqueue(n)
	}

	for want := 1; want <= 5; want++ {
		if n, ok := q.Dequeue(); !ok {
			t.Fatalf("expected successful dequeue")
		} else if n != want {
			t.Fatalf("have dequeued number %v, want %v", n, want)
		}
	}

	if _, ok := q.Dequeue(); ok {
		t.Fatalf("expected unsuccessful dequeue")
	}

}

func TestSize(t *testing.T) {
	q := doublyqueue.New()

	for n := 1; n <= 5; n++ {
		q.Enqueue(n)
		if s := q.Size(); s != n {
			t.Fatalf("have size %v, want %v", s, n)
		}
	}
}

func TestFront(t *testing.T) {
	q := doublyqueue.New()

	for n := 1; n <= 5; n++ {
		q.Enqueue(n)
	}

	for n := 1; n <= 5; n++ {
		if f := q.Front(); f != n {
			t.Fatalf("have front %v, want %v", f, n)
		}
		q.Dequeue()
	}
}
