package binary_tree

/*
输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
例如，给出
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// preorder : 根结点，左子树遍历，右子树遍历
// inorder : 左子树遍历，根结点，右子树遍历
func BuildTree(preOrder, inOrder []int) *TreeNode {
	if len(preOrder) == 0 || len(inOrder) == 0 {
		return nil
	}
	var root = &TreeNode{Val: preOrder[0]}
	var i int
	for i := 0; i < len(inOrder); i++ {
		if root.Val == inOrder[i] {
			break
		}
	}
	root.Left = BuildTree(preOrder[1:len(inOrder[:i])+1], inOrder[:i])
	root.Right = BuildTree(preOrder[len(inOrder[:i])+1:], inOrder[i+1:])
	return root
}
