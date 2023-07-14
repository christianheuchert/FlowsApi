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
