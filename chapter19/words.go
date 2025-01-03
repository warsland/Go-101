package main

import (
	"fmt"
	"strings"
)

var words = `As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshineto the solemn splendour of twilight in which he woalked. Whenever he felt able he ran again; the ground continued soft and springy, covered whith the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—— except the fact of wandering unprosvisioned and alone in a forest of unknown vegetation thousands or mollions of miles beyind the reach or knowledge of man.`

func trans(words string) []string {
	wordString := strings.Fields(strings.ToLower(words))
	wordSlice := make([]string, 0, len(wordString))
	for _, word := range wordString {
		word = strings.Trim(word, ",")
		word = strings.Trim(word, ".")
		word = strings.Trim(word, ";")
		word = strings.Trim(word, "——")
		wordSlice = append(wordSlice, word)
	}
	return wordSlice
}

func main() {
	wordString := trans(words)
	wordTimes := make(map[string]int, len(words))
	for _, word := range wordString {
		wordTimes[word]++
	}
	for word, times := range wordTimes {
		if times > 1 {
			fmt.Printf("%v : %v times\n", word, times)
		}
	}
}
