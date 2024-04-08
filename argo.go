package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	client.SetLanguage("fra")
	// client.SetImage("001-helloworld.png")
	client.SetImage("sample.jpg")
	text, _ := client.Text()
	fmt.Println(text)
	fmt.Println("text")
	// Hello, World!
}