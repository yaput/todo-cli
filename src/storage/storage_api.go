package storage_api

import (
	"context"

	"github.com/andygrunwald/go-jira"
)

type StorageService interface {
	GetAllTodoList() []TodoItem
	GetTodoListByID(id string) TodoItem
	CreateTodoList(title string, description string, status string, child []TodoItem)
}

type Storage struct {
	jiraClient JiraClient
}

type TodoItem struct {
	Title  string
	Status string
}

func NewStorage(jclient JiraClient) *Storage {
	return &Storage{
		jiraClient: jclient,
	}
}

func (s *Storage) GetAllTodoList() []TodoItem {
	return s.jiraClient.GetMyIssues()
}

func (s *Storage) GetTodoListByID(id string) TodoItem {
	return TodoItem{}
}

func (s *Storage) CreateTodoList(title string, description string, status string, child []TodoItem) {
}

type JiraClient struct {
	JiraURL   string
	Username  string
	Password  string
	client    *jira.Client
	SprintID  int
	accountId string
}

func NewJiraClient(jiraURL, username, password string, sprintId int) *JiraClient {
	tp := jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}
	client, err := jira.NewClient(tp.Client(), jiraURL)
	if err != nil {
		panic(err)
	}
	usr, _, errUsr := client.User.GetSelf()
	if errUsr != nil {
		panic(errUsr)
	}

	return &JiraClient{
		JiraURL:   jiraURL,
		Username:  username,
		Password:  password,
		client:    client,
		accountId: usr.AccountID,
		SprintID:  sprintId,
	}
}

func (j *JiraClient) GetMyIssues() []TodoItem {
	var issues []TodoItem
	sprints, _, err := j.client.Sprint.GetIssuesForSprintWithContext(context.Background(), j.SprintID)
	if err != nil {
		panic(err)
	}
	for _, sprint := range sprints {
		if sprint.Fields.Assignee != nil && sprint.Fields.Assignee.AccountID == j.accountId {
			issues = append(issues, TodoItem{
				Title:  sprint.Fields.Summary,
				Status: sprint.Fields.Status.Name,
			})
		}
	}
	return issues
}
