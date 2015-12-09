package download

import (
	"io/ioutil"
	"net/http"

	. "../log"
)

func DownloadHTML(url string) string {

	respond, err := http.Get(url)
	Check(err, "Cant get the url. URL: "+url)
	bytes, e := ioutil.ReadAll(respond.Body)
	Check(e, "Can't read the respond boddy URL: "+url)

	respond.Body.Close()
	return string(bytes)
}

func WriteFile(content string, fileName string) {
	err := ioutil.WriteFile(fileName, []byte(content), 0777)
	Check(err, "Error occured when writing the file: "+fileName)
}
