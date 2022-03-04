package main

import "fmt"

type element struct {
	data int
	next *element
}

type singleList struct {
	head *element
	len  int
	maxValue int
}

func main() {
	var sls []int
	numbers := []int{7, 2, 3, 4, 5, 6, 7, 8, 9}
	singList := initList()
	singList.maxValue = 3
	var setOfStacks []*element

	setOfStacks, sls = singList.Push(setOfStacks, singList, numbers, sls)
	singList.Pop(sls)
	setOfStacks = singList.popAt(setOfStacks, sls)

	singList.showAllElements(setOfStacks)
	//fmt.Print(sls)


}

func initList() *singleList {
	return &singleList{}
}

func (s *singleList) addElementInCurrentStack (data int) int{
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
	return data
}

func (s *singleList) Push (setOfStacks []*element, singList *singleList, numbers []int, sls []int) ([]*element, []int){

	for _, data := range numbers {
		//sls = append(sls, data)
		sls = addMinValue(sls, data)
		if singList.len < singList.maxValue {
			singList.addElementInCurrentStack(data)

		}else if singList.len == singList.maxValue {
			setOfStacks = append(setOfStacks, singList.head)
			singList.head = nil
			singList.len = 0
			singList.addElementInCurrentStack(data)
		}
	}

	setOfStacks = append(setOfStacks, singList.head)
	return setOfStacks, sls
}

func (s *singleList) fillArray(numbers [6]int) [6]int{
	for _, v := range numbers {
		s.addElementInCurrentStack(v)
	}
	return numbers
}

func (s *singleList) Pop(sls []int) {
		i := 0
		current := s.head
		for current.next != nil {
			current = current.next
			i++
		}
		deleteMinValue(sls, current.data)
		s.removeByIndex(i)
}

// Удаляет элемент i из списка.
func (s *singleList) removeByIndex(i int) {

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

func (s *singleList) showCurrentStackElements() error {
	if s.head == nil {
		return fmt.Errorf("single list is empty")
	}
	current := s.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
	return nil
}

func (s *singleList) showAllElements(setOfStacks []*element) {

	for i := 0; i < len(setOfStacks); i++ {
		current := setOfStacks[i]
		for current != nil {
			fmt.Println(current.data)
			current = current.next
		}
	}
}

func (s *singleList) popAt(setOfStacks []*element, sls []int) []*element{

	index := 5
	s.head = setOfStacks[index / 3]
	d := index % 3
	switch d {
		case 1:
			s.removeByIndex(0)
		case 2:
			s.removeByIndex(1)
		case 0:
			s.Pop(sls)
	default:
		fmt.Println("error")
		return nil
	}

	return setOfStacks
}

func addMinValue(sls []int, data int) []int{
	if len(sls) == 0{
		sls = append(sls, data)
	}

	if data < sls[len(sls)-1] {
		sls = append(sls, data)
	}
	return sls
}

func deleteMinValue(sls []int, data int) {

	for i, v := range sls {
		if data == v {
			sls = append(sls[:i], sls[i+1:]...)
		}
	}


}
