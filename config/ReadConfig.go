package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Configs struct {
	Questions []Question `json:"questions"`
}

type Question struct {
	Items []Item `json:"items"`
	Label string `json:"label"`
}

type Item struct {
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
