package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	
	"example/interface/note"
	"example/interface/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = saveData(todo)

	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}


//any value allowed types
// func printSomething(val interface{}){
// 	fmt.Println(val)
// }

// switch value.(type)

// func add(a,b interface{})interface{}{
// 	aint,aisint:=a.(int)
// 	bint,bisint:=b.(int)
// 	if aisint && bisint{
// 		return aint+bint
// 	}

// }

//generics
// func main(){
// 	result:=add(1,2)
// 	fmt.Print(result)
// }
// func add[T int|float64|string](a,b T)T{
// 	return a+b
// }