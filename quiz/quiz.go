package quiz

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func StartApp() {

	jsonFile, err := os.ReadFile("questions.json")

	if err != nil {
		fmt.Println(err)
	}

	var quiz Quiz

	err = json.Unmarshal(jsonFile, &quiz)

	isRunning := true
	var input string

	for isRunning {
		fmt.Scanln(&input)

		input = strings.ToLower(input)

		switch input {
		case "start":
			StartQuiz(quiz)
		case "exit":
			isRunning = false

		default:
			fmt.Println("Please Provide Input:")
		}
	}

}

func StartQuiz(quiz Quiz) {

	isRunning := true

	var input string
	var currentQuestion Question

	for isRunning {
		currentQuestion = quiz.Questions[0]

		fmt.Println(currentQuestion.Question)
		fmt.Println()
		fmt.Println("A. " + currentQuestion.Options.A)
		fmt.Println("B. " + currentQuestion.Options.B)
		fmt.Println("C. " + currentQuestion.Options.C)
		fmt.Println("D. " + currentQuestion.Options.D)
		fmt.Println()

		fmt.Scan(&input)

		switch input {
		case "next":
			continue
		case "exit":
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
		fmt.Scan()
	} else {

		fmt.Println("Wrong!")

		fmt.Println("Correct Anwser is:" + question.Answer)

		fmt.Println(question.Explanation)

		fmt.Println()
		fmt.Scan()
	}

}
