package main

import "fmt"

type color map[string]string

func main() {
	colors := color{
		"red":   "ff0000",
		"green": "4bf745",
		"white": "ffffff",
		"black": "000000",
	}

	colors.printMap()
	fmt.Println(colors)
	// colorName := "red"

	// colors := make(map[string]string)

	// colors["white"] = "ffffff"
	// colors["black"] = "000000"

	// delete(colors, "white")
}

func (c color) printMap() {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
