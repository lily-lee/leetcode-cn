package _0019

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := getNum(head)
	if i == n {
		return head.Next
	}
	deleteN(head, i-n+1)
	return head
}
func deleteN(list *ListNode, i int) {
	for i >= 0 {
		if i == 2 {
			list.Next = list.Next.Next
			break
		}
		list = list.Next
		i--
	}
}

func getNum(list *ListNode) int {
	i := 0
	for list != nil {
		list = list.Next
		i++
	}
	return i
}

type ListNode struct {
	Val  int
	Next *ListNode
}
