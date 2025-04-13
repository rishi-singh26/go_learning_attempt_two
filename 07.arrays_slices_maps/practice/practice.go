package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	// Task 1
	hobbies := [3]string{"Coding", "Reading", "Gymming"}
	fmt.Println(hobbies)

	// Task 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// Task 3
	fmt.Println(hobbies[:2])
	firstAndSecond := hobbies[0:2]
	fmt.Println(firstAndSecond)

	// Taks 4

	// Task 5
	goals := []string{"Learn GO", "Become an expert"}

	// Task 6
	goals[1] = "Become the expert"
	goals = append(goals, "Teach Others")
	fmt.Println(goals)

	// Task 7
	products := []Product{
		{
			title: "A Book",
			id:    "1",
			price: 9.99,
		},
		{
			title: "A Pen",
			id:    "2",
			price: 2.99,
		},
	}

	products = append(products, Product{
		title: "An Eraser",
		id:    "3",
		price: 0.99,
	})

	fmt.Println(products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
