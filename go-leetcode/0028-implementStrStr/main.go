package main

import "fmt"

func strStr(haystack string, needle string) int  {
	if len(needle) == 0{
		return 0
	}
	head ,tail := 0, len(needle)
	for tail <= len(haystack){
		fmt.Println(haystack[head:tail])
		if haystack[head:tail] == needle{
			return head
		}
		head++
		tail++
	}
	return -1
}

func main() {
	haystack := "a"
	needle := "a"
	fmt.Println(strStr(haystack,needle))
}
