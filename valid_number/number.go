package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isNumber("3.3e-2.3w"))
}

const NUMBER_STR = "0123456789"

func recurse(s_chars []rune, exp_seen bool) bool {
	if s_chars[0] == '+' || s_chars[0] == '-' {
		s_chars = s_chars[1:]
		if s_chars[0] == '+' || s_chars[0] == '-' {
			return false
		}
	}
	dot_seen := false

	followed_by_num := false
	i := 0
	for i < len(s_chars) {
		ch := s_chars[i]
		if !strings.Contains(NUMBER_STR, string(ch)) {
			if ch == '.' {
				if dot_seen {
					return false
				}
				if i < len(s_chars) && !strings.Contains(NUMBER_STR, string(s_chars[i+1])) {
					return false
				}
				dot_seen = true

			} else if strings.Contains("eE", string(ch)) {
				if !followed_by_num {
					return false
				}
				return recurse(s_chars[i+1:], true)
			} else {
				return false
			}
		}
		followed_by_num = true
		i++
	}

	return true
}

func isNumber(s string) bool {
	s_chars := []rune(s)

	return recurse(s_chars, false)

}
