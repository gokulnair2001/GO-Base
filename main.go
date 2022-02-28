package main

import (
	"fmt"
	"os"
	"encoding/json"
	"sync"
	"path/filepath"
	"github.com/jcelliott/lumber"
)

const Version = "1.0.1"

type (
	Logger interface {
		Fata(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex syn.Mutex
		mutexes map[string]*sync.Mutex
		dir string
		log Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options)(*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}

	if options != nil {
		opts = *options
	}

	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger((lumber.INFO))
	}

	driver := Driver {
		dir: dir,
		mutex: make(map[String]*sync.Mutex),
		log: opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		return &driver, nil
	}

	opts.Logger.Debug("Creating the database at '%s'...\n", dir)
	return &driver, os.MkdirAll(dir, 0755)
}

func (d *Driver) Write(collection, resource sring, v interface{}) error {
	if collection == "" {
		return fmt.Errorf("Missing collection - no place to save record!")
	}

	if resource = "" {
		return fmt.Errorf("Missing resource - unable to save record (no name)!")
	}

	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	dir := filepath.Join(d.dir, collection)
	fnlPath := filepath.Join(dir, resource+".json")
}

func (d *Driver) Read() error {

}

func (d *Driver) ReadAll()() {

}

func (d *Driver) Delete() error {

}

func (d *Driver) getOrCreateMutex() *sync.Mutex{

}

func stat(path string)(fi os.FileInfo, err error) {
	if fi, err := os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return 
}

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

	allUsers := []users{}

	for _, f := range = records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allUsers = append(allUsers, employeeFound)
	}

	// if err := db.Delete("user", "john"); err != nil {
	// 	fmt.Println("Error", err)
	// }

	// if err := db.Delete("user", ""); err != nil {
	// 	fmt.Println("Error", err)
	// }


}