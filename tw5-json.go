package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type tw5Tiddler struct {
	Title    string `json:"title,omitempty"`
	Created  string `json:"created,omitempty"`
	Creator  string `json:"creator,omitempty"`
	Modified string `json:"modified,omitempty"`
	Modifier string `json:"modifier,omitempty"`
	Type     string `json:"type,omitempty"`
	Text     string `json:"text,omitempty"`
	Revision string `json:"revision,omitempty"`
	Bag      string `json:"bag,omitempty"`
	Tags     string `json:"tags,omitempty"`
	List     string `json:"list,omitempty"`
}

// ProcessJson unmarshals the passed JSON data to extract
// the tiddlers, converts them to Markdown and saves them to
// the output folder
func ProcessJson(inputJson, outputDir string) error {
	tiddlers, err := unmarshalJson(inputJson)

	if err != nil {
		return fmt.Errorf("error loading JSON: %w", err)
	}

	for _, tid := range tiddlers {
		md, err := processTiddler(tid)

		if err != nil {
			return fmt.Errorf("error processing Tiddler %s: %w", tid.Title, err)
		} else {
			savePath := filepath.Join(outputDir, tid.Title)
			savePath = savePath + ".md"

			err := os.WriteFile(savePath, []byte(md), 0666)

			if err != nil {
				return fmt.Errorf("error saving Markdown %s: %w", savePath, err)
			}
		}
	}

	fmt.Printf("%d tiddlers processed.\n", len(tiddlers))

	return nil
}

// loadJson reads in the passed JSON file and returns an array
// of TW5Tiddler structs.
func unmarshalJson(jsondata string) ([]tw5Tiddler, error) {
	var tiddlers []tw5Tiddler

	err := json.Unmarshal([]byte(jsondata), &tiddlers)

	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	return tiddlers, nil
}
