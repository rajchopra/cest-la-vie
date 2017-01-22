package main

import (
	"net/http"
	"strings"
	"io/ioutil"

)

func main() {

	reader := strings.NewReader(`{"body":123}`)
	request, err := http.NewRequest("GET", "http://ifconfig.co", reader)
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		println(err.Error())
	}

	// print(resp.Body)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	println(body)



}
