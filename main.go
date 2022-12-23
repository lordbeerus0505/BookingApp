package main

import "fmt"

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

}
