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

	sqlFolderMap = make(map[string]string)
)

func main() {
	rootFolder := "deposit"
	folders, err := ioutil.ReadDir(rootFolder)
	if err != nil {
		fmt.Println(err)
		return
	}

	filePath := getAllFileInFolder(rootFolder, folders)
	initFolderMap(filePath)
	for _, file := range filePath {
		createFormatSQL(file)
	}
}

func initFolderMap(filePath []string) {
	for _, file := range filePath {
		fmt.Println(file)
		sqlFolderMap[file] = "new_" + file
	}
}

func getAllFileInFolder(rootFolder string, folders []os.FileInfo) []string {
	paths := []string{}
	newFolderRoot := "new_" + rootFolder
	os.MkdirAll(newFolderRoot, os.ModePerm)
	fn := func(newRootFolder, rootFolder string, folders []os.FileInfo) {}
	fn = func(newRootFolder, rootFolder string, folders []os.FileInfo) {
		for _, folder := range folders {
			if folder.IsDir() {
				subFolder, _ := ioutil.ReadDir(rootFolder + "/" + folder.Name())
				os.MkdirAll(newRootFolder+"/"+folder.Name(), os.ModePerm)
				fn(newRootFolder+"/"+folder.Name(), rootFolder+"/"+folder.Name(), subFolder)
			} else {
				paths = append(paths, rootFolder+"/"+folder.Name())
			}
		}
	}
	fn(newFolderRoot, rootFolder, folders)

	return paths
}

func createFormatSQL(file string) {
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

	fmt.Printf("create new %s sql \n", file)
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
			// 特殊字串轉成大寫
			if specialWordMap[strings.ToLower(word)] {
				sql += strings.ToUpper(word) + " "
				continue
			}
			// COALESCE(SUM(tmp.transfer_amount),
			match, _ := regexp.MatchString("([a-z_]+)\\(([a-z_]+)\\(([a-z_]+)\\.([a-z_]+)\\),", word)
			if match {
				r, _ := regexp.Compile("([a-z_]+)\\(([a-z_]+)\\(([a-z_]+)\\.([a-z_]+)\\),")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				if specialWordMap[strings.ToLower(matches[2])] {
					matches[1] = strings.ToUpper(matches[2])
				}
				sql += fmt.Sprintf("%s(%s(`%s`.`%s`), ", matches[1], matches[2], matches[3], matches[4])
				continue
			}

			// find_in_set(view_member_reseller_agents.share_login,
			// 解析 '(' , '.' , ',' 之間的英文
			match, _ = regexp.MatchString("(.*)\\((.*)\\.(.*),", word)
			if match {
				r, _ := regexp.Compile("(.*)\\((.*)\\.(.*),")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				sql += fmt.Sprintf("%s(`%s`.`%s`, ", matches[1], matches[2], matches[3])
				continue
			}
			// IFNULL(deposit.transfer_amount
			// 解析 '(' 與 '.' 兩側的英文
			match, _ = regexp.MatchString("(.*)\\((.*)\\.(.*)", word)
			if match {
				r, _ := regexp.Compile("(.*)\\((.*)\\.(.*)")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				sql += fmt.Sprintf("%s(`%s`.`%s` ", matches[1], matches[2], matches[3])
				continue
			}
			// view_member_agent.share_login
			// 主要解析 '.' 兩側的英文
			match, _ = regexp.MatchString("([a-z_]+)\\.([a-z_]+)", word)
			if match {
				r, _ := regexp.Compile("([a-z_]+)\\.([a-z_]+)")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("`%s`.`%s` ", matches[1], matches[2])
				continue
			}
			// (IFNULL({member_share_login},-99)=-99
			// 主要解析兩個 '(' 之間的英文 (ex IFNULL)
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

			// action_code,
			// 解析 ',' 前的英文
			match, _ = regexp.MatchString("([a-z_]+),", word)
			if match {
				r, _ := regexp.Compile("([a-z_]+),")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("`%s`,", matches[1])
				continue
			}
			// action_code
			// 普通英文的處理
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

	f, err := os.Create(sqlFolderMap[file])
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

func getNewFileName(file string) string {
	f := strings.Split(file, ".sql")
	return fmt.Sprintf("%s_new.sql", f[0])
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
