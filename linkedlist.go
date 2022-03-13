// Package linkedlist implements a single linked list.
//
// To iterate over a linkedlist (where list is a *LinkedList):
// for node := list.Head; node != nil; node = node.Next() {
//	// do something with node.Value
// }
//
package linkedlist

import (
	"encoding/json"
)

// Node is an element of a linked list
type Node struct {
	// Next pointer
	next *Node

	// The value stored with this node
	Value interface{}
}

func (node *Node) Next() *Node {
	return node.next
}

func (node *Node) String() ([]byte, error) {
	return json.Marshal(node)
}

type LinkedList struct {
	len  int
	head *Node
}

func (list *LinkedList) Head() *Node {
	return list.head
}

func (list *LinkedList) Len() int {
	return list.len
}

func (list *LinkedList) Add(value interface{}) *Node {
	node := &Node{Value: value}
	if list.head == nil {
		list.head = node
	} else {
		parent := list.head
		for parent.next != nil {
			parent = parent.next
		}
		parent.next = node
	}
	list.len++
	return node
}

func (list *LinkedList) Remove(value interface{}) *Node {
	removed := &Node{}
	current := list.head

	if current == nil {
		return removed
	}

	if current.Value == value {
		list.head = current.next
		removed = current
		list.len--
	} else {
		for current.next != nil {
			if current.next.Value == value {
				removed = current.next
				current.next = current.next.next
				list.len--
				break
			}
			current = current.next
		}
	}
	return removed
}

func (list *LinkedList) RemoveHead() *Node {
	head := list.head

	if head != nil {
		list.head = head.next
		list.len--
	}

	return head
}

func New() *LinkedList {
	//return &LinkedList{}
	return new(LinkedList)
}
