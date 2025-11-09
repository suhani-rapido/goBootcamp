package main

import "fmt"

type Products struct {
	Name     string
	Rating   int
	Comments []string
	User     string
}

func main() {
	product := []Products{
		{"Bag", 4, []string{"Nice Color and Design"}, "Suhani"},
		{"Shoes", 5, []string{"Very Comfortable"}, "John"},
		{"Shirt", 2, []string{"Bad Quality"}, "Alice"},
	}

	product[0].Comments = append(product[0].Comments, ", Good Product")
	productReview(product[0])
	fmt.Printf("\n")
	productReview(product[1])
	fmt.Printf("\n")
	productReview(product[2])

}

func productReview(product Products) {
	fmt.Printf("Name of the product: %v\n", product.Name)
	fmt.Printf("The rating for product: %v\n", product.Rating)
	fmt.Printf("Comments: %v \n", product.Comments)
	fmt.Printf("Thank you for your feedback, %v!\n", product.User)
}
