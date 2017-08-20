package main

import (
	"fmt"
	"math/rand"
)

func main(){
	p := fmt.Println

	//产生 0 - 100 的伪随机数
	p(rand.Intn(100))
	p(rand.Intn(100))

	//产生浮点随机数
	p(rand.Float64())
	p(rand.ExpFloat64()*10)
	p(rand.ExpFloat64()*100)


}
