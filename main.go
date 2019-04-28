package main

import (
	"fmt"
	"nlp/web"
)

func main() {
	fmt.Println("Launch in 127.0.0.1:8080 or localhost:8080")
	web.Run()

	// words, tler, grams, tags := tler.English("I am cooking chicken")
	// fmt.Println(words)
	// fmt.Println(tler)
	// fmt.Println(grams)
	// fmt.Println(tags)
}
