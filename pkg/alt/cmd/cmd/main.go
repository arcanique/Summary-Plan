package main

import (
	"sort"
	"fmt"
	


)

const INT_MAX = ^int(^uint(0) >> 1)

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
func main() {
	maxArea := INT_MAX
	input := []int{1, 9, 6, 6, 5, 2, 6, 6, 6, 6, 3, 4, 8, 9, 1, 2, 6, 6, 6, 6, 6, 3, 9, 6, 6, 1, 5, 9, 6, 6, 6, 6, 6, 6, 6, 6}
	input2 := []int{2, 1}
	input3 := []int{-1,0,1,2,-1,-4}
	ret := mergeSort2(input3)
	fmt.Println(input2, ret, maxArea, input3)
	sort.Ints(input)
	area := largestRectangleArea(input)

	matrix := [][]int{{1, 2, 3}, {1, 4, 3}}

	a := []int{2, 3, 4}
	b := [2]int{1, 5}
	fmt.Printf("%+x, %x, %d, %+x", a, b, area, matrix)
	fn1(a)
	fn2(b)
	fmt.Println(1, a, b, len(matrix), len(matrix[0]))
}

func fn1(a []int) {
	a[0] = 1
}

func fn2(b [2]int) {
	b[0] = 9
}

func insertSort(nums []int, s, e int) {
	if e-s <= 1 {
		return
	}
	midValue := nums[s]
	i := s     // left
	j := e - 1 // right
	for i < j {
		for j > i && nums[j] >= midValue {
			j--
			continue
		}
		if i < j {
			nums[i] = nums[j]
			i++
		}

		for i < j && nums[i] < midValue {
			i++
			continue
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}
	nums[i] = midValue
	insertSort(nums, s, i)
	insertSort(nums, i+1, e)
}

func mergeSort(nums []int) []int {
	arrayLen := len(nums)
	if arrayLen < 2 {
		return nums
	}

	midIdx := arrayLen >> 1
	left := nums[:midIdx]
	right := nums[midIdx:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0
	if len(left) == 3 && len(right) == 3 {
		fmt.Println(1)
	}
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func mergeSort2(nums []int) []int {
    w := len(nums)
    if w < 2 {
        return nums
    }
    mid := w>>1
    left := nums[:mid]
    right := nums[mid:]

    return merge2(mergeSort(left),mergeSort(right))
}

func merge2(left,right []int) []int {
    i, j := 0, 0
    lw,rw := len(left), len(right)
	ret := []int{}
	if len(left) == 3 && len(right) == 3 {
		fmt.Println(1)
	}
    for i < lw && j < rw {
        if left[i] <= right[j] {
            ret = append(ret, left[i])
            i++
        } else {
            ret = append(ret,right[j])
            j++
        }
    }

    for ;i<lw;i++ {
        ret = append(ret, left[i])
    }

    for ;j<rw; j++ {
        ret = append(ret, right[j])
    }
    fmt.Println(left, right, ret)
    return ret
}