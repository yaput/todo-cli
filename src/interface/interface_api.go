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
	fmt.Println("===============ToDo-CLI===============")
	for _, todoItem := range todoItems {
		fmt.Printf("=> (%s) -- %s\n", todoItem.Status, todoItem.Title)
	}
	fmt.Println("======================================")
}

func NewInterface(storage_api storage_api.Storage) *Interface {
	return &Interface{storage_service: &storage_api}
}
