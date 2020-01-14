package quotes_test

import (
	"testing"

	"github.com/jpmcb/wisdom/quotes"
)

func TestRandomQuote(t *testing.T) {
	t.Run("it returns a random quote", func(t *testing.T) {
		got := quotes.RandomQuote()

		if len(got) <= 0 {
			t.Errorf("Returned empty random quote")
		}
	})
}
