package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	fmt.Println(doctor.Intro())
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Q: ")
		userInput, _ := reader.ReadString('\n')
		// replace the enter key in Windows
		userInput = strings.ReplaceAll(userInput, "\r\n", "")
		//replace the enter key in linux/macOS
		userInput = strings.ReplaceAll(userInput, "\n", "")

		fmt.Printf("A: %s\n", doctor.Response(userInput))
		if userInput == "quit" {
			break
		} else {
			response := doctor.Response(userInput)
			println(response)
		}
	}
}
