package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

//! create newNode and insert at end of it.
func (ll *LinkedList) insertAtEnd(data int) {
	// create new node and store his address in newNode variable
	newNode := &Node{data: data}

	// check if ll.head == nil
	if ll.head == nil {
		ll.head = newNode
		return
	}
	// create new next to point in current node to make check on it.
	current := ll.head
	// loop over the list to reach at end the list.
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// ! insert at beginning of the list
func (ll *LinkedList) insertAtBeginning(data int) {
	newNode := &Node{data: data}
	newNode.next = ll.head
	ll.head = newNode
}

// ! insert at Middle of the list
func (ll *LinkedList) insertAtMiddle(data, key int) {
	newNode := &Node{data: data}
	current := ll.head
	for current != nil && current.data != key {
		current = current.next
	}
	if current == nil {
		fmt.Println("Key not found in the linked list")
	}
	newNode.next = current.next
	current.next = newNode
}

// ! print linked list data
func (ll *LinkedList) PrintData() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Printf("null")

}

func main() {
	ll := LinkedList{}

	ll.insertAtEnd(2)
	ll.insertAtBeginning(4)
	ll.insertAtEnd(1)
	ll.insertAtEnd(3)
	ll.insertAtMiddle(10 , 1)

	fmt.Printf("Linked-list : ") // Linked-list : 4 -> 2 -> 1 -> 10 -> 3 -> null

	ll.PrintData()

}
