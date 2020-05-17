package main

// func majorityElement(nums []int) int {
// 	insertSort(nums, 0, len(nums))

// 	target := judgeExist(nums)
// 	return target
// }

// func judgeExist(nums []int) int {
// 	arrayLen := len(nums)

// 	if arrayLen == 0 {
// 		return -1
// 	}

// 	if arrayLen == 1 {
// 		return nums[0]
// 	}

// 	if arrayLen == 2 {
// 		if nums[0] != nums[1] {
// 			return -1
// 		}
// 		return nums[0]
// 	}

// 	targetIdx := arrayLen >> 1
// 	target := nums[targetIdx]

// 	if target != nums[targetIdx-1] && target != nums[targetIdx+1] {
// 		return -1
// 	}

// 	idx := findRight(nums, target, targetIdx, arrayLen)

// 	if nums[idx-targetIdx] != target {
// 		return -1
// 	}
// 	return target
// }

// func findRight(nums []int, target, s, e int) int {
// 	for {
// 		midIdx := (e + s) / 2
// 		if nums[midIdx] == target {
// 			s = midIdx
// 			if e-s == 1 {
// 				return midIdx
// 			}
// 		} else {
// 			e = midIdx
// 		}
// 	}

// }

// func insertSort(nums []int, s, e int) {
// 	if e-s <= 1 {
// 		return
// 	}
// 	midValue := nums[s]
// 	i := s     // left
// 	j := e - 1 // right
// 	for i < j {
// 		for j > i && nums[j] >= midValue {
// 			j--
// 			continue
// 		}
// 		if i < j {
// 			nums[i] = nums[j]
// 			i++
// 		}

// 		for i < j && nums[i] < midValue {
// 			i++
// 			continue
// 		}
// 		if i < j {
// 			nums[j] = nums[i]
// 			j--
// 		}
// 	}
// 	nums[i] = midValue
// 	insertSort(nums, s, i)
// 	insertSort(nums, i+1, e)
// }

// // 方法二
// func majorityElement(nums []int) int {
// 	arrayLen := len(nums)
// 	if arrayLen == 0 {
// 		return -1
// 	}

// 	if arrayLen == 1 {
// 		return nums[0]
// 	}

// 	record := make(map[int]int)
// 	max := 0
// 	ret := -1
// 	for _, value := range nums {
// 		_, ok := record[value]
// 		if ok {
// 			record[value]++
// 		} else {
// 			record[value] = 1
// 		}

// 		if max < record[value] {
// 			max = record[value]
// 			ret = value
// 		}
// 	}

// 	if max <= (arrayLen >> 1) {
// 		return -1
// 	}

// 	return ret
// }
