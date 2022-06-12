package main

import (
	"Tutorial/tasks"
)

func main() {
	var numbers = []int{1, 2, 3, 4, 5, 6, 7}
	//How to make a deep copy of array or slice
	//How to remove some files from indexing

	var listNode *tasks.ListNode
	listNode = tasks.CreateLinkedList(numbers)
	listNode.PrintLinkedList()
	reversedListNode := listNode.ReverseList()
	reversedListNode.PrintLinkedList()
}
