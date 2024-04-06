package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

}

type Routes struct {
	route string `json:"path"`
}

func loadRoutes() {
	jsonFile, err := os.Open("routes.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Loaded routes.")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var routes Routes
	json.Unmarshal(byteValue, &routes)
}
