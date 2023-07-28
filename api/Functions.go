package api

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

//consoleLog
func PrintFunction(item interface{}){
	fmt.Println("Item to Print: ", item)
}


//http://52.45.17.177:802/XpertRestApi/api/MetaData/GetGroups?CustomerId=1
func RestCallGetGroups(IP string, customerId int, uname string, pword string ){
	fmt.Println("rest call")

	// Replace these with your actual credentials
	username := "your_username"
	password := "your_password"

	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "https://api.example.com/some/endpoint"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add basic authentication to the request header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

}