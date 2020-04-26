/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	for idx := 0; idx < len(matrix); idx++ {
		maxW, maxH := 0, 0
		for idy := 0; idy < len(matrix[0]); idy++ {
			if matrix[idx][idy] == '0' {
				continue
			}

			for i := idy; i < len(matrix[0]); i++ {
				
			}
		}
	}
}

// @lc code=end
