package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
)

// Reads in Flows stored in flows.json
func ReadFlows(flowPath string) []Flow{

	// Read the file
	contents, err := ioutil.ReadFile(flowPath)
	if err != nil {
	  fmt.Println(err)
	  return nil
	}

	// Unmarshal the JSON data into an object
	var flowArray []Flow
	errUnmarshal := json.Unmarshal([]byte(contents), &flowArray)
	if errUnmarshal != nil {
	  fmt.Println(errUnmarshal)
	  return nil
	}

	return flowArray
}

// Reads in Config Settings stored in Config.json
func ReadConfig(configPath string) Config{

	// Read the file
	contents, err := ioutil.ReadFile(configPath)
	if err != nil {
	  fmt.Println(err)
	  return Config{}
	}

	// Unmarshal the JSON data into an object
	var config Config
	errUnmarshal := json.Unmarshal([]byte(contents), &config)
	if errUnmarshal != nil {
	  fmt.Println(errUnmarshal)
	  return Config{}
	}

	return config
}

// make executable
func CreateExecutable(flow Flow) {
	//path to the Golang program
    pathToProgram := "./create-executable"

    //output file name for the executable
    outputFileName := flow.Name + ".exe"

    //Create the executable
    cmd := exec.Command("go", "build", "-o", outputFileName, pathToProgram)
    err := cmd.Run()
    if err != nil {
        panic(err)
    }

}

//Update Config.json FlowToBuild
func UpdateConfig(newConfig Config) {
	// Read in File
	configPath := "Config.json"
    contents, err := ioutil.ReadFile(configPath)
	if err != nil {
	  fmt.Println(err)
	  return
	}
	// Unmarshal the config JSON into an object
	var config Config
	errUnmarshal := json.Unmarshal([]byte(contents), &config)
	if errUnmarshal != nil {
	  fmt.Println(errUnmarshal)
	  return
	}
	// update config field if value
	if(newConfig.CurrentId != ""){
		config.CurrentId = newConfig.CurrentId
	}
	if(newConfig.FlowToBuild != ""){
		config.FlowToBuild = newConfig.FlowToBuild
	}

    // Serialize the `Config` object to JSON.
    jsonData, err := json.MarshalIndent(config, "", "    ")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Write the JSON data to the `Config.json` file.
    err = ioutil.WriteFile("Config.json", jsonData, 0644)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // fmt.Println("Config updated successfully!")
}
