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
	binary_tree.Mirror(tree)
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
	//合并
	fmt.Println(binary_tree.LevelPrint(binary_tree.Merge(tree, tree)))
	//对称
	fmt.Println(binary_tree.IsSymmetric(tree))
	//平衡
	fmt.Println(binary_tree.IsBalance(tree))
	//第K大的元素
	fmt.Println(binary_tree.KthLargest(tree, 1))
	//给定两个节点的公共祖先
	fmt.Println(binary_tree.LowestCommonAncestor(tree, 2, 8))
	//单值二叉树
	fmt.Println(binary_tree.IsSameValTree(tree))
	//二叉树的直径
	fmt.Println(binary_tree.Diameter(tree))
	//二叉树剪枝
	fmt.Println(binary_tree.LevelPrint(binary_tree.Prune(tree)))
}
