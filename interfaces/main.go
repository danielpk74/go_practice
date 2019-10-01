package main

import "fmt"

type bot interface {
	getGreeting() string
	getBotVersion() float64
}
type englishBot struct{}
type spanishBot struct{}
type rusianBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	rb := rusianBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(rb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// This is part of the problem that interfaces wants to solve
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (spanishBot) getGreeting() string {
	// Very custom logic for geneating an english greeting
	return "Hi There!"
}

func (englishBot) getGreeting() string {
	// Very custom logic for geneating an english greeting
	return "Hola!"
}

func (rusianBot) getGreeting() string {
	// Very custom logic for geneating an english greeting
	return "привет!"
}

func (spanishBot) getBotVersion() float64 {
	// Very custom logic for geneating an english greeting
	return 1.1
}

func (englishBot) getBotVersion() float64 {
	// Very custom logic for geneating an english greeting
	return 1.1
}

func (rusianBot) getBotVersion() float64 {
	// Very custom logic for geneating an english greeting
	return 1.1
}
