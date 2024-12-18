package hello

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// displayQuestion displays the question with the given question number.
func displayQuestion(questionNumber int) {
	switch questionNumber {
	case 1:
		fmt.Print("Enter your first name: ")
	case 2:
		fmt.Println("\nSelect your title prefix:")
		fmt.Println("a) Mr.")
		fmt.Println("b) Mrs.")
		fmt.Println("c) Dr.")
		fmt.Print("Enter your choice (a, b, or c): ")
	}
}

// collectAnswer collects and processes the answer for the given question number.
func collectAnswer(questionNumber int) string {
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(answer)
	return answer
}

// welcomeMessage generates the welcome message based on the title and first name.
func welcomeMessage(title, firstName string) {
	// Determine the title based on the choice
	switch title {
	case "a":
		title = "Mr."
	case "b":
		title = "Mrs."
	case "c":
		title = "Dr."
	default:
		title = "" // Handle invalid choice
	}

	fmt.Println("\nHello,", title, firstName)
}

func hello() {
	var firstName, title string

	for i := 1; i <= 2; i++ {
		displayQuestion(i)
		answer := collectAnswer(i)

		if i == 1 {
			firstName = answer
		} else if i == 2 {
			title = answer
		}
	}

	welcomeMessage(title, firstName)
}
