package models

type SlackArgs struct {
	RecordType    string `json:"RecordType"`
	Type          string `json:"Type"`
	TypeCode      int    `json:"TypeCode"`
	Name          string `json:"Name"`
	Tag           string `json:"Tag"`
	MessageStream string `json:"MessageStream"`
	Description   string `json:"Description"`
	Email         string `json:"Email"`
	From          string `json:"From"`
	BouncedAt     string `json:"BouncedAt"`
}

type CommonAPIResponse struct {
	IsError         bool
	ResponseMessage string
	ResponseData    interface{}
}
