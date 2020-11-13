package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	sqlFolderMap = make(map[string]string)
)

func main() {
	//rootFolder := "test"
	//folders, err := ioutil.ReadDir(rootFolder)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//filePath := GetAllFileInFolder(rootFolder, folders)
	//createNewFolder(rootFolder, folders)
	//initFolderMap(filePath)
	//for _, file := range filePath {
	//	createFormatSQL(file)
	//}

	//getSQLParameter()
	CreateNewFolder("test_", "sql")
	fmt.Println(len(getAllFile("sql")))

}

func getAllFile(rootFolder string) []string {
	paths := []string{}
	folders, err := ioutil.ReadDir(rootFolder)
	if err != nil {
		return nil
	}
	for _, f := range folders {
		subFolder := rootFolder + "/" + f.Name()
		if f.IsDir() {
			paths = append(paths, getAllFile(subFolder)...)
		} else {
			paths = append(paths, subFolder)
		}
	}
	return paths
}

//func GetAllFileInFolder(rootFolder string, folders []os.FileInfo) []string {
//	paths := []string{}
//	fn := func(rootFolder string, folders []os.FileInfo) []string { return []string{} }
//	fn = func(rootFolder string, folders []os.FileInfo) []string {
//		tmpPath := []string{}
//		for _, folder := range folders {
//			file := rootFolder + "/" + folder.Name()
//			if folder.IsDir() {
//				subFolder, _ := ioutil.ReadDir(file)
//				fn(file, subFolder)
//			} else {
//				tmpPath = append(tmpPath, file)
//			}
//		}
//	}
//
//	paths = append(paths, fn(rootFolder, folders)...)
//
//	return paths
//}

func CreateNewFolder(mean, rootFolder string) error {
	folders, err := ioutil.ReadDir(rootFolder)
	if err != nil {
		return err
	}
	//rootFolder = mean + rootFolder
	for _, f := range folders {
		subFolder := rootFolder + "/" + f.Name()
		if f.IsDir() {
			os.MkdirAll(mean+subFolder, os.ModePerm)
			CreateNewFolder(mean, subFolder)
		}
	}
	return nil
}

func getParameter() map[string]string {
	parameter := readFileToString("parameter.txt")
	parameters := strings.Split(parameter, "\n")
	parameterMap := make(map[string]string)
	fmt.Println(parameters)
	fmt.Println(len(parameters))

	for _, p := range parameters {
		ps := strings.Split(p, ":")
		parameterMap[ps[0]] = ps[1]
	}
	fmt.Println(parameterMap)

	return parameterMap
}

func getSQLParameter() {
	file := "new_test/test.sql"
	sql := readFileToString(file)
	r, _ := regexp.Compile("\\{([a-z_]+)}")
	matches := r.FindAllStringSubmatch(sql, -1)
	parameterFileContent := ""
	parameterMap := make(map[string]bool)
	for _, match := range matches {
		if !parameterMap[match[0]] {
			parameterFileContent += fmt.Sprintf("%s:\n", match[0])
			parameterMap[match[0]] = true
		}
	}
	writeToFile(parameterFileContent, "parameter.txt")
}

func initFolderMap(filePath []string) {
	for _, file := range filePath {
		fmt.Println(file)
		sqlFolderMap[file] = "new_" + file
	}
}

func createFormatSQL(file string) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	originSQL := string(buf)
	sql := sqlFormat(originSQL)
	if err := copyToNewFile(file, sql); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("create new %s sql \n", file)
}

func copyToNewFile(file, sql string) error {
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
			// COALESCE(SUM(tmp.transfer_amount),0)
			match, _ := regexp.MatchString("([a-zA-Z_]+)\\(([a-zA-Z_]+)\\(([a-zA-Z_]+)\\.([a-zA-Z_]+)\\),(.*)", word)
			if match {
				r, _ := regexp.Compile("([a-zA-Z_]+)\\(([a-zA-Z_]+)\\(([a-zA-Z_]+)\\.([a-zA-Z_]+)\\),(.*)")
				matches := r.FindStringSubmatch(word)
				if specialWordMap[strings.ToLower(matches[1])] {
					matches[1] = strings.ToUpper(matches[1])
				}
				if specialWordMap[strings.ToLower(matches[2])] {
					matches[2] = strings.ToUpper(matches[2])
				}
				sql += fmt.Sprintf("%s(%s(`%s`.`%s`),%s ", matches[1], matches[2], matches[3], matches[4], matches[5])
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
			// view_member_agent.share_login,
			// 主要解析 '.' ',' 之間的英文
			match, _ = regexp.MatchString("([a-z_]+)\\.([a-z_]+),", word)
			if match {
				r, _ := regexp.Compile("([a-z_]+)\\.([a-z_]+),")
				matches := r.FindStringSubmatch(word)
				sql += fmt.Sprintf("`%s`.`%s`, ", matches[1], matches[2])
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
				sql += fmt.Sprintf("`%s`, ", matches[1])
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

func checkFormat(word string) bool {
	match, _ := regexp.MatchString("`(.*)`", word)
	if match {
		return true
	}
	return false
}

func readFileToString(file string) string {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(buf)
}

func writeToFile(value, fileName string) error {
	sqlFile := []byte(value)

	f, err := os.Create(fileName)
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

func createNewFolder(rootFolder string, folders []os.FileInfo) {
	newFolderRoot := "new_" + rootFolder
	os.MkdirAll(newFolderRoot, os.ModePerm)
	fn := func(rootFolder string, folders []os.FileInfo) {}
	fn = func(rootFolder string, folders []os.FileInfo) {
		for _, folder := range folders {
			file := newFolderRoot + "/" + folder.Name()
			if folder.IsDir() {
				subFolder, _ := ioutil.ReadDir(file)
				os.MkdirAll(file, os.ModePerm)
				fn(file, subFolder)
			}
		}
	}
	fn(rootFolder, folders)
}
