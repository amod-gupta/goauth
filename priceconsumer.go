package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Traceableai/goagent/instrumentation/net/traceablehttp"
)

func getPriceURL() string {
	host, ok := os.LookupEnv("PRICESERVICE_HOST")
	url := ""

	if !ok {
		url = "http://localhost:5000/"
		fmt.Println("PRICESERVICE_HOST is not present")
	} else {
		url = "http://" + host + ":5000/"
		fmt.Println("PRICESERVICE_HOST: " + host)
	}

	return url
}

func getpricehome() []byte {

	url := getPriceURL()

	client := http.Client{
		Transport: traceablehttp.NewTransport(http.DefaultTransport),
	}

	//response, err := http.Get(url)
	response, err := client.Get(url)

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

func getbtcprice() []byte {

	baseurl := getPriceURL()
	url := baseurl + "price/"
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
