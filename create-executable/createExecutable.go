package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"airistaflows/api"
)

// This file is to generate the executable with passed Flow settings

//Global Output Channel
var outputChannel = make(chan interface{}) // Output Channel for all Triggers

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

	switch trigger := flowToBuild.Trigger.Type; trigger{
	case "mqtt":
		api.MqttTrigger(flowToBuild, outputChannel)
	default:
		fmt.Println("Set Trigger not found")
	}

	// Channel listener to pass data to Functions
	for{
		triggerOuput := <- outputChannel

		for _, function := range flowToBuild.Functions{
			
		}

	}
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