package _0021

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	l := &ListNode{}
	merge(l1, l2, l)
	return l
}

func merge(l1, l2, l *ListNode) {
	if l1 == nil && l2 == nil {
		return
	}
	if l1 == nil {
		l.Val = l2.Val
		if l2.Next != nil {
			l.Next = &ListNode{}
			merge(l1, l2.Next, l.Next)
		}
	} else if l2 == nil {
		l.Val = l1.Val
		if l1.Next != nil {
			l.Next = &ListNode{}
			merge(l1.Next, l2, l.Next)
		}
	} else if l1.Val <= l2.Val {
		l.Val = l1.Val
		l.Next = &ListNode{}
		merge(l1.Next, l2, l.Next)
	} else {
		l.Val = l2.Val
		l.Next = &ListNode{}
		merge(l1, l2.Next, l.Next)
	}
}
