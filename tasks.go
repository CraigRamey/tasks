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

	"github.com/beego/bee/logger/colors"
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
	todos := getTodos()
	for _, todo := range todos {
		if todo.IsComplete == true {
			fmt.Printf("%s is complete", colors.Blue(todo.Task))
		} else {
			fmt.Printf("%s is not yet complete", colors.Red(todo.Task))
		}
		fmt.Println(todo.toString())
	}
}
