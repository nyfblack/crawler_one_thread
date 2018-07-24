package main

import (
	"fmt"
	"regexp"
)

const test = `my email is 123456789@163.com
email1 is aaa@qq.com
email2 is dfjkd@fjkd.org
email3 is jfkjk@aaa.com.cn
`

func main() {
	//正则表达式是从外界传入的时候，使用compile方法，可以处理错误
	//re, err := regexp.Compile("123456789@163.com")

	//程序内部写的正则表达式，确认无误，可使用mustCompile方法
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(test, -1)
	for _, m := range match {
		fmt.Println(m)
	}

}
