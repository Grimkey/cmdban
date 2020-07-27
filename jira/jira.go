package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Grimkey/cmdban/todolist"
)

type Config struct {
	User         string `json:"user"`
	Key          string `json:"key"`
	URL          string `json:"url"`
	DefaultBoard string `json:"defaultBoard"`
}

type client struct {
	cfg  Config
	http *http.Client
}

func New(cfg Config) todolist.Reader {
	return &client{
		cfg:  cfg,
		http: http.DefaultClient,
	}
}

func (jira *client) URL() string {
	return jira.cfg.URL
}

func (jira *client) Board() string {
	return fmt.Sprintf("%s/jira/software/projects/%s", jira.cfg.URL, jira.cfg.DefaultBoard)
}

func (jira *client) Issue(ctx context.Context, id string) (todolist.Issue, error) {
	var issue todolist.Issue

	query := fmt.Sprintf("%s/rest/api/2/issue/%s", jira.cfg.URL, id)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, query, nil)
	if err != nil {
		return issue, err
	}

	if err = jira.do(req, &issue); err != nil {
		return issue, err
	}

	return issue, nil
}
func (jira *client) CurrentUser(ctx context.Context) (todolist.IssuePage, error) {
	var issues todolist.IssuePage

	const doneAndCurrentUser = `%20status%20!%3D%20"Done"%20AND%20assignee%20%3D%20currentUser()`
	query := fmt.Sprintf("%s/rest/api/2/search?jql=%s", jira.cfg.URL, doneAndCurrentUser)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, query, nil)
	if err != nil {
		return issues, err
	}

	if err = jira.do(req, &issues); err != nil {
		return issues, err
	}

	return issues, nil
}

func (jira *client) do(req *http.Request, o interface{}) error {
	req.SetBasicAuth(jira.cfg.User, jira.cfg.Key)
	req.Header.Add("accept", "application/json")

	resp, err := jira.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status result %d", resp.StatusCode)
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respByte, &o)
	if err != nil {
		return err
	}

	// err = json.NewDecoder(resp.Body).Decode(&issues)
	// if err != nil {
	// 	return err
	// }

	return nil
}
