package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func lowercaseFilterStream(done <-chan int, ch <-chan []string) <-chan []string {
	filterStream := make(chan []string)

	go func() {
		defer close(filterStream)

		for {
			select {
			case <-done:
				return
			case tokens := <-ch:
				r := lowercaseFilter(tokens)

				go func() {
					filterStream <- r
				}()
			}
		}
	}()

	return filterStream
}

func stopWordFilter(tokens []string) []string {
	stopWords := map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}
	r := make([]string, len(tokens))

	for _, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			r = append(r, token)
		}
	}

	return r
}

func stopWordFilterStream(done <-chan int, ch <-chan []string) <-chan []string {
	stopWordStream := make(chan []string)

	go func() {
		defer close(stopWordStream)

		for {
			select {
			case <-done:
				return
			case tokens := <-ch:
				r := stopWordFilter(tokens)

				go func() {
					stopWordStream <- r
				}()
			}
		}
	}()

	return stopWordStream
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))

	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}

	return r
}

func stemmerFilterStream(done <-chan int, ch <-chan []string) <-chan []string {
	stemmerStream := make(chan []string)

	go func() {
		defer close(stemmerStream)

		for {
			select {
			case <-done:
				return
			case tokens := <-ch:
				r := stemmerFilter(tokens)

				go func() {
					stemmerStream <- r
				}()
			}
		}
	}()

	return stemmerStream
}
