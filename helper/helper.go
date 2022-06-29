package helper

import "strings"

//When you're defining functions in the package, capitalize the first letter of the function name to make the function exported

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	isValidEmail := strings.Contains(email, "@")

	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	// isValidCity := (city == "singapore") || (city == "London")

	// if userTickets == 0 {
	// 	fmt.Printf("Booking cycle end. Thank ypu ^_^")
	// 	break

	// }
	return isValidName, isValidEmail, isValidTicketNumber

}

//You can also export variables by capitalizing the first letter of the variable name
var Myvar = "This is a global variable across all packages that imports the helper package"
