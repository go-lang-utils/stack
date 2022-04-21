package stack

type node[K any] struct {
	val  K
	next *node[K]
}

type Stack[K any] struct {
	top  *node[K]
	size int
}

func New[K any](items ...K) Stack[K] {
	s := Stack[K]{
		top:  nil,
		size: 0,
	}
	s.Push(items...)
	return s
}

func (s *Stack[K]) Push(items ...K) {
	for _, item := range items {
		s.size++
		newNode := node[K]{
			val:  item,
			next: nil,
		}

		if s.top != nil {
			newNode.next = s.top
		}
		s.top = &newNode
	}
}

func (s *Stack[K]) Pop() K {
	s.size--
	n := s.top
	s.top = n.next

	return n.val
}

func (s *Stack[K]) Peek() K {
	return s.top.val
}

func (s *Stack[K]) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack[K]) ToSlice() []K {
	currentNode := s.top
	var slice []K
	for currentNode != nil {
		slice = append(slice, currentNode.val)
		currentNode = currentNode.next
	}

	return slice
}

func (s *Stack[K]) Size() int {
	return s.size
}
