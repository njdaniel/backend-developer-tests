package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "error") {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
