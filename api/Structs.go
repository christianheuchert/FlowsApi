package api

import (
	"encoding/json"
)

// Flow Structs
type Flow struct {
	Name        string      `json:"Name"`
	Description string      `json:"Description"`
	Trigger     Trigger     `json:"Trigger"`
	Functions   []Functions `json:"Functions"`
}
type Trigger struct {
	Name         string       `json:"Name"`
	Description  string       `json:"Description"`
	Type         string       `json:"Type"`
	Input        string       `json:"Input"`
	Output       string       `json:"Output"`
	Settings 	 json.RawMessage	  `json:"Settings"`
}
type Functions struct {
	Name             string           `json:"Name"`
	Description      string           `json:"Description"`
	Type             string           `json:"Type"`
	Input            string           `json:"Input"`
	Output           string       	  `json:"Output"`
	Settings 		 json.RawMessage		  `json:"Settings"`
}

// Data json from Variables.json
type Config struct {
	Triggers    []string `json:"Triggers"`
	Functions   []string `json:"Functions"`
	OutputTypes []string `json:"OutputTypes"`
	InputTypes  []string `json:"InputTypes"`
	FlowToBuild string   `json:"FlowToBuild"`
}


// Unique Settings Structs
type RestCallSettings struct {
	Variables     []Variables `json:"Variables"`
	Method        string      `json:"Method"`
	URL           string      `json:"URL"`
	Authorization string      `json:"Authorization"`
	Username      string      `json:"Username"`
	Password      string      `json:"Password"`
}
type Variables struct {
	Name  string `json:"Name"`
	Input string `json:"Input"`
	Value string `json:"Value"`
}
type MqttSettings struct {
	Name     string `json:"Name"`
	Protocol string `json:"Protocol"`
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	Topic 	 string `json:"Topic"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}