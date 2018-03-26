package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func randBool() bool {
	return rand.Intn(2) == 0
}

func duplicateVowel(word string, i int) string {
	return word[:i+1] + word[i:]
}

func removeVowel(word string, i int) string {
	return word[:i] + word[i+1:]
}

func main() {
	rand.Seed(time.Now().UTC().Unix())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := s.Text()
		if randBool() {
			fmt.Println(word)
			continue
		}

		var vI = -1
		for i, char := range word {
			switch char {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				if randBool() {
					vI = i
				}
			}
		}

		if vI == -1 {
			fmt.Println(word)
		} else if randBool() {
			fmt.Println(duplicateVowel(word, vI))
		} else {
			fmt.Println(removeVowel(word, vI))
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalln(err)
	}
}
