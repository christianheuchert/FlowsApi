package main

import (
	"fmt"
	"net/http"

	"airistaflows/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Globals Variables
var storedFlows = api.ReadFlows() // Flows from flows.json
var outputChannel = make(chan string) // Output Channel for all Triggers

func main() {
	// API setup and Start
	if (false){ // API On/Off switch
		router := gin.Default()
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:4200"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowHeaders: []string{"Content-Type", "Authorization"},
		}))
		router.GET("/flows", getFlows) // return all flows
		router.POST("/flows", postFlow) // create flow
		router.GET("/flows/:name", getFlowByName) // return specific flow
		router.GET("/flowExec/:name", createFlowExec) // create executable for specific flow

		router.Run("localhost:8080")
	}else{
		fmt.Println("Testing Function")
		testingFunction()
	}
}

func testingFunction(){// Unmarshal the JSON data into an object
	// testData := storedFlows[1].Trigger.Settings
	// var mqttSettings api.MqttSettings
	// errUnmarshal := json.Unmarshal([]byte(testData), &mqttSettings)
	// if errUnmarshal != nil {
	//   fmt.Println(errUnmarshal)
	//   return
	// }
	// api.MqttTrigger(mqttSettings, outputChannel)
	// for{
	// 	time.Sleep(10 * time.Second)
	// }// Loop forever
	api.CreateExecutable(storedFlows[1])
}

// getFlows responds with the list of all Flows as JSON.
func getFlows(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, storedFlows)
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

    // Loop over the list of flows, looking for
    // a flow whose name matches the parameter.
    for _, flow := range storedFlows {
        if flow.Name == name {
			api.CreateExecutable(flow)
            c.IndentedJSON(http.StatusOK, flow)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "flow not found"})
}