package main

import (
	"sort"
)

func Manipulate(input string) []string {
	var runeInput = []rune(input)
	var result []string

	if len(runeInput) < 2 {
		result = append(result, string(runeInput))
		return result
	}

	for i := 0; i < len(runeInput); i++ {
		result = append(result, swapRune(runeInput)...)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return removeDuplicate(result)
}

func swapRune(runeInput []rune) (result []string) {
	for i := 0; i < len(runeInput); i++ {
		if i == len(runeInput)-1 {
			return result
		}
		runeInput[i+1], runeInput[i] = runeInput[i], runeInput[i+1]
		result = append(result, string(runeInput))
	}
	return result
}

func removeDuplicate[T string | int](sliceList []T) []T {
	var allKeys = make(map[T]bool)
	var list = []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

//func main() {
//	var input = "a"
//	fmt.Println(Manipulate(input))
//
//	input = "ab"
//	fmt.Println(Manipulate(input))
//
//	input = "abc"
//	fmt.Println(Manipulate(input))
//
//	input = "aabb"
//	fmt.Println(Manipulate(input))
//}
