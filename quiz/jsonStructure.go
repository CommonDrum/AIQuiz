package quiz

import (
	"fmt"
	"math/rand"
)

type Option struct {
	A string `json:"A"`
	B string `json:"B"`
	C string `json:"C"`
	D string `json:"D"`
}

type Question struct {
	Topic       string `json:"topic"`
	Question    string `json:"question"`
	Options     Option `json:"options"`
	Answer      string `json:"answer"`
	Explanation string `json:"explanation"`
}

type Quiz struct {
	Questions []Question `json:"questions"`
}

func (question Question) Print() {

	fmt.Println(question.Question)
	fmt.Println()
	fmt.Println("A. " + question.Options.A)
	fmt.Println("B. " + question.Options.B)
	fmt.Println("C. " + question.Options.C)
	fmt.Println("D. " + question.Options.D)
	fmt.Println()

}

func (quiz Quiz) GetRandomQuestion() Question {
	randomIndex := rand.Intn(len(quiz.Questions))
	return quiz.Questions[randomIndex]
}
