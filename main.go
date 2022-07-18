package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text := os.Args[1]
	fmt.Println(len(strings.Fields(text)))
}
