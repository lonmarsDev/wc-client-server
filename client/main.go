package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	url      string = "http://localhost:8081/"
	filePath string = "GoLang_Test.txt"
)

func main() {

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", url, file)

	req.Header.Add("content-type", "text/plain")
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	fmt.Println(string(body))

}
