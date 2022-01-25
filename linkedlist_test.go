package linkedlist

import "testing"

func checkLen(t *testing.T, list *LinkedList, len int) bool {
	if n := list.Len(); n != len {
		t.Errorf("l.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func TestLinkedList_Add(t *testing.T) {
	list := New()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	checkLen(t, list, 3)
}

func TestLinkedList_Head(t *testing.T) {

}
