package binary_tree

/*
如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
只有给定的树是单值二叉树时，才返回 true；否则返回 false。
示例 1：
输入：[1,1,1,1,1,null,1]
输出：true
示例 2：
输入：[2,2,2,5,2]
输出：false
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/univalued-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func IsSameValTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if (root.Left == nil || (root.Left.Val == root.Val && IsSameValTree(root.Left))) && (root.Right == nil || (root.Right.Val == root.Val && IsSameValTree(root.Right))) {
		return true
	}
	return false
}
