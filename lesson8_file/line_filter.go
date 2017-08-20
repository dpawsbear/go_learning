package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main(){
	//用缓冲扫描包装无缓冲的 os.Stdin 可以让我们方便的扫描方法使扫描仪进入下一标记，
	//这是默认扫描器中的下一行。
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan(){
		// `Text` 返回当前标记，从输入来说这里是下一行
		ucl := strings.ToUpper(scanner.Text())
		//写出上面的行
		fmt.Println(ucl)
	}

	//检查扫描期间的错误。文件结束是预期的，不会被 scan 报告为错误。
	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr,"error:",err)
		os.Exit(1)
	}
}
