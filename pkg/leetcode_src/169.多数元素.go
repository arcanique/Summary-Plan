/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	count := 0
	ret := 0
	for _, value := range nums {
		if count == 0 {
			ret = value
			count++
		} else if value == ret {
			count++
		} else {
			count--
		}
	}

	if count >= 1 {
		return ret
	}

	return -1

}

// @lc code=end
