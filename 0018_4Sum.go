

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    result := make([][]int,0)
    seen := make(map[string]bool)
    for i:=0;i<n-3;i++ {
       for j:=i+1;j<n-2;j++ {
        k:=j+1
        m:=n-1
        for k<m {
            switch {
            case nums[i]+nums[j]+nums[k]+nums[m] == target:
                key := fmt.Sprintf("%d%d%d%d", nums[i], nums[j], nums[k], nums[m])
                if !seen[key] {
                    seen[key]=true
                    result = append(result, []int{nums[i], nums[j], nums[k], nums[m]})
                }
                k++
            case nums[i]+nums[j]+nums[k]+nums[m] > target:
                m--
            default:
                k++
            }
        }
       }    
    }
    return result
}
