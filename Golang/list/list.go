package list

import "errors"

type ListNode struct {
	Val int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func (l *LinkedList)Size() int {
	return l.size
}

func (l *LinkedList)Head() *ListNode {
	return l.head
}
func (l *LinkedList)Add(val int) {
	if l.head == nil {
		l.head = &ListNode{ Val: val }
		l.size += 1
		return
	}

	current := l.head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = &ListNode{ Val: val }
	l.size += 1
}
func (l *LinkedList)Find(val int) (*ListNode, error) {
	if l.head == nil {
		return nil, errors.New("could not found")
	}

	current := l.head
	for current != nil {
		if current.Val == val {
			return current, nil
		}
		current = current.Next
	}

	return nil, errors.New("could not found")
}

func (l *LinkedList)ForEach(callback func(int)) {
	if l.head == nil {
		return
	}
	
	current := l.head

	for current != nil {
		callback(current.Val)
		current = current.Next
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}