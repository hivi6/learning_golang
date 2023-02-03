package main

import "fmt"

func main() {
	// define map
	emails := make(map[string]string)
	// assign key-values
	emails["Bob"] = "bob@domain.com"
	emails["Ken"] = "ken@domain.com"
	fmt.Println(emails, emails["Bob"], emails["Unnamed"], len(emails))
	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// initialize map
	newEmails := map[string]string{"Bob": "bob@domain.com", "Ken": "ken@domain.com"}
	fmt.Println(newEmails)
}
