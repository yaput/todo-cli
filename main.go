package main

import (
	"os"

	interface_api "github.com/yaput/todo-cli/src/interface"
	storage_api "github.com/yaput/todo-cli/src/storage"
)

func main() {
	token := os.Getenv("JIRA_TOKEN")
	jclient := storage_api.NewJiraClient("https://29022131.atlassian.net", "antonius.putra@traveloka.com", token, "847")
	todoList := interface_api.NewInterface(*storage_api.NewStorage(*jclient))

	todoList.StartTodoList()
}
