package tasks

import "fmt"

func CheckDuplicates(nums []int) bool {
	dict := map[int]bool{}

	for _, num := range nums {
		if _, ok := dict[num]; ok {
			return true
		}
		dict[num] = true
	}
	return false
}

func CheckIfAnagrams(s, t string) bool {
	letters := map[int32]int{}

	for _, letter := range s {
		_, ok := letters[letter]
		if ok {
			letters[letter] += 1
		} else {
			letters[letter] = 1
		}
	}

	for _, letter := range t {
		_, ok := letters[letter]
		if ok {
			letters[letter] -= 1
		} else {
			return false
		}
	}

	for _, value := range letters {
		if value != 0 {
			return false
		}
	}
	return true
}

type Void struct {
}

func CreateSet() {
	set := map[int]Void{}
	fmt.Println(set)
}
