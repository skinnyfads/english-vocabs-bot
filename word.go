package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/skinnyfads/english-vocabs-bot/config"
)

type Definition struct {
	Definition string `json:"definition"`
}

type Meaning struct {
	Definitions []Definition `json:"definitions"`
}

type Entry struct {
	Meanings []Meaning `json:"meanings"`
}

func GetMeaning(word string) (string, error) {
	url := fmt.Sprintf(config.Env.BaseURL+"%s", word)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("API error: %s", string(body))
	}

	var entries []Entry
	err = json.NewDecoder(resp.Body).Decode(&entries)
	if err != nil {
		return "", err
	}

	var definitions []string
	for _, entry := range entries {
		for _, meaning := range entry.Meanings {
			for _, def := range meaning.Definitions {
				definitions = append(definitions, def.Definition)
			}
		}
	}

	return strings.Join(definitions, "\n"), nil
}

