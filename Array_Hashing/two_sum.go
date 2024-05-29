package main

import "fmt"

func main() {
	nums := []int{1, 3, 3}
	target := 6
	result := twoSum(nums, target)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	MMAP := make(map[int]int)
	for index2, num := range nums {
		compiment := target - num
		if idex1, ok := MMAP[compiment]; ok {
			return []int{idex1, index2}
		}
		MMAP[num] = index2
		fmt.Println(MMAP)
	}
	return []int{0, 0}
}
