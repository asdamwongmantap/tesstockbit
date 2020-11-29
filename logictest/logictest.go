package main

import (
	"fmt"
	"sort"
	"strings"
)
var data = []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}

func createanagram(){

	//make map first before extract
	list := make(map[string][]string)

	//extract data
	for _, word := range data {
		s := strings.Split(word, "")
		sort.Strings(s)

		key := strings.Join(s, "")
		//add extracted data to array
		list[key] = append(list[key], word)
	}

	//looping array extracted data
	for _, words := range list {
		for _, w := range words {
			fmt.Print(w," ")
		}

		fmt.Println()
	}

}
func main() {

	createanagram()

}
