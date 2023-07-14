package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"airistaflows/api"
)

// This file is to generate the executable with passed Flow settings

func main(){
	storedFlows := api.ReadFlows()
	flowToBuildName := readSettings()
	var flowToBuild api.Flow
	for _, flow := range storedFlows{
		if (flow.Name == flowToBuildName){
			flowToBuild = flow
		}
	}
	fmt.Println(flowToBuild.Name)

	for{
		time.Sleep(10 * time.Second)
	}// Loop forever
}


func readSettings() string{
	// The file path
	filepath := "./Config.json"

	// Read the file
	contents, err := ioutil.ReadFile(filepath)
	if err != nil {
	  fmt.Println(err)
	  return ""
	}

	// Unmarshal the JSON data into an object
	var config api.Config
	errUnmarshal := json.Unmarshal([]byte(contents), &config)
	if errUnmarshal != nil {
	  fmt.Println(errUnmarshal)
	  return ""
	}

	return config.FlowToBuild
}