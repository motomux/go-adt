package linkedlist

import (
	"fmt"
	"io"
	"strconv"
)

// LinkedList is interface of LinkedList implementation
type LinkedList interface {
	Size() (size int)
	Head() (head *Node)
}

// SinglyLinkedList is LinkedList with Singly connection
type SinglyLinkedList struct {
	size int
	head *Node
}

// NewSinglyLinkedList represents singly linked list
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		size: 0,
		head: nil,
	}
}

// Size returns size of LinkedList
func (list *SinglyLinkedList) Size() int {
	return list.size
}

// Head returns head node of LinkedList
func (list *SinglyLinkedList) Head() *Node {
	return list.head
}

// Add adds node to top of LinkedList
func (list *SinglyLinkedList) Add(value int) {
	node := NewNode(value, list.head)
	list.head = node
	list.size++
}

// AddLast adds a node to the last of linked list
func (list *SinglyLinkedList) AddLast(value int) {
	list.size++
	node := NewNode(value, nil)
	if list.head == nil {
		list.head = node
		return
	}

	var last *Node
	for last = list.head; last.Next != nil; last = last.Next {
	}
	last.Next = node
}

// Print iterates all nodes in a linked list and print values
func (list *SinglyLinkedList) Print(w io.Writer) {
	for node := list.head; node != nil; node = node.Next {
		fmt.Fprintf(w, strconv.Itoa(node.Value))
	}
	fmt.Fprintf(w, "\n")
}
