package main

import (
	"fmt"
	"github.com/disac/leetcode/binary_tree"
)

func main() {
	var tree = &binary_tree.TreeNode{
		Val: 5,
		Left: &binary_tree.TreeNode{
			Val:   2,
			Left:  &binary_tree.TreeNode{Val: 1},
			Right: &binary_tree.TreeNode{Val: 3},
		},
		Right: &binary_tree.TreeNode{
			Val:   7,
			Left:  &binary_tree.TreeNode{Val: 6},
			Right: &binary_tree.TreeNode{Val: 8},
		},
	}
	fmt.Println(binary_tree.LevelPrint(tree))
	mirror := binary_tree.Mirror(tree)
	fmt.Println(binary_tree.LevelPrint(mirror))
	var pre = make([]int, 0)
	binary_tree.PreTraversal(tree, &pre)
	var in = make([]int, 0)
	binary_tree.InTraversal(tree, &in)
	var post = make([]int, 0)
	binary_tree.PostTraversal(tree, &post)
	fmt.Println(pre, in, post)
}
