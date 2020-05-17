/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	ret := 0
	dd := int64(dividend)
	ds := int64(divisor)

	if dd < 0 {
		dd = -dd
		if dd > 0x7fffffff {
			dd -= 1
		}
	}
	if ds < 0 {
		ds = -ds
		if ds > 0x7fffffff {
			ds -= 1
		}
	}

	for {
		if dd < ds {
			break
		}
		dd -= ds
		ret++
	}

	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		return ret
	}

	return 0 - ret
}

// @lc code=end
