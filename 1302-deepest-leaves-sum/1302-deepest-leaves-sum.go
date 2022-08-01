package deepestleavessum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var q []*TreeNode
	sum := 0

	q = append(q, root)

	for len(q) > 0 {
		sum = 0
		for currentLevelSize := len(q); currentLevelSize > 0; currentLevelSize-- {
			node := q[0]
			q = q[1:]

			if node.Left == nil && node.Right == nil {
				sum += node.Val
			} else {
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
		}
	}
	return sum
}
