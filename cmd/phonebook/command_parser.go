package main

import (
	"fmt"
	"os"
)

func parseInput(input string, phbook *phonebook) {
	switch input {
	case "ADD":
		phbook.addContact()
		fmt.Println("Add command...")
	case "REMOVE":
		fmt.Println("Remove command...")
	case "LIST":
		phbook.printContacts()
		return
	case "QUIT":
		fmt.Println("Quitting...")
		os.Exit(0)
	case "HELP", "h":
		fmt.Println("Print help command...")
	default:
		fmt.Printf("Command not recognized %q ignoring input.\nfor help enter 'HELP' or 'h'\n", input)
	}
}
