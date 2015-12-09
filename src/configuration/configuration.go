package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type WorkRequest struct {
	Name         string
	Dalay        time.Duration
	Matches      []MatchExpression
	PagePattern  string
	MinPageRange int64
	MaxPageRange int64
}

type MatchExpression struct {
	Name  string
	Regex string
}

func ReadConfiguration() []*WorkRequest {
	var configs []*WorkRequest
	relativeDir := "../configuration/"
	files, err := ioutil.ReadDir(relativeDir)
	check(err, "Can't find the configuration dir")
	for _, fileinfo := range files {
		name := fileinfo.Name()
		if !fileinfo.IsDir() && filepath.Ext(name) == ".config" {
			f, e := os.Open(filepath.Join(relativeDir, name))
			check(e, "Error occured when read file :"+name)
			decoder := json.NewDecoder(f)
			configuration := &WorkRequest{}
			decodeeError := decoder.Decode(&configuration)
			check(decodeeError, "Error when decode file :"+name)

			configs = append(configs, configuration)
		}
	}
	return configs
}

//local functions
func check(e error, m string) {

	if e != nil {
		panic(m + "\n" + e.Error())
		os.Exit(1)
	}
}
