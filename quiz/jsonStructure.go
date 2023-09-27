package quiz

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
