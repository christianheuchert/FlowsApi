package api

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

//consoleLog ---------------------------------------------------- consoleLog
func PrintFunction(item interface{}){
	fmt.Println("Item to Print: ", item)
}


//RestCallGetGroups ---------------------------------------------------- RestCallGetGroups
//http://52.45.17.177:802/XpertRestApi/api/MetaData/GetGroups?CustomerId=1
func RestCallGetGroups(IP string, customerId string, username string, password string )string{
	fmt.Println("rest call")

	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "http://"+IP+"/XpertRestApi/api/MetaData/GetGroups?CustomerId="+customerId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err.Error()
	}

	// Add basic authentication to the request header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return err.Error()
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	stringResponse := string(body)

	return stringResponse
}

// AddTen  ----------------------------------------------------  AddTen 
func AddTen(number int)int{
	return number + 10
}