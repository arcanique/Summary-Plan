/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
import (
	"sort"
)
func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return nil
	}
	sort.Sort(SliceArray(intervals))
	//fmt.Println(intervals)
	//insertSort(intervals)

	j := 1
	temp := intervals[0]
	ret := [][]int{}
	for j < len(intervals) {
		if temp[1] >= intervals[j][0] {
			maxR := intervals[j][1]
			if maxR < temp[1] {
				maxR = temp[1]
			}
			temp = []int{temp[0], maxR}
		} else {
			ret = append(ret, temp)
			temp = intervals[j]
		}
		j++
	}
	ret = append(ret, temp)
	return ret
}

func insertSort(nums [][]int) {
	w := len(nums)
	if w < 2 {
		return
	}

	i, j := 0, w-1
	target := nums[0]

	for i < j {
		for i < j && target[0] <= nums[j][0] {
			j--
		}

		if i < j {
			nums[i] = nums[j]
			i++
		}
		for i < j && nums[i][0] <= target[0] {
			i++
		}
		if i < j {
			nums[j] = nums[i]
			j--
		}
	}

	nums[i] = target

	insertSort(nums[:i])
	insertSort(nums[i+1:])
}

type SliceArray [][]int

func (s SliceArray)Less(i,j int) bool{
	return s[i][0] < s[j][0]
}

func (s SliceArray)Swap(i,j int){
	s[i],s[j] = s[j], s[i]
}

func (s SliceArray)Len() int {
	return len(s)
}
// @lc code=end

