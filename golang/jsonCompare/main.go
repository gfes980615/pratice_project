package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	getMapKeys2("test3.json")
	//keys := getMapKeys("test1.json")
	//fileOneMap := getFileDataToMap("test1.json", keys)
	//fileTwoMap := getFileDataToMap("test2.json", keys)
	//
	//count := 0
	//for key, _ := range fileOneMap {
	//	if fileTwoMap[key] {
	//		count++
	//	}
	//}
	//
	//if len(fileOneMap) == len(fileTwoMap) && count == len(fileOneMap) && count == len(fileTwoMap) {
	//	fmt.Println("equally")
	//	return
	//}
}

func getFileDataToMap(file string, keys []string) map[string]bool {
	buf, _ := ioutil.ReadFile(file)
	fileResult := []interface{}{}
	json.Unmarshal(buf, &fileResult)

	resultMap := make(map[string]bool)
	for _, result := range fileResult {
		v := reflect.ValueOf(result)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		key := ""
		for _, k := range keys {
			key += fmt.Sprintf("%v/", v.MapIndex(reflect.ValueOf(k)))
		}
		resultMap[key] = true
	}

	return resultMap
}

// 為了固定順序
func getMapKeys(file string) []string {
	buf, _ := ioutil.ReadFile(file)
	fileResult := []interface{}{}
	err := json.Unmarshal(buf, &fileResult)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	v := reflect.ValueOf(fileResult[0])
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	result := []string{}
	for _, value := range v.MapKeys() {
		result = append(result, value.String())
	}
	return result
}

func getMapKeys2(file string) {
	buf, _ := ioutil.ReadFile(file)
	var fileResult interface{}
	err := json.Unmarshal(buf, &fileResult)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileResult)
	v := reflect.ValueOf(fileResult)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	result := []string{}
	if v.Kind() == reflect.Map {
		
	}
	for _, value := range v.MapKeys() {
		result = append(result, value.String())
	}
}
