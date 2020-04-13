package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	TitleMap   = map[string][]string{}
	TitleArray = []string{}
	Brand      = []string{"lv", "ls", "c7"}
)

func main() {
	var option string
	var itemTitle string
	var items string
	fmt.Println("請選擇操作方式: \n1.新增 \n2.刪除")
	fmt.Scanln(&option)
	fmt.Println("請輸入要新增/刪除的title: ")
	fmt.Scanln(&itemTitle)
	fmt.Println("請輸入設定: (多項時用逗號隔開)")
	fmt.Scanln(&items)

	Items := strings.Split(items, ",")
	refresh_conf_func(itemTitle, Items, option)

}

func setTitleMap(title string) {
	if _, exist := TitleMap[title]; !exist {
		TitleMap[title] = []string{}
		TitleArray = append(TitleArray, title)
	}
}

func getFile(brand string) {
	r, _ := regexp.Compile("\\[([a-z]+)\\]")

	f, err := os.Open("./" + brand + "/test.conf")
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)
	title := ""
	for s.Scan() {
		lineString := s.Text()
		if r.MatchString(lineString) {
			title = lineString
			setTitleMap(title)
			continue
		} else if lineString == "" {
			continue
		}
		if check(title, lineString) {
			TitleMap[title] = append(TitleMap[title], lineString)
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

}

func check(title, newItem string) bool {
	for _, item := range TitleMap[title] {
		if item == newItem {
			return false
		}
	}
	return true
}

func refresh_conf_func(itemTitle string, items []string, option string) {
	for _, brand := range Brand {
		getFile(brand)
		if option == "2" {
			remove(itemTitle, items)
		} else {
			if _, exist := TitleMap[itemTitle]; !exist {
				setTitleMap(itemTitle)
			}
			for _, addItem := range items {
				if check(itemTitle, addItem) {
					TitleMap[itemTitle] = append(TitleMap[itemTitle], addItem)
				}
			}
		}

		var err error
		count := 0
		for _, title := range TitleArray {
			writeString := title + "\n"
			for _, item := range TitleMap[title] {
				writeString += item + "\n"
			}
			writeString += "\n"
			str := []byte(writeString)
			if count == 0 {
				err = ioutil.WriteFile("./"+brand+"/test.conf", str, 0644)
				count++
			} else {
				f, _ := os.OpenFile("./"+brand+"/test.conf", os.O_WRONLY, 0644)
				n, _ := f.Seek(0, os.SEEK_END)
				_, err = f.WriteAt([]byte(writeString), n)
			}
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func remove(removeItemTitle string, removeItems []string) {
	for _, s := range removeItems {
		for index, item := range TitleMap[removeItemTitle] {
			if s == item {
				TitleMap[removeItemTitle] = append(TitleMap[removeItemTitle][:index], TitleMap[removeItemTitle][index+1:]...)
				break
			}
		}
	}
}
