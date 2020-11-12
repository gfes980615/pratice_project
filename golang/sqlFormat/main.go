package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"
)

var (
	specialWordMap = map[string]bool{
		"select":      true,
		"count":       true,
		"coalesce":    true,
		"sum":         true,
		"as":          true,
		"from":        true,
		"and":         true,
		"or":          true,
		"ifnull":      true,
		"where":       true,
		"if":          true,
		"exists":      true,
		"inner":       true,
		"join":        true,
		"left":        true,
		"find_in_set": true,
		"union":       true,
		"all":         true,
		"in":          true,
	}
)

func main() {
	file := "test.sql"
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	originSQL := string(buf)
	sql := sqlFormat(originSQL)
	if err := writeTofile(file, sql); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}

func sqlFormat(originSQL string) string {
	sqlLineString := strings.Split(originSQL, "\r\n")
	sql := ""
	for _, line := range sqlLineString {
		for _, word := range strings.Split(line, " ") {
			if checkFormat(word) {
				sql += word + " "
				continue
			}
			if specialWordMap[strings.ToLower(word)] {
				sql += strings.ToUpper(word) + " "
				continue
			}

			match, _ := regexp.MatchString("(.*)\\((.*)\\.(.*),", word)
			if match {
				r, _ := regexp.Compile("(.*)\\((.*)\\.(.*),")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				sql += fmt.Sprintf("%s(`%s`.`%s`, ", matches[1], matches[2], matches[3])
				continue
			}

			match, _ = regexp.MatchString("(.*)\\((.*)\\.(.*)", word)
			if match {
				r, _ := regexp.Compile("(.*)\\((.*)\\.(.*)")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("%s(`%s`.`%s` ", matches[1], matches[2], matches[3])
				continue
			}

			match, _ = regexp.MatchString("([a-z_]+)\\.([a-z_]+)", word)
			if match {
				r, _ := regexp.Compile("([a-z_]+)\\.([a-z_]+)")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("`%s`.`%s` ", matches[1], matches[2])
				continue
			}

			match, _ = regexp.MatchString("^\\(([a-z_]+)\\((.*)", word)
			if match {
				r, _ := regexp.Compile("^\\(([a-z_]+)\\((.*)")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				sql += fmt.Sprintf("(%s(%s ", matches[1], matches[2])
				continue
			}

			match, _ = regexp.MatchString("([a-z]+),", word)
			if match {
				r, _ := regexp.Compile("([a-z]+),")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("`%s`,", matches[1])
				continue
			}

			match, _ = regexp.MatchString("^([a-z_]+)$", word)
			if match {
				r, _ := regexp.Compile("^([a-z_]+)$")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("`%s` ", matches[1])
				continue
			}

			sql += word + " "

		}
		sql += "\n"
	}
	return sql
}

func writeTofile(file, sql string) error {
	sqlFile := []byte(sql)
	f, err := os.Create(fmt.Sprintf("new_%s", file))
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(sqlFile)
	if err != nil {
		return err
	}
	return nil
}

func checkFormat(word string) bool {
	match, _ := regexp.MatchString("`(.*)`", word)
	if match {
		return true
	}
	return false
}

func handler1(sqlLineString []string) string {
	newSQL := ""
	for _, sqlLine := range sqlLineString {
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
			} else if len(v) == 0 {
				newSQL += " "
			} else {
				newSQL += "`" + v + "`" + " "
			}
			if index == len(lineWord)-1 {
				newSQL += "\n"
			}
		}

	}
	return newSQL
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
