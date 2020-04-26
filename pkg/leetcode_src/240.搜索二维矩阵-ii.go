/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for {
		targetValue := matrix[i][j]
		if targetValue == target {
			return true
		}
		if targetValue < target {
			i++
		}
		if targetValue > target {
			j--
		}

		if i == len(matrix) || j < 0 {
			return false
		}
	}
}

// @lc code=end
