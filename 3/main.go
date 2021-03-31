package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBrackets(str string) string {
	i := strings.IndexByte(str, '(')
	if i < 0 {
		return ""
	}
	i++
	j := strings.IndexByte(str[i:], ')')
	if j < 0 {
		return ""
	}
	return str[i : i+j]
}

func main(){
	result := findFirstStringInBrackets("Test (123)")
	fmt.Print(result) //123
}