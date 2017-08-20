package main

import (
	"crypto/sha1"
	"fmt"
)

func main(){
	s := "only test you"

	//new hash

	h := sha1.New()

	h.Write([]byte(s))

	//计算最终散列值，后面nil可以附件到现有的字节片段
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n",bs)
}
