package main

import (
	"fmt"
)

func main() {
	helloList := []string{
		"Witaj, świecie",
		"Καλημέρα κόσμε",
		"こんにちは世界",
		"سلام دنیا",
		"Привет, мир",
	}
	fmt.Println(len(helloList))
	fmt.Println(helloList[len(helloList)-1])
	// fmt.Println(helloList[len(helloList)])
}
