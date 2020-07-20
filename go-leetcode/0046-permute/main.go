package main

import "fmt"

var result [][]int


func backtrack(nums, pathNums []int, used []bool)   {
	if len(nums) == len(pathNums){
		var temp = make([]int,len(nums))
		copy(temp,pathNums)
		result = append(result,temp)
		return
	}

	for i:= 0; i < len(nums); i ++{
		if !used[i]{
			used[i] = true
			pathNums = append(pathNums,nums[i])
			backtrack(nums,pathNums,used)
			pathNums = pathNums[:len(pathNums)-1]
			used[i] = false
		}
	}
}

func permute(nums []int) [][]int  {
	var pathNums []int
	var used = make([]bool,len(nums))
	var result = [][]int{}
	backtrack(nums,pathNums,used)
	return result
}

func main() {
	var a=[]int{1,2,3}
	permute(a)
	fmt.Println(result)
}
