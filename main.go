package main

import (
	"booking-app/notmain"
	"fmt"
	"strconv"
	"time"
)

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


	fmt.Println("---------------------------------")

	//fn from helper.go

	ad,sub,mul,div := maths(x,y)
	fmt.Println(ad,sub,mul,div)
	
	fmt.Println("---------------------------------")

	//fn from diffpack.go

	fmt.Println(notmain.Notmainadd(10,20))

	//maps

	var bookingMap = make(map[string]int)

	bookingMap["user1"] = 10
	bookingMap["user2"] = 20

	fmt.Println(bookingMap)
	fmt.Println(bookingMap["user1"])

	fmt.Println("---------------------------------")
	// type conversions

	var a int = 10
	var b float64 = 20.5

	var c float64 = float64(a) + b

	fmt.Println(c)

	var d string = "10"
	var e int = 20


	f, _ := strconv.Atoi(d)
	fmt.Println(f+e)

	fmt.Println("---------------------------------")

	// structs

	type User struct {
		name string
		age int

	}

	var user1 User = User{"John", 25}
	fmt.Println(user1)

	var user2 User
	user2.name = "Doe"
	user2.age = 30

	fmt.Println("---------------------------------")


	// concurrency - Groutines


	for {
		var name string
		var age int

		fmt.Printf("Name: ")
		fmt.Scan((&name)) 

		fmt.Printf("Age: ")
		fmt.Scan(&age)

		go groutine(name, age)

	}




}

func groutine(name string, age int) {
	time.Sleep(5 * time.Second)
	fmt.Println("Name: ", name)
	fmt.Println("Age: ", age)
}


func greet(username string) (int, int) {
	fmt.Println("Hello ",username)
	var ret1 int  = 10
	var ret2 int = 20
	return ret1, ret2
}