package main

import (
	"fmt"

	go_english_words "github.com/pefish/go-english-words"
)

func main() {
	fmt.Println(go_english_words.IsWord("test"))
	fmt.Println(go_english_words.IsWord("test1"))
	fmt.Println(go_english_words.IsWord("test2"))
	fmt.Println(go_english_words.IsWord("my"))
	fmt.Println(go_english_words.IsWord("you"))
	fmt.Println(go_english_words.IsWord("fuck"))

}
