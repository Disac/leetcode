package binary_tree

//前序遍历
func PreTraversal(t *TreeNode, res *[]int) {
	if t == nil {
		return
	}
	*res = append(*res, t.Val)
	PreTraversal(t.Left, res)
	PreTraversal(t.Right, res)
}

//中序遍历
func InTraversal(t *TreeNode, res *[]int) {
	if t == nil {
		return
	}
	InTraversal(t.Left, res)
	*res = append(*res, t.Val)
	InTraversal(t.Right, res)
}

//后序遍历
func PostTraversal(t *TreeNode, res *[]int) {
	if t == nil {
		return
	}
	PostTraversal(t.Left, res)
	PostTraversal(t.Right, res)
	*res = append(*res, t.Val)
}
