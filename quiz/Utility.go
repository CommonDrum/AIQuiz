package quiz

import (
	"encoding/json"
	"log"
	"os"
)

func loadQuiz(fileName string) Quiz {

	jsonFile, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	var quiz Quiz

	err = json.Unmarshal(jsonFile, &quiz)

	if err != nil {
		log.Fatal(err)
	}

	return quiz

}

func saveQuiz(fileName string, quiz Quiz) {

	jsonData, err := json.MarshalIndent(quiz, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling the quiz: %v", err)
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

type RepetitionDetector struct {
	int map[int]bool
}

func (detector *RepetitionDetector) Add(i int) bool {
	if detector.int[i] {
		return true
	}
	detector.int[i] = true
	return false
}

func (detector *RepetitionDetector) Reset() {
	detector.int = make(map[int]bool)
}

func (detector *RepetitionDetector) Get(i int) bool {
	return detector.int[i]
}
