package config

import (
	"encoding/json"
	"os"

	"github.com/Grimkey/cmdban/jira"
)

// Resource contains configuration information for this run.
type Resource struct {
	Jira jira.Config `json:"jira"`
}

// LoadResourceFile loads type config.Resource from a file
func LoadResourceFile(s string) (Resource, error) {
	var res Resource

	f, err := os.Open(s)
	if err != nil {
		return res, err
	}

	err = json.NewDecoder(f).Decode(&res)
	if err != nil {
		return res, err
	}

	return res, nil
}
