package viewmodel

type Base struct {
	Title string
}

func CreateBase() Base {
	return Base{
		Title: "Limonada",
	}
}
