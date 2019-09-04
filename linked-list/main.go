package main

import "fmt"

// ListNode contains value and pointer to next node
type ListNode struct {
	val  int
	next *ListNode
}

// NewNode will create new node and return address
func NewNode(num int) *ListNode {
	return &ListNode{num, nil}
}

func (node *ListNode) clear() {
	node.val = 0
	node.next = nil
}

// LinkedList contains pointer to head of list
type LinkedList struct {
	head *ListNode
}

// NewList will create list by given array and return address
func NewList(array []int) *LinkedList {
	var list LinkedList
	for _, num := range array {
		node := NewNode(num)
		list.insertEnd(node)
	}
	return &list
}

// Insert to front of the list
func (list *LinkedList) insertFront(node *ListNode) {
	if list.head == nil {
		list.head = node
	} else {
		node.next = list.head
		list.head = node
	}
}

// Insert after specified node
func (list *LinkedList) insertAfter(node *ListNode, newNode *ListNode) {
	if node == nil {
		return
	}
	newNode.next = node.next
	node.next = newNode
}

// Insert to end of the list
func (list *LinkedList) insertEnd(node *ListNode) {
	if list.head == nil {
		list.head = node
	} else {
		tail := list.head
		for tail.next != nil {
			tail = tail.next
		}
		tail.next = node
	}
}

// Delete front of the list
func (list *LinkedList) deleteFront() {
	target := list.head
	list.head = list.head.next
	target.clear()
}

func (list *LinkedList) delete(key int) {
	if list.head == nil {
		return
	}
	prev := &ListNode{0, list.head}
	for prev.next.next != nil {
		if prev.next.val == key {
			target := prev.next
			prev.next = target.next
			target.clear()
			break
		}
		prev = prev.next
	}
}

// Delete end of the list
func (list *LinkedList) deleteEnd() {
	if list.head == nil {
		return
	}
	tailPrev := &ListNode{0, list.head}
	for tailPrev.next.next != nil {
		tailPrev = tailPrev.next
	}
	tailPrev.next.clear()
	tailPrev.next = nil
}

func (list *LinkedList) search(key int) *ListNode {
	node := list.head
	for node != nil {
		if node.val == key {
			break
		}
		node = node.next
	}
	return node
}

func (list *LinkedList) print() {
	node := list.head
	for node != nil {
		if node == list.head {
			fmt.Print(node.val)
		} else {
			fmt.Print(" -> ", node.val)
		}
		node = node.next
	}
	fmt.Print("\n")
}

func main() {
	array := []int{5, 4, 3}
	list := NewList(array)
	fmt.Println("Initialize with array:", array)
	list.print()
	fmt.Println("Insert Front 0")
	list.insertFront(NewNode(0))
	list.print()
	fmt.Println("Insert End 6")
	list.insertEnd(NewNode(6))
	list.print()
	fmt.Println("Delete front")
	list.deleteFront()
	list.print()
	fmt.Println("Delete end")
	list.deleteEnd()
	list.print()
	fmt.Println("Search 4 and insert 7 after it")
	node := list.search(4)
	list.insertAfter(node, &ListNode{7, nil})
	list.print()
	fmt.Println("Delete 1")
	list.delete(1)
	list.print()
	fmt.Println("Delete 4")
	list.delete(4)
	list.print()
}
