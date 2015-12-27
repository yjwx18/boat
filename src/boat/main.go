package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	. "../configuration"
	. "../download"
)

func main() {
	//Global variable
	configs := ReadConfiguration()
	var wg sync.WaitGroup
	for _, v := range configs {
		fmt.Println("--")
		fmt.Println(WorkRequest(*v).Name)
	}

	for _, s := range configs {

		if WorkRequest(*s).Name == "post2u" {

			dirName := "." + string(filepath.Separator) + "outputs" + string(filepath.Separator) + WorkRequest(*s).Name + string(filepath.Separator) + time.Now().Format("2006-Jan-02") + string(filepath.Separator)
			os.MkdirAll(dirName, 0777)
			workerNumber := WorkRequest(*s).NumberOfDownloaders
			jobChan := make(chan int, workerNumber)
			urlMap := GetUrl(s)
			delay := WorkRequest(*s).Delay
			for i, url := range urlMap {
				jobChan <- 1
				wg.Add(1)
				go DoJobs(url, dirName, i, jobChan, delay, &wg)
			}
		}
	}
	wg.Wait()

}

func DoJobs(url string, dirName string, i string, jobChan chan int, delay int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(delay))
	fmt.Println(i + " started")
	WriteFile(DownloadHTML(url), dirName+string(filepath.Separator)+i+".html")
	fmt.Println(i + " has been saved")
	<-jobChan
}
