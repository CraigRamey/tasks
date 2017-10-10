package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/fatih/color"
)

type Todo struct {
	Task       string `json:"task"`
	IsComplete bool   `json:"isComplete"`
}

func (t Todo) toString() string {
	bytes, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(bytes)
}

//TODO: updateTodo function
//TODO: deleteTodo function
func deleteTodo(n int) error {
	todos := getTodos()
	var updatedTodos []Todo
	for i, todo := range todos {
		num := i + 1
		if num != n {
			updatedTodos = append(updatedTodos, todo)
		}
	}

	bytes, err := json.Marshal(updatedTodos)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	path, err := getPath()
	if err != nil {
		log.Fatal(err)
	}

	file := filepath.Join(path, ".todos.json")
	ioutil.WriteFile(file, bytes, 0755)
	return nil
}

func addTodo(t string) error {
	todos := getTodos()
	todo := Todo{Task: t, IsComplete: false}
	todos = append(todos, todo)
	bytes, err := json.Marshal(todos)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	path, err := getPath()
	if err != nil {
		log.Fatal(err)
	}

	file := filepath.Join(path, ".todos.json")
	ioutil.WriteFile(file, bytes, 0755)
	return nil
}

func getTodos() []Todo {
	var todos []Todo
	path, err := getPath()
	if err != nil {
		log.Fatal(err)
	}

	file := filepath.Join(path, ".todos.json")
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &todos)
	return todos
}

func getPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", errors.New("Cannot get home directory")
	}

	path := usr.HomeDir
	return path, nil
}

func main() {
	path, err := getPath()
	if err != nil {
		log.Fatal(err)
	}

	file := filepath.Join(path, ".todos.json")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		ioutil.WriteFile(file, nil, 0755)
	}

	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	todos := getTodos()

	for i, todo := range todos {
		num := i + 1
		if todo.IsComplete == true {
			fmt.Printf("%d. %s %s\n", num, blue("[X]"), red(todo.Task))
		} else {
			fmt.Printf("%d. %s %s\n", num, blue("[ ]"), red(todo.Task))
		}
	}
}
