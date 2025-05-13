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
        if len(os.Args) < 2 {
                color.Red("❌ Error: Command is required.")
                return
        }

        prompt := strings.Join(os.Args[1:], " ")

        body := fmt.Sprintf("{\"prompt\":\"%s\"}", prompt)
        req, err := http.NewRequest("POST", "https://www.gitfluence.com/api/generate", bytes.NewBuffer([]byte(body)))
        if err != nil {
                color.Red("❌ Error creating request: %v", err)
                return
        }

        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("accept", "*/*")
        req.Header.Set("accept-language", "en")
        req.Header.Set("Content-Type", "application/json")
        req.Header.Set("origin", "https://www.gitfluence.com")
        req.Header.Set("priority", "u=1, i")
        req.Header.Set("Referer", "https://www.gitfluence.com/")
        req.Header.Set("sec-ch-ua", `"Chromium";v="136", "Google Chrome";v="136", "Not.A/Brand";v="99"`)
        req.Header.Set("sec-ch-ua-mobile", "?0")
        req.Header.Set("sec-ch-ua-platform", `"Windows"`)
        req.Header.Set("sec-fetch-dest", "empty")
        req.Header.Set("sec-fetch-mode", "cors")
        req.Header.Set("sec-fetch-site", "same-origin")
        req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/136.0.0.0 Safari/537.36")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
                color.Red("❌ Error sending request: %v", err)
                return
        }
        defer resp.Body.Close()

        if resp.StatusCode != 200 {
                color.Yellow("⚠️ Request failed with status: %d", resp.StatusCode)
                return
        }

        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
                color.Red("❌ Error reading response body: %v", err)
                return
        }

        bodyString := string(bodyBytes)

        start := strings.Index(bodyString, "```")
        end := strings.LastIndex(bodyString, "```")
        if start != -1 && end != -1 && end > start {
                command := bodyString[start+3 : end]
                command = strings.ReplaceAll(command, "\n", "")
                color.Green("%s", command)

                fmt.Print("⚡ Do you want to execute this command? (y/n): ")
                var input string
                fmt.Scanln(&input)

                if input == "y" {
                        cmd := exec.Command("sh", "-c", command)

                        cmd.Stdout = os.Stdout
                        cmd.Stderr = os.Stderr

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
