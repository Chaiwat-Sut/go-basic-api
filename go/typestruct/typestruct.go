package typestruct

type Student struct {
	Lastname  string  `json:"lastname"`
	Firstname string  `json:"firstname"`
	SSN       string  `json:"ssn"`
	Test1     float64 `json:"test1"`
	Test2     float64 `json:"test2"`
	Test3     float64 `json:"test3"`
	Test4     float64 `json:"test4"`
	Final     float64 `json:"final"`
	Grade     string  `json:"grade"`
}
