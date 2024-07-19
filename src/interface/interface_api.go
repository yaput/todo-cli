package interface_api

import (
	"fmt"

	storage_api "github.com/yaput/todo-cli/src/storage"
)

type Interface struct {
	storage_service storage_api.StorageService
}

func (i *Interface) StartTodoList() {
	todoItems := i.storage_service.GetAllTodoList()
	fmt.Println("==================================================ToDo-CLI==================================================")
	fmt.Println("Blocked:")
	for _, item := range todoItems["Blocked"] {
		printTodoItem(item.ID, item.Title)
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("✍️ To Do:")
	for _, item := range todoItems["✍️ To Do"] {
		printTodoItem(item.ID, item.Title)
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("In Progress:")
	for _, item := range todoItems["In Progress"] {
		printTodoItem(item.ID, item.Title)
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Ready to Production:")
	for _, item := range todoItems["Ready to Production"] {
		printTodoItem(item.ID, item.Title)
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}

func NewInterface(storage_api storage_api.Storage) *Interface {
	return &Interface{storage_service: &storage_api}
}

func printTodoItem(itemId, itemTitle string) {
	fmt.Printf("=> (%s) %s \n", itemId, itemTitle)
}
