/*
* @Description: Given an array of integers, return indices of the two numbers such that they add up to a specific target.
*               给定一个整数数组, 将两个数字相加得到特定的目标, 返回这两个数字的索引。
* @Example: Given nums = [2, 7, 11, 15], target = 9,
*           Because nums[0] + nums[1] = 2 + 7 = 9,
*           return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	lenNums := len(nums)
	for i := 0; i < lenNums; i++ {
		for j := i + 1; j < lenNums; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}
