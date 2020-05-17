/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
    w := len(s)

    if w < 2 {
        return s
    }

    left,right, maxL, maxR := 0,0,0,0

    for right < w {
        for right+1 < w && s[left] == s[right+1] {
            right++
        }

        for left-1 >= 0 && right+1 < w && s[left-1] == s[right+1] {
            left--
            right++
        }

        if right - left > maxR - maxL {
            maxL,maxR = left,right
        }
        left = (right + left)>>1 + 1
        right = left
    }
    return s[maxL:maxR+1]
}
// @lc code=end

