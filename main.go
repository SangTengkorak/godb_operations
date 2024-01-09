package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"mastengkorak.com/godatabasebasic/dboperations"
)

// DB access credentials
var DBusername string
var DBpassword string
var DBname string
var DBhost string

// DB instance

func main() {
	dboperations.DBconnections()

	//Find most expensive products
	dboperations.Mostexpensive()

	//Find cheapest products
	dboperations.Cheapest()

	//Display all products
	dboperations.Allitems()

	//Display specific products
	//Capture item id
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nEnter item ID you want to purchase: ")
	scanner.Scan()
	itemidstr := scanner.Text()
	itemidint, _ := strconv.Atoi(itemidstr)

	//Capture user savings
	fmt.Print("Enter your savings: ")
	scanner.Scan()
	savingsstr := scanner.Text()
	savingsflt, _ := strconv.ParseFloat(savingsstr, 64)

	dboperations.Buyitem(itemidint, savingsflt)

}
