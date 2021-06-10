package Queue


type QueueNode struct {
	key interface{}
	prev *QueueNode
	next *QueueNode
}