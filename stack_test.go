package stack

import (
	"log"
	"testing"
)

func TestStackCreation(t *testing.T) {
	s := New[int]()
	if s.top != nil {
		log.Fatal("s.top should be nil")
	}
	if s.size != 0 {
		log.Fatalf("s.size should be %d", 0)
	}
}

func TestStack_PushOne(t *testing.T) {
	s := New[int]()
	s.Push(1)
	if s.size != 1 {
		log.Fatalf("s.size should be %d", 1)
	}
	if s.top.val != 1 {
		log.Fatalf("s.top should be %d", 1)
	}
	if s.top.next != nil {
		log.Fatal("s.top.next should be nil")
	}
}

func TestStack_PushTwo(t *testing.T) {
	s := New[int]()
	s.Push(1, 2)
	if s.size != 2 {
		log.Fatalf("s.size should be %d", 2)
	}
	if s.top.val != 2 {
		log.Fatalf("s.top should be %d", 2)
	}
	if s.top.next.val != 1 {
		log.Fatalf("s.top.next should be %d", 1)
	}
	if s.top.next.next != nil {
		log.Fatal("s.top.next.next should be nil")
	}
}

func TestStack_Pop(t *testing.T) {
	s := New[int]()
	s.Push(1, 2)
	el := s.Pop()
	if el != 2 {
		log.Fatalf("s.pop should return %d", 2)
	}
	if s.size != 1 {
		log.Fatalf("s.size should be %d", 1)
	}
	if s.top.next != nil {
		log.Fatal("s.top.next should be nil")
	}
}

func TestStack_Peek(t *testing.T) {
	s := New[int]()
	s.Push(1, 2)
	top := s.Peek()
	if top != 2 {
		log.Fatalf("s.Peek() should return %d", 2)
	}
}

func TestStack_IsEmpty(t *testing.T) {
	s := New[int]()
	s.Push(1, 2)

	if s.IsEmpty() {
		log.Fatal("s.isEmpty should be false")
	}

	s2 := New[int]()
	if !s2.IsEmpty() {
		log.Fatalf("s.isEmpty should be true")
	}
}

func TestStack_Size(t *testing.T) {
	s := New[int]()
	s.Push(1, 2)

	if s.Size() != 2 {
		log.Fatalf("s.Size() should be %d", 2)
	}
}

func TestStack_ToSlice(t *testing.T) {
	s := New[int]()
	s.Push(5, 4, 3, 2, 1)
	expectedSlice := []int{1, 2, 3, 4, 5}

	actual := s.ToSlice()

	for index, item := range actual {
		if item != expectedSlice[index] {
			log.Fatalf("actual[%d] should match expectedSlice[%d]", index, index)
		}
	}
}
