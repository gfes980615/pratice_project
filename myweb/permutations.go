package main

type permValue struct {
	str            []string
	allCombination []string
}

// Permutations ...
func Permutations(str string) interface{} {
	tmp := &permValue{}
	for _, s := range str {
		tmp.str = append(tmp.str, string(s))
	}

	tmp.perm(0, len(tmp.str))

	return tmp.allCombination
}

func (s *permValue) perm(offset int, n int) {
	if offset == n {
		tmp := ""
		for j := 0; j < n; j++ {
			tmp += s.str[j]
		}
		s.allCombination = append(s.allCombination, tmp)
	} else {
		for j := offset; j < n; j++ {
			s.swap(&s.str[offset], &s.str[j])
			s.perm(offset+1, n)
			s.swap(&s.str[offset], &s.str[j])
		}
	}
}

func (s *permValue) swap(a, b *string) {
	tmp := *a
	*a = *b
	*b = tmp
}
