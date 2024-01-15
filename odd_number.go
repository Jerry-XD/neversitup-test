package main

func FindOddNumber(input []int) int {
	var countInput = make(map[int]int)
	for i := 0; i < len(input); i++ {
		countInput[input[i]] += 1
	}

	input = removeDuplicate(input)

	for i := 0; i < len(input); i++ {
		if (countInput[input[i]])%2 != 0 {
			return input[i]
		}
	}

	return 0
}

//func removeDuplicate[T string | int](sliceList []T) []T {
//	var allKeys = make(map[T]bool)
//	var list = []T{}
//	for _, item := range sliceList {
//		if _, value := allKeys[item]; !value {
//			allKeys[item] = true
//			list = append(list, item)
//		}
//	}
//	return list
//}

//func main() {
//	var input = []int{7}
//	fmt.Println(FindOddNumber(input))
//
//	input = []int{0}
//	fmt.Println(FindOddNumber(input))
//
//	input = []int{1, 1, 2}
//	fmt.Println(FindOddNumber(input))
//
//	input = []int{0, 1, 0, 1, 0}
//	fmt.Println(FindOddNumber(input))
//
//	input = []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
//	fmt.Println(FindOddNumber(input))
//}
