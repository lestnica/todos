package service

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo/pkg/repo"
)

type Service struct {
	repository repo.RepoTodo
}

func NewService() *Service {
	return &Service{
		repository: repo.NewMemoryRepo(),
	}
}

// Input Reader from the user
func InputReading(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

// Add TODO
func (s *Service) AddTodo(todo string) {
	s.repository.Add(todo)

	fmt.Println("TODO Added")
}

// Remove TODO
func (s *Service) RemoveTodo(todo string) {
	s.repository.Remove(todo)

	fmt.Println("TODO Removed")
}

// Get TODOs
func (s *Service) GetTodo() []string {
	return s.repository.Get()
}

func (s *Service) RunTodoCLI() {
	for {
		fmt.Println("1. Add TODO")
		fmt.Println("2. Remove TODO")
		fmt.Println("3. Get TODOs")
		fmt.Println("4. Exit")

		choice := InputReading("Choose action: ")

		switch choice {
		case "1":
			todo := InputReading("Add TODO: ")
			s.AddTodo(todo)
		case "2":
			todo := InputReading("Remove TODO: ")
			s.RemoveTodo(todo)
		case "3":
			todos := s.GetTodo()
			fmt.Println("Get list TOODs:")
			for _, todo := range todos {
				fmt.Println(todo)
			}
		case "4":
			fmt.Println("FINISHED")
			return
		default:
			fmt.Println("Incorrect choice or something...")
		}

		fmt.Println()
	}
}
