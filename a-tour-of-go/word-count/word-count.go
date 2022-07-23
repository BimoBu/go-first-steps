/*

Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.

*/

package wordCount

import (
	"fmt"
	"strings"
	// "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	wordCountMap := make(map[string]int)

	for _, word := range words {
		count := wordCountMap[word]
		wordCountMap[word] = count + 1
	}

	return wordCountMap
}

func Test() {
	fmt.Println("Word Count needs to be executed in the Go Playground")
	// wc.Test(WordCount)
}
