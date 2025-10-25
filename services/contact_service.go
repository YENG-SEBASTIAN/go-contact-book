package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"contact-book/models"
)

var contacts []models.Contact
const dataFile = "data/contacts.txt"

// LoadContacts reads existing contacts from the file
func LoadContacts() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("ℹ️ No previous contacts found — starting fresh.")
			return
		}
		fmt.Println("Error opening contacts file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) == 3 {
			contact := models.Contact{
				Name:  strings.TrimSpace(parts[0]),
				Phone: strings.TrimSpace(parts[1]),
				Email: strings.TrimSpace(parts[2]),
			}
			contacts = append(contacts, contact)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("❌ Error reading contacts file:", err)
	}
}

// SaveContacts writes all contacts to the file
func SaveContacts() error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, c := range contacts {
		line := fmt.Sprintf("%s | %s | %s\n", c.Name, c.Phone, c.Email)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}

// AddContact adds a new contact and saves to file
func AddContact(name, phone, email string) {
	newContact := models.Contact{Name: name, Phone: phone, Email: email}
	contacts = append(contacts, newContact)

	err := SaveContacts()
	if err != nil {
		fmt.Println("❌ Error saving contact:", err)
		return
	}

	fmt.Println("✅ Contact added successfully!")
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
