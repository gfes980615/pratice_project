package main

import (
	"fmt"
)

func main() {
	s := "dvdf"
	ans := []string{}

	tmp := 0
	for i := 0; i < len(s)-1; i++ {
		j := i + 1
		if s[i] == s[j] {
			ans = append(ans, s[tmp:i+1])
			tmp = j
			// fmt.Println("123: ", ans)
		}
		if j == len(s)-1 {
			ans = append(ans, s[tmp:j+1])
			// fmt.Println("finish: ", ans)
		}
	}

	newS := ""
	for _, t := range ans {
		if len(t) > len(newS) {
			newS = t
		}
	}
	fmt.Println(newS)

	tmp2 := []int{}
	for i := 1; i < len(newS); i++ {
		count := 0
		for j := i - 1; j >= 0; j-- {
			if newS[i] == newS[j] {
				break
			}
			count++
		}
		tmp2 = append(tmp2, count)
	}
	var big int = 0
	for _, b := range tmp2 {
		fmt.Println(b)
		if big < b {
			big = b
		}
	}
	big = big + 1

	fmt.Print(big)
}
