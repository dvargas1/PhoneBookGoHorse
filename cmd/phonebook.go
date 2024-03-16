package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type phonebook struct {
	contacts []contact
}

func (p *phonebook) addContact() {
	var c contact

	fmt.Println("Add new contact:")

	fmt.Println("Please add the contact Frist Name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.FirstName = scanner.Text()

	fmt.Println("Please add the contact Last name:")
	scanner.Scan()
	c.LastName = scanner.Text()

	fmt.Println("Please add the contact Age:")
	scanner.Scan()
	c.Age, err = strconv.Atoi(scanner.Text())
	if err != nil || c.Age <= 0 {
		fmt.Printf("Not a valid age, Try again\n")
		return
	}

	fmt.Println("Please add the contact Description:")
	scanner.Scan()
	c.Description = scanner.Text()

	fmt.Println("Finally, please add the contact Phone number:")
	scanner.Scan()
	c.PhoneNumber = scanner.Text()

	p.contacts = append(p.contacts, c)
}

func (p phonebook) printContacts() {
	for _, c := range p.contacts {
		fmt.Println(c)
	}
}
