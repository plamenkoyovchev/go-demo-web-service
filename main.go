package main

import (
	"fmt"

	"github.com/plamenkoyovchev/go-demo-web-service/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "plamen",
		LastName:  "yovchev",
	}

	fmt.Println(u)
}
