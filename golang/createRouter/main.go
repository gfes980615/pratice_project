package main

import (
	"fmt"
	"github.com/gfes980615/mytool-golang/filetool"
	"regexp"
	"strings"
)

func main() {
	content, err := filetool.ReadFileToString("test.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	content = strings.ReplaceAll(content, "\r\n", " ")
	//fmt.Println(content)
	methodContent := getMethodContent(content, "SetupRouter")
	methodContentlines := strings.Split(methodContent, "\t")
	for _, line := range methodContentlines {
		line = strings.ReplaceAll(line, " ", "")
		r := regexp.MustCompile(`(.*):=(.*).Group\("(.*)"\)`)
		subContents := r.FindStringSubmatch(line)
		if len(subContents) == 0 {
			continue
		}
		fmt.Println(subContents)
	}
}



func getMethodContent(content, method string) string {
	reStr := fmt.Sprintf("func(.*) %s(.*)", method)
	r := regexp.MustCompile(reStr)
	contents := r.FindStringSubmatch(content)
	return parseContent(contents[2])
}

func parseContent(content string) string {
	newContent := ""
	tmp := 0
	startFlag := false
	for _, str := range content {
		if str == '{' {
			startFlag = true
			tmp++
			continue
		}
		if str == '}' {
			tmp--
			continue
		}
		if startFlag {
			newContent += string(str)
		}
		if startFlag && tmp == 0 {
			break
		}
	}
	return newContent
}
