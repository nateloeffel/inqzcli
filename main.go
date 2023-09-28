package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

type IPResponse struct {
	IP string `json:"ip"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func getIPAddress() string {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	// Reading the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	// Unmarshalling the JSON response
	var ipResponse IPResponse
	err = json.Unmarshal(body, &ipResponse)
	if err != nil {
		return ""
	}

	return ipResponse.IP
}

func main() {

	args := os.Args
	if len(args) == 1 {
		fmt.Println("----------------------------------\nWelcome to the...\n----------------------------------\n /$$$$$$ /$$   /$$  /$$$$$$  /$$$$$$$$       /$$   /$$ /$$$$$$$$ /$$$$$$$$ /$$      /$$  /$$$$$$  /$$$$$$$  /$$   /$$\n|_  $$_/| $$$ | $$ /$$__  $$|_____ $$       | $$$ | $$| $$_____/|__  $$__/| $$  /$ | $$ /$$__  $$| $$__  $$| $$  /$$/\n  | $$  | $$$$| $$| $$  \\ $$     /$$/       | $$$$| $$| $$         | $$   | $$ /$$$| $$| $$  \\ $$| $$  \\ $$| $$ /$$/ \n  | $$  | $$ $$ $$| $$  | $$    /$$/        | $$ $$ $$| $$$$$      | $$   | $$/$$ $$ $$| $$  | $$| $$$$$$$/| $$$$$/  \n  | $$  | $$  $$$$| $$  | $$   /$$/         | $$  $$$$| $$__/      | $$   | $$$$_  $$$$| $$  | $$| $$__  $$| $$  $$  \n  | $$  | $$\\  $$$| $$/$$ $$  /$$/          | $$\\  $$$| $$         | $$   | $$$/ \\  $$$| $$  | $$| $$  \\ $$| $$\\  $$ \n /$$$$$$| $$ \\  $$|  $$$$$$/ /$$$$$$$$      | $$ \\  $$| $$$$$$$$   | $$   | $$/   \\  $$|  $$$$$$/| $$  | $$| $$ \\  $$\n|______/|__/  \\__/ \\____ $$$|________/      |__/  \\__/|________/   |__/   |__/     \\__/ \\______/ |__/  |__/|__/  \\__/\n                        \\__/                                                                                                                   ")
		os.Exit(0)
	}

	if args[1] == "create" {
		if len(args) > 2 {
			if args[2] == "discord" {
				path := "BaseDiscordBot"
				if len(args) > 3 {
					path = args[3]
				}
				fmt.Println("Creating Discord Bot")
				cmd := exec.Command("git", "clone", "https://github.com/nateloeffel/BaseDiscordBot.git", path)
				_, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				file, err := os.OpenFile(path+"/.env", os.O_WRONLY|os.O_CREATE, 0644)
				if err != nil {
					// Handle the error
					panic(err)
				}
				defer file.Close()

				error := os.Remove(path + "/.env.example")
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
		} else {
			fmt.Println("You must pass an argument with create.")
		}
	} else if args[1] == "ip" {

		fmt.Println("Your IP address is:", getIPAddress())
	} else if args[1] == "serve" {
		port := "8080"
		if len(args) > 2 {
			port = string(args[2])
		}
		http.HandleFunc("/", handler)
		fmt.Println("Serving on port " + port)
		http.ListenAndServe(":"+port, nil)

	} else if args[1] == "checkip" {
		ip := ""
		if len(args) < 3 {
			ip = getIPAddress()
		} else {
			ip = args[2]
		}

		resp, err := http.Get("http://api.ipstack.com/" + ip + "?access_key=c7ceecbdf42731fd5e0a58271ebf8b43")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		// Reading the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Unmarshalling the JSON response
		var data map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Printing the important information
		fmt.Println("IP Address:", data["ip"])
		fmt.Println("Country:", data["country_name"])
		fmt.Println("Region:", data["region_name"])
		fmt.Println("City:", data["city"])
		fmt.Println("Latitude:", data["latitude"], "Longitude:", data["longitude"])
		fmt.Println("Zip Code:", data["zip"])
	}
}
