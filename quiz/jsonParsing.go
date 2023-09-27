package quiz

import "math/rand"

func RandomizeAnswers(fileName string) {
	quiz := loadQuiz(fileName)

	for i := range quiz.Questions {
		options := []string{
			quiz.Questions[i].Options.A,
			quiz.Questions[i].Options.B,
			quiz.Questions[i].Options.C,
			quiz.Questions[i].Options.D}

		var answerIndex int

		switch quiz.Questions[i].Answer {
		case "A":
			answerIndex = 0
		case "B":
			answerIndex = 1
		case "C":
			answerIndex = 2
		case "D":
			answerIndex = 3
		}

		newAnswerIndex := rand.Intn(4)

		options[answerIndex], options[newAnswerIndex] = options[newAnswerIndex], options[answerIndex] // Wow I love this swap in go!

		quiz.Questions[i].Options.A = options[0]
		quiz.Questions[i].Options.B = options[1]
		quiz.Questions[i].Options.C = options[2]
		quiz.Questions[i].Options.D = options[3]

		switch newAnswerIndex {
		case 0:
			quiz.Questions[i].Answer = "A"
		case 1:
			quiz.Questions[i].Answer = "B"
		case 2:
			quiz.Questions[i].Answer = "C"
		case 3:
			quiz.Questions[i].Answer = "D"
		}
	}

	saveQuiz(fileName, quiz)
}
