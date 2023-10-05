package main

import (
	"quiz/quiz"
)

func main() {
	quiz.RandomizeAnswers("taylor.json")
	quiz.StartApp()

}
