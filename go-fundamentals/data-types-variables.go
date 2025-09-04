package main

import (
	"fmt"
)

func variableFuncWithGreeting(name string, country string) {
	userName := name
	message := fmt.Sprintf("Hello! My name is %v, from %v. I bring you love from %v.", userName, country, country)
	fmt.Println(message)

	/*
	* Different format specifiers
	* =============================================================
	* Verb   | Description
	* -------------------------------------------------------------
	* %v     | default format
	* %T    | type of the value
	* %d    | integers
	* %c    | characters
	* %q   | quoted characters/string
	* %s   | plain string
	* %t   | true or false
	* %f   | floating numbers
	* %.2f | floating numbers upto 2 decimal places
	* etc...
	* =============================================================
	 */
}

func userInputFunction() string {
	var name string
	var job string
	var age int
	var dob string

	fmt.Print("Enter your name, job,age, and date of birth. Use empty space to separate the info: ")
	_, err := fmt.Scanf("%s %s %d %s", &name, &job, &age, &dob)

	if err != nil {
		return fmt.Sprintf("%s", err)
	}

	return fmt.Sprintf("User ID Card\nName: %v\nJob: %v\nAge: %d\nDob: %v", name, job, age, dob)
}
