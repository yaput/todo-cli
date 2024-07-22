package main

import (
	"os"

	interface_api "github.com/yaput/todo-cli/src/interface"
	storage_api "github.com/yaput/todo-cli/src/storage"
)

func main() {
	email := os.Getenv("JIRA_EMAIL")
	token := os.Getenv("JIRA_TOKEN")
	jiraBoardID := os.Getenv("JIRA_BOARD_ID")
	jiraUrl := os.Getenv("JIRA_URL")
	jclient := storage_api.NewJiraClient(jiraUrl, email, token, jiraBoardID)
	todoList := interface_api.NewInterface(*storage_api.NewStorage(*jclient))

	todoList.StartTodoList()
}
