package main

import (
	"quiz/quiz"
)

func main() {
	//TODO: create a input system that handles quries for both the quiz and the wiki
	quiz.RandomizeAnswers("questions.json")
	quiz.StartApp()

}
