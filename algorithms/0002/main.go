package _002

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	sum(l1, l2)
	return l1
}

func sum(l1, l2 *ListNode) {
	for {
		if l1 == nil {
			break
		}

		if l2 != nil {
			l1.Val = l1.Val + l2.Val
			if l2 = l2.Next; l2 != nil && l1.Next == nil {
				l1.Next = &ListNode{}
			}
		}

		if l1.Val >= 10 {
			if l1.Next == nil {
				l1.Next = &ListNode{}
			}
			l1.Next.Val = l1.Next.Val + l1.Val/10
			l1.Val = l1.Val % 10
		}

		l1 = l1.Next
	}
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var l = &ListNode{}
	sum2(l1, l2, l)
	return l
}

func sum2(l1, l2, l *ListNode) {
	for l != nil {
		if l1 != nil {
			l.Val = l.Val + l1.Val
			if l1 = l1.Next; l1 != nil {
				l.Next = &ListNode{}
			}
		}
		if l2 != nil {
			l.Val = l.Val + l2.Val
			if l2 = l2.Next; l2 != nil {
				l.Next = &ListNode{}
			}
		}

		if l.Val >= 10 {
			val := l.Val % 10
			l.Next = &ListNode{Val: (l.Val - val) / 10}
			l.Val = val
		}
		l = l.Next
	}
}

func getLen(l *ListNode) int {
	length := 0
	for l != nil {
		length++
		l = l.Next
	}
	return length
}
func revertList(l *ListNode) *ListNode {
	if l == nil {
		return nil
	}
	var l2 *ListNode
	for l != nil {
		item := &ListNode{Val: l.Val}
		item.Next = l2
		l2 = item
		l = l.Next
	}
	return l2
}

func printList(l *ListNode) {
	fmt.Print("[")
	for l != nil {
		fmt.Printf(fmt.Sprintf("%d ", l.Val))
		l = l.Next
	}
	fmt.Print("]\n")
}
