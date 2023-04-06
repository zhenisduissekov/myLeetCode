/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {

    if root == nil {
        return true
    }

    left := maxDepth(root.Left)
    right := maxDepth(root.Right)

    if abs(left - right) > 1 {
        return false
    }

    return isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := maxDepth(root.Left)
    right := maxDepth(root.Right)

    return max(left, right)+1
}


func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}
