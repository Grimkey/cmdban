package main

import (
	"github.com/Grimkey/cmdban/config"
	"github.com/Grimkey/cmdban/tools/cmd"

	"github.com/Grimkey/cmdban/jira"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	cfg, err := config.LoadResourceFile(".cmdbanrc")
	check(err)

	client := jira.New(cfg.Jira)
	cmd.Execute(client)
}
