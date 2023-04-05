/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
   if root == nil {
       return 0
   }
    length := 0

    var dff func(root *TreeNode) int

    dff = func(root *TreeNode) int {
        if root == nil {
            return -1
        }
        left := dff(root.Left)
        right := dff(root.Right)

        length = max(length, left+right+2)
        return max(left, right)+1
    }
    dff(root)
    return length
}

func max(a, b int) int {
    if a>b {
        return a
    }
    return b
}
