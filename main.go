package main

import (
	"bufio"
	"fmt"
	"os"

	"contact-book/services"
	"contact-book/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin) // shared reader for all input

	for {
		fmt.Println("\n=== Contact Book CLI ===")
		fmt.Println("1. Add Contact")
		fmt.Println("2. List Contacts")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		// 🧹 Clean the leftover newline
		reader.ReadString('\n')

		switch choice {
		case 1:
			name := utils.ReadLine("Enter Name: ")
			phone := utils.ReadLine("Enter Phone: ")
			email := utils.ReadLine("Enter Email: ")
			services.AddContact(name, phone, email)

		case 2:
			services.ListContacts()

		case 3:
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("3Invalid choice. Please try again.")
		}
	}
}
