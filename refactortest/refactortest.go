package main

import (
	"strings"
)
func findFirstStringInBracket(str string) string {
	if (len(str) > 0) {
		indexFirstBracketFound := strings.Index(str,"(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket,")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound-1])
			}else{
				return ""
			}
		}else{
			return ""
		}
	}else{
		return ""
	}
	return ""
}
func main() {

	findFirstStringInBracket("tes")

}
