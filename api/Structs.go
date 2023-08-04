package api

import (
	"encoding/json"
)

// Flow Structs
type Flow struct {
	Id          string `json:"Id"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
	Variables   []Variable
	Triggers    []Trigger  `json:"Triggers"`
	Functions   []Function `json:"Functions"`
}
type Trigger struct {
	Id          string          `json:"Id"`
	Name        string          `json:"Name"`
	Description string          `json:"Description"`
	Type        string          `json:"Type"`
	Input       string          `json:"Input"`
	Output      string          `json:"Output"`
	Settings    json.RawMessage `json:"Settings"`
}
type Function struct {
	Id          string          `json:"Id"`
	Name        string          `json:"Name"`
	Description string          `json:"Description"`
	Type        string          `json:"Type"`
	Input       string          `json:"Input"`
	Output      string          `json:"Output"`
	Settings    json.RawMessage `json:"Settings"`
}

// Data json from Variables.json
type Config struct {
	Triggers    []string `json:"Triggers"`
	Functions   []string `json:"Functions"`
	OutputTypes []string `json:"OutputTypes"`
	InputTypes  []string `json:"InputTypes"`
	FlowToBuild string   `json:"FlowToBuild"`
	CurrentId   string   `json:"CurrentId"`
}

// Unique Settings Structs
type RestCallSettings struct {
	Variables     []Variable `json:"Variables"`
	Method        string     `json:"Method"`
	URL           string     `json:"URL"`
	Authorization string     `json:"Authorization"`
	Username      string     `json:"Username"`
	Password      string     `json:"Password"`
}
type Variable struct {
	Name  string      `json:"Name"`
	Type  string      `json:"Type"`
	Value interface{} `json:"Value"`
}
type MqttSettings struct {
	Protocol string `json:"Protocol"`
	Host     string `json:"Host"`
	Port     string `json:"Port"`
	Topic    string `json:"Topic"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// Airista Department Struct

type Department struct {
	CustomerID               int     `json:"CustomerId"`
	DateCreated              string  `json:"DateCreated"`
	DateUpdated              string  `json:"DateUpdated"`
	Description              string  `json:"Description"`
	EnableTenancy            bool    `json:"EnableTenancy"`
	Name                     string  `json:"Name"`
	TenantID                 string  `json:"TenantId"`
	ElapsedTimeInMillseconds float64 `json:"ElapsedTimeInMillseconds"`
	ErrorMessage             string  `json:"ErrorMessage"`
	SuccessMessage           string  `json:"SuccessMessage"`
	HasError                 bool    `json:"HasError"`
	ID                       int     `json:"Id"`
}

type RestCallGetDepartmentsResponse struct {
	List                     []Department `json:"List"`
	ElapsedTimeInMillseconds float64      `json:"ElapsedTimeInMillseconds"`
	EndDateTime              string       `json:"EndDateTime"`
	ErrorMessage             string       `json:"ErrorMessage"`
	IsDescending             bool         `json:"IsDescending"`
	NumberOfRecordsPerPage   int          `json:"NumberOfRecordsPerPage"`
	PageNumber               int          `json:"PageNumber"`
	SearchText               string       `json:"SearchText"`
	SortColumn               string       `json:"SortColumn"`
	StartDateTime            string       `json:"StartDateTime"`
	TotalRecordsCount        int          `json:"TotalRecordsCount"`
}

type Welcome4 struct {
	List                     []List      `json:"List"` //
	SortField                string      `json:"SortField"`
	StaffSettings            string      `json:"StaffSettings"`
	UseCase                  string      `json:"UseCase"`
	ElapsedTimeInMillseconds float64     `json:"ElapsedTimeInMillseconds"`
	EndDateTime              interface{} `json:"EndDateTime"`
	ErrorMessage             string      `json:"ErrorMessage"`
	IsDescending             bool        `json:"IsDescending"`
	NumberOfRecordsPerPage   int         `json:"NumberOfRecordsPerPage"`
	PageNumber               int         `json:"PageNumber"`
	SearchText               string      `json:"SearchText"`
	SortColumn               string      `json:"SortColumn"`
	StartDateTime            interface{} `json:"StartDateTime"`
	TotalRecordsCount        int         `json:"TotalRecordsCount"`
}

type List struct {
	Address                      interface{}   `json:"Address"`
	AlarisStatus                 interface{}   `json:"AlarisStatus"`
	AlertStatus                  AlertStatus   `json:"AlertStatus"`
	AssociatedDevices            []interface{} `json:"AssociatedDevices"`
	BatteryLevel                 int           `json:"BatteryLevel"`
	BedStatus                    string        `json:"BedStatus"`
	CurrentBuildingID            int           `json:"CurrentBuildingID"`
	CurrentBuildingName          string        `json:"CurrentBuildingName"`
	CurrentSiteName              string        `json:"CurrentSiteName"`
	CurrentFloorName             string        `json:"CurrentFloorName"`
	CurrentMapID                 int           `json:"CurrentMapId"`
	CurrentModelID               int           `json:"CurrentModelId"`
	CurrentSiteID                int           `json:"CurrentSiteID"`
	CurrentTimestamp             string        `json:"CurrentTimestamp"`
	CurrentX                     float64       `json:"CurrentX"`
	CurrentY                     float64       `json:"CurrentY"`
	CurrentZones                 string        `json:"CurrentZones"`
	DepartmentID                 int           `json:"DepartmentID"`
	DepartmentName               string        `json:"DepartmentName"`
	DeviceID                     int           `json:"DeviceID"`
	DeviceLogID                  int           `json:"DeviceLogID"`
	DeviceName                   string        `json:"DeviceName"`
	Email                        string        `json:"Email"`
	EnableAlerts                 bool          `json:"EnableAlerts"`
	EnableHygiene                bool          `json:"EnableHygiene"`
	EnableSDCT                   bool          `json:"EnableSDCT"`
	EventCountAcknowledged       int           `json:"EventCountAcknowledged"`
	EventCountClosed             int           `json:"EventCountClosed"`
	EventCountNew                int           `json:"EventCountNew"`
	EventCountOpen               int           `json:"EventCountOpen"`
	FromLDAP                     bool          `json:"FromLDAP"`
	GroupID                      int           `json:"GroupID"`
	GroupName                    string        `json:"GroupName"`
	HealthStatus                 string        `json:"HealthStatus"`
	Icon                         Icon          `json:"Icon"`
	ImageData                    string        `json:"ImageData"`
	ImageType                    string        `json:"ImageType"`
	IsTestMode                   bool          `json:"IsTestMode"`
	OldTamper                    *bool         `json:"OldTamper"`
	OldMotion                    bool          `json:"OldMotion"`
	Latitude                     int           `json:"Latitude"`
	Longitude                    int           `json:"Longitude"`
	LocationUpdated              *string       `json:"LocationUpdated"`
	ModelName                    string        `json:"ModelName"`
	OldBuildingID                int           `json:"OldBuildingID"`
	OldMapID                     int           `json:"OldMapId"`
	OldModelID                   int           `json:"OldModelId"`
	OldSiteID                    int           `json:"OldSiteID"`
	OldLocationUpdated           string        `json:"OldLocationUpdated"`
	OldX                         int           `json:"OldX"`
	OldY                         int           `json:"OldY"`
	OldZones                     string        `json:"OldZones"`
	PendingDepartmentDateUpdated string        `json:"PendingDepartmentDateUpdated"`
	PendingDepartmentID          int           `json:"PendingDepartmentId"`
	PhoneNumber                  string        `json:"PhoneNumber"`
	Portrait                     PortraitEnum  `json:"Portrait"`
	StaffID                      string        `json:"StaffID"`
	StaffSettings                interface{}   `json:"StaffSettings"`
	Temperature                  string        `json:"Temperature"`
	UseCases                     []interface{} `json:"UseCases"`
	MultiAssign                  bool          `json:"MultiAssign"`
	AssocItemID                  int           `json:"AssocItemID"`
	PendingDepartmentName        *string       `json:"PendingDepartmentName"`
	AssocItemName                *string       `json:"AssocItemName"`
	CustomerID                   int           `json:"CustomerId"`
	DateCreated                  string        `json:"DateCreated"`
	DateUpdated                  string        `json:"DateUpdated"`
	Description                  string        `json:"Description"`
	EnableTenancy                bool          `json:"EnableTenancy"`
	Name                         string        `json:"Name"`
	TenantID                     string        `json:"TenantId"`
	ElapsedTimeInMillseconds     int           `json:"ElapsedTimeInMillseconds"`
	ErrorMessage                 string        `json:"ErrorMessage"`
	SuccessMessage               string        `json:"SuccessMessage"`
	HasError                     bool          `json:"HasError"`
	ID                           int           `json:"Id"`
}

type Zones struct {
	ZoneID   int    `json:"ZoneID"`
	ZoneName string `json:"ZoneName"`
	ZoneType string `json:"ZoneType"`
}

type AlertStatus string

const (
	Blue AlertStatus = "Blue"
)

type Icon string

const (
	AssetsIconsPATIENTMPNG  Icon = "assets\\icons\\PATIENT_M.png"
	IconAssetsIconsSTAFFPNG Icon = "assets/icons/STAFF.png"
)

type PortraitEnum string

const (
	Portrait                    PortraitEnum = ""
	PortraitAssetsIconsSTAFFPNG PortraitEnum = "assets/icons/STAFF.png"
)
