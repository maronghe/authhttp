package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main1() {

	url := "http://localhost:8080/admin"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Authorization", "Basic bG9nYW46MTIz")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
