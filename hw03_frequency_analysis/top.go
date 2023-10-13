package hw03frequencyanalysis

import (
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Top10(str string) (result []string) {
	words := make(map[string]int)
	var cnt int
	var runs []rune

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		runs = append(runs, r)

		str = str[size:]
	}

	for _, symbol := range runs {
		cnt++
		if unicode.IsSpace(symbol) {
			var word string
			for j := 0; j < cnt; j++ {
				if !unicode.IsSpace(runs[j]) && !unicode.IsPunct(runs[j]) && !unicode.IsMark(runs[j]) && !unicode.IsPunct(runs[j]) {
					word += string([]rune(runs)[j])
				} /*  else if string(runs[j]) == "-" {
					word += string([]rune(runs)[j])
				} */
			}
			runs = runs[cnt:]
			cnt = 0
			if len([]rune(word)) != 0 {
				val, ok := words[strings.ToLower(word)]
				if ok {
					words[strings.ToLower(word)] = val + 1
				} else {
					words[strings.ToLower(word)] = 1
				}
			}
		}
	}

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
