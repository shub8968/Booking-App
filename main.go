package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"
	const ticket = 50
	var remainingTickets = ticket
	fmt.Printf("Welcome to  %v ticket Booking System\n", conferenceName)
	fmt.Println("Total Number of Tickets ", ticket)
	fmt.Println("Number of tickets remaining", remainingTickets)

	var userName string
	var userTickets int
	fmt.Scan(&userName)
	fmt.Scan(&userTickets)
	fmt.Printf("User %v booked %v tickets \n", userName, userTickets)
}
