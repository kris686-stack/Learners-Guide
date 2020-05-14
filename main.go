// excercise14.1 project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type msgData struct {
	Message string `json:"message"`
}

func getDataAndReturnResponse() string {
	r, err := http.Get("https://img1.exportersindia.com/product_images/bc-small/2019/9/1448281/wild-fish-1569326265-5093788.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)

	}
	return string(data)
}
func getStructDataAndRes() msgData {
	r, err := http.Get("http://localhost:8000")
	if err != nil {
		log.Fatal(err)

	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)

	}
	msg := msgData{}
	err = json.Unmarshal(data, &msg)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}

func main() {
	fmt.Println("Hello World!")
	//data := getDataAndReturnResponse()
	//fmt.Println(data)
	d := getStructDataAndRes()
	fmt.Println(d.Message)
}
