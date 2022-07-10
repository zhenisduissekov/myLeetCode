/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
   if head == nil {
        return nil
    }
    p := head
    m := make(map[*ListNode]struct{})

    for {
        if _, exists := m[p]; exists {
            return p
        }
        m[p] = struct{}{}
        if p.Next == nil {
            break
        }
        p = p.Next
    }
    
    
    return nil
}
