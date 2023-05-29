package main

import (
	"fmt"

	"github.com/pokpitch/packagedemo/book"
)

func main() {

	b := book.New()
	fmt.Printf("%T %v\n", b, b)

	b.Name = "Pokpitch"
	fmt.Printf("%T %v\n", b, b)

}
