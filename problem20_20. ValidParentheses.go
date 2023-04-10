func isValid(s string) bool {
	var openBrackets []rune
    m := map[rune]rune{')': '(',']': '[', '}': '{'}
	for _, br := range s {
        if strings.Contains("({[", string(br)) {
            openBrackets = append(openBrackets, rune(br))
        } else if strings.Contains(")}]", string(br)) {
            if len(openBrackets) == 0 || m[rune(br)] != openBrackets[len(openBrackets)-1]{
                return false
            } else if openBrackets[len(openBrackets)-1] == m[rune(br)] {
                openBrackets = openBrackets[:len(openBrackets)-1]
            }
        }
	}
	return len(openBrackets) == 0
}


//solution 2 with Stack
func isValid(s string) bool {
    closing := map[rune]rune{')': '(',']': '[', '}': '{'}
    st := New()
	for _, elem := range(s) {
        if val, ok := closing[elem]; ok {
            top := st.Pop()
            if top == nil || top != val{
                return false
            }
        } else {
            st.Push(elem)
        }
    }
    
    return st.len == 0
}


type stack struct {
    top *node
    len int
}

type node struct {
    val interface{}
    prev *node
}
func New() *stack{
    n := &node{nil, nil}
    return &stack{
        top: n,
        len: 0,
    }
}

func (s *stack) Push(val interface{}) {
    n := &node{val: val, prev: s.top }
    s.len ++
    s.top = n
}

func (s *stack) Pop() interface{} {
    if s.len == 0 {
        return nil
    }
    n := s.top
    s.len --
    s.top = n.prev
    return n.val
}


