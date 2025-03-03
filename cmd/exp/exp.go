package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
)

type User struct {
	Name     string
	Age      int
	Location string
	Hobbies  []string
	Contact  map[string]string
}

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		return fmt.Errorf("create org: %w", err)
	}
	return nil
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name:     "John Smith",
		Age:      123,
		Location: "UK",
		Hobbies:  []string{"running", "reading", "music"},
		Contact: map[string]string{
			"Home": "123",
			"Work": "456",
		},
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}

	err = CreateUser()
	if err != nil {
		log.Println(err)
	}

	err = CreateOrg()
	if err != nil {
		log.Println(err)
	}

	Demo()
	Demo(1)
	Demo(1, 2, 3)

	fmt.Println(Sum())
	fmt.Println(Sum(4))
	fmt.Println(Sum(4, 5, 6))

	fib := []int{1, 1, 2, 4, 5, 8}
	Demo(fib...)
	fmt.Println(Sum(fib...))
}

func Demo(numbers ...int) {
	for _, number := range numbers {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func Sum(numbers ...int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum
}
