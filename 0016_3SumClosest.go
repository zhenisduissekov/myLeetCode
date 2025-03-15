

func threeSumClosest(nums []int, target int) int {
    if len(nums) < 3 {
        return 0
    }
    sort.Ints(nums)
    best := abs(nums[0]+nums[1]+nums[2] - target)
    bestV := nums[0]+nums[1]+nums[2]
    for i:=0;i<len(nums);i++ {
        j:=i+1
        k:=len(nums)-1
        for j<k {
            if best > abs(nums[i]+nums[j]+nums[k]-target) {
                best = abs(nums[i]+nums[j]+nums[k]-target)
                bestV = nums[i]+nums[j]+nums[k]
            }
            switch {
            case nums[i]+nums[j]+nums[k] < target:
                j++
            default:
                k--
            }
        }
    }

    return bestV
}


func abs(val int) int {
    if val > 0 {
        return val
    }
    return -val
}
