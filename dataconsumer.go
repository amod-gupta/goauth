package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"github.com/Traceableai/goagent/instrumentation/net/traceablehttp"
)

/*
func getHyperClient() http.Client {

	client := http.Client{
		Transport: traceablehttp.NewTransport(http.DefaultTransport),
	}

	return client
}*/

func getURL() string {
	host, ok := os.LookupEnv("DATASERVICE_HOST")
	url := ""

	if !ok {
		url = "http://localhost:8080/"
		fmt.Println("DATASERVICE_HOST is not present")
	} else {
		url = "http://" + host + ":8080/"
		fmt.Println("DATASERVICE_HOST: " + host)
	}

	return url
}

func getdata(id string) []byte {

	baseurl := getURL()
	url := baseurl + "customer/test/" + id
	//client := getHyperClient()
	//response, err := client.Get(url)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		// Write error to response body
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func getallcustomers() []byte {

	baseurl := getURL()
	url := baseurl + "customer/all"
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		// Write error to response body
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}

func getcustomerbyid(id string) []byte {

	baseurl := getURL()
	url := baseurl + "customer/byid/" + id
	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		// Write error to response body
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
