package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasSeen(seen []int, val int) bool {
	for x := 0; x < len(seen); x++ {
		if seen[x] == val {
			return true
		}
	}
	return false
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	prev := head
	seen := make([]int, 0)

	for curr != nil {
		if hasSeen(seen, curr.Val) {
			// Delete link
			prev.Next = curr.Next
		} else {
			prev = curr
		}

		seen = append(seen, curr.Val)
		curr = curr.Next
	}

	return head
}

func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val, "->")
		curr = curr.Next
	}
	fmt.Println("")
}

func main() {

	head := &ListNode{
		1, &ListNode{
			1, &ListNode{
				2, &ListNode{
					3, &ListNode{
						3, nil,
					},
				},
			},
		},
	}
	printList(head)

	newHead := deleteDuplicates(head)
	printList(newHead)
}