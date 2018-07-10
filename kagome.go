package main

import (
	"bufio"
	"fmt"
	"github.com/ikawaha/kagome/tokenizer"
	"os"
	"strings"
)

func main() {
	fp, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		t := tokenizer.New()
		tokens := t.Tokenize(scanner.Text())
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
}
