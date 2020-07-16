package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

const (
	rUserURL     = `mvc.(.*?)\("(.*)", (.*)\).SetItem\(utilsMVC.API_Level, utilsMVC.API_Level_User\).SetWriter\(&biApiWriter\)`
	rMemberURL   = `mvc.(.*?)\("(.*)", (.*)\).SetItem\(utilsMVC.API_Level, utilsMVC.API_Level_Member\).SetWriter\(&biApiWriter\)`
	rResellerURL = `mvc.(.*?)\("(.*)", (.*)\).SetItem\(utilsMVC.API_Level, utilsMVC.API_Level_Agent\).SetWriter\(&biApiWriter\)`
	rfuncAPI     = `func (.{0,50})\(ctx \*mvc.Context\) \{(.*?)api := (.*?)ApisToAthenaV2\(ctx, api\)}`
)

func getFileString(file string) string {
	file, err := getFileResult(file)
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func setInitFuncURL(file string, folder string) (map[string]map[string]string, []string) {
	src := getFileString(file)
	re, _ := regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	var tmpRegexp string
	switch folder {
	case "admin":
		tmpRegexp = rUserURL
	case "member":
		tmpRegexp = rMemberURL
	case "reseller":
		tmpRegexp = rResellerURL
	}

	rpURL := regexp.MustCompile(tmpRegexp)
	itemsGetURL := rpURL.FindAllStringSubmatch(src, -1)

	sortFunc := []string{}
	urlMap := make(map[string]map[string]string)
	for _, item := range itemsGetURL {
		sortFunc = append(sortFunc, item[3])
		method := make(map[string]string)
		method = map[string]string{
			item[1]: item[2],
		}
		urlMap[item[3]] = method
	}

	return urlMap, sortFunc
}

func setFuncURL(file string) map[string]string {
	src := getFileString(file)
	re, _ := regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "")

	rpFunc := regexp.MustCompile(rfuncAPI)
	itemsFunc := rpFunc.FindAllStringSubmatch(src, -1)
	// fmt.Println(src)
	funcMap := make(map[string]string)
	for _, item := range itemsFunc {
		tmp := strings.Replace(item[3], "\" + ", "{", -1)
		tmp2 := strings.Replace(tmp, " + \"", "}", -1)
		tmp3 := strings.Replace(tmp2, "\"", "", -1)
		if len(strings.Split(tmp3, "}")) != len(strings.Split(tmp3, "{")) {
			tmp3 += "}"
		}
		funcMap[item[1]] = tmp3
	}
	return funcMap
}

func createYaml(file string, createFolder string, folder string) {
	initURL, sortFunc := setInitFuncURL(file, folder)
	funcURL := setFuncURL(file)
	r, _ := regexp.Compile("{(.*?)}")
	var level string
	switch folder {
	case "admin":
		level = "user"
	case "member":
		level = "member"
	case "reseller":
		level = "agent"
	}
	defaultString := "servive: athena\napi_level: " + level + "\napis:\n"
	urlString := ""
	for _, sf := range sortFunc {
		method := initURL[sf]
		for m, url := range method {
			admin := r.FindAllString(url, -1)
			athena := r.FindAllString(funcURL[sf], -1)
			for i := 0; i < len(admin); i++ {
				funcURL[sf] = strings.Replace(funcURL[sf], athena[i], admin[i], -1)
			}
			urlString += fmt.Sprintf("  - path: %s\n    redirect: %s\n    method: %s\n", url, funcURL[sf], m)
		}

	}
	resultString := defaultString + urlString

	d1 := []byte(resultString)
	f, err := os.Create(createFolder)
	err = ioutil.WriteFile(createFolder, d1, 0644)
	defer f.Close()
	fmt.Println(err)
}

func main() {
	folders, _ := ioutil.ReadDir("./bisystem")
	for _, folder := range folders {
		if folder.IsDir() {
			tmpFolderName := "./bisystem/" + folder.Name()
			files, _ := ioutil.ReadDir(tmpFolderName)
			for _, file := range files {
				if file.Name() == "adapter_athena.go" {
					continue
				}
				yamlFile := strings.Replace(file.Name(), ".go", ".yaml", -1)
				readFile := tmpFolderName + "/" + file.Name()
				createFolder := "./adapter.conf/athena/" + folder.Name() + "/" + yamlFile
				createYaml(readFile, createFolder, folder.Name())
			}
		}
	}

}

func getFileResult(fileName string) (string, error) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("read fail", err)
	}
	return string(f), nil
}
