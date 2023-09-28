package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	args := os.Args
	if len(args) == 1 {
		fmt.Println("----------------------------------\nWelcome to the...\n----------------------------------\n /$$$$$$ /$$   /$$  /$$$$$$  /$$$$$$$$       /$$   /$$ /$$$$$$$$ /$$$$$$$$ /$$      /$$  /$$$$$$  /$$$$$$$  /$$   /$$\n|_  $$_/| $$$ | $$ /$$__  $$|_____ $$       | $$$ | $$| $$_____/|__  $$__/| $$  /$ | $$ /$$__  $$| $$__  $$| $$  /$$/\n  | $$  | $$$$| $$| $$  \\ $$     /$$/       | $$$$| $$| $$         | $$   | $$ /$$$| $$| $$  \\ $$| $$  \\ $$| $$ /$$/ \n  | $$  | $$ $$ $$| $$  | $$    /$$/        | $$ $$ $$| $$$$$      | $$   | $$/$$ $$ $$| $$  | $$| $$$$$$$/| $$$$$/  \n  | $$  | $$  $$$$| $$  | $$   /$$/         | $$  $$$$| $$__/      | $$   | $$$$_  $$$$| $$  | $$| $$__  $$| $$  $$  \n  | $$  | $$\\  $$$| $$/$$ $$  /$$/          | $$\\  $$$| $$         | $$   | $$$/ \\  $$$| $$  | $$| $$  \\ $$| $$\\  $$ \n /$$$$$$| $$ \\  $$|  $$$$$$/ /$$$$$$$$      | $$ \\  $$| $$$$$$$$   | $$   | $$/   \\  $$|  $$$$$$/| $$  | $$| $$ \\  $$\n|______/|__/  \\__/ \\____ $$$|________/      |__/  \\__/|________/   |__/   |__/     \\__/ \\______/ |__/  |__/|__/  \\__/\n                        \\__/                                                                                                                   ")
		os.Exit(0)
	}

	if args[1] == "create" {
		if len(args) > 2 && args[2] == "discord" {
			fmt.Println("Creating Discord Bot")

			cmd := exec.Command("git", "clone", "https://github.com/nateloeffel/BaseDiscordBot.git")
			_, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			file, err := os.OpenFile("./BaseDiscordBot/.env", os.O_WRONLY|os.O_CREATE, 0644)
			if err != nil {
				// Handle the error
				panic(err)
			}
			defer file.Close()

			error := os.Remove("./BaseDiscordBot/.env.example")
			if error != nil {
				panic(err)
			}

			// Write string to file
			_, err = file.WriteString("GUILD_ID=\nCLIENT_ID=\nTOKEN=")
			if err != nil {
				// Handle the error
				panic(err)
			}
		}

	}
}
