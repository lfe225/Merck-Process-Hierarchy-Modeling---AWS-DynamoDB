package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Merck-Fall2024")
	fmt.Println("Choose an option:")
	fmt.Println("1. Populate")
	fmt.Println("2. Delete all items")
	fmt.Println("3. Get all stages")
	fmt.Println("4. Get Operations by stage")

	//User input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()

	//Definiting table
	tableName := "Merck-Fall2024"

	//Running the option
	switch strings.TrimSpace(choice) {
	case "1":
		PopulateDatabase(tableName)
	case "2":
		DeleteAllItemsFromTable(tableName)
	case "3":
		GetAllStages(tableName)
	case "4":
		GetOperationsByStage(tableName)
	}
}
