/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}
	for indexX := 0; indexX < len(nums); indexX++ {
		vX := nums[indexX]
		// if vX >= target {
		// 	continue
		// }

		for indexY := indexX + 1; indexY < len(nums); indexY++ {
			if vX + nums[indexY] == target {
				return []int{indexX, indexY}
			}
		}
	}

	return nil
}

// @lc code=end
