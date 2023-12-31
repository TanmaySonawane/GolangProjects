// made a simple linked list
package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
	prev *Node
}

type List struct {
	head *Node
}

func (l *List) add(value int, pos int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	if pos == 0 {
		newNode.next = l.head
		l.head = newNode
		return
	}

	curr := l.head
	prev := l.head

	for i := 0; i < pos && curr != nil; i++ {
		prev = curr
		curr = curr.next
	}

	prev.next = newNode
	newNode.next = curr
}

func (l *List) remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	curr := l.head
	prev := l.head

	for curr != nil && curr.data != value {
		prev = curr
		curr = curr.next
	}

	if curr != nil {
		prev.next = curr.next
	}
}

func main() {
	list := &List{}
	list.add(1, 0)
	list.add(5, 1)
	list.add(3, 2)
	list.add(2, 3)
	list.add(4, 4)
	list.add(8, 2)

	fmt.Println("Initial List:")
	printList(list)

	list.remove(2)
	fmt.Println("List after removing 2:")
	printList(list)

	list.remove(4)
	fmt.Println("List after removing 4:")
	printList(list)
}

func printList(l *List) {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}
