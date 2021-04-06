package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fff657",
	}
	colors["yellow"] = "#lowenfn"
	delete(colors, "red")
	delete(colors, "yellow")
	emre(colors)

}

func emre(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for color", color, "is", hex)
	}
}
