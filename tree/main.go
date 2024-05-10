package main

import "fmt"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

var root *TreeNode = nil

func insert(r *TreeNode, node *TreeNode) *TreeNode {
	if r == nil {
		return node
	}
	if r.val < node.val {
		r.right = insert(r.right, node)
	} else {
		r.left = insert(r.left, node)
	}
	return r
}

func (t *TreeNode) Add(val int) bool {
	temp := new(TreeNode)
	temp.val = val
	root = insert(root, temp)
	return true
}

func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}
	t.left.PrintInorder()
	fmt.Println(t.val)
	t.right.PrintInorder()
}

func main() {
	fmt.Println(root)
	root.Add(2)
	fmt.Println(root)
	root.Add(3)
	root.Add(1)
	root.Add(4)
	root.Add(6)
	fmt.Println(root)
	root.PrintInorder()

}
