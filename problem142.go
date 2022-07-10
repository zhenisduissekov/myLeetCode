/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    p := head
    m := make(map[*ListNode]struct{})

    for p !=nil && p.Next != nil {
        if _, exists := m[p]; exists {
            return p
        }
        m[p] = struct{}{}
        p = p.Next
    }
    
    
    return nil
}
