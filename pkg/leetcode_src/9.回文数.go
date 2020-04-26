/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	xList := make([]int, 0)
	a := x
	for {
		xList = append(xList, a%10)
		a = a / 10
		if 0 == a {
			break
		}
	}

	y := 0

	for _, value := range xList {
		y = y*10 + value
	}

	if y == x {
		return true
	}

	return false
}

// @lc code=end
