package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Configs struct {
	Questions      []Question `json:"questions"`
	TemplateCommit string     `json:"templateCommit"`
}

type Question struct {
	Id       string   `json:"id"`
	Type     string   `json:"type"`
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
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exePath := filepath.Dir(ex)

	jsonFile, err := os.Open(exePath + "/config.json")

	if err != nil {
		fmt.Println("need a file config.json", err)
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var configs Configs

	json.Unmarshal(byteValue, &configs)

	return configs
}
