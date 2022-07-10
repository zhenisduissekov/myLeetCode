/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    m := make(map[*ListNode]struct{})

    for head !=nil && head.Next != nil {
        if _, exists := m[head]; exists {
            return head
        }
        m[head] = struct{}{}
        head = head.Next
    }
    
    return nil
}
