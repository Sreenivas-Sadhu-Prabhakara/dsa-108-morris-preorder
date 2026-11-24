package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// MorrisPreorder returns root-left-right order in O(1) space using the same
// predecessor threading as Morris inorder — but the visit happens when the
// thread is CREATED (first arrival), not when it is removed.
func MorrisPreorder(root *TreeNode) []int {
	out := []int{}
	cur := root
	for cur != nil {
		if cur.Left == nil {
			out = append(out, cur.Val)
			cur = cur.Right
			continue
		}
		pred := cur.Left
		for pred.Right != nil && pred.Right != cur {
			pred = pred.Right
		}
		if pred.Right == nil {
			out = append(out, cur.Val) // visit on first arrival: preorder
			pred.Right = cur
			cur = cur.Left
		} else {
			pred.Right = nil
			cur = cur.Right
		}
	}
	return out
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	root := tn(4, tn(2, tn(1, nil, nil), tn(3, nil, nil)), tn(6, tn(5, nil, nil), tn(7, nil, nil)))
	fmt.Println(MorrisPreorder(root))
}
