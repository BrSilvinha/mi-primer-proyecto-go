package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Presiona Enter para continuar...")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
