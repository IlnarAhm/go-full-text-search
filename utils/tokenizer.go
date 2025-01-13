package utils

import (
	"strings"
	"unicode"
)

var (
	done            chan int
	pipeline        chan []string
	lowercaseStream <-chan []string
	stopWordStream  <-chan []string
	stemmerStream   <-chan []string
)

func RunPipeline() {
	pipeline = make(chan []string)
	done = make(chan int)
	lowercaseStream = lowercaseFilterStream(done, pipeline)
	stopWordStream = stopWordFilterStream(done, lowercaseStream)
	stemmerStream = stemmerFilterStream(done, stopWordStream)
}

func ClosePipeline() {
	close(pipeline)
	close(done)
}

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func analyze(text string) []string {
	pipeline <- tokenize(text)

	return <-stemmerStream
}
