package cmd

import (
	"embed"
	"encoding/json"
)

var commandMapping map[string]string

//go:embed config.json
var configFile embed.FS

func init() {
	if len(commandMapping) == 0 {
		data, _ := configFile.ReadFile("config.json")
		json.Unmarshal([]byte(data), &commandMapping)
	}
}
