package controllers

import (
	"fmt"
	"atm-simulator/db"
	"atm-simulator/utils"
)

func Register() {
	fmt.Println("\n=== Register ===")

	username := utils.GetInput("Enter username: ")
	password := utils.GetInput("Enter password: ")

	// Cek apakah username sudah ada
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM accounts WHERE username = ?)", username).Scan(&exists)
	if err != nil {
		fmt.Println("Error checking username:", err)
		return
	}

	if exists {
		fmt.Println("Username already exists. Please choose another.")
		return
	}

	// Insert user baru
	_, err = db.DB.Exec("INSERT INTO accounts (username, password, balance) VALUES (?, ?, ?)", username, password, 0)
	if err != nil {
		fmt.Println("Error registering user:", err)
		return
	}

	fmt.Println("Registration successful!")
}

func Login() {
	fmt.Println("\n=== Login ===")

	username := utils.GetInput("Enter username: ")
	password := utils.GetInput("Enter password: ")

	var dbPassword string
	err := db.DB.QueryRow("SELECT password FROM accounts WHERE username = ?", username).Scan(&dbPassword)
	if err != nil {
		fmt.Println("Invalid username or password.")
		return
	}

	if password != dbPassword {
		fmt.Println("Invalid username or password.")
		return
	}

	fmt.Println("Login successful!")
}

func CheckBalance() {
	fmt.Println("\n=== Check Balance ===")

	username := utils.GetInput("Enter username: ")

	var balance float64
	err := db.DB.QueryRow("SELECT balance FROM accounts WHERE username = ?", username).Scan(&balance)
	if err != nil {
		fmt.Println("Error fetching balance:", err)
		return
	}

	fmt.Printf("Your balance: %.2f\n", balance)
}

func Deposit() {
	fmt.Println("\n=== Deposit ===")

	username := utils.GetInput("Enter username: ")
	amount := utils.GetInputInt("Enter amount to deposit: ")

	// Update balance
	_, err := db.DB.Exec("UPDATE accounts SET balance = balance + ? WHERE username = ?", amount, username)
	if err != nil {
		fmt.Println("Error depositing:", err)
		return
	}

	fmt.Println("Deposit successful!")
}

func Withdraw() {
	fmt.Println("\n=== Withdraw ===")

	username := utils.GetInput("Enter username: ")
	amount := utils.GetInputInt("Enter amount to withdraw: ")

	// Cek saldo dulu
	var balance float64
	err := db.DB.QueryRow("SELECT balance FROM accounts WHERE username = ?", username).Scan(&balance)
	if err != nil {
		fmt.Println("Error fetching balance:", err)
		return
	}

	if float64(amount) > balance {
		fmt.Println("Insufficient balance.")
		return
	}

	// Update balance
	_, err = db.DB.Exec("UPDATE accounts SET balance = balance - ? WHERE username = ?", amount, username)
	if err != nil {
		fmt.Println("Error withdrawing:", err)
		return
	}

	fmt.Println("Withdraw successful!")
}

func Transfer() {
	fmt.Println("\n=== Transfer ===")

	fromUsername := utils.GetInput("Enter your username: ")
	toUsername := utils.GetInput("Enter recipient username: ")
	amount := utils.GetInputInt("Enter amount to transfer: ")

	// Cek saldo pengirim
	var balance float64
	err := db.DB.QueryRow("SELECT balance FROM accounts WHERE username = ?", fromUsername).Scan(&balance)
	if err != nil {
		fmt.Println("Error fetching sender balance:", err)
		return
	}

	if float64(amount) > balance {
		fmt.Println("Insufficient balance.")
		return
	}

	// Mulai transaksi database (biar aman)
	tx, err := db.DB.Begin()
	if err != nil {
		fmt.Println("Error starting transaction:", err)
		return
	}

	// Kurangi saldo pengirim
	_, err = tx.Exec("UPDATE accounts SET balance = balance - ? WHERE username = ?", amount, fromUsername)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error deducting sender balance:", err)
		return
	}

	// Tambah saldo penerima
	_, err = tx.Exec("UPDATE accounts SET balance = balance + ? WHERE username = ?", amount, toUsername)
	if err != nil {
		tx.Rollback()
		fmt.Println("Error adding recipient balance:", err)
		return
	}

	// Commit transaksi
	err = tx.Commit()
	if err != nil {
		fmt.Println("Error committing transaction:", err)
		return
	}

	fmt.Println("Transfer successful!")
}

