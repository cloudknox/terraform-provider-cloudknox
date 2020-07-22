package models

// RolePolicyData is the struct that contains data that will be populated for the body sent to the API
type RolePolicyData struct {
	AuthSystemInfo AuthSystemInfo `json:"authSystemInfo"`
	IdentityType   string         `json:"identityType"`
	IdentityIds    interface{}    `json:"identityIds"`
	Filter         Filter         `json:"filter"`
	RequestParams  *RequestParams `json:"requestParams,omitempty"`
}

// HistoryDuration is the struct that contains data that will be populated for the body sent to the API
type HistoryDuration struct {
	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`
}

// RequestParams is the struct that contains data that will be populated for the body sent to the API
type RequestParams struct {
	Scope     interface{} `json:"scope,omitempty"`
	Resource  interface{} `json:"resource,omitempty"`
	Resources interface{} `json:"resources,omitempty"`
	Condition interface{} `json:"condition,omitempty"`
}

// Filter is a struct used to encompass filter fields for the api
type Filter struct {
	HistoryDays     interface{}      `json:"historyDays,omitempty"`
	PreserveReads   bool             `json:"preserveReads"`
	HistoryDuration *HistoryDuration `json:"historyDuration,omitempty"`
}

// AuthSystemInfo is a struct used to encompass the AuthSystemInfo fields for the api
type AuthSystemInfo struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
