package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

//////////////////////////////////////////////////////
//////////////Package Level Variables/////////////////
//////////////////////////////////////////////////////

///Note: They can not be assigned using ":="
var conferenceName = "Go Conference" // this is not dynamic typing as you provide the value with the definition
//Go have the feature named Type Inference which means Go can infer the type when you assign a value to it.

// var remainingTickets	//This can't work as you didn't provide a type neither init value

var remainingTickets uint = 50 //uint has +ve values. So, we can't have negative number of tickets

const conferenceTickets = 50

//var bookings = [50]string{"Ehab, Wael, Osama"} //define the size of the array, then type pf elements, then minimum number of values
//var bookings [50]string //(define empty array with size 50)

// //Here we are creating an empty list of maps
// //maps only allow for one data type for the keys. And one data type for the values.
// var bookings = make([]map[string]string, 0) //slice is an abstraction of an Array. Slices are more flexible and more powerful. Support dynamic size
// //bookings := []string{}  //equivelant to define a slice with dynamic unkonwn size. You can add values either.

//Here we are creating a list of struct UserData
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

//Waits for the launched goroutine to finish (So that if the main program ends but there are threads that didn't finish, we wait for them)
var wg = sync.WaitGroup{}

func main() {

	greatUsers()

	//All loop types (while loop, for loop, for-each loop) are implemented in go in "for" syntax only
	//This is an infinite loop
	// for {

	// }

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			//As this functions takes 10 seconds to execute, we use goroutines
			// just right the keyword "go" before the function name that you want to create in a new thread.
			//Go handles the creation of the threads (goroutines)
			wg.Add(1) //Increase the number of goroutines we're waiting for before ending the program.
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			// fmt.Println("Your input data is not valid, try again")
			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
	}
	wg.Wait() //Wait for all goroutines to exit
}

func greatUsers() {

	fmt.Println("Welcome to our", conferenceName, "booking application") //Spaces are added automatically between string and variables

	fmt.Printf("Conference tickets is %T, remaining tickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have a total of %v tickets. There is %v tickets available.\n", conferenceName, remainingTickets) //using formatted strings
	fmt.Println("Get your ticket right now to attend")

}

func getUserInput() (string, string, string, uint) {

	//How to define types
	var firstName string
	var lastName string
	var email string
	// var city string
	var userTickets uint

	//ask user for their name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	// fmt.Println("Enter where you want to attend the conference")
	// fmt.Scan(&city)

	fmt.Println("Enter your userTickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

//Type of the function is a slice of strings and comes after parameter definition
func getFirstNames() []string {

	firstNames := []string{}

	//This is a for-each loop. When range is used with slices, it returns back the index and the entry in the slice.
	for _, booking := range bookings {

		//firstNames = append(firstNames, strings.Fields(booking)[0]) //Fields function splits the white space separated string to a slice. Eg: "Osama Magdy" --> ["Osama", "Magdy"]
		//firstNames = append(firstNames, booking["firstName"]) //In case of bookings is a list of maps
		firstNames = append(firstNames, booking.firstName) //In case of bookings is a list of UserData struct
	}

	return firstNames
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {

	remainingTickets -= userTickets

	//bookings[0] = firstName + " " + lastName

	//create an empty map for a user
	//Note: in Go, all keys have to be of the same data type and all values have to be of the same data type
	// var userData = make(map[string]string)

	// //add data to store information in the userData map
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	//Create an object of struct UserData
	var userData = UserData{firstName: firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("Remaining number of tickets is %v for conference %v \n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {

	//Assume that the client that sends the mail will take 10 seconds to do it
	time.Sleep(10 * time.Second)
	//Sprintf is the same as Printf but instead of printing to the terminal, it returns the formatted string to be stored in a variable
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("###################")

	wg.Done() //This indicates that the goroutine created for this function should end and DECREMENTS the counter of goroutines we're waiting for.
}

//reading the time
// today := time.Now().Weekday()
// fmt.Println("today is :", today, "When's Saturday?")
// switch time.Saturday {
// case today + 0:
// 	fmt.Println("Today.")
// case today + 2:
// 	fmt.Println("In two days.")
// case today + 1:
// 	fmt.Println("Tomorrow.")
// default:
// 	fmt.Println("Too far away.")
// }

//switch case syntax
// city := "London"
// switch city {

// case "New York":
// 	// some code here
// case "Singapore", "Hong Kong":
// 	//some code here
// case "London", "Berlin":
// 	//some code here
// case "Mexico City":
// 	//some code here
// default:
// 	fmt.Println("No valid city is selected")
// }
