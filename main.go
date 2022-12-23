package main

import (
	"fmt"
	"strings"
)

// This is the entry point for the file. Need to put things under a function
func main() {
	var nameApp = "Go Conference Booking Application" // Compiler error when variable, package is unused!
	fmt.Println("Welcome to the ", nameApp, ". Get your tickets here.")
	const totalNumberOfTickets = 50
	var numberOfTicketsRemaining = 50
	// fmt.Println("Currently we have ", numberOfTicketsRemaining, "remaining")
	fmt.Printf("Currently we have %v remaining tickets\n", numberOfTicketsRemaining)

	/* Data Types */
	// Strings, Integers, Booleans, Maps, Arrays

	// Go can infer the datatype when you provide a value, if not; explicitly provide a type.
	var userName string // variable userName of type String
	fmt.Scanf("%v", &userName)
	fmt.Printf("Hello %v\n", userName)

	// To declare a variable - can also use the := syntax
	var age, numTickets int
	fmt.Scan(&age)
	fmt.Println("The user's age is", age)

	fmt.Println("Enter the number of tickets")
	fmt.Scan(&numTickets)
	numberOfTicketsRemaining -= numTickets
	fmt.Println("Number of tickets left", numberOfTicketsRemaining)

	// Arrays and slices
	// var bookings = [50]string{"User1", "User2"} // Size of the array and datatype, can add some default values, else just add {} for empty.

	// Alternative Syntax
	// var bookings [50]string

	// Dynamic array size -> Slice; abstraction of array.
	var bookings []string
	// Can run these infinite times with a loop
	for {
		var firstName, lastName string
		fmt.Println("Enter first and last name")
		fmt.Scan(&firstName)
		fmt.Scan(&lastName)
		// bookings[0] = firstName + " " + lastName # This is for array syntax
		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("First Name %v LastName %v Array Length %v\n", firstName, lastName, len(bookings))
		// Using a for-each
		var firstNames []string
		for _, val := range bookings { // _ used for blank identifiers that we dont want to use.
			var name = strings.Fields(val) // slices into first and last name based on space
			firstNames = append(firstNames, name[0])
		}
		fmt.Println("The booked users so far are", firstNames)

		// Now should terminate the application if all tickets sold
		if numberOfTicketsRemaining == 0 {
			fmt.Println("All tickets sold out")
			break
		}
		numberOfTicketsRemaining -= 1
	}
}
