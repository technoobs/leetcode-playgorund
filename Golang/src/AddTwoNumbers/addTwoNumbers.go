package main

import (
	"flag"
	"fmt"
)

// ListNode is the definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

var numberA int
var numberB int

func main() {
	fmt.Println("Add Two Numbers")
	flag.IntVar(&numberA, "IntegerA", 0, "IntegerA")
	flag.IntVar(&numberB, "IntegerB", 0, "IntegerB")
	flag.Parse()

	nodeA := generateLinkedList(numberA)
	nodeB := generateLinkedList(numberB)

	readLinkedList(nodeA)
	fmt.Println("")
	readLinkedList(nodeB)
	fmt.Println("")

	nodeResult := solutionOne(nodeA, nodeB)
	readLinkedList(nodeResult)
}

// --------------------------- Solution --------------------------//
func solutionOne(nodeA *ListNode, nodeB *ListNode) *ListNode {
	resultNode := ListNode{Val: 0}
	var addOne bool
	value := nodeA.Val + nodeB.Val
	if value >= 10 {
		addOne = true
	} else {
		addOne = false
	}
	resultNode.Val = value % 10
	nextNode := ListNode{Val: 0}
	if nodeA.Next == nil {
		tempNode := ListNode{Val: 0}
		nodeA.Next = &tempNode
	}
	if nodeB.Next == nil {
		tempNode := ListNode{Val: 0}
		nodeB.Next = &tempNode
	}

	resultNode.Next = &nextNode
	return solutionOneRecur(&resultNode, &nextNode, nodeA.Next, nodeB.Next, addOne)
}

func solutionOneRecur(startNode *ListNode, currentNode *ListNode,
	nodeA *ListNode, nodeB *ListNode, addOne bool) *ListNode {
	var value int
	nextNode := ListNode{Val: 0}

	if nodeA.Next == nil && nodeB.Next == nil {
		v := nodeA.Val + nodeB.Val
		if addOne {
			currentNode.Val = (v + 1) % 10
		} else {
			currentNode.Val = v % 10
		}
		if v >= 10 {
			nextNode.Val = 1
			currentNode.Next = &nextNode
		}
		return startNode
	}
	if addOne {
		value = nodeA.Val + nodeB.Val + 1
	} else {
		value = nodeA.Val + nodeB.Val
	}
	currentNode.Val = value % 10
	currentNode.Next = &nextNode
	if value >= 10 {
		addOne = true
	} else {
		addOne = false
	}
	if nodeA.Next == nil && nodeB.Next != nil {
		tempNode := ListNode{Val: 0}
		nodeA.Next = &tempNode
	}
	if nodeA.Next != nil && nodeB.Next == nil {
		tempNode := ListNode{Val: 0}
		nodeB.Next = &tempNode
	}

	return solutionOneRecur(startNode, &nextNode, nodeA.Next, nodeB.Next, addOne)
}

// --------------------------- Helper Function --------------------------//

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
