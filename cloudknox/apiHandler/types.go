package apiHandler

type PolicyData struct {
	AuthSystemInfo struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"authSystemInfo"`
	IdentityType string      `json:"identityType"`
	IdentityIds  interface{} `json:"identityIds"`
	Filter       struct {
		HistoryDays     interface{}      `json:"historyDays, omitempty"`
		PreserveReads   bool             `json:"preserveReads"`
		HistoryDuration *HistoryDuration `json:"historyDuration, omitempty"`
	} `json:"filter"`
	RequestParams *RequestParams `json:"requestParams, omitempty"`
}

type HistoryDuration struct {
	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`
}

type RequestParams struct {
	Scope     interface{} `json:"scope, omitempty"`
	Resource  interface{} `json:"resource, omitempty"`
	Resources interface{} `json:"resources, omitempty"`
	Condition interface{} `json:"condition, omitempty"`
}
