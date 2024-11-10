package datastructures

type DynamicQueue struct {
	Front *Node
	Back  *Node
}

type Queue interface {
	Enqueue(item int)
}

func NewQueue() *DynamicQueue {
	return &DynamicQueue{}
}

func (q *DynamicQueue) Enqueue(item int) {
	node := NewNode(item)
	if q.Back == nil {
		q.Front = node
		q.Back = node
		return
	}
	q.Back.Next = node
	node.Prev = q.Back
	q.Back = node
}
