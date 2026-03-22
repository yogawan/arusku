package main

import (
	"log"

	"belajar-golang/internal/app/belajargolang"
)

func main() {
	if err := belajargolang.Run(); err != nil {
		log.Fatal(err)
	}
}
