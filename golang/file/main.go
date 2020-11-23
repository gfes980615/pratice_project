package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func main() {
	filePaths := getAllFileInFolder("test")
	goFileArray := []string{}
	sqlFileArray := []string{}
	for _, fileName := range filePaths {
		// 檢查是否式go檔或sql檔
		if extension, check := checkFileExtension(fileName); check {
			if file := getChineseFileName(fileName); file != "" {
				switch extension {
				case "go":
					goFileArray = append(goFileArray, file)
				case "sql":
					sqlFileArray = append(sqlFileArray, file)
				}
			}
		}
		break
	}
	//createFileByExtension(goFileArray, "athenaGo.txt")
	//createFileByExtension(sqlFileArray, "athenaSQL.txt")
}

func createFileByExtension(files []string, fileName string) {
	WriteToFile(strings.Join(files, "\n"), fileName)
}

func checkFileExtension(fileName string) (string, bool) {
	r, _ := regexp.MatchString("(.*).(go|sql)", fileName)
	re := regexp.MustCompile("(.*).(go|sql)")
	e := re.FindStringSubmatch(fileName)
	if len(e) != 3 {
		return "", r
	}
	return e[2], r
}

func getChineseFileName(fileName string) string {
	content, err := readFile(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	lines := strings.Split(content, "\n\t")
	for index, line := range lines {
		fmt.Println(index, line)
		//r := regexp.MustCompile("(.*)(\n)")
		//fmt.Println(r.FindStringSubmatch(line))
		// 排除註解
		//subLine := strings.Split(line, "//")
		//if len(subLine) > 1 {
		//	continue
		//}
		//// 檢查有沒有中文
		//if checkChinese(line) {
		//
		//}
	}
	return ""
}

func getAllFileInFolder(rootFolder string) []string {
	paths := []string{}
	folders, err := ioutil.ReadDir(rootFolder)
	if err != nil {
		return nil
	}
	for _, f := range folders {
		subFolder := rootFolder + "/" + f.Name()
		if f.IsDir() {
			paths = append(paths, getAllFileInFolder(subFolder)...)
		} else {
			paths = append(paths, subFolder)
		}
	}
	return paths
}

func checkChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

func readFile(file string) (string, error) {
	sqlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(sqlFile), nil
}

func WriteToFile(value, fileName string) error {
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
