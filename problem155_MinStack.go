type MinStack struct {
    stack []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    vals := this.stack
    this.stack = append(vals, val)
}


func (this *MinStack) Pop()  {
    vals := this.stack
    this.stack = vals[:len(vals)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    min := this.stack[0]
    for i:=1;i<len(this.stack);i++ {
        if min > this.stack[i] {
            min = this.stack[i]
        }
    }
    return min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
