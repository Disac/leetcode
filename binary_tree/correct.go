package binary_tree

/*
你有一棵二叉树，这棵二叉树有个小问题，其中有且只有一个无效节点，它的右子节点错误地指向了与其在同一层且在其右侧的一个其他节点。
给定一棵这样的问题二叉树的根节点 root ，将该无效节点及其所有子节点移除（除被错误指向的节点外），然后返回新二叉树的根结点。
自定义测试用例：
测试用例的输入由三行组成：
TreeNode root
int fromNode （在 correctBinaryTree 中不可见）
int toNode （在 correctBinaryTree 中不可见）
当以 root 为根的二叉树被解析后，值为 fromNode 的节点 TreeNode 将其右子节点指向值为 toNode 的节点 TreeNode 。然后， root 传入 correctBinaryTree 的参数中。
示例 1:
输入: root = [1,2,3], fromNode = 2, toNode = 3
输出: [1,null,3]
解释: 值为 2 的节点是无效的，所以移除之。
示例 2:
输入: root = [8,3,1,7,null,9,4,2,null,null,null,5,6], fromNode = 7, toNode = 4
输出: [8,3,1,null,null,9,4,null,null,5,6]
解释: 值为 7 的节点是无效的，所以移除这个节点及其子节点 2。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/correct-a-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
var father map[int]*TreeNode

func Correct(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	if t.Right != nil {
		if father[t.Right.Val] == nil {
			father[t.Right.Val] = t
		} else {
			return nil
		}
		t.Right = Correct(t.Right)
	}
	if t.Left != nil {
		if father[t.Left.Val] == nil {
			father[t.Left.Val] = t
		} else {
			return nil
		}
		t.Left = Correct(t.Left)
	}
	return t
}
