package main

import (
	"fmt"
	"strings"
)

var phrases = []string{
	"radar",
	"subi no onibus",
	"socorram me subi no onibus em marrocos",
	"roma ate cubanos",
	"No final da noite, Oto voltou para casa em seu Civic, pensando em todas as aventuras malucas que havia vivido em Natan e arara. Ele sorriu, sabendo que sua vida era cheia de palíndromos e diversão.",
	"Era uma vez, em uma cidade chamada Natan, um rapaz chamado Oto. Oto tinha um carro antigo, um Civic, que ele adorava. Todos os dias, ele dirigia seu Civic pelas ruas de Natan, olhando para o radar de trânsito para evitar multas.",
}

var totals map[string]int64

func main() {
	totals = make(map[string]int64)

	for _, phrase := range phrases {
		phraseRunes := []rune(strings.ToLower(phrase))

		if isPalindrome(phraseRunes) {
			totals[string(phraseRunes)]++
			fmt.Println("The whole phrase is a palindrome:", phrase)
		}

		for _, word := range strings.Fields(phrase) {
			wordRunes := []rune(strings.ToLower(word))
			if isPalindrome(wordRunes) {
				totals[string(wordRunes)]++
				fmt.Println("The word", word, "is a palindrome")
			}
		}
	}

	fmt.Println(totals)
}

func isPalindrome(runes []rune) bool {
	length := len(runes)

	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-i-1] {
			return false
		}
	}
	return true
}