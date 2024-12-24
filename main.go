package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to generate a response based on user input
func generateResponse(input string) string {
	input = strings.ToLower(input)

	// Simple keyword-based responses
	if strings.Contains(input, "hello") || strings.Contains(input, "hey") {
		return "Hi there! How can I help you today?"
	} else if strings.Contains(input, "how are you") {
		return "I'm just a bot, but I'm doing great! How about you?"
	} else if strings.Contains(input, "bye") || strings.Contains(input, "exit") {
		return "Goodbye! It was nice talking to you."
	} else if strings.Contains(input, "your name") {
		return "I am a simple chatbot created with Go."
	}

	// Default response for unrecognized input
	return "I'm sorry, I didn't understand that. Could you please rephrase?"
}

func main() {
	// Create a reader for user input
	reader := bufio.NewReader(os.Stdin)

	// Greet the user
	fmt.Println("Hello! I am a chatbot. Type 'exit' or 'bye' to end the conversation.")

	// Start the conversation loop
	for {
		// Get user input
		fmt.Print("You: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput) // Remove newline character

		// Check if the user wants to exit
		if userInput == "exit" || userInput == "bye" {
			fmt.Println("Chatbot: Goodbye!")
			break
		}

		// Get the chatbot's response and print it
		response := generateResponse(userInput)
		fmt.Println("Chatbot:", response)
	}
}
