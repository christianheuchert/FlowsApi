package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"

	"airistaflows/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Globals Variables
var storedFlows = api.ReadFlows("flows.json") // Flows from flows.json
var storedConfig = api.ReadConfig("Config.json") // Flows from flows.json

func main() {
	// API setup and Start
	if (true){ // API On/Off switch
		router := gin.Default()
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:4200"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders: []string{"Content-Type", "Authorization"},
		}))
		router.GET("/flows", getFlows) // return all flows
        router.GET("/config", getConfig) // return config
		router.POST("/flows", postFlow) // create flow
		router.GET("/flows/:name", getFlowByName) // return specific flow
		router.GET("/flowExec/:name", createFlowExec) // create executable for specific flow

		router.Run("localhost:8080")
	}else{
		fmt.Println("Testing Function")
		testingFunction()
	}
}

func testingFunction(){
	api.CreateExecutable(storedFlows[1])
}

// getFlows responds with the list of all Flows as JSON.
func getFlows(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, storedFlows)
}

// getConfig responds with the Config as JSON.
func getConfig(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, storedConfig)
}

// postFlow adds an flow from JSON received in the request body.
func postFlow(c *gin.Context) {
    var newFlow api.Flow

    // Call BindJSON to bind the received JSON to
    // newFlow.
    if err := c.BindJSON(&newFlow); err != nil {
        return
    }

    // Add the new Flow to the slice.
    storedFlows = append(storedFlows, newFlow)
    c.IndentedJSON(http.StatusCreated, newFlow)
}

// getFlowByName locates the Flow whose ID value matches the name
// parameter sent by the client, then returns that Flow as a response.
func getFlowByName(c *gin.Context) {
    name := c.Param("name")

    // Loop over the list of flows, looking for
    // a flow whose name matches the parameter.
    for _, a := range storedFlows {
        if a.Name == name {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "flow not found"})
}

// creates an executable for requested Flow
func createFlowExec(c *gin.Context) {
    name := c.Param("name")

    // Update Config.json FlowToBuild to passed name
    api.UpdateConfig(name)

    // get absolute path to json data
	_, path, _, _ := runtime.Caller(0)
	// step up two levels
	path = filepath.Dir(path)

    // Loop over the list of flows, looking for
    // a flow whose name matches the parameter.
    for _, flow := range storedFlows {
        if flow.Name == name {
			api.CreateExecutable(flow)
            c.IndentedJSON(http.StatusOK, path)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "flow not found"})
}