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
	goFileMap := make(map[string][]ChineseRow)
	sqlFileMap := make(map[string][]ChineseRow)
	for _, fileName := range filePaths {
		// 檢查是否式go檔或sql檔
		if extension, check := checkFileExtension(fileName); check {
			if file, rows := getChineseFileName(fileName); len(rows) > 0 {
				switch extension {
				case "go":
					goFileMap[file] = rows
				case "sql":
					sqlFileMap[file] = rows
				}
			}
		}
		break
	}
	createFileByExtension(goFileMap, "athenaGo.txt")
	createFileByExtension(sqlFileMap, "athenaSQL.txt")
}

func createFileByExtension(files map[string][]ChineseRow, writeToFileName string) {
	resetContent := []string{}
	for file, rows := range files {
		str := fmt.Sprintf("file: %s\n", file)
		for _, row := range rows {
			str += fmt.Sprintf("%d\t%s\n", row.row, row.chinese)
		}
		resetContent = append(resetContent, str)
	}
	WriteToFile(strings.Join(resetContent, "\n--------------------------------------------------\n"), writeToFileName)
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

type ChineseRow struct {
	row     int
	chinese string
}

func getChineseFileName(fileName string) (string, []ChineseRow) {
	content, err := readFile(fileName)
	if err != nil {
		fmt.Println(err)
		return "", []ChineseRow{}
	}
	lines := strings.Split(content, "\n")
	chineseRows := []ChineseRow{}
	for index, line := range lines {
		// 排除註解
		subLine := strings.Split(line, "//")
		if len(subLine) == 2 {
			tmp := getChinese(line, index+1)
			if len(tmp.chinese) > 0 {
				chineseRows = append(chineseRows, tmp)
			}
			continue
		}
		tmp := getChinese(line, index+1)
		if len(tmp.chinese) > 0 {
			chineseRows = append(chineseRows, tmp)
		}
	}
	return fileName, chineseRows
}

func getChinese(str string, row int) ChineseRow {
	chinese := ChineseRow{
		row:     row,
		chinese: "",
	}
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			chinese.chinese += string(r)
		}
	}
	return chinese
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
