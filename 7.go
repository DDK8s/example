package main

import "fmt"

type element struct {
	data int
	next *element
}


type sinList struct {
	len  int
	head *element
}

func main() {

	numbers := [6]int{8, 9, 7, 1, 2, 3}
	Act(numbers)
}

func Act(numbers [6]int) int {
	k := 2
	sinList := initiationList()
	numbers = sinList.fillArray(numbers)

	p1 := sinList.head
	p2 := sinList.head

	for i := 0; i < k-1; i++ {
		p2 = p2.next
	}

	for p2.next != nil {
		p1 = p1.next
		p2 = p2.next
	}

	// Вывести весь список.
	/*err := sinList.showList()
	if err != nil {
		fmt.Println(err.Error())
	}*/

	return p1.data
}

func initiationList() *sinList {
	return &sinList{}
}

func (s *sinList) fillArray(numbers [6]int) [6]int {

	for _, v := range numbers {
		s.addInEnd(v)
	}
	return numbers
}

func (s *sinList) addInEnd(data int) {
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

func (s *sinList) addInFront(data int) {
	element := &element{
		data: data,
	}
	if s.head == nil {
		s.head = element
	} else {
		element.next = s.head
		s.head = element
	}
	s.len++
	return
}

// Удаляет элемент i из списка.
func (s *sinList) removeElement(i int) {

	current := s.head
	j := 0
	for j < i-1 {
		j++
		current = current.next
	}
	remove := current.next
	current.next = remove.next
	s.len--
	return
}

func (s *sinList) showList() error {
	if s.head == nil {
		return fmt.Errorf("sinList is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	return nil
}
