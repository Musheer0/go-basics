package main

import "fmt"

func main() {
	// nums_2d := [2][2]int{{1, 2}, {4, 5}}
	// unintialized size
	// var nums []int
	// fmt.Print(nums==nil)
	// var  i int =0
	// fmt.Print(i)
	// var nums= make([]int,2)
	// fmt.Print(nums)
	//copy 
	// var nums = make([]int,0,5)
	// nums = append(nums,6)
	// var nums2 =make([]int,len(nums))
	// copy(nums2,nums)
	// fmt.Print(nums2,nums)

	//slices
	var nums = []int{1,2,4}
	fmt.Print(nums[0:2])

}