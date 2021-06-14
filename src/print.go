package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func main()  {
	fmt.Println("hello world")
	fmt.Println("runtime goos: ", runtime.GOOS)

	fmt.Println(path.Base("http://www.baidu.com/file/aa.jpg"))

	//matches, err := fmt.Println(filepath.Glob("./*.exe"))
	matches, err := filepath.Glob("./*.go")
	fmt.Println(err)
	for k, v := range matches {
		fmt.Println(k, v)
	}

	log.Println("log println")

	//fmt.Println(strings.TrimSpace("hello      world77,,    "))
	for _, value := range strings.TrimSpace(" adfaklf       hello,      world77,,    "){
		fmt.Printf("%c\n", value)
	}

	fmt.Println("***************")
	notALetter := func(char rune) bool {
	//	return !unicode.IsLetter(char)
		return false
	}
	str := strings.FieldsFunc("__hello8989, world&&34yyyy", notALetter)

	for _, v := range str{
		fmt.Println("str:", v)
	}

}





