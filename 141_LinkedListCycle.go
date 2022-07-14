/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    m := make(map[int]struct{})
    if head == nil {
        return false
    }
    p := head
    for {
        if _, exists := m[p.Val]; exists {
            return true
        }    
        m[p.Val] = struct{}{}
        if p.Next == nil {
            break
        }
        p = p.Next
    }
    
    return false
}
