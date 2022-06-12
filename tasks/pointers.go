package tasks

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

type Integer int

func (value Integer) PrintInt() {
	fmt.Println(value)
}

// Head 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> nil Tail
func CreateLinkedList(numbers []int) *ListNode {
	var listNode = new(ListNode)
	listNode = nil

	for i := len(numbers) - 1; i >= 0; i-- {
		currentListNode := new(ListNode)
		currentListNode.Value = numbers[i]
		currentListNode.Next = listNode
		listNode = currentListNode
	}

	return listNode
}

func (head *ListNode) PrintLinkedList() {
	for head != nil {
		fmt.Println(head.Value)
		head = head.Next
	}
}

func (head *ListNode) ReverseList() *ListNode {
	var previousNode *ListNode = nil
	var currentNode = head
	for currentNode != nil {
		var savedNode = currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = savedNode
	}
	return previousNode
}

func PlayingWithPointers() {
	var c map[int]bool
	c = make(map[int]bool)
	c[6] = true
	var a int
	var b *int
	b = &a
	fmt.Println(b)
}
