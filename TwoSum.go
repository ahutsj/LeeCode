package main

func twoSum(nums []int,target int) []int {
	for i := 0;i <= len(nums)-1;i ++ {
		for j := i+1;j <= len(nums)-1;j ++ {
				if nums[i] + nums[j] == target {
					return []int{i,j}
				}else {
					continue
				}
		}
	}
	return []int{}
}
