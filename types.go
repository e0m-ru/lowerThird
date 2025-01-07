package main

type LowerThird struct {
	Surname, Name, Title string
}

// Возвращает поля как слайс строк
func (L *LowerThird) likeSliceOfStrings() []string {
	return []string{
		L.Surname,
		L.Name,
		L.Title,
	}
}
