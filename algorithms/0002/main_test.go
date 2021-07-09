package _002

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("", func(t *testing.T) {
		l1 := &ListNode{1, &ListNode{3, &ListNode{3, &ListNode{5, nil}}}}
		l2 := &ListNode{Val: 2, Next: &ListNode{Val: 1}}
		sum := addTwoNumbers(l1, l2)
		printList(sum)
	})
	t.Run("", func(t *testing.T) {
		l1 := &ListNode{}
		l2 := &ListNode{}
		sum := addTwoNumbers(l1, l2)
		printList(sum)
	})
	t.Run("", func(t *testing.T) {
		l1 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
		l2 := &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
		sum := addTwoNumbers(l1, l2)
		printList(sum)
	})
	t.Run("", func(t *testing.T) {
		l1 := &ListNode{}
		l2 := &ListNode{7, &ListNode{3, nil}}
		sum := addTwoNumbers(l1, l2)
		printList(sum)
	})
}
