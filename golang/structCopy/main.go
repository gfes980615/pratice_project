package main

import (
	"fmt"
	"reflect"
)

type StructInterface interface {
	GetStruct() interface{}
	GetStructPtr() interface{}
}

type Original struct {
	Title string
	URL   string
	Name  string
	AAA   int
	BBB   int
}

type Result struct {
	Title string
	AAA   int
	Name  string
	CCC   int
}

func (r Result) GetStruct() interface{} {
	return Result{}
}
func (r Result) GetStructPtr() interface{} {
	return &Result{}
}

func main() {
	ori := &Original{
		Title: "123",
		URL:   "http://",
		Name:  "name",
		AAA:   123,
		BBB:   456,
	}
	res := structCopy(ori, Result{})
	fmt.Println(res.(*Result))
}

func structCopy(ori interface{}, resInterface StructInterface) interface{} {
	res := resInterface.GetStruct()
	resValue := reflect.ValueOf(res)

	oriValue := reflect.ValueOf(ori)
	oriType := reflect.TypeOf(ori)
	fieldMap := make(map[string]interface{})
	for i := 0; i < resValue.NumField(); i++ {
		fieldName := oriType.Elem().Field(i).Name
		if _, ok := fieldMap[fieldName]; !ok {
			fieldMap[fieldName] = oriValue.Elem().FieldByName(fieldName).Interface()
		}
	}

	result := resInterface.GetStructPtr()
	rf := reflect.ValueOf(result)
	s := rf.Elem()
	if s.Kind() == reflect.Struct {
		for field, v := range fieldMap {
			f := s.FieldByName(field)
			if f.IsValid() && f.CanSet() {
				switch f.Kind() {
				case reflect.String:
					f.SetString(v.(string))
				case reflect.Int:
					f.SetInt(int64(v.(int)))
				case reflect.Float64:
					f.SetFloat(v.(float64))
				case reflect.Bool:
					f.SetBool(v.(bool))
				}
			}
		}
	}

	return result
}
