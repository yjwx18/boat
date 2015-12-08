package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func main() {
	//global variable
	var configs []*WorkRequest
	files, err := ioutil.ReadDir("./")
	check(err, "Can't find the configuration dir")
	for _, fileinfo := range files {
		name := fileinfo.Name()
		if !fileinfo.IsDir() && filepath.Ext(name) == ".config" {
			f, e := os.Open(name)
			check(e, "Error occured when read file :"+name)
			decoder := json.NewDecoder(f)
			configuration := &WorkRequest{}
			decodeeError := decoder.Decode(&configuration)
			check(decodeeError, "Error when decode file :"+name)
			fmt.Println(configuration.Name)
			for _, match := range configuration.Matches {
				fmt.Println(match.Name + "|" + match.Regex)
			}
			configs = append(configs, configuration)
		}
	}
}

type WorkRequest struct {
	Name    string
	Dalay   time.Duration
	Matches []MatchExpression
}

type MatchExpression struct {
	Name  string
	Regex string
}

//local functions
func check(e error, m string) {

	if e != nil {
		panic(m + "\n" + e.Error())
		os.Exit(1)
	}
}
