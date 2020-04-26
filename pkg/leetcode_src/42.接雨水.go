/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	ret := 0

	for idx := 1; idx < len(height)-1; idx++ {
		value := height[idx]
		maxLeft, maxRight := 0, 0
		for i := idx; i >= 0; i-- {
			maxLeft = max(maxLeft, height[i])
		}

		for i := idx; i < len(height); i++ {
			maxRight = max(maxRight, height[i])
		}

		if value == maxLeft || value == maxRight {
			continue
		}

		ret += (min(maxLeft, maxRight) - height[idx])
	}

	return ret
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

// @lc code=end
