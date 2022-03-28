package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	// Условные новые данные.
	data := enterNewData()

	// Инициализирует связанный список с
	// несколькими начальными значениями.
	singleList := createLinkedList()

	// Добавляет элемент в setOfStacks.
	singleList.Push(data)

	// Удаляет последний элемент в связанном списке.
	singleList.Pop()

	// Выводит в терминал элементы из связанного списка.
	singleList.showAllElements()

}

// Функция вызывается в тесте для
// сравнения []expected и []result

func (s *singleList) TestingFor() []int{
	current := s.head
	var result []int
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}

func createLinkedList()  *singleList{
	c := &element{
		data: 3,
		next: nil,
	}

	b := &element{
		data: 2,
		next: c,
	}

	a := &element{
		data: 1,
		next: b,
	}

	singleList := &singleList{a, 2, 1}
	singleList.maxValue = 3
	return singleList
}

func enterNewData() int {
	fmt.Println("Enter data: ")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	t := sc.Text()
	data, err := strconv.Atoi(t)

	if err != nil {
		panic(err)
	}
	return data
}

func (s *singleList) Push (data int) {
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

func (s *singleList) Pop() {
	current := s.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

func (s *singleList) showAllElements() error{
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