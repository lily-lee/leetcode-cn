package partition

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}
	fmt.Println(partition(node, 3))
	fmt.Println(partition(&ListNode{}, 0))
	fmt.Println(partition(&ListNode{Val: 1}, 0))
	fmt.Println(partition(&ListNode{Val: 1, Next: &ListNode{Val: 1}}, 0))
}
