package apiHandler

type PolicyData struct {
	AuthSystemInfo struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"authSystemInfo"`
	IdentityType string      `json:"identityType"`
	IdentityIds  interface{} `json:"identityIds"`
	Filter       struct {
		HistoryDays     int  `json:"historyDays"`
		PreserveReads   bool `json:"preserveReads"`
		HistoryDuration *HD  `json:"historyDuration, omitempty"`
	} `json:"filter"`
	RequestParams *RP `json:"requestParams, omitempty"`
}

type HD struct {
	StartTime int `json:"startTime"`
	EndTime   int `json:"endTime"`
}

type RP struct {
	Scope     interface{} `json:"scope, omitempty"`
	Resource  interface{} `json:"resource, omitempty"`
	Resources interface{} `json:"resources, omitempty"`
	Condition interface{} `json:"condition, omitempty"`
}
