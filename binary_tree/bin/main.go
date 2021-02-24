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
	//层级遍历
	fmt.Println(binary_tree.LevelPrint(tree))
	//镜像后层级遍历
	fmt.Println(binary_tree.LevelPrint(binary_tree.Mirror(tree)))
	//前序遍历
	var pre = make([]int, 0)
	binary_tree.PreTraversal(tree, &pre)
	//中序遍历
	var in = make([]int, 0)
	binary_tree.InTraversal(tree, &in)
	//后序遍历
	var post = make([]int, 0)
	binary_tree.PostTraversal(tree, &post)
	fmt.Println(pre, in, post)
	//深度
	fmt.Println(binary_tree.MaxDepth(tree))
}
