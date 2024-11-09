package linkedlist

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertRear(head *ListNode, data int) *ListNode {
	newNode := &ListNode{
		Val: data,
	}
	if head == nil {
		head = newNode
		return head
	}
	if head.Next == nil {
		head.Next = newNode
		return head
	}
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	return head
}

func Print(head *ListNode) {
	curr := head
	if curr == nil {
		fmt.Println("Head is nil")
	}
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

/*
Reverse a SLL using recursion
*/
func ReverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newNode := ReverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newNode
}

/*
Reverse a SLL iteratively
*/
func ReverseListIterative(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	cur := head.Next
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	head.Next = nil
	head = prev
	return head
}

/*
Swap Nodes in Pairs
*/
// Given a linked list, swap every two adjacent nodes and return its head.
// You must solve the problem without modifying the values in the list's nodes
// (i.e., only nodes themselves may be changed.)
func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	Print(head)
	first := head
	second := head.Next
	first.Next = SwapPairs(second.Next)
	second.Next = first
	return second
}

/*
Linked List Cycle 1
*/
// Given head, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached
// again by continuously following the next pointer. Internally, pos is used to denote the
// index of the node that tail's next pointer is connected to. Note that pos is
// not passed as a parameter.
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fptr := head
	sptr := head
	for fptr != nil || fptr.Next != nil {
		fptr = fptr.Next.Next
		sptr = sptr.Next
		if fptr == sptr {
			return true
		}
	}
	return false
}

/*
Palindrome linked list
*/
func MiddleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fastPtr := head
	slowPtr := head
	for fastPtr != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}
	return slowPtr
}

// Can directly call ReverseListIterative or ReverseListRecursive
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newNode := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newNode
}

func IsPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	firstHalfStart := head
	firstHalfEnd := MiddleNode(head)
	secondHalfStart := reverseList(firstHalfEnd)
	for secondHalfStart != nil {
		if secondHalfStart.Val != firstHalfStart.Val {
			return false
		}
		secondHalfStart = secondHalfStart.Next
		firstHalfStart = firstHalfStart.Next
	}
	return true
}

/*
Swapping Nodes in a Linked List
*/
// You are given the head of a linked list, and an integer k.
// Return the head of the linked list after swapping the values of the kth node
// from the beginning and the kth node from the end (the list is 1-indexed).
func SwapNodes(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	size := 0
	curX := head
	for curX != nil {
		size++
		curX = curX.Next
	}
	curX = head
	curY := head
	for i := 1; i < k; i++ {
		curX = curX.Next
	}
	for i := 1; i < size-k+1; i++ {
		curY = curY.Next
	}
	// Swapping directly will not work for odd number of nodes as xPtr==yPtr second step gives 0
	if curX == curY {
		return head
	}
	curX.Val = curX.Val + curY.Val
	curY.Val = curX.Val - curY.Val
	curX.Val = curX.Val - curY.Val
	return head
}

/*
Remove Linked List Elements
*/
// Given the head of a linked list and an integer val, remove all the nodes of the linked list
// that has Node.val == val, and return the new head.
func RemoveElements(head *ListNode, value int) *ListNode {
	// if head == nil || (head.Next == nil && head.Val == value){
	// 	return nil
	// }
	// if head.Next == nil {
	// 	return head
	// }
	temp := &ListNode{}
	temp.Next = head
	cur := temp
	for cur.Next != nil {
		if cur.Next.Val == value {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return temp.Next
}

/*
Intersection of Two Linked Lists
*/
// Given the heads of two singly linked-lists headA and headB,
// return the node at which the two lists intersect. If the two linked lists have
// no intersection at all, return null.
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	aPtr := headA
	bPtr := headB
	for aPtr != bPtr {
		if aPtr == nil {
			aPtr = headB
		}else{
			aPtr = aPtr.Next
		}
		if bPtr == nil {
			bPtr = headA
		}else{
			bPtr = bPtr.Next
		}
	}
	return aPtr
}
