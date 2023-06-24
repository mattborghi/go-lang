package main

import "fmt"

func main() {
	// First way of declaring a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors) // prints map[blue:#0000ff green:#00ff00 red:#ff0000]

	// Second way of declaring a map
	var colors2 map[string]string
	fmt.Println(colors2) // print empty initialization map[]

	// Third way of declaring a map
	colors3 := make(map[string]string)
	colors3["white"] = "#ffffff"
	fmt.Println(colors3) // print map[white:#ffffff]

	delete(colors3, "white")
	fmt.Println(colors3) // print empty initialization map[]

	// print map looped
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Color %s with hex %s\n", color, hex)
	}
}
