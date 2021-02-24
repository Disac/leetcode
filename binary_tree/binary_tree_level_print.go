package binary_tree

/*
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func LevelPrint(t *TreeNode) [][]int {
	if t == nil {
		return [][]int{}
	}
	var res = make([][]int, 0)
	var queue = []*TreeNode{t}
	for len(queue) != 0 {
		tmp := make([]int, 0)
		for _, v := range queue {
			queue = queue[1:]
			tmp = append(tmp, v.Val)
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}
