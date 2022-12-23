package main

import (
	"fmt"
	"strconv"
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

	// Datatypes processing
	// basic(&numberOfTicketsRemaining) // Passing by reference

	// For loops:
	// loops(&numberOfTicketsRemaining)

	// Maps
	// maps(&numberOfTicketsRemaining)

	// Structures
	structs(&numberOfTicketsRemaining)

}

type users struct { // type used to create a custom type
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func structs(numberOfTicketsRemaining *int) {
	var bookings = make([]users, 0)
	var firstName, lastName, email string
	var tickets int
	for *numberOfTicketsRemaining > 0 {

		fmt.Scan(&firstName)
		fmt.Scan(&lastName)
		fmt.Scan(&email)
		fmt.Scan(&tickets)
		*numberOfTicketsRemaining -= tickets
		var userData = users{
			firstName:       firstName,
			lastName:        lastName,
			email:           email,
			numberOfTickets: tickets,
		}
		fmt.Println("The map is ", userData)
		bookings = append(bookings, userData)
	}
	fmt.Println(bookings)
}

func maps(numberOfTicketsRemaining *int) {
	// Multiple key values for input
	var firstName, lastName, email string
	var tickets int
	var bookings = make([]map[string]string, 0) // need to initialize size of the slice.

	for *numberOfTicketsRemaining > 0 {
		var userData = make(map[string]string) // key and value; make will create an empty object
		fmt.Scan(&firstName)
		fmt.Scan(&lastName)
		fmt.Scan(&email)
		fmt.Scan(&tickets)
		*numberOfTicketsRemaining -= tickets
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["email"] = email
		// Convert the number of tickets from int to string
		userData["tickets"] = strconv.FormatInt(int64(tickets), 10)
		fmt.Println("The map is ", userData)
		bookings = append(bookings, userData)
	}
	fmt.Println(bookings)

}

func loops(numberOfTicketsRemaining *int) {
	// Arrays and slices
	// var bookings = [50]string{"User1", "User2"} // Size of the array and datatype, can add some default values, else just add {} for empty.

	// Alternative Syntax
	// var bookings [50]string
	// Dynamic array size -> Slice; abstraction of array.
	var bookings []string
	// Can run these infinite times with a loop = for {} or for some condition
	for *numberOfTicketsRemaining > 0 {
		var firstName, lastName string
		fmt.Println("Enter first and last name")
		fmt.Scan(&firstName)
		fmt.Scan(&lastName)
		// Validate user input
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		if !isValidName {
			fmt.Println("Invalid input")
			continue
		}
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
		if *numberOfTicketsRemaining == 0 {
			fmt.Println("All tickets sold out")
			break
		}
		*numberOfTicketsRemaining -= 1
	}
}

func basic(numberOfTicketsRemaining *int) {
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
	if numTickets > *numberOfTicketsRemaining {
		fmt.Println("Impossible")
	} else if numTickets < 0 {
		fmt.Println("Again, impossible")
	} else {
		*numberOfTicketsRemaining -= numTickets
		fmt.Println("Number of tickets left", *numberOfTicketsRemaining)
	}
}
