package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	phonebook := phonebook{}

	for {
		fmt.Println("Enter your command:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}
		parseInput(scanner.Text(), &phonebook)
	}
}
