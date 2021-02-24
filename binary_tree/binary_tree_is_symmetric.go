package binary_tree

/*
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func IsSymmetric(t *TreeNode) bool {
	return Symmetric(t.Left, t.Right)
}

func Symmetric(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil && t1.Val == t2.Val {
		return Symmetric(t1.Left, t2.Right) && Symmetric(t1.Right, t2.Left)
	}
	return false
}
