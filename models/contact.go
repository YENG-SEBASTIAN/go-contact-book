package models

import "fmt"

// Contact represents a contact entry/ a class in python
type Contact struct {
	Name  string
	Phone string
	Email string
}

// Display prints a formatted contact/ a method in python
func (c Contact) Display() {
	fmt.Printf("Name: %s | Phone: %s | Email: %s\n", c.Name, c.Phone, c.Email)
}
