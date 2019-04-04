package main

import (
	"fmt"
)

// ListNode is the definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	fmt.Println("Add Two Numbers")

	node := generateLinkedList(52111)
	readLinkedList(node)
}

// function to generate linked list from the provided integer
func generateLinkedList(num int) *ListNode {
	v1 := num / 10
	v2 := num % 10
	node := ListNode{Val: v2}
	if v1 == 0 {
		return &node
	}
	return combineNode(&node, &node, v1)
}

func combineNode(startNode *ListNode, currentNode *ListNode, nextVal int) *ListNode {
	v1 := nextVal / 10
	v2 := nextVal % 10
	nextNode := ListNode{Val: v2}
	currentNode.Next = &nextNode
	if v1 == 0 {
		return startNode
	} else {
		return combineNode(startNode, &nextNode, v1)
	}
}

// function to read linked list
func readLinkedList(node *ListNode) {
	if node != nil {
		fmt.Printf("%d", node.Val)
		if node.Next != nil {
			fmt.Print(" -> ")
			readLinkedList(node.Next)
		}
	} else {
		return
	}
}
