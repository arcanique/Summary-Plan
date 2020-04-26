/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k {
		return nil
	}

	if len(nums) == k {
		return nums
	}

	maxWinSize := =0 
	for i := 0,; i < k; i++ {
		maxWinSize += nums[i]
	}

	lastIdx := 0
	lastSize := maxWinSize
	max 
	for i := 1; i < len(nums); i++ {
		maxWinSize = max(maxWinSize, lastSize - nums[lastIdx] + )
	}

}
// @lc code=end

