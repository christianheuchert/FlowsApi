package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"airistaflows/api"
)

// This file is to generate the executable with passed Flow settings

//Global Output Channel
var outputChannel = make(chan interface{}) // Output Channel for all Triggers

// Main function gets turned into an executable. 
func main(){
	// get absolute path to json data
	_, path, _, _ := runtime.Caller(0)
	// step up two levels
	path = filepath.Dir(path)
	path = filepath.Dir(path)
	// join absolute path and relative path
	pathFlow:= (path + "\\flows.json")
	pathConfig:= (path + "\\Config.json")

	storedFlows := api.ReadFlows(pathFlow)
	flowToBuildName := readSettings(pathConfig)
	var flowToBuild api.Flow
	for _, flow := range storedFlows{
		if (flow.Name == flowToBuildName){
			flowToBuild = flow
		}
	}

	switch trigger := flowToBuild.Trigger.Type; trigger{
	case "mqtt":
		api.MqttTrigger(flowToBuild, outputChannel)
	default:
		fmt.Println("Trigger Type not found")
	}

	// Channel listener to pass trigger output to Functions
	for{
		triggerOuput := <- outputChannel

		for _, function := range flowToBuild.Functions{
			switch functionName := function.Type; functionName{
			case "consoleLog":
				api.PrintFunction(triggerOuput)
			case "restCall":
				fmt.Println("restCall to come")
			default:
				fmt.Println("Function Type not found")
			}
		}
	}

}


func readSettings(configPath string) string{
	// Read the file
	contents, err := ioutil.ReadFile(configPath)
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