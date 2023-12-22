package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Configs struct {
	Questions      []Question `json:"questions"`
	TemplateCommit string     `json:"templateCommit"`
}

type Question struct {
	Id       string   `json:"id"`
	Label    string   `json:"label"`
	Options  []Option `json:"options"`
	ErrorMsg string   `json:"errorMsg"`
	Min      int      `json:"min"`
	Max      int      `json:"max"`
}

type Option struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func ReadConfigs() Configs {
	jsonFile, err := os.Open("./config.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var configs Configs

	json.Unmarshal(byteValue, &configs)

	return configs
}
