/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	max := nums[0]
	for idx := 1; idx < len(nums); idx++ {
		sum := nums[idx] + nums[idx-1]
		if sum > nums[idx] && sum > 0 {
			nums[idx] = sum
		}

		if max < nums[idx] {
			max = nums[idx]
		}
	}

	return max
}

// @lc code=end
