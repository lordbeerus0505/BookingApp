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
}
