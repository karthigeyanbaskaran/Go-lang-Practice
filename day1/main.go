package main // Package declaration, necessary for every Go program. Everything in Go is organized into packages.

import "fmt" // Importing the fmt package for formatted I/O operations.

func main() { // Main function: the entry point of the Go program.
	totalTickets := 50                                             // Using shorthand declaration for variables. Equivalent to var totalTickets = 50.
	var remainingTickets uint = 50                                 // no negative number will get at.
	fmt.Println("Hello Peoples, Welcome to the Online booking's")  // Print a welcome message.
	fmt.Printf("Total tickets are %v\n", totalTickets)             // Print total tickets using formatted string.
	fmt.Printf("The remaining tickets are %v\n", remainingTickets) // Print remaining tickets using formatted string.

	for {
		// Declare variables for user input.
		var firstName string
		var lastName string
		var email string
		var usertickets uint
		//var bookings []string
		bookings := []string{}
		// Prompt the user for their first name.
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName) // Scan for input and store it in firstName.

		// Prompt the user for their last name.
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName) // Scan for input and store it in lastName.

		// Prompt the user for their email.
		fmt.Println("Enter your email ID:")
		fmt.Scan(&email) // Scan for input and store it in email.

		// Prompt the user for the number of tickets.
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&usertickets) // Scan for input and store it in tickets. Using & to get the variable's memory address.

		bookings = append(bookings, firstName+" "+lastName)
		remainingTickets = remainingTickets - usertickets
		// Print a confirmation message with the user's details.
		fmt.Printf("Hi %v %v, thank you for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, usertickets, email)
		fmt.Printf("The Remaining Tickets are %v\n", remainingTickets)
		fmt.Println(bookings)
	}
}
