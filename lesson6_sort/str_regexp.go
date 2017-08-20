package main

import (
	"regexp"
	"fmt"
	"bytes"
)

func main(){
	//测试规则是否能匹配字符串
	match,_:=regexp.MatchString("p([a-z]+)ch","peach")
	fmt.Println(match)

	//使用regexp其他函数的时候需要先 使用 compile 函数进行优化
	r,_:= regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	//找寻符合规则的字符串
	fmt.Println(r.FindString("peach punch"))
	//找寻符合规则的字符串所在位置 这里返回时开始和结束的位置
	fmt.Println(r.FindStringIndex("peach punch"))
	//匹配 "p([a-z]+)ch" 和 "([a-z]+)" 两个表达式
	fmt.Println(r.FindStringSubmatch("peach punch"))
	//匹配上面两个表达式找到的字符所在的位置
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	//找寻所有符合的字符串
	fmt.Println(r.FindAllString("peach punch pinch",-1))
	//找寻所有子规则符合的字符串所在位置
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch",-1))
	//限制只找两个
	fmt.Println(r.FindAllString("peach punch pinch",2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach","<fruit>"))

	in  := []byte("a peach")
	out := r.ReplaceAllFunc(in,bytes.ToUpper)
	fmt.Println(string(out))
}
