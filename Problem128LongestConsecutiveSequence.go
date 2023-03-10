func longestConsecutive(nums []int) int {
	sort.Ints(nums)

	counter := 0
	maxCounter := 0

    if len(nums) == 1 || len(nums) == 0 {
        return len(nums)
    }

    for i := 0; i < len(nums)-1; i++ {
        displ := nums[i+1]-nums[i]
        if displ==0 || displ==1 {
            counter += displ
            if maxCounter < counter {
                maxCounter = counter
            } 
        } else {
            counter = 0
        }
    }
    
    return maxCounter + 1
}
