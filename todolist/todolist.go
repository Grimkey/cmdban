package todolist

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type IssueFields struct {
	Statuscategorychangedate string `json:"statuscategorychangedate"`
	Issuetype                struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Description string `json:"description"`
		IconURL     string `json:"iconUrl"`
		Name        string `json:"name"`
		Subtask     bool   `json:"subtask"`
		AvatarID    int    `json:"avatarId"`
		EntityID    string `json:"entityId"`
	} `json:"issuetype"`
	Timespent interface{} `json:"timespent"`
	Project   struct {
		Self           string `json:"self"`
		ID             string `json:"id"`
		Key            string `json:"key"`
		Name           string `json:"name"`
		ProjectTypeKey string `json:"projectTypeKey"`
		Simplified     bool   `json:"simplified"`
	} `json:"project"`
	FixVersions        []interface{} `json:"fixVersions"`
	Aggregatetimespent interface{}   `json:"aggregatetimespent"`
	Resolution         interface{}   `json:"resolution"`
	Resolutiondate     interface{}   `json:"resolutiondate"`
	Workratio          int           `json:"workratio"`
	Watches            struct {
		Self       string `json:"self"`
		WatchCount int    `json:"watchCount"`
		IsWatching bool   `json:"isWatching"`
	} `json:"watches"`
	LastViewed interface{} `json:"lastViewed"`
	Created    string      `json:"created"`
	Priority   struct {
		Self    string `json:"self"`
		IconURL string `json:"iconUrl"`
		Name    string `json:"name"`
		ID      string `json:"id"`
	} `json:"priority"`
	Labels                        []string      `json:"labels"`
	Aggregatetimeoriginalestimate interface{}   `json:"aggregatetimeoriginalestimate"`
	Timeestimate                  interface{}   `json:"timeestimate"`
	Versions                      []interface{} `json:"versions"`
	Issuelinks                    []interface{} `json:"issuelinks"`
	Assignee                      struct {
		Self         string `json:"self"`
		AccountID    string `json:"accountId"`
		EmailAddress string `json:"emailAddress"`
		DisplayName  string `json:"displayName"`
		Active       bool   `json:"active"`
		TimeZone     string `json:"timeZone"`
		AccountType  string `json:"accountType"`
	} `json:"assignee"`
	Updated string `json:"updated"`
	Status  struct {
		Self           string `json:"self"`
		Description    string `json:"description"`
		IconURL        string `json:"iconUrl"`
		Name           string `json:"name"`
		ID             string `json:"id"`
		StatusCategory struct {
			Self      string `json:"self"`
			ID        int    `json:"id"`
			Key       string `json:"key"`
			ColorName string `json:"colorName"`
			Name      string `json:"name"`
		} `json:"statusCategory"`
	} `json:"status"`
	Components            []interface{} `json:"components"`
	Timeoriginalestimate  interface{}   `json:"timeoriginalestimate"`
	Description           string        `json:"description"`
	Security              interface{}   `json:"security"`
	Aggregatetimeestimate interface{}   `json:"aggregatetimeestimate"`
	Summary               string        `json:"summary"`
	Creator               struct {
		Self        string `json:"self"`
		AccountID   string `json:"accountId"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
		TimeZone    string `json:"timeZone"`
		AccountType string `json:"accountType"`
	} `json:"creator"`
	Subtasks []interface{} `json:"subtasks"`
	Reporter struct {
		Self        string `json:"self"`
		AccountID   string `json:"accountId"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
		TimeZone    string `json:"timeZone"`
		AccountType string `json:"accountType"`
	} `json:"reporter"`
	Aggregateprogress struct {
		Progress int `json:"progress"`
		Total    int `json:"total"`
	} `json:"aggregateprogress"`
	Environment interface{} `json:"environment"`
	Duedate     interface{} `json:"duedate"`
	Progress    struct {
		Progress int `json:"progress"`
		Total    int `json:"total"`
	} `json:"progress"`
	Votes struct {
		Self     string `json:"self"`
		Votes    int    `json:"votes"`
		HasVoted bool   `json:"hasVoted"`
	} `json:"votes"`
}

type Issue struct {
	Expand string      `json:"expand"`
	ID     string      `json:"id"`
	Self   string      `json:"self"`
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields,omitempty"`
}

func (iss Issue) ToString() {
	a := iss.Key
	b := iss.Fields.Summary
	title := color.New(color.FgWhite, color.Bold)
	title.Printf("%s - %s\n", a, b)
}

func (iss Issue) Description() {
	label := color.New(color.FgYellow)

	iss.ToString()
	label.Print("Description: ")
	fmt.Println(iss.Fields.Issuetype.Description)
	label.Print("Labels: ")
	fmt.Println(strings.Join(iss.Fields.Labels, ","))
	label.Print("CreatedAt: ")
	fmt.Println(iss.Fields.Created)
	label.Print("Status: ")
	fmt.Println(iss.Fields.Status.Name)
}

// Walk pages: http://localhost:8080/rest/api/2/issue/createmeta/TEST/issuetypes?startAt=0&maxResults=50
// Source: https://developer.atlassian.com/server/jira/platform/jira-rest-api-examples/
type IssuePage struct {
	Expand     string  `json:"expand"`
	StartAt    int     `json:"startAt"`
	MaxResults int     `json:"maxResults"`
	Total      int     `json:"total"`
	Issues     []Issue `json:"issues"`
}

type Reader interface {
	CurrentUser(ctx context.Context) (IssuePage, error)
	Issue(ctx context.Context, id string) (Issue, error)
	URL() string
	Board() string
}

type Writer interface {
}

type Backend interface {
	Reader
	Writer
}
