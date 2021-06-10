package Queue


type Queue struct{
	root *QueueNode
	tail *QueueNode
	size int
}

func(q *Queue) Add(e interface{}){
	new_node := &QueueNode{key: e}
	if q.root == nil{
		q.root = new_node
		q.tail = new_node
		q.size++
		return
	}
	temp := q.tail
	for temp.next != nil{
		temp = temp.next
	}
	new_node.prev = temp
	temp.next = new_node
	q.tail = new_node
	q.size++
}

func(q *Queue) Pop() interface{}{
	if q.root == nil{
		panic("cannot pop queue any more")
	}
	value := q.root.key
	q.root = q.root.next
	q.size--
	return value
}

func(q *Queue) IsEmpty() bool{
	return q.root == nil
}



