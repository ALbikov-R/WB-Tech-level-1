package main

import "fmt"

//Объявление структуры Human
type Human struct {
	Name string
	Age  uint8
}

//Метод структуры Human
func (h *Human) Hello() {
	fmt.Println("Hi, my name is", h.Name, "and I'm", h.Age)
}

//Объявление структуры Action
type Action struct {
	Human      //Встраивание в структуру Action структуру Human
	actionType string
}

//Метод, использующие поля структуры Human
func (a *Action) Do(actionType string) {
	fmt.Println(a.Name, "is going to", actionType, "now!")
}

func main() {
	//Создание объекта Action со встроенной структурой Human
	person := Action{
		Human: Human{
			Name: "Rasim",
			Age:  22,
		},
		actionType: "work",
	}
	person.Do(person.actionType) //Вызов метода структуры Action, использующие поля Human
	person.Hello()               //Вызов метода структуры Human, демонстрируя встраивание
}
