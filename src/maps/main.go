package main

import "fmt"

func main() {
	var colorPool map[string]string
	anotherColors := make(map[string]string)
	fmt.Println(colorPool)
	fmt.Println(anotherColors)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)
	delete(colors, "green")
	fmt.Println(colors)

	colors["green"] = "#00ff00"
	fmt.Println(colors)
	print(colors)
}

func print(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
