package main

import (
	interface_api "github.com/yaput/todo-cli/src/interface"
	storage_api "github.com/yaput/todo-cli/src/storage"
)

func main() {
	jclient := storage_api.NewJiraClient("https://29022131.atlassian.net", "antonius.putra@traveloka.com", "ATATT3xFfGF0WawsusrVopIfRqyNsxrXvtvowPNqsEb_tuqu-O_ojR8lxvSu-C2cc2tfSRGN9srovv9zkenNY-RnnTIKgC7fI7vjjswHCLPgKjGu3dfJrzyA2cMorv85TWnCD-9qJ7buuLxzM-g3dMBhAXSpz2zCuYk61-X7wFWpvx04WbrriTs=73C74C19", 20812)
	todoList := interface_api.NewInterface(*storage_api.NewStorage(*jclient))

	todoList.StartTodoList()
}
