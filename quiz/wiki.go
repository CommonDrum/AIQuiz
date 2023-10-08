package quiz

type Wiki struct {
	Definitons []Definition `json:"definitions"`
}

type Definition struct {
	Name       string `json:"name"`
	Definition string `json:"definition"`
}

func (wiki Wiki) Search(name string) Definition {
	for _, definition := range wiki.Definitons {
		if definition.Name == name {
			return definition
		}
	}
	return Definition{}
}
