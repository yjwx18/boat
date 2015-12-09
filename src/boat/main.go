package main

import (
	"fmt"
	"os"

	"../configuration"
)

func main() {
	//Global variable
	configs := configuration.ReadConfiguration()
	for _, v := range configs {
		fmt.Println("--")
		fmt.Println(configuration.WorkRequest(*v).Name)
	}

}

//local functions
func check(e error, m string) {

	if e != nil {
		panic(m + "\n" + e.Error())
		os.Exit(1)
	}
}
