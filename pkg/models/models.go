package models

type Agify struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}
type Genderize struct {
	Count       int    `json:"count"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Probability int    `json:"probability"`
}
type Nationalize struct {
	Count   int       `json:"count"`
	Name    string    `json:"name"`
	Country []country `json:"country"`
}
type country struct {
	Country_id  string
	Probability int
}
type Email struct {
	Email string `json:"email"`
}
type AddingHumanStruct struct {
	Name      string `json:"name"`
	Secname   string `json:"secname"`
	Thirdname string `json:"thirdname"`
	Gender    string `json:"gender"`
	Nation    string `json:"nation"`
	Age       int    `json:"age"`
}
