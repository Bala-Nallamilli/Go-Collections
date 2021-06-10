package Stack


type StackGo struct {
	root *StackNode
	tail *StackNode
	size int
}

func(s *StackGo) Push(e interface{}){
	new_node := &StackNode{data: e}
	if s.root == nil{
		s.root = new_node
		s.tail = new_node
		s.size++
	}else{
		temp := s.root
		for temp.next != nil{
			temp = temp.next
		}
		new_node.prev = temp
		temp.next = new_node
		s.tail = new_node
		s.size++
	}
}

func(s *StackGo)Pop() interface{}{
	vl := s.tail.data
	s.tail = s.tail.prev
	s.size--
	return vl
}

func(s *StackGo) IsEmpty() bool{
	return s.tail == nil
}










