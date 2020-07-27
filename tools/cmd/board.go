package cmd

import (
	"github.com/Grimkey/cmdban/todolist"
	"github.com/spf13/cobra"
)

func boardCmd(reader todolist.Reader) *cobra.Command {
	command := cobra.Command{
		Use:   "board",
		Short: "open a tab to the jira board.",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			openbrowser(reader.Board())
			return nil
		},
	}

	return &command
}