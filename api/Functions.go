package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

//RestCallGetStaffByDepartment ---------------------------------------------------- RestCallGetStaffByDepartment
//http://52.45.17.177:802/XpertRestApi/api/MetaData/GetDepartments?CustomerId=1
func RestCallGetAllByDepartment(IP string, customerId string, username string, password string, departmentItem interface{})string{
	fmt.Println("RestCallGetAllByDepartment")

	departmentId := ""

	switch item:= departmentItem.(type){
	case string:
		Response := RestCallGetDepartments(IP, customerId, username, password)
		for _, element := range Response.List{
			if (element.Name == item){// find Department with Name and use Id
				departmentId = strconv.Itoa(element.ID)
				break
			}
		}
	case int:
		departmentId = strconv.Itoa(item)
	case Department: 
		fmt.Println("Department: ", item)
		departmentId = strconv.Itoa(item.ID)
	default:
		fmt.Println("Unknown Input Type:", item)
	}
	if (departmentId == ""){
		return "Error: no Department Info provided, RestCallGetAllByDepartment"
	}

	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "http://"+IP+"/XpertRestApi/api/Staff/GetAllByDepartment?CustomerId="+customerId+"&DepartmentId="+departmentId
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

func RestCallGetDepartments(IP string, customerId string, username string, password string )RestCallGetDepartmentsResponse{

	//Declare response struct object
	var response RestCallGetDepartmentsResponse

	// Create an HTTP client
	client := &http.Client{}

	// Create the request
	url := "http://"+IP+"/XpertRestApi/api/MetaData/GetDepartments?CustomerId="+customerId
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		response.TotalRecordsCount = 0
		response.ErrorMessage = err.Error()
		return response
	}

	// Add basic authentication to the request header
	auth := username + ":" + password
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", basicAuth)

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		response.TotalRecordsCount = 0
		response.ErrorMessage = err.Error()
		return response
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// Unmarshal the config JSON into an object
	errUnmarshal := json.Unmarshal([]byte(body), &response)
	if errUnmarshal != nil {
	 	fmt.Println(errUnmarshal)
	  	response.TotalRecordsCount = 0
		response.ErrorMessage = err.Error()
		return response
	}

	return response
}