package main

import (
	"quiz/quiz"
)

func main() {
	quiz.RandomizeAnswers("questions.json")
	quiz.StartApp()

}
