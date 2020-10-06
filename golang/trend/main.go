package main

import "fmt"

func main() { // 0	1  2  3   4   5   6   7   8   9
	req := []int{5, 7, 9, 10, 11, 12, 13, 14, 28, 42}

	seq := []int{}
	// tag := 0
	for i := 1; i < len(req); i++ {
		seq = append(seq, req[i]-req[i-1])
	}

	fmt.Println(seq)
	type Set struct {
		begin int
		end   int
	}
	result := []Set{}
	for i := 0; i < len(seq)-1; i++ {
		j := i + 1
		if seq[i] == seq[j] {
			tmp := Set{
				begin: i,
				end:   j + 1,
			}
			result = append(result, tmp)
		}
	}

	for i := 0; i < len(seq)-1; i++ {
		count := 0
		for j := i + 1; j < len(seq); j++ {
			if seq[i] == seq[j] {
				count++
			} else {
				break
			}
			if count >= 2 {
				tmp := Set{
					begin: i,
					end:   j + 1,
				}
				result = append(result, tmp)
			}
		}
	}

	fmt.Println(result)

}
