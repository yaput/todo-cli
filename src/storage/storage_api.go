package storage_api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/andygrunwald/go-jira"
)

type StorageService interface {
	GetAllTodoList() map[string][]TodoItem
	GetTodoListByID(id string) TodoItem
	CreateTodoList(title string, description string, status string, child []TodoItem)
}

type Storage struct {
	jiraClient JiraClient
}

type TodoItem struct {
	ID     string
	Title  string
	Status string
}

func NewStorage(jclient JiraClient) *Storage {
	return &Storage{
		jiraClient: jclient,
	}
}

func (s *Storage) GetAllTodoList() map[string][]TodoItem {
	todoItems, err := s.jiraClient.GetMyIssues()
	if err != nil {
		fmt.Println("Error getting issues from Jira")
		fmt.Println(err)
		return make(map[string][]TodoItem)
	}
	return todoItems
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
	BoardID   string
	accountId string
	sprintID  int
}

func NewJiraClient(jiraURL, username, password, boardID string) *JiraClient {
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

	sprint, err := getActiveSprint(client, boardID)
	if err != nil {
		panic(err)
	}

	return &JiraClient{
		JiraURL:   jiraURL,
		Username:  username,
		Password:  password,
		client:    client,
		accountId: usr.AccountID,
		BoardID:   boardID,
		sprintID:  sprint.ID,
	}
}

func (j *JiraClient) GetMyIssues() (map[string][]TodoItem, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/sprint/%d/issue", j.sprintID)

	// Parse the URL
	u, _ := url.Parse(apiEndpoint)

	// Create URL values
	q := url.Values{}
	q.Add("jql", "(status != 'Completed' AND status != 'Ready for Production' AND status != 'DELIVERED TO PRODUCTION') AND assignee = "+j.accountId)
	q.Add("validateQuery", "true")
	q.Add("fields", "summary")
	q.Add("fields", "status")

	// Add the query parameters to the URL
	u.RawQuery = q.Encode()

	parsedEndpoint := fmt.Sprint(u.String())
	req, err := j.client.NewRequestWithContext(context.Background(), "GET", parsedEndpoint, nil)
	if err != nil {
		return nil, err
	}
	issues := new(jira.IssuesInSprintResult)
	_, err = j.client.Do(req, issues)
	if err != nil {
		return nil, err
	}

	mappedIssues := make(map[string][]TodoItem)
	for _, issue := range issues.Issues {
		todoItem := TodoItem{
			ID:     issue.ID,
			Title:  issue.Fields.Summary,
			Status: issue.Fields.Status.Name,
		}
		if mappedIssues[todoItem.Status] != nil {
			mappedIssues[todoItem.Status] = append(mappedIssues[todoItem.Status], todoItem)
		} else {
			mappedIssues[todoItem.Status] = []TodoItem{todoItem}
		}
	}
	return mappedIssues, nil
}

func getActiveSprint(jiraClient *jira.Client, boardID string) (*jira.Sprint, error) {
	apiEndpoint := fmt.Sprintf("rest/agile/1.0/board/%s/sprint?state=active", boardID)
	req, err := jiraClient.NewRequestWithContext(context.Background(), "GET", apiEndpoint, nil)
	if err != nil {
		return nil, err
	}

	sprint := new(jira.SprintsList)
	_, err = jiraClient.Do(req, sprint)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &sprint.Values[0], nil
}
