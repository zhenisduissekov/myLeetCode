func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    type keyStr struct{
        a int
        b int
        c int
        d int
    }
    result := make([][]int,0)
    seen := make(map[keyStr]bool)
    for i:=0;i<n-3;i++ {
       for j:=i+1;j<n-2;j++ {
        k:=j+1
        m:=n-1
        for k<m {
            switch {
            case nums[i]+nums[j]+nums[k]+nums[m] == target:
                key := keyStr{a:nums[i], b:nums[j], c:nums[k], d: nums[m]}
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


// без хеш мапы

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    n := len(nums)

    result := make([][]int,0)
    
    for i:=0;i<n-3;i++ {
        if i>0 && nums[i] == nums[i-1] {
            continue
        }
       for j:=i+1;j<n-2;j++ {
        if j>i+1 && nums[j]==nums[j-1] {
            continue
        }
        k:=j+1
        m:=n-1
        for k<m {
            switch {
            case nums[i]+nums[j]+nums[k]+nums[m] == target:
                result = append(result, []int{nums[i], nums[j], nums[k], nums[m]})

                for k<m && nums[k]==nums[k+1] {
                    k++
                }
                k++
                for k<m && nums[m]==nums[m-1] {
                    m--
                }
               m--

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
