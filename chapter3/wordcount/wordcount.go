// Sample program to show how to show you how to briefly work with io.
package main

import (
	"fmt"
	"io/ioutil"
	//"os"

	"../../chapter3/words"
)

// main is the entry point for the application.
func main() {
	//filename := os.Args[1]
	filename := "/Users/gongyao/workspace/go_code1/code/chapter3/wordcount/gowords.txt"

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents)
	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
