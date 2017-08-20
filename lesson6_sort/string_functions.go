package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println


func main(){

	//在 http://golang.org/pkg/strings/ 中有更多的解释
	p("Contains:  ",s.Contains("test","es"))

	p("Count:     ",s.Count("Test","t"))
	p("HasPrefix: ",s.HasPrefix("test","te"))
	p("HasSuffix: ",s.HasSuffix("test","st"))
	p("Index:     ",s.Index("test","e"))
	p("Join:      ",s.Join([]string{"a","b"},"-"))
	p("Repeat:    ",s.Repeat("a",5))
	p("Replace:   ",s.Replace("foo","o","0",-1))
	p("Replace:   ",s.Replace("foo","o","0",1))
	p("Split:     ",s.Split("a-b-c-d-e","-"))
	p("ToLoer:    ",s.ToLower("TEST"))
	p("ToUpper:   ",s.ToUpper("test"))
	p()

	p("Len:       " ,len("Hello"))
	p("Char:       ","Hello"[1])

}
