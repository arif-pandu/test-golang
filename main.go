package main

import "fmt"

func main() {
	//
	appName := "Go Booking"
	const allTicket int = 50
	var remainingTicket uint = 50

	fmt.Println("===========================")
	fmt.Println("== Welcome To", appName, "==")
	fmt.Println("===========================")
	println("We Have quota of total", allTicket, "ticket")
	println("And remaining ticket :", remainingTicket)

	bookings := []string{}
	var userName string
	var userTicket uint
	var email string

	fmt.Println("Enter your name :")
	fmt.Scan(&userName)

	fmt.Println("How many ticket :")
	fmt.Scan(&userTicket)

	fmt.Println("Your Email addres : :")
	fmt.Scan(&email)

	remainingTicket = remainingTicket - userTicket

	bookings = append(bookings, userName)

	fmt.Println("Your name is", userName, "booking for", userTicket, "tickets, we'll inform you through", email)

	fmt.Println("Remaining Ticket is", remainingTicket)

	fmt.Println("The Slices", bookings)
	fmt.Println("Slices Length", len(bookings))

}
