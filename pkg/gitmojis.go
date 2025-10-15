package pkg

type Gitmoji struct {
	Emoji       string `json:"emoji"`
	Entity      string `json:"entity"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Scope       string `json:"scope"`
}

type Gitmojis struct {
	Gitmojis []Gitmoji `json:"gitmojis"`
}
