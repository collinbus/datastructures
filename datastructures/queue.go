package datastructures

type DynamicQueue struct {
	front *Node
	back  *Node
}

type Queue interface {
	Enqueue(item int)
	Dequeue() int
}

func NewQueue() *DynamicQueue {
	return &DynamicQueue{}
}

func (q *DynamicQueue) Enqueue(item int) {
	node := NewNode(item)
	if q.back == nil {
		q.front = node
		q.back = node
		return
	}
	q.back.Next = node
	node.Prev = q.back
	q.back = node
}

func (q *DynamicQueue) Dequeue() int {
	node := q.front
	q.front = q.front.Next

	if q.front == nil {
		q.back = nil
	}
	return node.Value
}
