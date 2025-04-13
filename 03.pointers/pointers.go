package main

import "fmt"

func main() {
	age := 32
	fmt.Println(age)
	editAgeToAdultYears(&age)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
