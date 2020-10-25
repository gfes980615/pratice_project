package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Int struct {
	int
}

func (t *TreeNode) addValue(val *Int) {
	if val == nil {
		return
	}
	t.Val = val.int
}

func (t *TreeNode) addLeftNode(node *TreeNode) {
	t.Left = node
}

func (t *TreeNode) addRightNode(node *TreeNode) {
	t.Right = node
}

func createNode(val *Int) *TreeNode {
	node := &TreeNode{}
	node.addValue(val)
	return node
}

func createTree(values []*Int) *TreeNode {
	nodes := []*TreeNode{}
	for _, value := range values {
		nodes = append(nodes, createNode(value))
	}
	j := 0
	for i := 1; i < len(nodes); i++ {
		root := nodes[j]
		if i%2 != 0 {
			root.addLeftNode(nodes[i])
		} else {
			root.addRightNode(nodes[i])
			j++
		}
	}

	return nodes[0]
}

func main() {
	tree := createTree([]*Int{&Int{1}, &Int{7}, &Int{0}, &Int{7}, &Int{-8}, nil, nil, &Int{10}, &Int{800}})
	getNodes(tree, 0)
	max := -1000000
	d := 0
	for depth, nodes := range nodeMap {
		sum := 0
		for _, node := range nodes {
			sum += node
		}
		if sum > max {
			max = sum
			d = depth
		}
	}
	fmt.Println(d)
}

var (
	nodeMap = make(map[int][]int)
)

func getNodes(tree *TreeNode, depth int) {
	if tree == nil {
		return
	}
	depth++
	nodeMap[depth] = append(nodeMap[depth], tree.Val)
	getNodes(tree.Left, depth)
	getNodes(tree.Right, depth)

}
