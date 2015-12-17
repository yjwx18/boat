package configuration

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	. "../log"
)

type WorkRequest struct {
	Name                string
	Delay               int
	Matches             []MatchExpression
	PagePattern         string
	MinPageRange        int
	MaxPageRange        int
	NumberOfDownloaders int
}

type MatchExpression struct {
	Name  string
	Regex string
}

func ReadConfiguration() []*WorkRequest {
	var configs []*WorkRequest
	relativeDir := "../configuration/"
	files, err := ioutil.ReadDir(relativeDir)
	Check(err, "Can't find the configuration dir")
	for _, fileinfo := range files {
		name := fileinfo.Name()
		if !fileinfo.IsDir() && filepath.Ext(name) == ".config" {
			f, e := os.Open(filepath.Join(relativeDir, name))
			Check(e, "Error occured when read file :"+name)
			decoder := json.NewDecoder(f)
			configuration := &WorkRequest{}
			decodeeError := decoder.Decode(&configuration)
			Check(decodeeError, "Error when decode file :"+name)

			configs = append(configs, configuration)
		}
	}
	return configs
}

func GetUrl(wr *WorkRequest) map[string]string {

	result := make(map[string]string)
	//checkings
	min := strconv.Itoa(int(wr.MinPageRange))
	max := strconv.Itoa(int(wr.MaxPageRange))
	if max < min {
		panic("The max value is smaller than the min value in the configuration of " + wr.Name)
	}

	for i := wr.MinPageRange; i <= wr.MaxPageRange; i++ {
		istring := strconv.Itoa(int(i))
		result[istring] = strings.Replace(wr.PagePattern, "[]", istring, -1)
		//result = append(result, element)
	}

	return result
}
