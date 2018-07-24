package model

type Profile struct {
	NickName   string
	Age        int
	Gender     string
	Height     int
	Weight     int
	Salary     string
	Marriage   string //婚况
	Education  string
	Occupation string //职业
	WorkCity   string
	Location   string //籍贯
	House      string
	Car        string
}

type Joke struct {
	Title   string
	Context string
}
