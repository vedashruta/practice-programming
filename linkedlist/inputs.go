package linkedlist

import "fmt"

func Input() {
	// n := 5
	// var node *ListNode
	// for i := range n {
	// 	val := i + 1
	// 	node = InsertRear(node, val)
	// }
	// Print(node)
	// swapped := SwapPairs(node)
	// Print(swapped)

	// Simulate Cycle
	// temp := node
	// temp2 := node
	// for temp2.Next != nil {
	// 	temp2 = temp2.Next
	// }
	// temp2.Next = temp

	// ok := HasCycle(node)
	// if !ok {
	// 	panic(fmt.Errorf("cycle does not exists"))
	// }
	// ok := IsPalindrome(node)
	// if !ok {
	// 	fmt.Printf("not a palindrome")
	// }
	// k := 2
	// head := SwapNodes(node, k)
	// Print(head)
	// head := RemoveElements(node, 2)
	// Print(head)

	var node1 *ListNode
	node1 = InsertRear(node1, 1)
	node1 = InsertRear(node1, 2)
	node1 = InsertRear(node1, 3)

	var node2 *ListNode
	node2 = InsertRear(node2, 4)
	node2 = InsertRear(node2, 5)

	var node3 *ListNode
	node3 = InsertRear(node3, 6)
	node3 = InsertRear(node3, 7)
	node3 = InsertRear(node3, 8)

	temp1 := node1
	for temp1.Next != nil {
		temp1 = temp1.Next
	}
	temp2 := node2
	for temp2.Next != nil {
		temp2 = temp2.Next
	}
	temp1.Next = node3
	temp2.Next = node3
	Print(node1)
	Print(node2)
	Print(node3)
	node := GetIntersectionNode(node1, node2)
	fmt.Println(node.Val)
}
