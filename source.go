package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func main() {

	fmt.Println("Enter the Webhook that you want to delete : ")
	var url string
	fmt.Scanln(&url)
	req, _ := http.NewRequest("DELETE", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

}