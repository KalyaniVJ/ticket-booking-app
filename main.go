package main

import "fmt"

func main() {
	var eventname string = "Ticket Booking Application"
	fmt.Printf("Welcome to %v\n", eventname)
	var totalTickets int = 50
	fmt.Printf("We are having Total of %v tickets.\n", totalTickets)
	fmt.Printf("Please Enter the below details to book a ticket:\n")

	var firstName string
	fmt.Printf("Enter your first name:\n")
	fmt.Scan(&firstName)
	var lastName string
	fmt.Printf("Enter your last name:\n")
	fmt.Scan(&lastName)
	var email string
	fmt.Printf("Enter your email address:\n")
	fmt.Scan(&email)

	var bookings = [50]string{}
	bookings[0] = firstName + "" + lastName + "" + email + ""
	fmt.Printf("All bookings:%v\n", bookings)
	fmt.Printf("First customer is:%v\n", bookings[0])
	fmt.Printf("Array type:%T\n", bookings)
	fmt.Printf("Array length:%v\n", len(bookings))

	var userTickets int
	fmt.Printf("Enter the number of Tickets you want to book:\n")
	fmt.Scan(&userTickets)

	totalTickets = totalTickets - userTickets
	fmt.Printf("The number of Available tickets are:%v\n", totalTickets)

	fmt.Printf("Thank you for booking %v tickets! . Please visit again.\n", userTickets)

}
