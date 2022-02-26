package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Address struct {
	City string
	State string
	Country string
	Pincode json.number
}

type User struct {
	Name string
	Age json.number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User {
		{"John", "23", "2345678654", "Evo", Address{"Vellore", "Tamil Nadu", "India", "234343"}},
		{"Rahul", "25", "67634787645", "Slack", Address{"Banglore", "Karnataka", "India", "348789"}},
		{"Pratyush", "20", "8786736256", "Intuit", Address{"Chennai", "Tamil Nadu", "India", "457633"}},
		{"Kevin", "23", "9876787234", "EduPro", Address{"Delhi", "Delhi", "India", "232989"}},
		{"Sarthak", "29", "0987898724", "Amazon", Address{"Mysore", "Karnataka", "India", "215433"}},
		{"Loky", "28", "2370982435", "Apple", Address{"Banglore", "Karnataka", "India", "987834"}},
	}


for _, value := range employees {
	db.Write("users", value.Name, User {
		Name: value.Name,
		Age: value.Age,
		Contact: value.Contact,
		Company: value.Company,
		Address: value.Address,
	})
}

	records, err := db.ReadAll("Users")
	if err != nil {
	fmt.Println("Error", err)
	}

	fmt.Println(records)
}