/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{Val:0}
    cur := head
    carry := 0
    for l1 != nil || l2 != nil || carry != 0 {
        v1, v2 := 0, 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
        
        sum := v1 + v2 + carry
        carry = sum / 10
        cur.Next = &ListNode{Val: sum % 10}
        cur = cur.Next
    }
    return head.Next
}


or

**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := addNode(l1, l2, 0)
    return head
}

func addNode(l1, l2 *ListNode, keep int) *ListNode {
    p := &ListNode{}
    switch {
        case l1 == nil && l2 == nil && keep == 1:
            p.Val = 1
        case l1 == nil && l2 == nil && keep == 0:
            p = nil
        case l1 != nil && l2 != nil:
            p.Val = (l1.Val + l2.Val + keep) % 10
            p.Next = addNode(l1.Next, l2.Next, (l1.Val + l2.Val + keep)/10)
        case l1 != nil && l2 == nil:
            p.Val = (l1.Val + keep) % 10
            p.Next = addNode(l1.Next, nil, (l1.Val + keep)/10)
        case l1 == nil && l2 != nil:
            p.Val = (l2.Val + keep) % 10
            p.Next = addNode(nil, l2.Next, (l2.Val + keep)/10)
    }
    return p
}
