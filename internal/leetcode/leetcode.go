package leetcode

func similarPairs(words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if haveSameUniqueChars(words[i], words[j]) {
				count++
			}
		}
	}
	return count
}

func haveSameUniqueChars(word1, word2 string) bool {

	set1 := make(map[rune]bool)
	set2 := make(map[rune]bool)
	for _, ch := range word1 {
		set1[ch] = true
	}
	for _, ch := range word2 {
		set2[ch] = true
	}
	if len(set1) != len(set2) {
		return false
	}
	for ch := range set1 {
		if !set2[ch] {
			return false
		}
	}
	return true
}
