package main

import (
	"net/url"
	"fmt"
	"net"
)

func main(){
	//示例网址
	s := "postgres://user:pass@host.com:1314/path?k=v#f"

	//解析网址，确保无错
	u,err := url.Parse(s)
	if err != nil{
		panic(err)
	}

	//访问 scheme
	fmt.Println(u.Scheme)

	//User 包含了所有验证信息 用户名和密码都可以单独来。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p,_:=u.User.Password()
	fmt.Println(p)

	//Host 包含了 hostname 和 port 可以使用 SplitHostPort解析他们
	fmt.Println(u.Host)
	host ,port ,_:=net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	//提取路径和片段
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	//要使用K=V格式的字符串获取查询参数，需要使用RawQuery。也可以解析到map中。
	//解析查询参数映射是从字符串到字符串的片段。因此索引为0
	fmt.Println(u.RawQuery)
	m,_:= url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
