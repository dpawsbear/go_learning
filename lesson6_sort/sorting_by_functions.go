package main

import (
	"sort"
	"fmt"
)

//定义类型
type ByLength []string

//重写下面几个接口函数
func (s ByLength) Len()int{
	return len(s)
}

func (s ByLength) Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}

func (s ByLength) Less(i,j int) bool {
	return len(s[i]) < len(s[j])
}


//正常调用方法
func main(){
	fruits := []string{"peach","banana","kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}