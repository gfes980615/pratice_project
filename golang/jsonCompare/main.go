package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	fileOneMap := getFileDataToMap("test1.json")
	fileTwoMap := getFileDataToMap("test2.json")

	count := 0
	for key, _ := range fileOneMap {
		if fileTwoMap[key] {
			count++
		} else {

		}
	}

	if len(fileOneMap) == len(fileTwoMap) && count == len(fileOneMap) && count == len(fileTwoMap) {
		fmt.Println("equally")
		return
	}
}

func getFileDataToMap(file string) map[string]bool {
	buf, _ := ioutil.ReadFile(file)
	fileResult := []OrderSettlement{}
	json.Unmarshal(buf, &fileResult)

	resultMap := make(map[string]bool)

	for _, result := range fileResult {
		v := reflect.ValueOf(result)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		key := ""
		for i := 0; i < v.NumField(); i++ {
			key += fmt.Sprintf("%v/", v.Field(i).Interface())
		}
		resultMap[key] = true
	}
	return resultMap
}

type OrderSettlement struct {
	MemberLogin interface{} `json:"member_login"`
	ChannelCode interface{} `json:"channel_code"`
	ProductCode interface{} `json:"product_code"`
	Point       interface{} `json:"point"`
	Total       interface{} `json:"total"`
	Payout      interface{} `json:"payout"`
}
