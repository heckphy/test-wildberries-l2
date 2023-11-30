package main

import (
	"os"

	"github.com/heckphy/wb-l2/develop/dev09/wget"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// accepting url from user
	rawURL := os.Args[1]
	//

	c := wget.NewClient()

	err := c.Get(rawURL, wget.Options{
		Recursive: true,
	})

	if err != nil {
		panic(err)
	}
}
