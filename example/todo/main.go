package main

import (
	"fmt"

	"github.com/worbridg/gonumber"
)

const (
	TODO = iota
	DOING
	DONE
)

type ToDoList []*ToDo

func (todos ToDoList) Statuses() gonumber.Numbers {
	numbers := gonumber.Numbers{}
	for i := 0; i < len(todos); i++ {
		numbers = append(numbers, todos[i].status)
	}
	return numbers
}

func (todos ToDoList) StartAll() {
	for i := 0; i < len(todos); i++ {
		todos[i].Start()
	}
}

func (todos ToDoList) DoneAll() {
	for i := 0; i < len(todos); i++ {
		todos[i].Done()
	}
}

type ToDo struct {
	status *gonumber.Number
	text   string
}

func NewToDo(text string) *ToDo {
	status, _ := gonumber.NewInt(TODO).ShouldBe(TODO, DOING, DONE)
	return &ToDo{
		status: status,
		text:   text,
	}
}

func (todo *ToDo) Start() error {
	todo.status.WillBe(DOING)
	if err := todo.status.Increment(); err != nil {
		return fmt.Errorf("invalid status")
	}
	fmt.Printf("%s start\n", todo.text)
	return nil
}

func (todo *ToDo) Suspend() error {
	todo.status.WillBe(TODO)
	if err := todo.status.Decrement(); err != nil {
		return fmt.Errorf("%s isn't doing", todo.text)
	}
	fmt.Printf("%s suspend\n", todo.text)
	return nil
}

func (todo *ToDo) Done() error {
	todo.status.WillBe(DONE)
	if err := todo.status.Increment(); err != nil {
		return fmt.Errorf("%s doesn't start yet", todo.text)
	}
	fmt.Printf("%s done\n", todo.text)
	return nil
}

func oneOf(todos ToDoList) *gonumber.OneOfNumbersState {
	return gonumber.OneOf(todos.Statuses())
}

func all(todos ToDoList) *gonumber.AllNumbersState {
	return gonumber.All(todos.Statuses())
}

func main() {
	todo := NewToDo("todo1")
	fmt.Println(todo.Done())
	fmt.Println(todo.Suspend())
	todo.Start()
	todo.Suspend()
	todo.Start()
	todo.Done()
	todos := ToDoList{todo, NewToDo("todo2"), NewToDo("todo3")}
	if _, exist := oneOf(todos).Is(TODO); exist {
		fmt.Println("there are a todo in TODO")
	}
	todos.StartAll()
	todos.DoneAll()
	if all(todos).Are(DONE) {
		fmt.Println("all tasks done")
	}
}
