package main

import (
	"fmt"
)


type todo struct {
	id int
	name string
	state bool
}

var todoList map[int]todo

func main() {
	id_track := 0

	todoList = make(map[int] todo)
	var (
		c1 = ""
		c2 = ""
		cid = 0
	)

	fmt.Println("use 'help' if braindead")
	
	for true {

		fmt.Print(">> ")
		fmt.Scan(&c1)
		if(c1 == "add") {
			fmt.Scan(&c2)
			addTodo(c2, id_track)
			id_track++
			fmt.Println("Todo added succesfully")
		} else if(c1 == "show") {
			displayTodo()
		} else if(c1 == "rem-i") { //Remove by Id
			fmt.Scan(&cid)
			removeTodo(cid)
			fmt.Println("Todo removed successfully")
		} else if(c1 == "exit") {
			fmt.Println("Exiting...")
			break
		} else if(c1 == "rem-n") { //Remove by Name
			fmt.Scan(&c2) 
			for id, value := range todoList {
				if(value.name == c2) {
					removeTodo(id)
					fmt.Println("Todo removed successfully")
					break
				}
			}
		} else if(c1 == "cngst-i") { //Change state by id
			fmt.Scan(&cid)
			stateChange(cid)
			fmt.Println("State changed")
		} else if(c1 == "cngst-n") {  //change state by name
			fmt.Scan(&c2)
			for id, value := range todoList {
				if(value.name == c2) {
					stateChange(id)
					fmt.Println("State changed")
					break
				}
			}
		} else if(c1 == "help") {
			fmt.Println("add <name> :- to add todo")
			fmt.Println("show :- to display todolist")
			fmt.Println("rem-i <id> :- to remove todo by id")
			fmt.Println("rem-n <name> :- to remove todo by name")
			fmt.Println("cngst-i <id> :- to change state(✅, ⬜) of todo by id")
			fmt.Println("cngst-n <name> :- to change state(✅, ⬜) of todo by name") 
			fmt.Println("exit :- to exit")
		} else {
			fmt.Println("Invalid command")
		}
		
	}
}

func ternary[T any](cond bool, a, b T) T {
    if cond {
        return a
    }
    return b
}

func addTodo(n string, idt int) {
	todoList[idt] = todo{name: n, state: false, id: idt}
}

func removeTodo(idt int) {
	delete(todoList, idt)
}

func displayTodo() {

	for _, value := range todoList {
		fmt.Println(value.id, ". ", value.name, "\t", ternary(value.state, "✅", "⬜"))
	}

}

func stateChange(idt int) {
	t := todoList[idt]

	t.state = !t.state

	todoList[idt] = t
}

