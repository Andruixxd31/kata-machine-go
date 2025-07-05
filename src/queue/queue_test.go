package queue_test

import (
	"reflect"
	"testing"

	"github.com/andruixxd31/kata-machine-go/src/queue"
)

func TestQueueEnqueue(t *testing.T) {
	tests := []struct {
		name       string
		got        *queue.Queue
		want       *queue.Queue
		valueToAdd int
	}{
		{
			name:       "Enqueue to empty Queue",
			got:        &queue.Queue{},
			want:       &queue.Queue{Head: &queue.Node{Val: 1}, Tail: &queue.Node{Val: 1}, Length: 1},
			valueToAdd: 1,
		},
		{
			name:       "Enqueue to Queue with value",
			got:        &queue.Queue{Head: &queue.Node{Val: 1}, Tail: &queue.Node{Val: 1}, Length: 1},
			want:       &queue.Queue{Head: &queue.Node{Val: 1}, Tail: &queue.Node{Val: 2}, Length: 2},
			valueToAdd: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.got.Enqueue(&queue.Node{Val: tt.valueToAdd})

			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("%#v got %+v - want %+v", tt.got, tt.got, tt.want)
			}
		})
	}
}

func TestQueueDequeue(t *testing.T) {
	tests := []struct {
		name       string
		qGot       *queue.Queue
		qWant      *queue.Queue
		nGot       queue.Node
		nWant      queue.Node
		valueToAdd int
	}{
		{
			name:  "Dequeue empty Queue",
			qGot:  &queue.Queue{},
			qWant: &queue.Queue{},
			nWant: queue.Node{},
		},
		{
			name:  "Dequeue Queue with one value",
			qGot:  &queue.Queue{Head: &queue.Node{Val: 1}, Tail: &queue.Node{Val: 1}, Length: 1},
			qWant: &queue.Queue{},
			nWant: queue.Node{Val: 1},
		},
		{
			name: "Dequeue Queue with values",
			qGot: &queue.Queue{
				Head:   &queue.Node{Val: 1, Next: &queue.Node{Val: 2, Next: nil}},
				Tail:   &queue.Node{Val: 2},
				Length: 2},
			qWant: &queue.Queue{Head: &queue.Node{Val: 2}, Tail: &queue.Node{Val: 2}, Length: 1},
			nWant: queue.Node{Val: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.nGot = tt.qGot.Dequeue()

			if !reflect.DeepEqual(tt.qGot, tt.qWant) {
				t.Errorf("Eror with given queue %#v got %+v - want %+v", tt.qGot, tt.qGot, tt.qWant)
			}
			if !reflect.DeepEqual(tt.nGot, tt.nWant) {
				t.Errorf("Error with node returned got %+v - want %+v", tt.nGot, tt.nWant)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		name  string
		queue *queue.Queue
		got   *queue.Node
		want  *queue.Node
	}{
		{
			name:  "Peek empty Queue",
			queue: &queue.Queue{},
			want:  nil,
		},
		{
			name:  "Peek Queue with value",
			queue: &queue.Queue{Head: &queue.Node{Val: 1}, Tail: &queue.Node{Val: 1}, Length: 1},
			want:  &queue.Node{Val: 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.got = tt.queue.Peek()

			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("%#v got %+v - want %+v", tt.queue, tt.got, tt.want)
			}
		})
	}
}
