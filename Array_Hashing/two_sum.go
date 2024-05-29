package main

import "fmt"

func main() {
	nums := []int{15, 7, 11, 2}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result) // Output: [0, 1]
}

func twoSum(nums []int, target int) []int {
	nMap := make(map[int]int)
	for i, num := range nums {
		com := target - num
		// fmt.Println(com)
		// fmt.Println(nMap)
		if index, ok := nMap[com]; ok {
			return []int{index, i}
		}
		nMap[num] = i
	}
	return []int{}
}
