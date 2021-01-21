package partition

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
	var nodeLeft, nodeRight *ListNode
	for head != nil {
		if head.Val < x {
			nodeLeft = build(nodeLeft, head.Val)
		} else {
			nodeRight = build(nodeRight, head.Val)
		}
		head = head.Next
	}
	if nodeLeft == nil {
		return nodeRight
	}
	getTail(nodeLeft).Next = nodeRight
	return nodeLeft
}

func build(l *ListNode, val int) *ListNode {
	if l == nil {
		l = &ListNode{Val: val}
	} else {
		getTail(l).Next = &ListNode{Val: val}
	}
	return l
}

func getTail(l *ListNode) *ListNode {
	if l.Next == nil {
		return l
	}
	return getTail(l.Next)
}
