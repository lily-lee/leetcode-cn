package main

/**
面试题 02.04. 分割链表

编写程序以 x 为基准分割链表，使得所有小于 x 的节点排在大于或等于 x 的节点之前。如果链表中包含 x，x 只需出现在小于 x 的元素之后(如下所示)。分割元素 x 只需处于“右半部分”即可，其不需要被置于左右两部分之间。

示例:

输入: head = 3->5->8->5->10->2->1, x = 5
输出: 3->1->2->10->5->5->8

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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

func partition(head *ListNode, x int) *ListNode {
	nodeLeft := &ListNode{}
	nodeRight := &ListNode{}
	for {
		if head == nil {
			break
		}
		if head.Val < x {
			nodeLeft.getTail().Next = &ListNode{Val: head.Val}
		} else {
			nodeRight.getTail().Next = &ListNode{Val: head.Val}
		}
		head = head.Next
	}
	nodeLeft.getTail().Next = nodeRight
	return nodeLeft
}

func (l *ListNode) getTail() *ListNode {
	if l.Next == nil {
		return l
	}
	return l.Next.getTail()
}
