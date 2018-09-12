package doublyqueue_test

import (
	"testing"

	"github.com/techmexdev/algos/container/queue/doublyqueue"
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
		n, err := q.Dequeue()
		if err != nil {
			t.Fatalf("have error %s, want nil", err)
		} else if n != want {
			t.Fatalf("have dequeued number %v, want %v", n, want)
		}
	}

	_, err := q.Dequeue()
	if err == nil {
		t.Fatal("have nil error, want non-nil")
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
