package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

var (
	specialWordMap = map[string]bool{
		"select":   true,
		"count":    true,
		"coalesce": true,
		"sum":      true,
		"as":       true,
		"from":     true,
	}
)

func main() {
	buf, err := ioutil.ReadFile("C:/Users/fred_chen/Desktop/pratice_project/golang/sqlFormat/test.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	originSQL := string(buf)
	newSQL := ""
	sqlLineString := strings.Split(originSQL, "\r\n")
	for _, sqlLine := range sqlLineString {
		fmt.Println(sqlLine)
		lineWord := strings.Split(sqlLine, " ")
		for index, v := range lineWord {
			if strings.ToLower(v) == "select" {
				newSQL += v + "\n\t"
			} else if !IsEnglishWord(v) {
				if hasComma(v) {
					newSQL += "`" + strings.Split(v, ",")[0] + "`" + ",\n\t"
				} else {
					newSQL += v + " "
				}
			} else if specialWordMap[strings.ToLower(v)] {
				newSQL += strings.ToUpper(v) + " "
			} else {
				newSQL += "`" + v + "`" + " "
			}
			if index == len(lineWord)-1 {
				newSQL += "\n"
			}
		}
		break
	}
	fmt.Println(newSQL)
	//fmt.Println(string(buf))
}

func IsEnglishWord(word string) bool {
	for _, w := range word {
		if !unicode.IsLetter(w) {
			return false
		}
	}
	return true
}
func hasComma(word string) bool {
	if word[len(word)-1] == ',' {
		return true
	}
	return false
}
