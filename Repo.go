package main

import "fmt"

var currentId int

var todos Todos

// Give us some seed data
func init() {
  RepoCreateTodo(Todo{Name: "Write presentation"})
  RepoCreateTodo(Todo{Name: "Host meetup"})
}

func RepoFindTodo(id int) Todo {
  for _, t := range todos {
	if t.Id == id {
	  return t
	}
  }
  // return empty Todo it not foundf
  return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
  currentId += 1
  t.Id = currentId
  todos = append(todos, t)
  return t
}

func RepoDestroyTodo(id int) error {
  for i, t := range todos {
	if t.Id == id {
	  todos = append(todos[:i], todos[i+1:]...)
	  return nil
	}
  }
  return fmt.Errorf("Cound not find Todo with id of %d to delete", id)
}