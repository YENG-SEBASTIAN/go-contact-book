package services

import (
	"fmt"
	"contact-book/models"
)

var contacts []models.Contact

// adds a new contact to the list
func AddContact(name, phone, email string) {
	newContact := models.Contact{Name: name, Phone: phone, Email: email}
	contacts = append(contacts, newContact)
	fmt.Println("Contact added successfully!")
}

// ListContacts displays all stored contacts
func ListContacts() {
	if len(contacts) == 0 {
		fmt.Println("No contacts yet.")
		return
	}

	fmt.Println("\n📒 Contact List:")
	for i, c := range contacts {
		fmt.Printf("%d. ", i+1)
		c.Display()
	}
}
