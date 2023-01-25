package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var previous, current *ListNode = nil, head
	for current != nil {
		previous, current, current.Next = current, current.Next, previous
	}
	return previous
}

func main() {
	reverseList(nil)
}
