package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

//Function  - write CSV Rating   given to  Product
//Function -  Rating Process

func reviewSystem() rating {
	var r rating

	fmt.Println("Enter product id:")
	fmt.Scanln(&r.Id)

	fmt.Println("Enter stars  (1-5):")
	fmt.Scanln(&r.stars)

	fmt.Println("Enter comment If any:")
	fmt.Scanln(&r.comment)

	fmt.Println("Enter your userName:")
	fmt.Scanln(&r.userName)

	if r.stars >= 3 {
		color.Green("Thanks For the review 😊")
	} else {
		color.Red("We’ll look into your feedback 😔")
	}

	return r
}

func storeReviewInDataBase(r rating) error {

	filepath := "ratingsDb.csv"

	_, err := os.Stat(filepath)
	if err != nil {
		color.Red("Memory doesn't Exist : Resolving It ")
	} else {
		color.Green("Memory  Exist : Lets update the Review")
	}
	return nil

	if !filepath {

	} else {

	}
}
