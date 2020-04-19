package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var (
	compareWeb = make(map[string]int)
)

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("web %d", i)
		x := rand.Intn(1000)
		compareWeb[key] = x
	}
}

func main() {
	valueChan := make(chan int, len(compareWeb))

	for web, _ := range compareWeb {
		go getWebValue(web, valueChan)
	}

	valueArray := []int{}
	for i := 0; i < 10; i++ {
		valueArray = append(valueArray, <-valueChan)
	}

	fmt.Println(valueArray)
	sort.Ints(valueArray)
	fmt.Println(valueArray)
}

func getWebValue(web string, value chan int) {
	value <- compareWeb[web]
}
