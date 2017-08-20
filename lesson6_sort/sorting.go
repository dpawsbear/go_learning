package main

import (
	"sort"
	"fmt"
)

func main(){
	// example for string
	strs := []string{"c","a","b"}
	sort.Strings(strs)

	fmt.Println("Strings:",strs)

	//example of sorting
	ints := []int{7,2,4}
	sort.Ints(ints)
	fmt.Println("Ints: ",ints)

	//使用 sort 检查是相关 slice 是否已经被排好序
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted: ",s)

}
