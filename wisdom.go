package main

import (
	"fmt"

	"github.com/jpmcb/wisdom/quotes"
)

func main() {
	q := quotes.RandomQuote()

	fmt.Printf("%s\n", q)
}
