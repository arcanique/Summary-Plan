/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	if x < 4 {
		return 1
	}

	left := 2
	right := (x >> 1)

	if left == right {
		return left
	}

	mid := (right >> 1)
	for {

		value := mid * mid
		if value == x {
			return mid
		}

		if mid*mid > x {
			right = mid
		} else {
			left = mid
		}

		if right-left == 1 || right == left {
			return left
		}
		mid = (right + left) >> 1
	}
}

// @lc code=end
