package binary_tree

import "math"

/*
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
示例 :
给定二叉树
          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
var ans int

func Diameter(t *TreeNode) int {
	ans = 1
	Depth(t)
	return ans - 1
}

func Depth(t *TreeNode) int {
	if t == nil {
		return 0
	}
	left := Depth(t.Left)
	right := Depth(t.Right)
	ans = int(math.Max(float64(ans), float64(left+right+1)))
	return int(math.Max(float64(left), float64(right))) + 1
}
