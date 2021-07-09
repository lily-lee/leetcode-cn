package xam1wr

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var newHead *Node
	if head == nil {
		newHead = nil
		return newHead
	}
	headMap := map[*Node]int{}
	newArr := make([]*Node, 0)
	getNodeArr(head, headMap, &newArr, 0)
	newHead = newArr[0]
	copyList(head, newHead, headMap, newArr, 0)
	return newHead
}

func copyList(head, newHead *Node, headMap map[*Node]int, newArr []*Node, i int) {
	if head == nil {
		return
	}
	newHead = newArr[i]
	if i < len(newArr)-1 {
		newHead.Next = newArr[i+1]
	}
	if index, ok := headMap[head.Random]; ok {
		newHead.Random = newArr[index]
	}
	copyList(head.Next, newHead.Next, headMap, newArr, i+1)
}

func getNodeArr(head *Node, headMap map[*Node]int, newArr *[]*Node, i int) {
	if head == nil {
		return
	}
	headMap[head] = i
	item := &Node{Val: head.Val}
	*newArr = append(*newArr, item)
	getNodeArr(head.Next, headMap, newArr, i+1)
}

func printNode(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%+v\n", node)
	printNode(node.Next)
}

func buildArr(arr [][]int) [][]*int {
	arrNew := make([][]*int, len(arr))
	for i, item := range arr {
		arrNew[i] = make([]*int, 2)
		arrNew[i][0] = &item[0]
		if item[1] < 0 {
			arrNew[i][1] = nil
		} else {
			arrNew[i][1] = &item[1]
		}
	}
	return arrNew
}

func buildList(arr [][]*int) *Node {
	nodeSlice := make([]*Node, len(arr))
	for i, item := range arr {
		nodeSlice[i] = &Node{Val: *item[0]}
	}
	head := nodeSlice[0]
	buildNodeList(head, nodeSlice, 0, arr)
	return head
}

func buildNodeList(head *Node, nodes []*Node, i int, arr [][]*int) {
	if i > len(nodes)-1 {
		return
	}
	if index := arr[i][1]; index != nil {
		head.Random = nodes[*index]
	}
	if i < len(nodes)-1 {
		head.Next = nodes[i+1]
	}

	buildNodeList(head.Next, nodes, i+1, arr)
}
