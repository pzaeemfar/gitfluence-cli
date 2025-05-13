package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// Check if enough arguments are passed
	if len(os.Args) < 2 {
		color.Red("❌ Error: Command is required.")
		return
	}

	// Join the arguments to form the prompt
	prompt := strings.Join(os.Args[1:], " ")

	// Prepare the request body
	body := fmt.Sprintf("{\"prompt\":\"%s\"}", prompt)
	req, err := http.NewRequest("POST", "https://www.gitfluence.com/api/generate", bytes.NewBuffer([]byte(body)))
	if err != nil {
		color.Red("❌ Error creating request: %v", err)
		return
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		color.Red("❌ Error sending request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is 200
	if resp.StatusCode != 200 {
		color.Yellow("⚠️ Request failed with status: %d", resp.StatusCode)
		return
	}

	// Read the response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		color.Red("❌ Error reading response body: %v", err)
		return
	}

	// Convert the body to string
	bodyString := string(bodyBytes)

	// Extract command if it's in the expected format
	start := strings.Index(bodyString, "```")
	end := strings.LastIndex(bodyString, "```")
	if start != -1 && end != -1 && end > start {
		command := bodyString[start+3 : end]
		command = strings.ReplaceAll(command, "\n", "")
		color.Green("%s", command)

		// Ask the user if they want to execute the command
		fmt.Print("⚡ Do you want to execute this command? (y/n): ")
		var input string
		fmt.Scanln(&input)

		// Execute the command if user says "y"
		if input == "y" {
			// Execute the command and print output live
			cmd := exec.Command("sh", "-c", command)

			// Set the output and error streams to print to the terminal
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			// Run the command
			err := cmd.Run()
			if err != nil {
				color.Red("❌ Error executing command: %v", err)
			}
		} else {
			color.Magenta("💤 Exiting without executing command. 👋")
		}
	} else {
		color.Yellow("⚠️ No executable command found in the response.")
	}
}
