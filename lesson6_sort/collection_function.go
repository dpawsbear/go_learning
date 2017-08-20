package main

import (
	"fmt"
	"strings"
)

//找寻单词排第几
func Index(vs []string,t string ) int{
	for i,v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

//判断是否包含有该单词
func Include(vs []string,t string) bool{
	return Index(vs,t)>=0
}

//当含有目标单词，返回true
func Any(vs []string, f func(string) bool) bool{
	for _,v := range vs{
		if f(v) {
			return true
		}
	}
	return false
}
//当所有单词都在里面时，才返回true
func All(vs [] string,f func(string)bool)bool{
	for _,v := range vs{
		if !f(v){
			return false
		}
	}
	return true
}

//过滤包含某单词的
func Filter(vs []string, f func(string)bool) []string{
	vsf := make([]string , 0)
	for _,v := range vs {
		if f(v){
			vsf = append(vsf,v)
		}
	}
	return vsf
}
//新切片
func Map(vs []string, f func(string)string)[]string{
	vsm := make([]string,len(vs))
	for i,v := range vs{
		vsm[i] = f(v)
	}
	return vsm
}

func main(){
	var strs = []string{"peach","apple","pear","plum"}

	fmt.Println(Index(strs,"pear"))

	fmt.Println(Include(strs,"grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v,"p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v,"p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v,"e")
	}))

	fmt.Println(Map(strs,strings.ToUpper))

}