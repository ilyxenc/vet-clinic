package app

// Структура, описывающая поля данных каждого животного
type Animal struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Age  int    `json:"age"`
}