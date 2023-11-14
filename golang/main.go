package main

import (
	"fmt"
	"os"
	"time"
)

var users = map[string]string{
	"admin": "admin",
	"uslu" : "uslu",
}

func main() {
	var (
		username  string
		password  string
		loginTries = 5
	)

	for loginTries > 0 {
		fmt.Println("Please select your login type:")
		fmt.Println("0 - Admin Login")
		fmt.Println("1 - Student Login")

		var userType int
		fmt.Scanln(&userType)

		if userType != 0 && userType != 1 {
			fmt.Println("Invalid option. Try again.")
			continue
		}

		fmt.Print("Username: ")
		fmt.Scanln(&username)
		fmt.Print("Password: ")
		fmt.Scanln(&password)

		if verification(username, password, userType) {
			fmt.Println("\x1b[32mLogin Succesful!\x1b[0m")
			if userType == 0 {
				adminPage()
			} else {
				studentPage()
			}
			break
		} else {
			loginTries--
			if loginTries > 0 {
				fmt.Printf("\x1b[33mLogin information is incorrect. Your remaining entry rights are: %d\x1b[0m\n", loginTries)
			} else {
				fmt.Println("You have run out of incorrect login attempts. The program is terminating.")
			}
		}
	}

	fmt.Println("The program has been terminated.")
}

func verification(username, password string, userType int) bool {
	expectedPassword, exists := users[username]
	if !exists || expectedPassword != password {
		logEntry := fmt.Sprintf("Username: %s\nCheck-in Date and Time: %s\n\x1b[31mLogin Status : Failed!\x1b[0m", username, time.Now().Format("2006-01-02 15:04:05"))
		writeLog(logEntry)
		return false
	}

	logEntry := fmt.Sprintf("Username: %s\nCheck-in Date and Time: %s\n\x1b[32mLogin Status : Succesful!\x1b[0m", username, time.Now().Format("2006-01-02 15:04:05"))
	writeLog(logEntry)

	return true
}

func writeLog(logEntry string) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Log file could not be opened:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(logEntry + "\n\n"); err != nil {
		fmt.Println("Failed to log:", err)
	}
}

func adminPage() {
	for {
		fmt.Println("Please select an action:")
		fmt.Println("0 - View Logs")
		fmt.Println("1 - Logout")

		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			viewLogs()
		} else if choice == 1 {
			fmt.Println("Checking out.")
			break
		} else {
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func studentPage() {
	fmt.Println("\x1b[32mStudent Login Successful!\x1b[0m")
}

func viewLogs() {
	data, err := os.ReadFile("logs.txt")
	if err != nil {
		fmt.Println("Log file could not be read:", err)
		return
	}
	fmt.Println("Log Records:")
	fmt.Println(string(data))
}
