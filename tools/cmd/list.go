package cmd

import (
	"context"
	"fmt"

	"github.com/Grimkey/cmdban/todolist"
	"github.com/spf13/cobra"
)

func addListCmd(reader todolist.Reader) *cobra.Command {
	var id string
	var openBrowser bool
	command := cobra.Command{
		Use:   "list",
		Short: "list of todo issues owned by user.",
		Long:  `print a list of issues owned from the backend that are currently owned by the active user..`,
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			if openBrowser {
				return open(reader, id)
			}
			if id == "" {
				return listAll(reader)
			}

			return issue(reader, id)
		},
	}

	command.Flags().StringVarP(&id, "id", "i", "", "id to list")
	command.Flags().BoolVarP(&openBrowser, "open", "o", false, "open in a browser")

	return &command
}

func issue(reader todolist.Reader, id string) error {
	issue, err := reader.Issue(context.Background(), id)
	if err != nil {
		return err
	}

	issue.Description()

	return nil
}

func listAll(reader todolist.Reader) error {
	page, err := reader.CurrentUser(context.Background())
	if err != nil {
		return err
	}

	for _, field := range page.Issues {
		field.ToString()
	}

	return nil
}

func open(reader todolist.Reader, id string) error {
	issue, err := reader.Issue(context.Background(), id)
	if err != nil {
		return err
	}

	tab := fmt.Sprintf("%s/browse/%s", reader.URL(), issue.Key)
	openbrowser(tab)

	return nil
}
