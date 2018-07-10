package main

import (
	"fmt"
	"github.com/ikawaha/kagome/tokenizer"
	"strings"
)

func main() {
	t := tokenizer.New()
	tokens := t.Tokenize("寿司が食べたい。")
	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			fmt.Printf("%s\n", token.Surface)
			continue
		}
		features := strings.Join(token.Features(), ",")
		fmt.Printf("%s\t%v\n", token.Surface, features)
	}
}
