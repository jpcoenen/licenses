package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// GodepsFile is the godeps file
const GodepsFile = "Godeps/Godeps.json"

// Godeps represents a godeps file
type Godeps struct {
	ImportPath   string
	GoVersion    string
	GodepVersion string
	Packages     []string
	Deps         []Dep
}

// Dep is a go dep
type Dep struct {
	ImportPath string
	Comment    string
	Rev        string
}

// LoadGodeps loads all of the go deps into a struct
func LoadGodeps() (*Godeps, error) {
	_, err := os.Stat(GodepsFile)
	if err != nil {
		return nil, nil
	}

	data, err := ioutil.ReadFile(GodepsFile)
	if err != nil {
		return nil, err
	}

	godeps := &Godeps{}
	err = json.Unmarshal(data, godeps)
	if err != nil {
		return nil, err
	}

	return godeps, nil
}
