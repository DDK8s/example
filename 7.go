package main

import (
	"fmt"
)

type element struct {
	data string
	next *element
}

type singleList struct {
	len  int
	head *element
}

func main() {
	singleList := createLinkedList()
	singleList.fill()
	singleList.makeLoop()
	singleList.findLoop()

}

func (s *singleList) findLoop() (string, bool){
	current := s.head
	m := make(map[*element]interface{})

	for current.next != nil {
		if _, ok := m[current.next]; ok {
			loopPoint := current.next.data
			//fmt.Println("Loop found", loopPoint)
			return loopPoint, true
		}
		m[current.next] = nil
		current = current.next
	}

	//fmt.Println("There is no loop")
	return "Nothing", false

}

func (s *singleList )makeLoop() {
	current := s.head
	for current.next != nil {
		current = current.next
	}
	current.next = s.head.next.next
}

func createLinkedList() *singleList {
	singleList := &singleList{}
	return singleList
}

func (s *singleList) Push (data string) {

	element := &element{
		data: data,
	}
	if s.head == nil {
		s.head = element
	} else {
		current := s.head
		for current.next != nil {
			current = current.next
		}
		current.next = element
	}
	s.len++

}

func (s *singleList) fill() [6]string{
	words := [6]string{"A", "B", "C", "D", "E"}
	for i := range words {
		s.Push(words[i])
	}
	return words
}


func (s *singleList) showAllElements() error {
	if s.head == nil {
		return fmt.Errorf("list is empty")
	}

	current := s.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	return nil
}