package quiz

import (
	"fmt"
	"strings"
)

func StartApp() { // change thais to be run in the loop as the code for exiting now is wierd

	quiz := loadQuiz("questions.json")

	isRunning := true
	var input string

	for isRunning {
		fmt.Println("Main Menu:")
		fmt.Scanln(&input)

		input = strings.ToLower(input)

		switch input {
		case "start":
			StartQuiz(quiz, &isRunning)

		case "exit":
			isRunning = false

		default:
			fmt.Println("Wrong Input")
		}
	}

}

func StartQuiz(quiz Quiz, isAppLoopRunning *bool) {

	isRunning := true

	var input string
	var currentQuestion Question

	for isRunning {
		currentQuestion = quiz.GetRandomQuestion()

		currentQuestion.Print()

		fmt.Scan(&input)
		input = strings.ToLower(input)

		switch input {

		case "next":
			continue

		case "exit":
			*isAppLoopRunning = false
			isRunning = false

		case "menu":
			isRunning = false

		case "a", "b", "c", "d":

			questionCheck(input, currentQuestion)

		default:
			fmt.Println("Please Provide Input:")
			fmt.Println()

		}
	}

}

func questionCheck(input string, question Question) {

	input = strings.ToUpper(input)

	if input == question.Answer {

		fmt.Println("Correct!")
		fmt.Println()

	} else {

		fmt.Println("Wrong!")

		fmt.Println("Correct Anwser is:" + question.Answer)

		fmt.Println(question.Explanation)

		fmt.Println()

	}

}
