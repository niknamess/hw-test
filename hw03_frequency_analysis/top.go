package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Top10(str string) (result []string) {
	words := make(map[string]int)
	var runs []rune
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		if unicode.IsSpace(r) {
			word := CheckWord(runs)
			val, ok := words[word]
			if ok {
				words[word] = val + 1
			} else {
				words[word] = 1
			}
			runs = nil
		}
		runs = append(runs, r)
		str = str[size:]
	}
	delete(words, "")
	resmap := make(map[int][]string)
	for k := range words {
		resmap[words[k]] = append(resmap[words[k]], k)
	}
	keys := make([]int, 0, len(resmap))
	for k := range resmap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for _, k := range keys {
		sort.Strings(resmap[k])
		result = append(result, resmap[k]...)
		if len(result) >= 10 {
			return result[:10]
		}
	}
	return result
}

func CheckWord(runs []rune) (word string) {
	for j := 0; j < len(runs); j++ {
		if !unicode.IsSpace(runs[j]) && !unicode.IsPunct(runs[j]) && !unicode.IsMark(runs[j]) {
			word += string(runs[j])
		}
	}
	return strings.ToLower(word)
}
