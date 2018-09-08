package queue

import (
	"testing"

	"github.com/techmexdev/algos/queue"
)

func TestEnqueue(t *testing.T) {
	q := queue.New()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)
}

func TestDequeue(t *testing.T) {
	q := queue.New()
	q.Enqueue(0)
	q.Enqueue(1)
	q.Enqueue(2)

	nn := []int{0, 1, 2}
	for i := 0; i < len(nn); i++ {
		n, err := q.Dequeue()
		if err != nil {
			t.Fatalf("have error %s, want nil", err)
		} else if n != nn[i] {
			t.Fatalf("have dequeued number %v, want %v", n, nn[i])
		}
	}

	_, err := q.Dequeue()
	if err == nil {
		t.Fatal("have nil error, want non-nil")
	}

}
