/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	maxArea := 0
	s := &Stack{
		data:   make([]int, 0),
		length: 0,
	}

	s.push(-1)

	for idx := 0; idx < len(heights); idx++ {
		for s.peek() != -1 && heights[s.peek()] >= heights[idx] {
			maxArea = max(maxArea, heights[s.pop()]*(idx-s.peek()-1))
		}
		s.push(idx)
	}

	for s.peek() != -1 {
		maxArea = max(maxArea, heights[s.pop()]*(len(heights)-s.peek()-1))
	}
	return maxArea
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

type Stack struct {
	data   []int
	length int
}

func (s *Stack) push(v int) {
	s.data = append(s.data, v)
	s.length = len(s.data)
}

func (s *Stack) peek() int {
	if s.length == 0 {
		return -2
	}
	return s.data[s.length-1]
}

func (s *Stack) pop() int {
	if s.length == 0 {
		return -2
	}
	ret := s.data[s.length-1]
	if s.length > 1 {
		s.data = s.data[:s.length-1]
	} else {
		s.data = make([]int, 0)
	}
	s.length = s.length - 1
	return ret
}

// @lc code=end
