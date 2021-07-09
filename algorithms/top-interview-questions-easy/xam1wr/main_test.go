package xam1wr

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("[[7,null],[13,0],[11,4],[10,2],[1,0]]", func(t *testing.T) {
		head := &Node{
			Val:  7,
			Next: &Node{Val: 13, Next: &Node{Val: 11, Next: &Node{Val: 10, Next: &Node{Val: 1, Random: &Node{}}, Random: &Node{Val: 2}}, Random: &Node{Val: 4}}, Random: &Node{}},
		}
		newNode := copyRandomList(head)
		printNode(head)
		fmt.Println()
		printNode(newNode)
		fmt.Println()
		fmt.Println(head)
		fmt.Println(newNode)
	})
	t.Run("buidList", func(t *testing.T) {
		arr := buildArr([][]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})
		nodes := buildList(arr)
		printNode(nodes)

		newHead := copyRandomList(nodes)
		fmt.Println("--------")
		printNode(newHead)
	})

	t.Run("[[3,null], [3,0], [3,null]]", func(t *testing.T) {
		arr := buildArr([][]int{{3, -1}, {3, 0}, {3, -1}})
		nodes := buildList(arr)
		printNode(nodes)

		newHead := copyRandomList(nodes)
		fmt.Println("--------")
		printNode(newHead)
	})
}
