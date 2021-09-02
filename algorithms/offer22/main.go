package offer22

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}
	var kth *ListNode
	for kth = head; n > k; n-- {
		kth = kth.Next
	}
	return kth
}

func getKthFromEndSolution2(head *ListNode, k int) *ListNode {
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	return arr[len(arr)-k]
}

type ListNode struct {
	Val  int
	Next *ListNode
}
