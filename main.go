package main

import "fmt"

func main() {
	var confName = "Go lang Conference"
	var username string

	var totalSeats uint = 100
	var bookedSeats uint = 0

	fmt.Printf("Enter your name: ")
	fmt.Scan(&username)

	fmt.Printf("%v Welcome to %v",username, confName)
	fmt.Printf("Total Seats: %v\n Booked Seats; %v\n", totalSeats, bookedSeats)

	fmt.Println("---------------------------------")
	//2 types of arr declaration
	var bookingNameArr [100]string //name of users who booked tickets
	var bookingNoArr = [100]int{} //booking number of users who booked tickets

	bookingNameArr[0] = username
	bookingNoArr[0] = 10

	fmt.Printf("Whole Array Name: %v\n", bookingNameArr)
	fmt.Printf("Whole Array Number: %v\n", bookingNoArr)
	fmt.Printf("Array Type: %T\n", bookingNoArr)
	fmt.Printf("Array Length: %v\n", len(bookingNoArr))

	fmt.Println("---------------------------------")
	//Slice - dynamic array

	var bookingNameSlice []string
	bookingNameSlice = append(bookingNameSlice, username)
	fmt.Printf("Slice Name: %v\n", bookingNameSlice)
	fmt.Printf("Slice Type: %T\n", bookingNameSlice)
	fmt.Printf("Slice Length: %v\n", len(bookingNameSlice))

	fmt.Println("---------------------------------")



	//loops

	for{
		fmt.Println("Infinite Loop")
		break
	}

	for i:=0; i<10; i++ {
		fmt.Println(i)
	}

	//for each 
	for index, value := range bookingNameArr {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}    
     
	fmt.Println("---------------------------------")

	// if else

	if totalSeats > 25 {
		fmt.Println("Seats Available")
	} else if totalSeats <=25 && totalSeats > 0 {
		fmt.Println("Seats Limited")
	} else {
		fmt.Println("Seats Not Available")
	}

	fmt.Println("---------------------------------")

	//switch case

	switch totalSeats {
	case 100:
		fmt.Println("Seats Full")
	case 50:
		fmt.Println("Seats Half Full")
	default:
		fmt.Println("Seats Not Available")
	}

	fmt.Println("---------------------------------")

	// input validation

	var age int
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	isvalidAge := age > 18

	if isvalidAge {
		fmt.Println("You are eligible")
	} else {
		fmt.Println("You are not eligible")
	}

	fmt.Println("---------------------------------")

	//functions

	x,y := greet(username)
	x+=1
	y+=2


	fmt.Println("---------------------------------")







}


func greet(username string) (int, int) {
	fmt.Println("Hello ",username)
	var ret1 int  = 10
	var ret2 int = 20
	return ret1, ret2
}