package main

import (
	"fmt"
	"os"
	"time"

	. "../configuration"
	. "../download"
)

func main() {
	//Global variable
	configs := ReadConfiguration()
	for _, v := range configs {
		fmt.Println("--")
		fmt.Println(WorkRequest(*v).Name)
	}

	for _, s := range configs {

		if WorkRequest(*s).Name == "post2u" {
			dirName := "outputs\\" + WorkRequest(*s).Name + "\\" + time.Now().Format("2006-Jan-02")
			os.MkdirAll(dirName, 0777)
			for i, url := range GetUrl(s) {
				content := DownloadHTML(url)
				WriteFile(content, dirName+"\\"+i+".html")
				fmt.Println(i + " has been saved")
			}

		}
	}

}
