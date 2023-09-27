package main

import (
	"fmt"
	"strings"
)

func main() {
	var phrases = []string{
		"radar",
		"subi no onibus",
		"socorram me subi no onibus em marrocos",
		"roma ate cubanos",
		"No final da noite, Oto voltou para casa em seu Civic, pensando em todas as aventuras malucas que havia vivido em Natan e arara. Ele sorriu, sabendo que sua vida era cheia de palíndromos e diversão.",
		"Era uma vez, em uma cidade chamada Natan, um rapaz chamado Oto. Oto tinha um carro antigo, um Civic, que ele adorava. Todos os dias, ele dirigia seu Civic pelas ruas de Natan, olhando para o radar de trânsito para evitar multas.",
	}

	totals := run(phrases)
	fmt.Println(totals)
}

func run(phrases []string) map[string]int64{
	totals := make(map[string]int64)
	for _, phrase := range phrases {
		pLower := strings.ToLower(phrase)
		pWithoutSpaces := strings.ReplaceAll(pLower, " ", "")

		if isPalindrome(pWithoutSpaces) {
			totals[pWithoutSpaces]++
			fmt.Println("The whole phrase is a palindrome:", phrase)
		}

		for _, word := range strings.Fields(pLower) {
			if isPalindrome(word) {
				totals[word]++
				fmt.Println("The word", word, "is a palindrome")
			}
		}

	}

	return totals
}

func isPalindrome( s string ) bool {
	length := len(s)

	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}