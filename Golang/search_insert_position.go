package problemsolving

func searchInsert(nums []int, target int) int {
    var res int
    for i := range nums {
        switch target {
            case nums[i]:
            res = i 
            return res
        }
        if target < nums[i] {
            res = i
            return res
        }
        res = len(nums)
    }
    
    return res
}