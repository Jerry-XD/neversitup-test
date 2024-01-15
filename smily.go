package main

import (
	"strings"
)

var smilyFace = []string{":)", ";)", ":-)", ";-)", ":~)", ";~)", ":D", ";D", ":-D", ";-D", ":~D", ";~D"}

func CountSmilyFace(text []string) int {
	var count = 0
	for _, text := range text {
		for _, smilyFace := range smilyFace {
			if strings.Contains(smilyFace, text) {
				count += 1
				continue
			}
		}
	}
	return count
}

//func main() {
//	var input = []string{":)", ";(", ";}", ":-D"}
//	var res = CountSmilyFace(input)
//	fmt.Println(res)
//
//	input = []string{";D", ":-(", ":-)", ";~)"}
//	res = CountSmilyFace(input)
//	fmt.Println(res)
//
//	input = []string{";]", ":[", ";*", ":$", ";-D"}
//	res = CountSmilyFace(input)
//	fmt.Println(res)
//}
