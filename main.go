package main

import (
	"fmt"
	"atm-simulator/controllers"
	"atm-simulator/db"
	"atm-simulator/utils"
)

func main() {
	db.Init()

	for {
		fmt.Println("\n=== ATM Simulator ===")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Check Balance")
		fmt.Println("4. Deposit")
		fmt.Println("5. Withdraw")
		fmt.Println("6. Transfer")
		fmt.Println("7. Exit")

		choice := utils.GetInputInt("Enter choice: ")

		switch choice {
		case 1:
			controllers.Register()
		case 2:
			controllers.Login()
		case 3:
			controllers.CheckBalance()
		case 4:
			controllers.Deposit()
		case 5:
			controllers.Withdraw()
		case 6:
			controllers.Transfer()
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
